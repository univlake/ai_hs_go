-- 第三方管理（Third-Party Management）模块 — 4 张表
--
-- design_units      设计单位
-- evaluation_units  施工/评价单位（前端路径 /third-party/construction-units）
-- skill_units       技术服务单位（主表）
-- skill_unit_kids   技术服务单位 — 资质明细子表（前端 kids 数组整体替换）
--
-- 风格沿用 0015_process_modules.up.sql：company_id NOT NULL + audit + 软删除。
-- 不挂字典外键（Laravel 原表也没建外键），少一组 ON DELETE 决策。

CREATE TABLE design_units (
    id                  BIGSERIAL PRIMARY KEY,
    "name"              VARCHAR(255) NOT NULL,
    qualification_name  VARCHAR(255),
    contact             VARCHAR(255),
    contact_tel         VARCHAR(255),
    annex_url           TEXT,
    company_id          BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id             BIGINT NOT NULL DEFAULT 0,
    last_user_id        BIGINT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_design_units_company ON design_units (company_id) WHERE deleted_at IS NULL;

CREATE TABLE evaluation_units (
    id                       BIGSERIAL PRIMARY KEY,
    "name"                   VARCHAR(255) NOT NULL,
    is_qualification         SMALLINT NOT NULL DEFAULT 2,
    qualification_name       VARCHAR(255),
    qualification_annex_url  TEXT,
    evaluation_area          TEXT,
    contact                  VARCHAR(255),
    contact_tel              VARCHAR(255),
    is_agreement             SMALLINT NOT NULL DEFAULT 2,
    agreement_annex_url      VARCHAR(255),
    supervision_url          TEXT,
    annex_url                TEXT,
    company_id               BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id                  BIGINT NOT NULL DEFAULT 0,
    last_user_id             BIGINT NOT NULL DEFAULT 0,
    created_at               TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at               TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at               TIMESTAMPTZ
);
CREATE INDEX idx_evaluation_units_company ON evaluation_units (company_id) WHERE deleted_at IS NULL;

CREATE TABLE skill_units (
    id                BIGSERIAL PRIMARY KEY,
    "name"            VARCHAR(255) NOT NULL,
    serving_type      VARCHAR(255) NOT NULL,
    is_qualification  SMALLINT NOT NULL DEFAULT 2,
    evaluation_area   TEXT,
    contact           VARCHAR(255),
    contact_tel       VARCHAR(255),
    company_id        BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id           BIGINT NOT NULL DEFAULT 0,
    last_user_id      BIGINT NOT NULL DEFAULT 0,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at        TIMESTAMPTZ
);
CREATE INDEX idx_skill_units_company ON skill_units (company_id) WHERE deleted_at IS NULL;

CREATE TABLE skill_unit_kids (
    id                       BIGSERIAL PRIMARY KEY,
    qualification_name       VARCHAR(255),
    qualification_annex_url  TEXT,
    skill_unit_id            BIGINT NOT NULL,
    company_id               BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id                  BIGINT NOT NULL DEFAULT 0,
    last_user_id             BIGINT NOT NULL DEFAULT 0,
    created_at               TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at               TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at               TIMESTAMPTZ
);
CREATE INDEX idx_skill_unit_kids_unit    ON skill_unit_kids (skill_unit_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_skill_unit_kids_company ON skill_unit_kids (company_id);
