package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"ai_hs_go/internal/api"
	"ai_hs_go/internal/auth"
	"ai_hs_go/internal/config"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/rbac"
	"ai_hs_go/internal/repository"
	"ai_hs_go/internal/storage"
)

func main() {
	configPath := flag.String("config", "", "path to config yaml (default: ./configs/config.yaml or env)")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load config: %v\n", err)
		os.Exit(1)
	}

	logger, err := newLogger(cfg.App.LogLevel, cfg.App.Env)
	if err != nil {
		fmt.Fprintf(os.Stderr, "init logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	db, err := repository.Open(cfg, logger)
	if err != nil {
		logger.Fatal("connect db", zap.Error(err))
	}

	rsaSvc, err := auth.NewRSAService(cfg.RSA.PrivateKeyPath, cfg.RSA.PublicKeyPath)
	if err != nil {
		logger.Fatal("load rsa keys", zap.Error(err))
	}
	jwtSvc := auth.NewJWTService(cfg.JWT.Secret, cfg.JWT.TTL(), cfg.JWT.Issuer)

	enforcer, err := rbac.NewEnforcer(db, "./configs/casbin_model.conf")
	if err != nil {
		logger.Fatal("init casbin", zap.Error(err))
	}

	// In dev (env=local) every browser request lands as 127.0.0.1 through the
	// vite proxy, which makes the per-IP guest quota trivial to exhaust. Skip
	// limiting in dev; production keeps the quotas defined in config.yaml.
	if cfg.App.Env == "local" {
		middleware.SetRateLimitEnabled(false)
		logger.Info("rate limiting disabled (dev environment)")
	}

	var ossClient *storage.Client
	if cfg.OSS.AccessKeyID != "" && cfg.OSS.AccessKeySecret != "" && cfg.OSS.Bucket != "" {
		ossClient, err = storage.New(cfg.OSS)
		if err != nil {
			logger.Warn("oss init", zap.Error(err))
		}
	} else {
		logger.Info("oss not configured; upload endpoints will return 503")
	}

	router := api.New(api.Deps{
		Config:   cfg,
		Logger:   logger,
		DB:       db,
		RSA:      rsaSvc,
		JWT:      jwtSvc,
		Enforcer: enforcer,
		OSS:      ossClient,
	})

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.App.Port),
		Handler:           router.Engine(),
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		logger.Info("server listening", zap.Int("port", cfg.App.Port))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("listen", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("shutdown", zap.Error(err))
	}
}

func newLogger(level, env string) (*zap.Logger, error) {
	var cfg zap.Config
	if env == "local" {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}
	if level != "" {
		var lv zap.AtomicLevel
		if err := lv.UnmarshalText([]byte(level)); err == nil {
			cfg.Level = lv
		}
	}
	return cfg.Build()
}
