// Command risk-calc enqueues a single risk:calculate task on demand. Use
// when you want to re-run today's risk recalculation outside the 02:00
// scheduler tick (e.g. after fixing a data issue), without waiting until
// tomorrow.
//
//	go run ./cmd/risk-calc -config ./configs/config.yaml -content "manual run"
//
// The task lands on the same default queue the worker consumes, so the
// scheduler's Unique(24h) window applies — manual runs only succeed if no
// run has already fired today. Pass -force to override the dedup window.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/hibiken/asynq"

	"ai_hs_go/internal/config"
	"ai_hs_go/internal/queue"
)

func main() {
	configPath := flag.String("config", "", "path to config yaml")
	content := flag.String("content", "手动触发", "trigger label persisted on each risk_logs row")
	force := flag.Bool("force", false, "bypass the daily Unique(24h) dedup window")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load config: %v\n", err)
		os.Exit(1)
	}

	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: cfg.Redis.Addr, Password: cfg.Redis.Password, DB: cfg.Redis.DB,
	})
	defer client.Close()

	task, err := queue.NewCalculateRiskTask(queue.CalculateRiskPayload{Content: *content})
	if err != nil {
		fmt.Fprintf(os.Stderr, "build task: %v\n", err)
		os.Exit(1)
	}

	opts := []asynq.Option{}
	if *force {
		// `Unique(0)` disables the dedup key so this enqueue always sticks.
		opts = append(opts, asynq.Unique(0))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	info, err := client.EnqueueContext(ctx, task, opts...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "enqueue: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("enqueued risk:calculate id=%s queue=%s content=%q\n", info.ID, info.Queue, *content)
}
