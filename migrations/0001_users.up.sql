-- users (Phase 2)
-- Differences from Laravel users table:
--   * company_id is nullable (NULL = platform super admin)
--   * is_platform_admin flag (cross-tenant)
--   * must_change_credentials flag (force first-login setup for c{id}_admin)
--   * legacy fields preserved for downstream API compatibility

CREATE TABLE users (
    id                       BIGSERIAL PRIMARY KEY,
    username                 VARCHAR(30) NOT NULL,
    password                 VARCHAR(255) NOT NULL,
    company_id               BIGINT NULL,
    nickname                 VARCHAR(100) NOT NULL DEFAULT '',
    header_img               TEXT,
    phone                    VARCHAR(64),
    department_name          VARCHAR(255),
    email                    VARCHAR(255),
    is_expert                SMALLINT NOT NULL DEFAULT 0,
    role_id                  BIGINT NOT NULL DEFAULT 0,
    status                   SMALLINT NOT NULL DEFAULT 1,
    is_platform_admin        BOOLEAN NOT NULL DEFAULT FALSE,
    must_change_credentials  BOOLEAN NOT NULL DEFAULT FALSE,
    weapp_openid             VARCHAR(128),
    weixin_session_key       VARCHAR(255),
    user_id                  BIGINT NOT NULL DEFAULT 0,
    last_user_id             BIGINT NOT NULL DEFAULT 0,
    login_times              VARCHAR(64),
    created_at               TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at               TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at               TIMESTAMPTZ
);

CREATE UNIQUE INDEX uq_users_username ON users (username) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX uq_users_phone    ON users (phone)    WHERE phone IS NOT NULL AND deleted_at IS NULL;
CREATE UNIQUE INDEX uq_users_email    ON users (email)    WHERE email IS NOT NULL AND deleted_at IS NULL;
CREATE UNIQUE INDEX uq_users_weapp    ON users (weapp_openid) WHERE weapp_openid IS NOT NULL AND deleted_at IS NULL;
CREATE INDEX idx_users_company_id     ON users (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_users_status         ON users (status);
CREATE INDEX idx_users_is_platform    ON users (is_platform_admin) WHERE is_platform_admin = TRUE;
