DROP TABLE IF EXISTS information_drafts;
DROP TABLE IF EXISTS meeting_sign_ups;
DROP TABLE IF EXISTS meeting_notices;
DROP TABLE IF EXISTS meetings;
DROP TABLE IF EXISTS new_notify_scopes;
DROP TABLE IF EXISTS new_notifies;
DROP TABLE IF EXISTS office_notify_scopes;
DROP TABLE IF EXISTS office_notifies;

-- 重建 0007 的占位表，保持 down 后跟 0007 之后状态一致。
CREATE TABLE new_notifies (
    id          BIGSERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    body        TEXT,
    data        JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE new_notify_scopes (
    id              BIGSERIAL PRIMARY KEY,
    notify_id       BIGINT NOT NULL REFERENCES new_notifies(id) ON DELETE CASCADE,
    user_id         BIGINT NOT NULL,
    is_read         BOOLEAN NOT NULL DEFAULT FALSE,
    read_at         TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_new_notify_scopes_user ON new_notify_scopes (user_id, is_read);
CREATE INDEX idx_new_notify_scopes_notify ON new_notify_scopes (notify_id);
