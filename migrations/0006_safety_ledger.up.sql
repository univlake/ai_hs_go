-- Safety ledger MVP tables: special_equipment, special_equipment_staff,
-- dangerous_equipment, hazardous_chemicals, chemical_ingredients.
--
-- Schema-optimization vs Laravel:
--   * company_id NOT NULL with foreign key to companies (was loose int)
--   * boolean flags via SMALLINT NOT NULL DEFAULT 0 (Laravel parity)
--   * comments use POSTGRES native COMMENT ON COLUMN
--   * deleted_at + indexes scoped to undeleted rows where useful

CREATE TABLE special_equipment (
    id                BIGSERIAL PRIMARY KEY,
    equipment_name    VARCHAR(255),
    special_type_id   BIGINT,
    utilized_location VARCHAR(255),
    norm_model1       VARCHAR(255),
    norm_model2       VARCHAR(255),
    norm_model3       VARCHAR(255),
    norm_model4       VARCHAR(255),
    serial_number     VARCHAR(255),
    use_time          DATE,
    valid_period      DATE,
    inspect_unit      VARCHAR(255),
    maintain_info     VARCHAR(255),
    "comment"         TEXT,
    company_id        BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id           BIGINT NOT NULL DEFAULT 0,
    last_user_id      BIGINT NOT NULL DEFAULT 0,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at        TIMESTAMPTZ
);
CREATE INDEX idx_special_equipment_company ON special_equipment (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_special_equipment_type    ON special_equipment (special_type_id);

CREATE TABLE special_equipment_staff (
    id            BIGSERIAL PRIMARY KEY,
    "name"        VARCHAR(255),
    cert_title    VARCHAR(255),
    cert_number   VARCHAR(255),
    valid_period  DATE,
    company_id    BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id       BIGINT NOT NULL DEFAULT 0,
    last_user_id  BIGINT NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ
);
CREATE INDEX idx_special_equipment_staff_company ON special_equipment_staff (company_id) WHERE deleted_at IS NULL;

CREATE TABLE dangerous_equipment (
    id                BIGSERIAL PRIMARY KEY,
    equipment_name    VARCHAR(255),
    utilized_location VARCHAR(255),
    norm_model        VARCHAR(255),
    serial_number     VARCHAR(255),
    use_time          DATE,
    automation_type   VARCHAR(255),
    maintain_info     VARCHAR(255),
    "comment"         TEXT,
    company_id        BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id           BIGINT NOT NULL DEFAULT 0,
    last_user_id      BIGINT NOT NULL DEFAULT 0,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at        TIMESTAMPTZ
);
CREATE INDEX idx_dangerous_equipment_company ON dangerous_equipment (company_id) WHERE deleted_at IS NULL;

CREATE TABLE hazardous_chemicals (
    id                BIGSERIAL PRIMARY KEY,
    "name"            VARCHAR(255),
    alias             VARCHAR(255),
    is_supervision    SMALLINT NOT NULL DEFAULT 0,
    is_toxin          SMALLINT NOT NULL DEFAULT 0,
    is_explosion      SMALLINT NOT NULL DEFAULT 0,
    is_easy_explode   SMALLINT NOT NULL DEFAULT 0,
    is_narcotics      SMALLINT NOT NULL DEFAULT 0,
    is_monitor        SMALLINT NOT NULL DEFAULT 0,
    is_mixing         SMALLINT NOT NULL DEFAULT 0,
    cas               VARCHAR(64),
    "position"        VARCHAR(255),
    temperature       VARCHAR(64),
    storage           VARCHAR(64),
    year_use          VARCHAR(64),
    year_use_unit     VARCHAR(32),
    storage_unit      VARCHAR(32),
    pressure          VARCHAR(64),
    pressure_unit     VARCHAR(32),
    fire_class        VARCHAR(32),
    matter_state      VARCHAR(32),
    msds              VARCHAR(512),
    company_id        BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id           BIGINT NOT NULL DEFAULT 0,
    last_user_id      BIGINT NOT NULL DEFAULT 0,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at        TIMESTAMPTZ
);
CREATE INDEX idx_hazardous_chemicals_company ON hazardous_chemicals (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_hazardous_chemicals_cas     ON hazardous_chemicals (cas);

CREATE TABLE chemical_ingredients (
    id                    BIGSERIAL PRIMARY KEY,
    "name"                VARCHAR(255),
    content               VARCHAR(64),
    cas                   VARCHAR(64),
    hazardous_chemical_id BIGINT NOT NULL REFERENCES hazardous_chemicals(id) ON DELETE CASCADE,
    company_id            BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id               BIGINT NOT NULL DEFAULT 0,
    last_user_id          BIGINT NOT NULL DEFAULT 0,
    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at            TIMESTAMPTZ
);
CREATE INDEX idx_chemical_ingredients_chemical ON chemical_ingredients (hazardous_chemical_id);
CREATE INDEX idx_chemical_ingredients_company  ON chemical_ingredients (company_id);
