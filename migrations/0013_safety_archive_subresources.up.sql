-- 安全档案模块的 6 张子资源表
--
-- 起因：人员管理页面里的 4 个统计卡 + 安全资质 / 安全投入子页 / 安全人员
-- 管理 3 个子表 — 这些都是各自独立的资源，前端拆得很碎。一并补齐。

-- 1) 持证情况：三类人员（type 1-3）+ 消防控制室（type 4-6）共用一张表，
--    通过 type 区分页签
CREATE TABLE certificate_statuses (
    id                       BIGSERIAL PRIMARY KEY,
    name                     VARCHAR(255),
    type                     SMALLINT NOT NULL DEFAULT 1,
    should_obtain_people     INTEGER,
    already_obtain_people    INTEGER,
    no_obtain_people         VARCHAR(255), -- 说明性文字，不是数字
    train_people             INTEGER,
    company_id               BIGINT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    user_id                  BIGINT NOT NULL DEFAULT 0,
    last_user_id             BIGINT NOT NULL DEFAULT 0,
    created_at               TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at               TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at               TIMESTAMPTZ
);
CREATE INDEX idx_certificate_statuses_company_type ON certificate_statuses (company_id, type) WHERE deleted_at IS NULL;

-- 2) 安全生产标准化体系：每企业 0~1 条
CREATE TABLE production_standardizations (
    id            BIGSERIAL PRIMARY KEY,
    grade         VARCHAR(255),
    get_time      DATE,
    valid_period  DATE,
    annex_url     TEXT,
    company_id    BIGINT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    user_id       BIGINT NOT NULL DEFAULT 0,
    last_user_id  BIGINT NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ
);
CREATE INDEX idx_production_standardizations_company ON production_standardizations (company_id) WHERE deleted_at IS NULL;

-- 3) 职业健康安全管理体系：每企业 0~1 条
CREATE TABLE occupational_health_safeties (
    id                  BIGSERIAL PRIMARY KEY,
    certification_body  VARCHAR(255),
    standard_edition    VARCHAR(255),
    get_time            DATE,
    valid_period        DATE,
    annex_url           TEXT,
    company_id          BIGINT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    user_id             BIGINT NOT NULL DEFAULT 0,
    last_user_id        BIGINT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_occupational_health_safeties_company ON occupational_health_safeties (company_id) WHERE deleted_at IS NULL;

-- 4) 经费执行情况：列表 per company，关联 security_investments，列出时
--    后端要把同年总预算下该执行项的占比一同算出来
CREATE TABLE funding_executions (
    id                       BIGSERIAL PRIMARY KEY,
    year                     VARCHAR(8),
    security_investment_id   BIGINT REFERENCES security_investments(id) ON DELETE SET NULL,
    simple_content           VARCHAR(255),
    content                  TEXT,
    budget                   NUMERIC(14, 2),
    company_id               BIGINT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    user_id                  BIGINT NOT NULL DEFAULT 0,
    last_user_id             BIGINT NOT NULL DEFAULT 0,
    created_at               TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at               TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at               TIMESTAMPTZ
);
CREATE INDEX idx_funding_executions_company_year ON funding_executions (company_id, year) WHERE deleted_at IS NULL;
CREATE INDEX idx_funding_executions_investment   ON funding_executions (security_investment_id);

-- 5) 安全管理人员：三大类岗位人员台账 (与 safety_management_types 字典关联)
CREATE TABLE safety_management_staff (
    id                          BIGSERIAL PRIMARY KEY,
    name                        VARCHAR(255),
    safety_management_type_id   BIGINT NOT NULL DEFAULT 1,
    tel                         VARCHAR(64),
    cert_title                  VARCHAR(255),
    cert_number                 VARCHAR(255),
    valid_period                DATE,
    company_id                  BIGINT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    user_id                     BIGINT NOT NULL DEFAULT 0,
    last_user_id                BIGINT NOT NULL DEFAULT 0,
    created_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at                  TIMESTAMPTZ
);
CREATE INDEX idx_safety_management_staff_company ON safety_management_staff (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_safety_management_staff_type    ON safety_management_staff (safety_management_type_id);

-- 6) 消防控制室人员：含 type (后续 Laravel 加的列), 简单台账
CREATE TABLE elimination_control_staff (
    id            BIGSERIAL PRIMARY KEY,
    name          VARCHAR(255),
    type          SMALLINT NOT NULL DEFAULT 1,
    cert_title    VARCHAR(255),
    cert_number   VARCHAR(255),
    valid_period  DATE,
    company_id    BIGINT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    user_id       BIGINT NOT NULL DEFAULT 0,
    last_user_id  BIGINT NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ
);
CREATE INDEX idx_elimination_control_staff_company ON elimination_control_staff (company_id) WHERE deleted_at IS NULL;
