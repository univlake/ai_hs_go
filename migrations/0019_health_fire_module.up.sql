-- 职业健康 + 消防安全模块 — 9 张表
--
-- occupational_hazards    职业危害因素清单
-- health_systems          职业健康体系（singleton per company，前端暂未挂菜单）
-- health_files            健康档案
-- health_reports          职业病报告
-- health_checkups         职业健康体检
-- hygiene_evaluations     职业卫生评价（关联 skill_units / skill_unit_kids）
-- fire_safeties           消防安全（关联 skill_units / skill_unit_kids）
-- fire_equipment          消防器材
-- fire_facilities         消防设施
--
-- 风格沿用 0017_third_party_module.up.sql / 0018_meeting_training_module.up.sql：
-- company_id NOT NULL + audit + 软删除。fire_equipment / fire_facilities 跟
-- Laravel 一致：use_time 是 DATE，automation_type 是 VARCHAR。fire_safeties.valid_period
-- 沿用 Laravel string 类型（前端按 date 渲染，不强解析）。

CREATE TABLE occupational_hazards (
    id              BIGSERIAL PRIMARY KEY,
    "name"          VARCHAR(255) NOT NULL,
    "type"          VARCHAR(64),
    "number"        INTEGER,
    is_declare      VARCHAR(32),
    declare_date    DATE,
    company_id      BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id         BIGINT NOT NULL DEFAULT 0,
    last_user_id    BIGINT NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);
CREATE INDEX idx_occupational_hazards_company ON occupational_hazards (company_id) WHERE deleted_at IS NULL;

CREATE TABLE health_systems (
    id                  BIGSERIAL PRIMARY KEY,
    certification_time  DATE,
    annex_url           TEXT,
    company_id          BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id             BIGINT NOT NULL DEFAULT 0,
    last_user_id        BIGINT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_health_systems_company ON health_systems (company_id) WHERE deleted_at IS NULL;

CREATE TABLE health_files (
    id                  BIGSERIAL PRIMARY KEY,
    "name"              VARCHAR(255) NOT NULL,
    "number"            VARCHAR(64),
    department          VARCHAR(255),
    post                VARCHAR(255),
    checkup_date        DATE,
    checkup_expiry      TEXT,
    review_situation    VARCHAR(255),
    handling_opinions   VARCHAR(255),
    process_result      VARCHAR(255),
    annex_url           TEXT,
    company_id          BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id             BIGINT NOT NULL DEFAULT 0,
    last_user_id        BIGINT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_health_files_company ON health_files (company_id) WHERE deleted_at IS NULL;

CREATE TABLE health_reports (
    id              BIGSERIAL PRIMARY KEY,
    find_date       DATE,
    "name"          VARCHAR(255) NOT NULL,
    post            VARCHAR(255),
    "type"          VARCHAR(255),
    is_true         SMALLINT,
    mechanism       VARCHAR(255),
    identify_date   DATE,
    annex_url       TEXT,
    company_id      BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id         BIGINT NOT NULL DEFAULT 0,
    last_user_id    BIGINT NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);
CREATE INDEX idx_health_reports_company ON health_reports (company_id) WHERE deleted_at IS NULL;

CREATE TABLE health_checkups (
    id              BIGSERIAL PRIMARY KEY,
    "type"          VARCHAR(64) NOT NULL,
    department      VARCHAR(255),
    "number"        INTEGER,
    checkup_date    DATE,
    checkup_expiry  TEXT,
    measure         TEXT,
    annex_url       TEXT,
    company_id      BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id         BIGINT NOT NULL DEFAULT 0,
    last_user_id    BIGINT NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);
CREATE INDEX idx_health_checkups_company ON health_checkups (company_id) WHERE deleted_at IS NULL;

CREATE TABLE hygiene_evaluations (
    id                  BIGSERIAL PRIMARY KEY,
    "name"              VARCHAR(255) NOT NULL,
    "type"              VARCHAR(64),
    skill_unit_id       BIGINT,
    skill_unit_kid_id   BIGINT,
    valid_period        DATE,
    annex_url           TEXT,
    company_id          BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id             BIGINT NOT NULL DEFAULT 0,
    last_user_id        BIGINT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_hygiene_evaluations_company   ON hygiene_evaluations (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_hygiene_evaluations_unit      ON hygiene_evaluations (skill_unit_id);
CREATE INDEX idx_hygiene_evaluations_unit_kid  ON hygiene_evaluations (skill_unit_kid_id);

CREATE TABLE fire_safeties (
    id                  BIGSERIAL PRIMARY KEY,
    "name"              VARCHAR(255),
    "type"              VARCHAR(64),
    skill_unit_id       BIGINT,
    skill_unit_kid_id   BIGINT,
    valid_period        VARCHAR(64),
    annex_url           TEXT,
    company_id          BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id             BIGINT NOT NULL DEFAULT 0,
    last_user_id        BIGINT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_fire_safeties_company   ON fire_safeties (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_fire_safeties_unit      ON fire_safeties (skill_unit_id);
CREATE INDEX idx_fire_safeties_unit_kid  ON fire_safeties (skill_unit_kid_id);

CREATE TABLE fire_equipment (
    id                  BIGSERIAL PRIMARY KEY,
    equipment_name      VARCHAR(255),
    utilized_location   VARCHAR(255),
    special_type        VARCHAR(255),
    use_time            DATE,
    automation_type     VARCHAR(255),
    maintain_info       VARCHAR(255),
    "comment"           TEXT,
    company_id          BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id             BIGINT NOT NULL DEFAULT 0,
    last_user_id        BIGINT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_fire_equipment_company ON fire_equipment (company_id) WHERE deleted_at IS NULL;

CREATE TABLE fire_facilities (
    id                  BIGSERIAL PRIMARY KEY,
    equipment_name      VARCHAR(255),
    utilized_location   VARCHAR(255),
    special_type        VARCHAR(255),
    use_time            DATE,
    automation_type     VARCHAR(255),
    maintain_info       VARCHAR(255),
    "comment"           TEXT,
    company_id          BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id             BIGINT NOT NULL DEFAULT 0,
    last_user_id        BIGINT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_fire_facilities_company ON fire_facilities (company_id) WHERE deleted_at IS NULL;
