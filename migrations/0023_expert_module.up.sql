-- 专家检查 + 专家库 模块 — 4 张表覆盖三方协作（计划方 / 受检企业 / 专家）
--
-- expert_inspections        检查计划主表（创建方为 company_id；title 必填；start_date/end_date）
-- expert_companies          检查计划下挂的「受检企业 + 时间表」子表（每条对应一家受检企业）
-- expert_inspection_users   专家指派（每个 expert_companies 行可指派多名专家）
-- bloc_expert_users         专家库 — 集团把系统用户登记为本企业可用专家
--
-- 风格沿用 0022_notification_modules.up.sql：company_id NOT NULL + audit + 软删除。
-- expert_companies / expert_inspection_users 采用 ON DELETE CASCADE 跟着主表删除，
-- expert_user_id / check_company_id 用 RESTRICT 避免误删用户/企业时清空检查记录。
--
-- 时间字段：notify_company_time / check_company_time / confirm_time 在 Laravel 是
-- timestamp 字符串列（前端 datetime-local 输入）；Postgres 改为原生 TIMESTAMPTZ，
-- service 层负责把 "YYYY-MM-DDTHH:MM" 解析成 time.Time。

CREATE TABLE expert_inspections (
    id            BIGSERIAL PRIMARY KEY,
    title         VARCHAR(255) NOT NULL,
    description   TEXT,
    "comment"     TEXT,
    start_date    DATE,
    end_date      DATE,
    company_id    BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id       BIGINT NOT NULL DEFAULT 0,
    last_user_id  BIGINT NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ
);
CREATE INDEX idx_expert_inspections_company    ON expert_inspections (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_expert_inspections_title      ON expert_inspections (title)      WHERE deleted_at IS NULL;
CREATE INDEX idx_expert_inspections_end_date   ON expert_inspections (end_date)   WHERE deleted_at IS NULL;

CREATE TABLE expert_companies (
    id                       BIGSERIAL PRIMARY KEY,
    expert_inspection_id     BIGINT NOT NULL REFERENCES expert_inspections(id) ON DELETE CASCADE,
    check_company_id         BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    notify_company_time      TIMESTAMPTZ,
    check_company_time       TIMESTAMPTZ,
    confirm_time             TIMESTAMPTZ,
    inspection_notice_url    TEXT,
    rectification_sheet_url  TEXT,
    status                   INTEGER NOT NULL DEFAULT 1,
    company_id               BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id                  BIGINT NOT NULL DEFAULT 0,
    last_user_id             BIGINT NOT NULL DEFAULT 0,
    created_at               TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at               TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at               TIMESTAMPTZ
);
CREATE INDEX idx_expert_companies_inspection      ON expert_companies (expert_inspection_id)            WHERE deleted_at IS NULL;
CREATE INDEX idx_expert_companies_check_company   ON expert_companies (check_company_id)                WHERE deleted_at IS NULL;
CREATE INDEX idx_expert_companies_company         ON expert_companies (company_id)                      WHERE deleted_at IS NULL;
CREATE INDEX idx_expert_companies_pending_notify  ON expert_companies (check_company_id, confirm_time)  WHERE deleted_at IS NULL;

CREATE TABLE expert_inspection_users (
    id                  BIGSERIAL PRIMARY KEY,
    expert_user_id      BIGINT NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    expert_company_id   BIGINT NOT NULL REFERENCES expert_companies(id) ON DELETE CASCADE,
    notify_time         TIMESTAMPTZ,
    status              INTEGER NOT NULL DEFAULT 1,
    confirm_time        TIMESTAMPTZ,
    "comment"           TEXT,
    company_id          BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id             BIGINT NOT NULL DEFAULT 0,
    last_user_id        BIGINT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_expert_inspection_users_company  ON expert_inspection_users (expert_company_id)                  WHERE deleted_at IS NULL;
CREATE INDEX idx_expert_inspection_users_expert   ON expert_inspection_users (expert_user_id)                     WHERE deleted_at IS NULL;
CREATE INDEX idx_expert_inspection_users_pending  ON expert_inspection_users (expert_user_id, confirm_time)       WHERE deleted_at IS NULL;

CREATE TABLE bloc_expert_users (
    id              BIGSERIAL PRIMARY KEY,
    expert_user_id  BIGINT NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    company_id      BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id         BIGINT NOT NULL DEFAULT 0,
    last_user_id    BIGINT NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);
CREATE INDEX idx_bloc_expert_users_company  ON bloc_expert_users (company_id)     WHERE deleted_at IS NULL;
CREATE INDEX idx_bloc_expert_users_expert   ON bloc_expert_users (expert_user_id) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX uniq_bloc_expert_users_pair
    ON bloc_expert_users (company_id, expert_user_id)
    WHERE deleted_at IS NULL;
