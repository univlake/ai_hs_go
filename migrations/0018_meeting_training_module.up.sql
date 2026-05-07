-- 会议培训（Meetings & Training）模块 — 5 张表
--
-- internal_meetings       内部会议（主表）
-- internal_meeting_kids   参会人员子表（join_user_id → users.id）
-- internal_trainings      内部培训
-- external_trainings      外部培训
-- other_trainings         相关方培训
--
-- 内部会议子表整体替换；三类培训均无子表。
-- join_user_id ON DELETE SET NULL — 用户被清除时不影响会议历史。

CREATE TABLE internal_meetings (
    id                  BIGSERIAL PRIMARY KEY,
    "name"              VARCHAR(255),
    theme               VARCHAR(255),
    "type"              VARCHAR(64),
    meeting_date        TIMESTAMPTZ,
    summary             TEXT,
    annex_url           TEXT,
    notice_annex_url    TEXT,
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
CREATE INDEX idx_internal_meetings_company ON internal_meetings (company_id) WHERE deleted_at IS NULL;

CREATE TABLE internal_meeting_kids (
    id                   BIGSERIAL PRIMARY KEY,
    internal_meeting_id  BIGINT NOT NULL REFERENCES internal_meetings(id) ON DELETE CASCADE,
    join_user_id         BIGINT REFERENCES users(id) ON DELETE SET NULL,
    company_id           BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id              BIGINT NOT NULL DEFAULT 0,
    last_user_id         BIGINT NOT NULL DEFAULT 0,
    created_at           TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at           TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at           TIMESTAMPTZ
);
CREATE INDEX idx_internal_meeting_kids_meeting ON internal_meeting_kids (internal_meeting_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_internal_meeting_kids_user    ON internal_meeting_kids (join_user_id);

CREATE TABLE internal_trainings (
    id                  BIGSERIAL PRIMARY KEY,
    "name"              VARCHAR(255),
    "comment"           TEXT,
    "object"            VARCHAR(255),
    "number"            INTEGER,
    training_date       DATE,
    training_type       VARCHAR(255),
    happening           TEXT,
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
CREATE INDEX idx_internal_trainings_company ON internal_trainings (company_id) WHERE deleted_at IS NULL;

CREATE TABLE external_trainings (
    id                  BIGSERIAL PRIMARY KEY,
    "name"              VARCHAR(255),
    "comment"           TEXT,
    "object"            VARCHAR(255),
    "number"            INTEGER,
    training_date       DATE,
    training_unit       VARCHAR(255),
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
CREATE INDEX idx_external_trainings_company ON external_trainings (company_id) WHERE deleted_at IS NULL;

CREATE TABLE other_trainings (
    id                  BIGSERIAL PRIMARY KEY,
    "name"              VARCHAR(255),
    "comment"           TEXT,
    "object"            VARCHAR(255),
    "number"            INTEGER,
    training_date       DATE,
    task_content        TEXT,
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
CREATE INDEX idx_other_trainings_company ON other_trainings (company_id) WHERE deleted_at IS NULL;
