-- 风险监测栏目（集团视图）— 两张支撑表
--
-- focus_companies   集团重点关注企业关系（company_id 集团 → focus_id 被关注企业）
-- risk_logs         集团下属企业每日风险指数快照，type=1 固有 / 2 浮动，date 为 YYYYMMDD 整数
--
-- 集团视图 5 模块（企业分布图 / 重大危险源分布图 / 在线巡检 / 风险监控 / 特种设备数据统计）
-- 中前两者复用现有 companies / major_hazards 表，后三者依赖本迁移 + 现有 special_equipment* /
-- monthly_investigations 表做跨企业聚合，不再新增其它表。

-- focus_id ON DELETE CASCADE：被关注企业删了关系自然清掉；company_id 主体不能误删。
CREATE TABLE focus_companies (
    id            BIGSERIAL PRIMARY KEY,
    focus_id      BIGINT REFERENCES companies(id) ON DELETE CASCADE,
    company_id    BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id       BIGINT NOT NULL DEFAULT 0,
    last_user_id  BIGINT NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ
);
CREATE INDEX idx_focus_companies_company ON focus_companies (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_focus_companies_focus   ON focus_companies (focus_id)   WHERE deleted_at IS NULL;

CREATE TABLE risk_logs (
    id              BIGSERIAL PRIMARY KEY,
    "type"          SMALLINT NOT NULL DEFAULT 0,
    old_risk_num    NUMERIC(10, 2),
    risk_num        NUMERIC(10, 2) NOT NULL DEFAULT 0,
    "date"          INTEGER NOT NULL DEFAULT 0,
    "content"       VARCHAR(255),
    company_name    VARCHAR(255),
    old_risk_grade  VARCHAR(32),
    risk_grade      VARCHAR(32),
    company_id      BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id         BIGINT NOT NULL DEFAULT 0,
    last_user_id    BIGINT NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);
CREATE INDEX idx_risk_logs_company      ON risk_logs (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_risk_logs_company_date ON risk_logs (company_id, "date") WHERE deleted_at IS NULL;
CREATE INDEX idx_risk_logs_date_type    ON risk_logs ("date", "type") WHERE deleted_at IS NULL;

COMMENT ON COLUMN risk_logs.company_name IS 'Snapshot at row creation; does not track later companies.title renames. Join via company_id for current names.';
