-- 重大危险源（Major Hazard Sources）模块 — 两张表
--
-- major_hazards         主表：危险源基本信息 + 安全联锁字段
-- major_constitutes     子表：危险源构成（指向 hazardous_chemicals）
--
-- 与 0006_safety_ledger.up.sql 风格保持一致：company_id NOT NULL + audit + 软删除。

CREATE TABLE major_hazards (
    id                BIGSERIAL PRIMARY KEY,
    "name"            VARCHAR(255),
    "level"           VARCHAR(32),
    "position"        VARCHAR(255),
    r_value           VARCHAR(64),
    is_sis            SMALLINT NOT NULL DEFAULT 0,
    sil_level         VARCHAR(32),
    detect_parameter  VARCHAR(255),
    company_id        BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id           BIGINT NOT NULL DEFAULT 0,
    last_user_id      BIGINT NOT NULL DEFAULT 0,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at        TIMESTAMPTZ
);
CREATE INDEX idx_major_hazards_company ON major_hazards (company_id) WHERE deleted_at IS NULL;

CREATE TABLE major_constitutes (
    id                    BIGSERIAL PRIMARY KEY,
    "name"                VARCHAR(255),
    hazardous_chemical_id BIGINT REFERENCES hazardous_chemicals(id) ON DELETE SET NULL,
    "number"              VARCHAR(64),
    major_hazard_id       BIGINT NOT NULL REFERENCES major_hazards(id) ON DELETE CASCADE,
    company_id            BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id               BIGINT NOT NULL DEFAULT 0,
    last_user_id          BIGINT NOT NULL DEFAULT 0,
    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at            TIMESTAMPTZ
);
CREATE INDEX idx_major_constitutes_hazard   ON major_constitutes (major_hazard_id);
CREATE INDEX idx_major_constitutes_chemical ON major_constitutes (hazardous_chemical_id);
CREATE INDEX idx_major_constitutes_company  ON major_constitutes (company_id);
