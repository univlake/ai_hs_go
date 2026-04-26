CREATE TABLE personnels (
    id                          BIGSERIAL PRIMARY KEY,
    total_people                INTEGER NOT NULL DEFAULT 0,
    staff                       INTEGER NOT NULL DEFAULT 0,
    service                     INTEGER NOT NULL DEFAULT 0,
    manager                     INTEGER NOT NULL DEFAULT 0,
    high_school_diploma         INTEGER NOT NULL DEFAULT 0,
    security_officer            INTEGER NOT NULL DEFAULT 0,
    safety_engineer             INTEGER NOT NULL DEFAULT 0,
    safety_engineering_major    INTEGER NOT NULL DEFAULT 0,
    fire_engineer               INTEGER NOT NULL DEFAULT 0,
    other                       INTEGER NOT NULL DEFAULT 0,
    company_id                  BIGINT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    user_id                     BIGINT NOT NULL DEFAULT 0,
    last_user_id                BIGINT NOT NULL DEFAULT 0,
    created_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at                  TIMESTAMPTZ
);
CREATE INDEX idx_personnels_company ON personnels (company_id);
