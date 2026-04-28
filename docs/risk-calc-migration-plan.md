# 风险重算定时任务迁移计划（Laravel → Go）

把 aifix_hs 的 `risk:calculate` Artisan 命令搬到 ai_hs_go 的 worker 进程，每天定时跑全量企业风险重算。

## 来源（Laravel 规范，逐字翻译目标）
- `routes/console.php:12-16` — `dailyAt('02:00')->withoutOverlapping()->onOneServer()->runInBackground()`
- `app/Console/Commands/CalculateCompanyRisk.php` — 薄壳，调 trait
- `app/Traits/CalculateRiskValue.php` — **真正的规范**，约 600 行；重算所有企业固有/浮动风险并写 `inherent_risk_logs`、`float_risk_logs`、`risk_logs`，再算集团平均

## 已确认的决策

| # | 决策 |
|---|---|
| 1 | **A1 + B1**：在 `cmd/worker/main.go` 里加 `asynq.Scheduler`；重算逻辑放新建的 `internal/risk/` 包 |
| 2 | 三处 SQL 类型偏差走**内联 cast**，不动 migration |
| 3 | 事务语义**照搬 Laravel**（仅 `batchUpdateCompanies` 在事务里，bulk inserts / bloc updates 不在），保证结果正确即可 |
| 4 | **时区可配**：复用现有的 `database.timezone`（默认 `Asia/Shanghai`，env `DB_TIMEZONE`），Scheduler 与 Calculator 都从这里读取，保证调度器的"今天"和 SQL 的"今天"一致 |
| 5 | 当前单 worker；`asynq.Unique(24h)` 足够，无需 leader election |

## 阶段总览

每个阶段都有 verification gate，过了再开下一阶段。

---

## Phase 0 — 覆盖审计（已完成）

Trait 引用的 30+ 个表 / 模型在 Go 侧全部已存在。仅 3 处 SQL 类型偏差需要在 Go 查询里显式 cast：

| 列 | 位置 | Laravel 用法 | Go 需要 |
|---|---|---|---|
| `major_hazardous_processes.operating_people_num` | `model/process.go` | `> 3`（数字比较） | `CAST(NULLIF(...,'') AS INTEGER) > 3` |
| `hazardous_processes.operating_people_num` | `model/process.go` | 同上 | 同上 |
| `emergency_teams.team_number` | `model/accident_emergency.go` | `SUM(team_number)` | `SUM(COALESCE(NULLIF(...,'')::numeric,0))` |
| `three_togethers.status` | `model/risk.go` | `where status = 4` | `WHERE "status" = '4'` |

**Verification gate 0**：上面四处在 migration 里确实是 `VARCHAR`（开工时再 grep `migrations/0015_process_modules.up.sql`、`0021_accident_emergency_module.up.sql`、`0020_risk_module.up.sql` 确认一次）。

---

## Phase 1 — 架构（已定）

- **A1**：scheduler 进 worker。`cmd/worker/main.go` 加 ~40 行：`NewScheduler` → `Register("0 2 * * *", task)` → 与 `srv.Run(mux)` 并行运行，SIGTERM 时双双 graceful shutdown。
- **B1**：重算逻辑落 `internal/risk/calculator.go`（目前 `internal/risk/` 是空的）。`internal/queue/handlers/risk.go` 只是个薄 handler，`Calculator{DB, Logger, TZ}` 才是主角。
- **B 决策细化**：CRUD 风格的 `internal/service/risk_service.go` 不动；`internal/service/risk_monitor_service.go` 也不动（它只读重算后的列，不会和 Calculator 冲突）。

---

## Phase 2 — asynq 接线（不含业务逻辑）

目标：scheduler / 队列 / handler stub 全通，能验证 02:00 的 tick 触发了一次空跑。

### 2.1 任务类型 — `internal/queue/tasks.go`
```go
const TypeCalculateRisk = "risk:calculate"

type CalculateRiskPayload struct {
    Content string `json:"content"`
}

func NewCalculateRiskTask(p CalculateRiskPayload) (*asynq.Task, error) {
    body, err := json.Marshal(p)
    if err != nil { return nil, err }
    return asynq.NewTask(TypeCalculateRisk, body,
        asynq.Queue(QueueDefault),
        asynq.MaxRetry(0),                    // 不重试，避免风暴
        asynq.Unique(24*time.Hour),           // ≈ withoutOverlapping
    ), nil
}
```

### 2.2 Handler stub — `internal/queue/handlers/risk.go`（新文件）
方法 `HandleCalculateRisk(ctx, t) error`，先只 `logger.Info("risk:calculate received", payload)` 返回 nil。在 `internal/queue/handlers.go` 的 `Register` 里挂上去。

### 2.3 Scheduler — `cmd/worker/main.go`
```go
loc, err := time.LoadLocation(cfg.Database.TimeZone)  // 复用现有配置
if err != nil { logger.Fatal("load timezone", zap.Error(err)) }

sched := asynq.NewScheduler(redisOpt, &asynq.SchedulerOpts{
    Logger:   zapAsynqLogger{logger},
    Location: loc,
})
task, _ := queue.NewCalculateRiskTask(queue.CalculateRiskPayload{Content: "每日定时计算"})
if _, err := sched.Register("0 2 * * *", task,
    asynq.MaxRetry(0), asynq.Unique(24*time.Hour),
); err != nil { logger.Fatal("register risk:calculate", zap.Error(err)) }

go func() {
    if err := sched.Run(); err != nil { logger.Error("scheduler run", zap.Error(err)) }
}()
// 现有 srv.Run(mux) 保持不变；signal.NotifyContext 处理两边的 shutdown
```

### 2.4 Laravel → asynq 语义映射
| Laravel | asynq | 备注 |
|---|---|---|
| `dailyAt('02:00')` | `"0 2 * * *"` + `SchedulerOpts.Location` | **必须显式设 Location**，否则按 UTC 跑（=北京 10:00） |
| `withoutOverlapping()` | `asynq.Unique(24*time.Hour)` | Redis dedup；窗口 ≤ cron 间隔即可 |
| `onOneServer()` | 自动等价 | scheduler 只 enqueue，Redis 让一个 worker 出队；多副本 + Unique 也只会有一次 |
| `runInBackground()` | 天然等价 | scheduler enqueue 立返，worker 异步消费 |

### Verification gate 2
- `go build ./...` 通过
- 本地 hack `Register` 的 cron 到近 1 分钟，确认 stub 日志打出
- asynq Redis key 里能看到 scheduled entry

---

## Phase 3 — 逻辑搬运（按子步骤拆 commit，逐字翻译，不重构）

`Calculator` 方法 ↔ trait 方法对照：

| Calculator | trait |
|---|---|
| `Run(ctx, content)` | `countRisk` |
| `resetTodayLogs(ctx, today)` | `resetTodayLogs` |
| `loadBulk(ctx, ids, year, tenYearsAgo)` | `loadBulkMetrics` |
| `computePhase1(...)` | `computePhase1Metrics` |
| `computeInherent(...)` | `computeInherentRisk` |
| `computeFloat(...)` | `computeFloatRisk` |
| `computeBlocs(...)` | `calculateBlocRisk` |
| `loadCompanyTree`, `subtreeDepth` | 同名 |
| `batchUpdateCompanies`（事务） | 同名 |
| `bulkInsertChunked(table, rows, 500)` | 同名 |
| `normalized` / `normalizedInverse` / `industryWeight` / `riskGrade` | 同名 |

### 时区注入
`Calculator` 持 `tz *time.Location`（构造时由 handler 从 `cfg.Database.TimeZone` 解析后传入）。下面所有"今天 / 当前年"的取数都走 `time.Now().In(c.tz)`，`whereDate` / `whereYear` 的 SQL 也用 `tz` 算出的时间区间。

### 子步骤
1. **3.1 常量与辅助函数** — `industryWeights`、`INHERENT_WEIGHTS`、`FLOAT_WEIGHTS`、`riskGrade` 6 段桶分级；枚举字符串值（`"一级"`、`"已整改"` 等）原样保留为字符串常量
2. **3.2 Bulk loader** — 30 个 GORM 查询打到 `bulk` 结构里，aggregation 用小 struct + `Scan`；三处 SQL cast 在这步实现
3. **3.3 Phase 1 metrics** — 纯内存循环，构 `riskMaxByCid` 和 `globalMax`
4. **3.4 Inherent (Phase 2)** — 21 个指标 / 公司，输出 `inherentLogRows` + `riskLogRows`，再 `batchUpdateCompanies` 写 `inherent_risk` / `inherent_risk_grade`
5. **3.5 Float (Phase 3)** — 17 个指标 / 公司，`accident_bulletin_severe_scale` 的优先级 ladder 原样翻译
6. **3.6 Bulk inserts + risk_logs 合并** — 列对齐做"统一列归一化"（trait 里的 line 837-852）；`risk_logs` 是 inherent + float 合并后一次性插入
7. **3.7 Bloc averaging** — `loadCompanyTree` 一查到底，`subtreeDepth` 递归；按 depth 升序 update bloc，把"集团均值覆盖 bloc 自身生产分"那条注释 **逐字保留** 在 doc comment 里
8. **3.8 接线 handler** — `HandleCalculateRisk` 解 payload → `risk.NewCalculator(db, logger, tz).Run(ctx, p.Content)` → 返错

### Verification gate 3
- 每一步 `go build ./...` 通过
- Phase 4 的对账 diff 通过

---

## Phase 4 — 测试 & 对账

### 单测（`internal/risk/calculator_test.go`，go-sqlmock 已在 go.mod）
- `TestRiskGrade` — 6 个边界（4.99/5/9.99/10/14.99/15）
- `TestNormalized` / `TestNormalizedInverse` — 零分母 / 零分子 / 正常
- `TestIndustryWeight` — `"1"/"2"/"3"` + 缺失 key
- `TestSubtreeDepth` — leaf / 单子 / 钻石图
- `TestComputePhase1Metrics` — 3 公司，含无 personnel、零 acreage 边界
- `TestRiskLogRowShape` — 列数 / 列名与 migration 对齐

### 集成测（`calculator_integration_test.go`，build tag `integration`）
testcontainers Postgres → 全部 migration → seed（5 公司，1 bloc + 2 生产子）→ `Calculator.Run` → 断言：
- `companies.inherent_risk` 已写
- 当日 `inherent_risk_logs` / `float_risk_logs` 各 5 行
- `risk_logs` 行数 = 2 × 生产公司 + 2 × bloc

### Laravel 对账（手动一次性跑，**这是抓翻译漂移的关键**）
1. 拉一份 staging 快照存为 fixture
2. 对 fresh Postgres 跑 Laravel `php artisan risk:calculate --content=parity`，导出 4 张表为 CSV
3. truncate 4 张风险表 + 重置 `companies` 4 个风险列为 NULL
4. 跑 Go `Calculator.Run(ctx, "parity")`，再次导出
5. **diff 4 个 CSV，可接受差异只有 `created_at` / `updated_at` 时间戳**；其它差异 = 翻译 bug

### Verification gate 4
- 单测全绿
- 集成测对 fresh DB 通过
- 对账 diff 仅时间戳（fixture ≥ 5 生产公司 + ≥ 1 bloc）

---

## Phase 5 — 运维加固

非阻塞，每条一行：
1. **日志** — 起止打 `zap.Info`，含 `company_count`、`duration_ms`、各表写入行数；错误带 `phase` 名称
2. **超时** — handler 包 `context.WithTimeout(ctx, 30*time.Minute)`，防止 runaway 阻塞 default queue
3. **手动触发器** — 新建 `cmd/risk-calc/main.go`，加载 config → 起 asynq client → enqueue 一次（带 `--content`），方便运维不等 02:00

---

## 关键文件清单

| 文件 | 状态 |
|---|---|
| `internal/risk/calculator.go` | 新建 |
| `internal/queue/handlers/risk.go` | 新建 |
| `internal/queue/tasks.go` | 修改（加 `TypeCalculateRisk` + `NewCalculateRiskTask`） |
| `internal/queue/handlers.go` | 修改（注册新 handler） |
| `cmd/worker/main.go` | 修改（加 Scheduler，从 cfg 读 tz） |
| `cmd/risk-calc/main.go` | 新建（Phase 5，可后置） |
| `internal/risk/calculator_test.go` | 新建 |
| `internal/risk/calculator_integration_test.go` | 新建 |
| `app/Traits/CalculateRiskValue.php`（aifix_hs） | 规范源，每个子步骤都要回去对照 |

## 备注 / 注意事项

- `change_log` / `activity_log` —— Laravel 没写，Go **不要主动加**
- 时间戳列直接传 `time.Time`，让 GORM 序列化 `TIMESTAMPTZ`；不要传格式化字符串
- "今天"过滤用 `created_at >= today_midnight AND created_at < tomorrow_midnight`，**不要** `created_at::date = ?`（吃不到索引）
- bulk insert 不要加 `ON CONFLICT DO NOTHING`，trait 已先删今日行，重复 = bug
