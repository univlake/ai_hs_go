-- 事故应急管理模块 — 6 张表，覆盖生产视图 4 个页面 + 集团视图 2 个页面
--
-- accident_bulletins        事故快报主表（生产视图 /accident/report；集团视图 /management/accidents/report 共用同一表）
-- accident_bulletin_kids    伤亡人员明细（整体替换语义；按 accident_bulletin_id 级联删除）
-- emergency_drills          应急演练
-- emergency_supplies        应急物资（无子表，跟 Laravel 保持一致）
-- emergency_teams           应急队伍主表
-- emergency_team_kids       应急队伍成员明细（整体替换语义；按 emergency_team_id 级联删除）
--
-- 风格沿用 0020_risk_module.up.sql：company_id NOT NULL + audit + 软删除。
-- 集团视图（/bloc-accident-bulletins、/annual-stats）走相同的 accident_bulletins 表，
-- 通过 company subtree 在 service 层 whereIn(company_id) 跨公司聚合。

CREATE TABLE accident_bulletins (
    id                    BIGSERIAL PRIMARY KEY,
    "type"                VARCHAR(255),
    direct_cause          VARCHAR(255),
    indirect_cause        VARCHAR(255),
    accident_level        VARCHAR(255),
    department            VARCHAR(255),
    happen_time           DATE,
    happen_area           VARCHAR(255),
    "content"             TEXT,
    degree_injury         VARCHAR(255),
    direct_loss_property  VARCHAR(64),
    gauge_loss_property   VARCHAR(64),
    staff_die             INTEGER,
    staff_hurt            INTEGER,
    not_staff_die         INTEGER,
    not_staff_hurt        INTEGER,
    accident_report_url   TEXT,
    accident_photo_url    TEXT,
    company_id            BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id               BIGINT NOT NULL DEFAULT 0,
    last_user_id          BIGINT NOT NULL DEFAULT 0,
    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at            TIMESTAMPTZ
);
CREATE INDEX idx_accident_bulletins_company       ON accident_bulletins (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_accident_bulletins_happen_time   ON accident_bulletins (happen_time) WHERE deleted_at IS NULL;
CREATE INDEX idx_accident_bulletins_type          ON accident_bulletins ("type") WHERE deleted_at IS NULL;
CREATE INDEX idx_accident_bulletins_level         ON accident_bulletins (accident_level) WHERE deleted_at IS NULL;

CREATE TABLE accident_bulletin_kids (
    id                      BIGSERIAL PRIMARY KEY,
    "name"                  VARCHAR(64),
    degree_injury           VARCHAR(64),
    department              VARCHAR(128),
    age                     INTEGER,
    sex                     VARCHAR(8),
    work_type               VARCHAR(64),
    work_age                INTEGER,
    education_level         VARCHAR(64),
    days_off                INTEGER,
    accident_bulletin_id    BIGINT NOT NULL REFERENCES accident_bulletins(id) ON DELETE CASCADE,
    company_id              BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id                 BIGINT NOT NULL DEFAULT 0,
    last_user_id            BIGINT NOT NULL DEFAULT 0,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at              TIMESTAMPTZ
);
CREATE INDEX idx_accident_bulletin_kids_parent ON accident_bulletin_kids (accident_bulletin_id) WHERE deleted_at IS NULL;

CREATE TABLE emergency_drills (
    id                       BIGSERIAL PRIMARY KEY,
    "name"                   VARCHAR(255),
    drill_time               DATE,
    drill_local              VARCHAR(255),
    organization_department  VARCHAR(255),
    join_department          VARCHAR(255),
    drill_number             VARCHAR(255),
    drill_type               VARCHAR(255),
    drill_form               VARCHAR(255),
    drill_kind               VARCHAR(255),
    "event"                  TEXT,
    annex_url                TEXT,
    company_id               BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id                  BIGINT NOT NULL DEFAULT 0,
    last_user_id             BIGINT NOT NULL DEFAULT 0,
    created_at               TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at               TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at               TIMESTAMPTZ
);
CREATE INDEX idx_emergency_drills_company ON emergency_drills (company_id) WHERE deleted_at IS NULL;

CREATE TABLE emergency_supplies (
    id            BIGSERIAL PRIMARY KEY,
    "name"        VARCHAR(255),
    "number"      VARCHAR(255),
    save_user     VARCHAR(255),
    tel           VARCHAR(255),
    "comment"     TEXT,
    company_id    BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id       BIGINT NOT NULL DEFAULT 0,
    last_user_id  BIGINT NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ
);
CREATE INDEX idx_emergency_supplies_company ON emergency_supplies (company_id) WHERE deleted_at IS NULL;

CREATE TABLE emergency_teams (
    id            BIGSERIAL PRIMARY KEY,
    "name"        VARCHAR(255),
    team_number   VARCHAR(255),
    save_user     VARCHAR(255),
    tel           VARCHAR(255),
    company_id    BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id       BIGINT NOT NULL DEFAULT 0,
    last_user_id  BIGINT NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ
);
CREATE INDEX idx_emergency_teams_company ON emergency_teams (company_id) WHERE deleted_at IS NULL;

CREATE TABLE emergency_team_kids (
    id                  BIGSERIAL PRIMARY KEY,
    "name"              VARCHAR(255),
    department          VARCHAR(255),
    "content"           VARCHAR(255),
    tel                 VARCHAR(255),
    emergency_team_id   BIGINT NOT NULL REFERENCES emergency_teams(id) ON DELETE CASCADE,
    company_id          BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id             BIGINT NOT NULL DEFAULT 0,
    last_user_id        BIGINT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_emergency_team_kids_parent ON emergency_team_kids (emergency_team_id) WHERE deleted_at IS NULL;
