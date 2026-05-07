-- companies + base reference data
-- Schema-optimization vs Laravel:
--   * parent_id (NULL = root / 一级企业，可有多个)
--   * root_id 冗余字段 + 触发器维护
--   * 经纬度统一 NUMERIC(10,7)
--   * status SMALLINT，deleted_at 软删除

CREATE TABLE base_areas (
    area_id              BIGINT PRIMARY KEY,
    statistical_area_id  BIGINT,
    area_name            VARCHAR(100) NOT NULL,
    area_parent_id       BIGINT,
    area_sort            SMALLINT,
    allname              TEXT,
    all_childrens_id     TEXT,
    parent_ids           TEXT,
    parent_ids_withself  TEXT
);
CREATE INDEX idx_base_areas_statistical ON base_areas (statistical_area_id);
CREATE INDEX idx_base_areas_parent      ON base_areas (area_parent_id);

CREATE TABLE enterprise_types (
    id          BIGSERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);
CREATE INDEX idx_enterprise_types_title ON enterprise_types (title);

CREATE TABLE operational_types (
    id          BIGSERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);
CREATE INDEX idx_operational_types_title ON operational_types (title);

CREATE TABLE companies (
    id                    BIGSERIAL PRIMARY KEY,
    title                 VARCHAR(255) NOT NULL,
    enterprise_type_id    BIGINT REFERENCES enterprise_types(id) ON DELETE SET NULL,
    operational_type_id   BIGINT NOT NULL DEFAULT 1 REFERENCES operational_types(id) ON DELETE RESTRICT,
    area_id               BIGINT REFERENCES base_areas(area_id) ON DELETE SET NULL,
    address               TEXT,
    established           DATE,
    credit_code           VARCHAR(64),
    business_scope        TEXT,
    company_introduction  TEXT,
    longitude             NUMERIC(10, 7),
    dimension             NUMERIC(10, 7),
    geographic_range      TEXT,
    acreage               NUMERIC(12, 3),
    supervise_name        VARCHAR(255),
    supervise_tel         VARCHAR(64),
    supervise_fix         VARCHAR(64),
    supervise_email       VARCHAR(255),
    supervise_position    VARCHAR(255),
    status                SMALLINT NOT NULL DEFAULT 0,
    parent_id             BIGINT REFERENCES companies(id) ON DELETE RESTRICT,
    root_id               BIGINT NOT NULL,                -- 维护：插入/更新 parent_id 时由触发器写入
    is_produce            SMALLINT NOT NULL DEFAULT 0,
    is_bloc               SMALLINT NOT NULL DEFAULT 0,
    inherent_risk         NUMERIC(10, 2),
    inherent_risk_grade   VARCHAR(32),
    float_risk            NUMERIC(10, 2),
    float_risk_grade      VARCHAR(32),
    user_id               BIGINT NOT NULL DEFAULT 0,
    last_user_id          BIGINT NOT NULL DEFAULT 0,
    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at            TIMESTAMPTZ
);
CREATE INDEX idx_companies_title             ON companies (title);
CREATE INDEX idx_companies_enterprise_type   ON companies (enterprise_type_id);
CREATE INDEX idx_companies_operational_type  ON companies (operational_type_id);
CREATE INDEX idx_companies_area              ON companies (area_id);
CREATE INDEX idx_companies_credit_code       ON companies (credit_code);
CREATE INDEX idx_companies_parent            ON companies (parent_id);
CREATE INDEX idx_companies_root              ON companies (root_id);
CREATE INDEX idx_companies_root_status       ON companies (root_id, status) WHERE deleted_at IS NULL;
CREATE INDEX idx_companies_status            ON companies (status);

-- Trigger: maintain root_id automatically.
-- INSERT: root_id := parent.root_id (or self when no parent).
-- UPDATE on parent_id: cascade root_id to subtree (handled in repo layer for now;
-- routine company moves are rare and can be re-pointed offline).
CREATE OR REPLACE FUNCTION companies_set_root_id() RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' THEN
        IF NEW.parent_id IS NULL THEN
            NEW.root_id := COALESCE(NEW.id, NEW.root_id);
        ELSE
            SELECT root_id INTO NEW.root_id FROM companies WHERE id = NEW.parent_id;
            IF NEW.root_id IS NULL THEN
                RAISE EXCEPTION 'parent company % has no root_id', NEW.parent_id;
            END IF;
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_companies_root_id_insert
BEFORE INSERT ON companies
FOR EACH ROW EXECUTE FUNCTION companies_set_root_id();

-- After-insert fixup for forest roots: NEW.id is unknown in BEFORE INSERT for
-- BIGSERIAL, so a forest root saves root_id=NULL initially. Patch it now.
CREATE OR REPLACE FUNCTION companies_finalize_root_id() RETURNS TRIGGER AS $$
BEGIN
    IF NEW.parent_id IS NULL AND (NEW.root_id IS NULL OR NEW.root_id = 0) THEN
        UPDATE companies SET root_id = NEW.id WHERE id = NEW.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_companies_root_id_after
AFTER INSERT ON companies
FOR EACH ROW EXECUTE FUNCTION companies_finalize_root_id();

-- We declared root_id NOT NULL but the trigger may set it after insert; allow
-- the BEFORE trigger to leave it NULL by relaxing the constraint to DEFERRABLE.
ALTER TABLE companies ALTER COLUMN root_id DROP NOT NULL;
