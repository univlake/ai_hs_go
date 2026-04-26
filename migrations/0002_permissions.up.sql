-- Spatie permission schema (slimmed to api guard) + Casbin rule table.

CREATE TABLE permissions (
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    guard_name  VARCHAR(64) NOT NULL DEFAULT 'api',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX uq_permissions_name_guard ON permissions (name, guard_name);

CREATE TABLE roles (
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    guard_name  VARCHAR(64) NOT NULL DEFAULT 'api',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX uq_roles_name_guard ON roles (name, guard_name);

CREATE TABLE role_has_permissions (
    permission_id BIGINT NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
    role_id       BIGINT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    PRIMARY KEY (permission_id, role_id)
);

CREATE TABLE model_has_roles (
    role_id     BIGINT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    model_type  VARCHAR(255) NOT NULL,
    model_id    BIGINT NOT NULL,
    PRIMARY KEY (role_id, model_id, model_type)
);
CREATE INDEX idx_model_has_roles_model ON model_has_roles (model_id, model_type);

CREATE TABLE model_has_permissions (
    permission_id BIGINT NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
    model_type    VARCHAR(255) NOT NULL,
    model_id      BIGINT NOT NULL,
    PRIMARY KEY (permission_id, model_id, model_type)
);
CREATE INDEX idx_model_has_permissions_model ON model_has_permissions (model_id, model_type);

-- Casbin policy persistence (gorm-adapter v3 schema).
CREATE TABLE casbin_rule (
    id     BIGSERIAL PRIMARY KEY,
    ptype  VARCHAR(100) NOT NULL DEFAULT '',
    v0     VARCHAR(100) NOT NULL DEFAULT '',
    v1     VARCHAR(100) NOT NULL DEFAULT '',
    v2     VARCHAR(100) NOT NULL DEFAULT '',
    v3     VARCHAR(100) NOT NULL DEFAULT '',
    v4     VARCHAR(100) NOT NULL DEFAULT '',
    v5     VARCHAR(100) NOT NULL DEFAULT ''
);
CREATE UNIQUE INDEX uq_casbin_rule ON casbin_rule (ptype, v0, v1, v2, v3, v4, v5);
