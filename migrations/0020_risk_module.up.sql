-- 风险管理模块 — 6 张表，覆盖 5 个前端页面
--
-- risk_distribution_maps           风险源管理（per-company singleton：风险分布图上传）
-- monthly_investigations           隐患管理主表
-- monthly_investigation_reviews    隐患整改附件子表（整体替换语义）
-- monthly_investigation_views      隐患复查附件子表（整体替换语义）
-- job_management                   作业管理（外协单位 evaluation_units 关联）
-- three_togethers                  三同时管理（design_units / evaluation_units 关联）
-- assess_detects                   安全评价与检测（skill_units / skill_unit_kids 关联）
--
-- 风格沿用 0019_health_fire_module.up.sql：company_id NOT NULL + audit + 软删除。
-- 子表 monthly_investigation_id ON DELETE CASCADE：主表删除时附件自然清掉。
-- evaluation_unit_id / design_unit_id / skill_unit* 不挂外键，跟 Laravel 一致：
-- 这些关联表也在同一租户内，业务删除路径已经在 service 层校验，硬外键反而会
-- 在 evaluation_unit 软删除场景制造级联噪音。

CREATE TABLE risk_distribution_maps (
    id              BIGSERIAL PRIMARY KEY,
    annex_url       TEXT,
    company_id      BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id         BIGINT NOT NULL DEFAULT 0,
    last_user_id    BIGINT NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);
CREATE INDEX idx_risk_distribution_maps_company ON risk_distribution_maps (company_id) WHERE deleted_at IS NULL;

CREATE TABLE monthly_investigations (
    id                  BIGSERIAL PRIMARY KEY,
    "type"              VARCHAR(255),
    is_fire_fighting    SMALLINT,
    is_emergency        SMALLINT,
    "content"           TEXT,
    inv_time            DATE,
    inv_name            VARCHAR(255),
    "level"             VARCHAR(255),
    proposal            TEXT,
    legalbasis          TEXT,
    correct_date        DATE,
    review_discription  TEXT,
    review_date         DATE,
    review_content      TEXT,
    review_user         VARCHAR(255),
    correct_result      VARCHAR(255),
    is_govern           SMALLINT,
    govern_title        VARCHAR(255),
    govern_supplies     VARCHAR(255),
    govern_user         VARCHAR(255),
    govern_date         VARCHAR(255),
    govern_text         VARCHAR(255),
    govern_money        VARCHAR(255),
    company_id          BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id             BIGINT NOT NULL DEFAULT 0,
    last_user_id        BIGINT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_monthly_investigations_company ON monthly_investigations (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_monthly_investigations_type    ON monthly_investigations ("type") WHERE deleted_at IS NULL;

CREATE TABLE monthly_investigation_reviews (
    id                          BIGSERIAL PRIMARY KEY,
    monthly_investigation_id    BIGINT NOT NULL REFERENCES monthly_investigations(id) ON DELETE CASCADE,
    annex_url                   TEXT,
    company_id                  BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id                     BIGINT NOT NULL DEFAULT 0,
    last_user_id                BIGINT NOT NULL DEFAULT 0,
    created_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at                  TIMESTAMPTZ
);
CREATE INDEX idx_monthly_investigation_reviews_parent ON monthly_investigation_reviews (monthly_investigation_id) WHERE deleted_at IS NULL;

CREATE TABLE monthly_investigation_views (
    id                          BIGSERIAL PRIMARY KEY,
    monthly_investigation_id    BIGINT NOT NULL REFERENCES monthly_investigations(id) ON DELETE CASCADE,
    annex_url                   TEXT,
    company_id                  BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id                     BIGINT NOT NULL DEFAULT 0,
    last_user_id                BIGINT NOT NULL DEFAULT 0,
    created_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at                  TIMESTAMPTZ
);
CREATE INDEX idx_monthly_investigation_views_parent ON monthly_investigation_views (monthly_investigation_id) WHERE deleted_at IS NULL;

CREATE TABLE job_management (
    id                  BIGSERIAL PRIMARY KEY,
    department          VARCHAR(255),
    "content"           TEXT,
    "type"              VARCHAR(255),
    special_type        VARCHAR(255),
    "position"          VARCHAR(255),
    job_time            DATE,
    unit                VARCHAR(255),
    is_outsourcing      SMALLINT,
    evaluation_unit_id  BIGINT,
    job_people_number   INTEGER,
    job_people_name     VARCHAR(255),
    risk_analysis       TEXT,
    risk_protection     TEXT,
    company_id          BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id             BIGINT NOT NULL DEFAULT 0,
    last_user_id        BIGINT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_job_management_company    ON job_management (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_job_management_eval_unit  ON job_management (evaluation_unit_id);

CREATE TABLE three_togethers (
    id                      BIGSERIAL PRIMARY KEY,
    "name"                  VARCHAR(255),
    three_together_type     VARCHAR(255),
    establish_time          DATE,
    establish_section       VARCHAR(255),
    "status"                VARCHAR(255),
    design_unit_id          BIGINT,
    design_annex_url        TEXT,
    evaluation_unit_id      BIGINT,
    write_time              DATE,
    construction_comment    TEXT,
    safety_requirements     TEXT,
    complete_comment        TEXT,
    acceptance              TEXT,
    company_id              BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id                 BIGINT NOT NULL DEFAULT 0,
    last_user_id            BIGINT NOT NULL DEFAULT 0,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at              TIMESTAMPTZ
);
CREATE INDEX idx_three_togethers_company   ON three_togethers (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_three_togethers_design    ON three_togethers (design_unit_id);
CREATE INDEX idx_three_togethers_eval_unit ON three_togethers (evaluation_unit_id);

CREATE TABLE assess_detects (
    id                  BIGSERIAL PRIMARY KEY,
    "name"              VARCHAR(255) NOT NULL,
    assess_detect_type  VARCHAR(255),
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
CREATE INDEX idx_assess_detects_company   ON assess_detects (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_assess_detects_unit      ON assess_detects (skill_unit_id);
CREATE INDEX idx_assess_detects_unit_kid  ON assess_detects (skill_unit_kid_id);
