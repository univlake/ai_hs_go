-- 考核评分模块 — 5 个维度的填报/评分 + 模板管理 + 集团汇总。
--
-- 数据模型对齐 Laravel inspections.php：
--   daily_check_templates / daily_checks                 日常目标管理工作（filing 流：草稿→已提交→已评分）
--   accident_check_templates / accident_checks           安全生产事故考核（直接评分）
--   add_points_templates / add_points                    加分项（直接评分）
--   system_checks                                        活动参与及材料上报（自动计算 v1..v5 + 总分）
--   norm_checks                                          安全生产标准化（单条 score per company+year）
--
-- 模板表（daily/accident/add_points）共用 pid 自引用结构，level=1 是一级指标，level=2 是二级。
-- 每年一份，前端按 ?year= 拉树。考核模板汇总接口 /assessment-templates/years 取并集。

CREATE TABLE daily_check_templates (
    id           BIGSERIAL PRIMARY KEY,
    title        TEXT         NOT NULL,
    level        INTEGER      NOT NULL DEFAULT 1,
    sort         INTEGER      NOT NULL DEFAULT 0,
    "year"       INTEGER      NOT NULL DEFAULT 2022,
    pid          BIGINT       NOT NULL DEFAULT 0,
    score_range  VARCHAR(255) NOT NULL DEFAULT '1',
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMPTZ
);
CREATE INDEX idx_daily_check_templates_year ON daily_check_templates ("year") WHERE deleted_at IS NULL;
CREATE INDEX idx_daily_check_templates_pid  ON daily_check_templates (pid)    WHERE deleted_at IS NULL;

CREATE TABLE daily_checks (
    id                       BIGSERIAL PRIMARY KEY,
    daily_check_template_id  BIGINT       NOT NULL REFERENCES daily_check_templates(id) ON DELETE RESTRICT,
    path                     TEXT,
    "number"                 NUMERIC(8,1) NOT NULL DEFAULT 0,
    "comment"                TEXT,
    status                   SMALLINT     NOT NULL DEFAULT 0,
    "year"                   INTEGER      NOT NULL DEFAULT 2022,
    company_id               BIGINT       NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id                  BIGINT       NOT NULL DEFAULT 0,
    last_user_id             BIGINT       NOT NULL DEFAULT 0,
    created_at               TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at               TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at               TIMESTAMPTZ
);
CREATE INDEX idx_daily_checks_company  ON daily_checks (company_id)            WHERE deleted_at IS NULL;
CREATE INDEX idx_daily_checks_status   ON daily_checks (status)                WHERE deleted_at IS NULL;
CREATE INDEX idx_daily_checks_lookup   ON daily_checks (company_id, "year")    WHERE deleted_at IS NULL;
CREATE INDEX idx_daily_checks_template ON daily_checks (daily_check_template_id) WHERE deleted_at IS NULL;

CREATE TABLE accident_check_templates (
    id           BIGSERIAL PRIMARY KEY,
    title        TEXT         NOT NULL,
    level        INTEGER      NOT NULL DEFAULT 1,
    sort         INTEGER      NOT NULL DEFAULT 0,
    "year"       INTEGER      NOT NULL DEFAULT 2022,
    pid          BIGINT       NOT NULL DEFAULT 0,
    score_range  VARCHAR(255) NOT NULL DEFAULT '1',
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMPTZ
);
CREATE INDEX idx_accident_check_templates_year ON accident_check_templates ("year") WHERE deleted_at IS NULL;
CREATE INDEX idx_accident_check_templates_pid  ON accident_check_templates (pid)    WHERE deleted_at IS NULL;

CREATE TABLE accident_checks (
    id                          BIGSERIAL PRIMARY KEY,
    accident_check_template_id  BIGINT       NOT NULL REFERENCES accident_check_templates(id) ON DELETE RESTRICT,
    path                        TEXT,
    "number"                    NUMERIC(8,1) NOT NULL DEFAULT 0,
    "comment"                   TEXT,
    "year"                      INTEGER      NOT NULL DEFAULT 2022,
    company_id                  BIGINT       NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id                     BIGINT       NOT NULL DEFAULT 0,
    last_user_id                BIGINT       NOT NULL DEFAULT 0,
    created_at                  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at                  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at                  TIMESTAMPTZ
);
CREATE INDEX idx_accident_checks_company  ON accident_checks (company_id)         WHERE deleted_at IS NULL;
CREATE INDEX idx_accident_checks_lookup   ON accident_checks (company_id, "year") WHERE deleted_at IS NULL;
CREATE INDEX idx_accident_checks_template ON accident_checks (accident_check_template_id) WHERE deleted_at IS NULL;

CREATE TABLE add_points_templates (
    id           BIGSERIAL PRIMARY KEY,
    title        TEXT         NOT NULL,
    level        INTEGER      NOT NULL DEFAULT 1,
    sort         INTEGER      NOT NULL DEFAULT 0,
    "year"       INTEGER      NOT NULL DEFAULT 2022,
    pid          BIGINT       NOT NULL DEFAULT 0,
    score_range  VARCHAR(255) NOT NULL DEFAULT '1',
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMPTZ
);
CREATE INDEX idx_add_points_templates_year ON add_points_templates ("year") WHERE deleted_at IS NULL;
CREATE INDEX idx_add_points_templates_pid  ON add_points_templates (pid)    WHERE deleted_at IS NULL;

CREATE TABLE add_points (
    id                       BIGSERIAL PRIMARY KEY,
    add_points_template_id   BIGINT       NOT NULL REFERENCES add_points_templates(id) ON DELETE RESTRICT,
    path                     TEXT,
    "number"                 NUMERIC(8,1) NOT NULL DEFAULT 0,
    "comment"                TEXT,
    "year"                   INTEGER      NOT NULL DEFAULT 2022,
    company_id               BIGINT       NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id                  BIGINT       NOT NULL DEFAULT 0,
    last_user_id             BIGINT       NOT NULL DEFAULT 0,
    created_at               TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at               TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at               TIMESTAMPTZ
);
CREATE INDEX idx_add_points_company  ON add_points (company_id)         WHERE deleted_at IS NULL;
CREATE INDEX idx_add_points_lookup   ON add_points (company_id, "year") WHERE deleted_at IS NULL;
CREATE INDEX idx_add_points_template ON add_points (add_points_template_id) WHERE deleted_at IS NULL;

-- 系统考核 — 活动参与及材料上报情况（按企业 + 年份单条记录，由 /system-checks/compute 自动计算）。
--   v1 月报覆盖率（满分 4） / v2 年报（停用） / v3 通知反馈（停用）
--   v4 任务反馈率（满分 4） / v5 信息稿（满分 2，每 8 篇满分）
CREATE TABLE system_checks (
    id            BIGSERIAL PRIMARY KEY,
    number_v1     REAL         NOT NULL DEFAULT 0,
    number_v2     REAL         NOT NULL DEFAULT 0,
    number_v3     REAL         NOT NULL DEFAULT 0,
    number_v4     REAL         NOT NULL DEFAULT 0,
    number_v5     REAL         NOT NULL DEFAULT 0,
    "number"      INTEGER      NOT NULL DEFAULT 0,
    "year"        INTEGER      NOT NULL DEFAULT 2022,
    company_id    BIGINT       NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id       BIGINT       NOT NULL DEFAULT 0,
    last_user_id  BIGINT       NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ
);
CREATE INDEX idx_system_checks_company ON system_checks (company_id)         WHERE deleted_at IS NULL;
CREATE INDEX idx_system_checks_lookup  ON system_checks (company_id, "year") WHERE deleted_at IS NULL;

-- 标准化考核 — 大中型企业安全生产标准化管理体系构建（按企业 + 年份单条记录）。
--   number 是 0..100 的原始得分；汇总时按年份折算（year>2025 系数 30，否则 20）。
--   type ∈ {'危化','机械','造纸','其他'}。
CREATE TABLE norm_checks (
    id            BIGSERIAL PRIMARY KEY,
    path          TEXT,
    "number"      INTEGER      NOT NULL DEFAULT 0,
    "year"        INTEGER      NOT NULL DEFAULT 2022,
    "type"        VARCHAR(50)  NOT NULL,
    company_id    BIGINT       NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id       BIGINT       NOT NULL DEFAULT 0,
    last_user_id  BIGINT       NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ
);
CREATE INDEX idx_norm_checks_company ON norm_checks (company_id)         WHERE deleted_at IS NULL;
CREATE INDEX idx_norm_checks_lookup  ON norm_checks (company_id, "year") WHERE deleted_at IS NULL;
