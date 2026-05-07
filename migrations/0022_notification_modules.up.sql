-- 通知公告 / 会议通知 / 信息稿管理 — 8 张表，覆盖三个模块的生产视图 + 集团视图
--
-- 通知公告（任务型 office_notifies + 通知型 new_notifies）— 主表 + 整体替换的接收企业子表
--   office_notifies / office_notify_scopes        任务型（带反馈截止 + 反馈附件）
--   new_notifies   / new_notify_scopes            通知型（仅需阅读确认）
--
-- 会议通知 — 会议主表 + 接收企业子表（含 notify/confirm 时间）+ 报名子表
--   meetings              会议主表
--   meeting_notices       会议接收企业（整体替换 kids）
--   meeting_sign_ups      接收企业的报名记录
--
-- 信息稿管理 — 起草、上报、发布的稿件，集团视图直接复用同一张表通过 next_reviewer/status 过滤
--   information_drafts    信息稿主表（next_reviewer + status 驱动上报/发布流程）
--
-- 风格沿用 0021_accident_emergency_module.up.sql：company_id NOT NULL + audit + 软删除。
-- 集团视图（receiveOnly 前端、bloc dashboard）跟生产视图共用同一批表，仅在 service 层
-- 通过 (a) scopes.company_id = 当前公司 (b) ancestor chain over information_drafts 来切片。
--
-- 0007 注册的 new_notifies / new_notify_scopes 是更早期的队列侧通知占位（Title/Body/JSONB
-- schema），从未接入业务路径，迁移在此一并 DROP 重建，避免双语义并存。

DROP TABLE IF EXISTS new_notify_scopes;
DROP TABLE IF EXISTS new_notifies;

-- ── 任务型通知（office_notifies）─────────────────────────────────────

CREATE TABLE office_notifies (
    id            BIGSERIAL PRIMARY KEY,
    "name"        VARCHAR(255),
    feedback_time DATE,
    alert_others  VARCHAR(255),
    "comment"     TEXT,
    annex_url     TEXT,
    excel_url     TEXT,
    company_id    BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id       BIGINT NOT NULL DEFAULT 0,
    last_user_id  BIGINT NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ
);
CREATE INDEX idx_office_notifies_company ON office_notifies (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_office_notifies_name    ON office_notifies ("name")    WHERE deleted_at IS NULL;

CREATE TABLE office_notify_scopes (
    id                  BIGSERIAL PRIMARY KEY,
    company_name        VARCHAR(255),
    feedback_time       DATE,
    feedback_status     VARCHAR(64) NOT NULL DEFAULT '未反馈',
    feedback_annex_url  TEXT,
    office_notify_id    BIGINT NOT NULL REFERENCES office_notifies(id) ON DELETE CASCADE,
    company_id          BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    company_group_name  VARCHAR(255),
    company_group_id    BIGINT REFERENCES company_groups(id) ON DELETE SET NULL,
    user_id             BIGINT NOT NULL DEFAULT 0,
    last_user_id        BIGINT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_office_notify_scopes_parent  ON office_notify_scopes (office_notify_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_office_notify_scopes_company ON office_notify_scopes (company_id)       WHERE deleted_at IS NULL;
CREATE INDEX idx_office_notify_scopes_pending ON office_notify_scopes (company_id, feedback_status) WHERE deleted_at IS NULL;

-- ── 通知型通知（new_notifies）───────────────────────────────────────

CREATE TABLE new_notifies (
    id            BIGSERIAL PRIMARY KEY,
    "name"        VARCHAR(255),
    alert_others  VARCHAR(255),
    "comment"     TEXT,
    annex_url     TEXT,
    company_id    BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id       BIGINT NOT NULL DEFAULT 0,
    last_user_id  BIGINT NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ
);
CREATE INDEX idx_new_notifies_company ON new_notifies (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_new_notifies_name    ON new_notifies ("name")    WHERE deleted_at IS NULL;

CREATE TABLE new_notify_scopes (
    id                  BIGSERIAL PRIMARY KEY,
    company_name        VARCHAR(255),
    read_time           DATE,
    read_status         VARCHAR(64) NOT NULL DEFAULT '未读',
    new_notify_id       BIGINT NOT NULL REFERENCES new_notifies(id) ON DELETE CASCADE,
    company_id          BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    company_group_name  VARCHAR(255),
    company_group_id    BIGINT REFERENCES company_groups(id) ON DELETE SET NULL,
    user_id             BIGINT NOT NULL DEFAULT 0,
    last_user_id        BIGINT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_new_notify_scopes_parent  ON new_notify_scopes (new_notify_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_new_notify_scopes_company ON new_notify_scopes (company_id)    WHERE deleted_at IS NULL;
CREATE INDEX idx_new_notify_scopes_pending ON new_notify_scopes (company_id, read_status) WHERE deleted_at IS NULL;
-- partial unique 让 (notify, user) 软删后能重新添加；普通 UNIQUE 会永久卡住
CREATE UNIQUE INDEX uniq_new_notify_scope
    ON new_notify_scopes (new_notify_id, user_id)
    WHERE deleted_at IS NULL;

-- ── 会议通知（meetings + meeting_notices + meeting_sign_ups）─────────

CREATE TABLE meetings (
    id                  BIGSERIAL PRIMARY KEY,
    "name"              VARCHAR(255),
    meeting_date        TIMESTAMPTZ,
    sign_up_date        TIMESTAMPTZ,
    sign_end_date       TIMESTAMPTZ,
    alert_others        VARCHAR(255),
    theme               VARCHAR(255),
    place               VARCHAR(255),
    summary             TEXT,
    annex_url           TEXT,
    summary_annex_url   TEXT,
    sign_annex_url      TEXT,
    material_annex_url  TEXT,
    photo_annex_url     TEXT,
    company_id          BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id             BIGINT NOT NULL DEFAULT 0,
    last_user_id        BIGINT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_meetings_company      ON meetings (company_id)   WHERE deleted_at IS NULL;
CREATE INDEX idx_meetings_meeting_date ON meetings (meeting_date) WHERE deleted_at IS NULL;
CREATE INDEX idx_meetings_name         ON meetings ("name")       WHERE deleted_at IS NULL;

CREATE TABLE meeting_notices (
    id                  BIGSERIAL PRIMARY KEY,
    company_name        VARCHAR(255),
    meeting_id          BIGINT NOT NULL REFERENCES meetings(id) ON DELETE CASCADE,
    company_id          BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    company_group_id    BIGINT REFERENCES company_groups(id) ON DELETE SET NULL,
    company_group_name  VARCHAR(255),
    notify_time         TIMESTAMPTZ,
    confirm_time        TIMESTAMPTZ,
    user_id             BIGINT NOT NULL DEFAULT 0,
    last_user_id        BIGINT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_meeting_notices_meeting ON meeting_notices (meeting_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_meeting_notices_company ON meeting_notices (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_meeting_notices_pending ON meeting_notices (company_id, confirm_time) WHERE deleted_at IS NULL;

CREATE TABLE meeting_sign_ups (
    id            BIGSERIAL PRIMARY KEY,
    company_name  VARCHAR(255),
    "name"        VARCHAR(255),
    sex           VARCHAR(16),
    job_title     VARCHAR(255),
    tel           VARCHAR(64),
    card_number   VARCHAR(64),
    is_driving    VARCHAR(32),
    meeting_id    BIGINT NOT NULL REFERENCES meetings(id) ON DELETE CASCADE,
    company_id    BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id       BIGINT NOT NULL DEFAULT 0,
    last_user_id  BIGINT NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ
);
CREATE INDEX idx_meeting_sign_ups_meeting ON meeting_sign_ups (meeting_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_meeting_sign_ups_company ON meeting_sign_ups (company_id) WHERE deleted_at IS NULL;

-- ── 信息稿管理（information_drafts）──────────────────────────────────

CREATE TABLE information_drafts (
    id                BIGSERIAL PRIMARY KEY,
    "name"            VARCHAR(255) NOT NULL,
    company_name      VARCHAR(255) NOT NULL,
    draft_department  VARCHAR(255),
    draft_others      VARCHAR(255),
    "comment"         TEXT,
    annex_url         VARCHAR(512),
    next_reviewer     BIGINT REFERENCES users(id) ON DELETE SET NULL,
    "status"          INTEGER NOT NULL DEFAULT 0,
    is_top            INTEGER NOT NULL DEFAULT 0,
    company_id        BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id           BIGINT NOT NULL DEFAULT 0,
    last_user_id      BIGINT NOT NULL DEFAULT 0,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at        TIMESTAMPTZ
);
CREATE INDEX idx_information_drafts_company       ON information_drafts (company_id)    WHERE deleted_at IS NULL;
CREATE INDEX idx_information_drafts_next_reviewer ON information_drafts (next_reviewer) WHERE deleted_at IS NULL;
CREATE INDEX idx_information_drafts_status        ON information_drafts ("status")      WHERE deleted_at IS NULL;
CREATE INDEX idx_information_drafts_name          ON information_drafts ("name")        WHERE deleted_at IS NULL;

-- 快照列：插入时拷贝，不跟随 companies / company_groups 后续改名。
COMMENT ON COLUMN office_notify_scopes.company_name       IS 'Snapshot at row creation; for display when notification was issued.';
COMMENT ON COLUMN office_notify_scopes.company_group_name IS 'Snapshot at row creation; for display when notification was issued.';
COMMENT ON COLUMN new_notify_scopes.company_name          IS 'Snapshot at row creation; for display when notification was issued.';
COMMENT ON COLUMN new_notify_scopes.company_group_name    IS 'Snapshot at row creation; for display when notification was issued.';
COMMENT ON COLUMN meeting_notices.company_name            IS 'Snapshot at row creation; for display in meeting recipient list.';
COMMENT ON COLUMN meeting_notices.company_group_name      IS 'Snapshot at row creation; for display in meeting recipient list.';
COMMENT ON COLUMN meeting_sign_ups.company_name           IS 'Snapshot at sign-up; for historical sign-up sheet display.';
COMMENT ON COLUMN information_drafts.company_name         IS 'Snapshot at row creation; does not track later companies.title renames.';
