package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"

	"ai_hs_go/internal/config"
	"ai_hs_go/internal/queue"
	"ai_hs_go/internal/repository"
)

func main() {
	configPath := flag.String("config", "", "path to config yaml")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load config: %v\n", err)
		os.Exit(1)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Fprintf(os.Stderr, "init logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	db, err := repository.Open(cfg, logger)
	if err != nil {
		logger.Fatal("connect db", zap.Error(err))
	}

	loc, err := time.LoadLocation(cfg.Database.TimeZone)
	if err != nil {
		logger.Fatal("load scheduler timezone", zap.String("tz", cfg.Database.TimeZone), zap.Error(err))
	}

	handlers := &queue.Handlers{DB: db, Logger: logger, TZ: loc}
	mux := asynq.NewServeMux()
	handlers.Register(mux)

	redisOpt := asynq.RedisClientOpt{Addr: cfg.Redis.Addr, Password: cfg.Redis.Password, DB: cfg.Redis.DB}

	sched := asynq.NewScheduler(redisOpt, &asynq.SchedulerOpts{
		Logger:   zapAsynqLogger{logger},
		Location: loc,
	})
	riskTask, err := queue.NewCalculateRiskTask(queue.CalculateRiskPayload{Content: "每日定时计算"})
	if err != nil {
		logger.Fatal("build risk:calculate task", zap.Error(err))
	}
	if _, err := sched.Register("0 2 * * *", riskTask); err != nil {
		logger.Fatal("register risk:calculate", zap.Error(err))
	}
	if err := sched.Start(); err != nil {
		logger.Fatal("scheduler start", zap.Error(err))
	}
	defer sched.Shutdown()

	srv := asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: 10,
		Queues: map[string]int{
			queue.QueueDefault: 6,
			queue.QueueExports: 4,
		},
		Logger: zapAsynqLogger{logger},
	})

	logger.Info("worker started", zap.String("scheduler_tz", loc.String()))
	if err := srv.Run(mux); err != nil {
		logger.Fatal("worker run", zap.Error(err))
	}
}

type zapAsynqLogger struct {
	z *zap.Logger
}

func (l zapAsynqLogger) Debug(args ...any) { l.z.Sugar().Debugln(args...) }
func (l zapAsynqLogger) Info(args ...any)  { l.z.Sugar().Infoln(args...) }
func (l zapAsynqLogger) Warn(args ...any)  { l.z.Sugar().Warnln(args...) }
func (l zapAsynqLogger) Error(args ...any) { l.z.Sugar().Errorln(args...) }
func (l zapAsynqLogger) Fatal(args ...any) { l.z.Sugar().Fatalln(args...) }
