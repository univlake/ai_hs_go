package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"

	"ai_hs_go/internal/config"
	"ai_hs_go/internal/rbac"
	"ai_hs_go/internal/repository"
	"ai_hs_go/internal/service"
)

func main() {
	configPath := flag.String("config", "", "config yaml path (default: ./configs/config.yaml or env)")
	username := flag.String("admin", "superadmin", "platform super admin username")
	password := flag.String("password", "", "platform super admin initial password (default: Aifix@2026)")
	modelPath := flag.String("model", "./configs/casbin_model.conf", "casbin model path")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load config: %v\n", err)
		os.Exit(1)
	}
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	db, err := repository.Open(cfg, logger)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open db: %v\n", err)
		os.Exit(1)
	}

	abs, _ := filepath.Abs(*modelPath)
	enforcer, err := rbac.NewEnforcer(db, abs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "casbin: %v\n", err)
		os.Exit(1)
	}

	if err := service.Seed(context.Background(), db, enforcer, service.SuperAdminCredentials{
		Username: *username,
		Password: *password,
	}); err != nil {
		fmt.Fprintf(os.Stderr, "seed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("seed: ok")
}
