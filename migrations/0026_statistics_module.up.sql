-- 数据统计模块（集团视图 + 生产视图 月/年度数据上报） — 跟 Laravel
-- 2026_04_14_000007..000011 5 张迁移对齐：
--
--   month_statistics            月度数据上报主表（每企业每月一行；数值统计 30 列）
--   month_statistic_controls    重大隐患治理计划子表（整体替换语义）
--   month_statistic_accidents   事故快报子表（整体替换语义）
--   year_statistics             年度数据上报主表（每企业每年一行）
--   year_statistic_kids         三类人员 / 消防控制室人员持证情况子表（整体替换）
--
-- 集团视图 /statistics/month-statistics + /management/statistics/monthly-stats
-- 跟生产视图 /tasks/monthly-report 共用同一组路由 / 权限（month-statistics.*）；
-- 视角差异仅在前端表格的 pid vs company_id 入参。

CREATE TABLE month_statistics (
    id                                   BIGSERIAL PRIMARY KEY,
    average_employee                     INTEGER NOT NULL DEFAULT 0,
    safety_meeting_count                 INTEGER NOT NULL DEFAULT 0,
    fire_meeting_count                   INTEGER NOT NULL DEFAULT 0,
    safety_check_count                   INTEGER NOT NULL DEFAULT 0,
    fire_check_count                     INTEGER NOT NULL DEFAULT 0,
    safety_funding                       INTEGER NOT NULL DEFAULT 0,
    fire_funding                         INTEGER NOT NULL DEFAULT 0,
    hydrating_seedlings                  INTEGER NOT NULL DEFAULT 0,
    retrofit_electrical                  INTEGER NOT NULL DEFAULT 0,
    fire_extinguisher                    INTEGER NOT NULL DEFAULT 0,
    smart_fire                           INTEGER NOT NULL DEFAULT 0,
    safety_training_count                INTEGER NOT NULL DEFAULT 0,
    safety_training_number               INTEGER NOT NULL DEFAULT 0,
    new_people_training_number           INTEGER NOT NULL DEFAULT 0,
    fire_training_count                  INTEGER NOT NULL DEFAULT 0,
    fire_training_number                 INTEGER NOT NULL DEFAULT 0,
    safety_emergency_count               INTEGER NOT NULL DEFAULT 0,
    safety_emergency_number              INTEGER NOT NULL DEFAULT 0,
    fire_emergency_count                 INTEGER NOT NULL DEFAULT 0,
    fire_emergency_number                INTEGER NOT NULL DEFAULT 0,
    some_hidden_danger                   INTEGER NOT NULL DEFAULT 0,
    some_rectify_hidden_danger           INTEGER NOT NULL DEFAULT 0,
    no_rectify_info                      TEXT,
    big_hidden_danger                    INTEGER NOT NULL DEFAULT 0,
    big_rectify_hidden_danger            INTEGER NOT NULL DEFAULT 0,
    govern_hidden_danger_count           INTEGER NOT NULL DEFAULT 0,
    fire_fighting_danger                 INTEGER NOT NULL DEFAULT 0,
    fire_fighting_hidden_danger          INTEGER NOT NULL DEFAULT 0,
    fire_fighting_hidden_danger_rectify  INTEGER NOT NULL DEFAULT 0,
    accident_bulletin_count              INTEGER NOT NULL DEFAULT 0,
    company_id                           BIGINT      NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id                              BIGINT      NOT NULL DEFAULT 0,
    last_user_id                         BIGINT      NOT NULL DEFAULT 0,
    created_at                           TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at                           TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at                           TIMESTAMPTZ
);
CREATE INDEX idx_month_statistics_company ON month_statistics (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_month_statistics_month   ON month_statistics (company_id, created_at) WHERE deleted_at IS NULL;

CREATE TABLE month_statistic_controls (
    id                  BIGSERIAL PRIMARY KEY,
    govern_title        VARCHAR(255),
    govern_supplies     VARCHAR(255),
    govern_user         VARCHAR(255),
    govern_date         VARCHAR(255),
    govern_text         VARCHAR(255),
    govern_money        VARCHAR(255),
    month_statistic_id  BIGINT      NOT NULL REFERENCES month_statistics(id) ON DELETE CASCADE,
    company_id          BIGINT      NOT NULL REFERENCES companies(id)        ON DELETE RESTRICT,
    user_id             BIGINT      NOT NULL DEFAULT 0,
    last_user_id        BIGINT      NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_month_statistic_controls_parent  ON month_statistic_controls (month_statistic_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_month_statistic_controls_company ON month_statistic_controls (company_id)         WHERE deleted_at IS NULL;

CREATE TABLE month_statistic_accidents (
    id                    BIGSERIAL PRIMARY KEY,
    happen_time           DATE,
    department            VARCHAR(255),
    name                  VARCHAR(255),
    sex                   VARCHAR(255),
    age                   INTEGER,
    work_type             VARCHAR(255),
    work_age              INTEGER,
    education_level       VARCHAR(255),
    degree_injury         VARCHAR(255),
    "type"                VARCHAR(255),
    direct_cause          VARCHAR(255),
    content               TEXT,
    days_off              INTEGER,
    direct_loss_property  VARCHAR(255),
    month_statistic_id    BIGINT      NOT NULL REFERENCES month_statistics(id) ON DELETE CASCADE,
    company_id            BIGINT      NOT NULL REFERENCES companies(id)        ON DELETE RESTRICT,
    user_id               BIGINT      NOT NULL DEFAULT 0,
    last_user_id          BIGINT      NOT NULL DEFAULT 0,
    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at            TIMESTAMPTZ
);
CREATE INDEX idx_month_statistic_accidents_parent  ON month_statistic_accidents (month_statistic_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_month_statistic_accidents_company ON month_statistic_accidents (company_id)         WHERE deleted_at IS NULL;

CREATE TABLE year_statistics (
    id                BIGSERIAL PRIMARY KEY,
    company_name      VARCHAR(255),
    base_area         VARCHAR(255),
    enterprise_type   VARCHAR(255),
    safety_funding    VARCHAR(255),
    fire_fighting     VARCHAR(255),
    compliance_time   TIMESTAMPTZ,
    revaluation_time  TIMESTAMPTZ,
    company_id        BIGINT      NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id           BIGINT      NOT NULL DEFAULT 0,
    last_user_id      BIGINT      NOT NULL DEFAULT 0,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at        TIMESTAMPTZ
);
CREATE INDEX idx_year_statistics_company ON year_statistics (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_year_statistics_year    ON year_statistics (company_id, created_at) WHERE deleted_at IS NULL;

CREATE TABLE year_statistic_kids (
    id                     BIGSERIAL PRIMARY KEY,
    name                   VARCHAR(255),
    "type"                 INTEGER     NOT NULL DEFAULT 1,
    should_obtain_people   INTEGER,
    already_obtain_people  INTEGER,
    no_obtain_people       VARCHAR(255),
    train_people           INTEGER,
    year_statistic_id      BIGINT      NOT NULL REFERENCES year_statistics(id) ON DELETE CASCADE,
    company_id             BIGINT      NOT NULL REFERENCES companies(id)       ON DELETE RESTRICT,
    user_id                BIGINT      NOT NULL DEFAULT 0,
    last_user_id           BIGINT      NOT NULL DEFAULT 0,
    created_at             TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at             TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at             TIMESTAMPTZ
);
CREATE INDEX idx_year_statistic_kids_parent  ON year_statistic_kids (year_statistic_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_year_statistic_kids_company ON year_statistic_kids (company_id)        WHERE deleted_at IS NULL;

-- 快照列：插入时拷贝，不跟随 companies 后续改名/换地区/换企业类型。
COMMENT ON COLUMN year_statistics.company_name    IS 'Snapshot at row creation; does not track later companies.title renames.';
COMMENT ON COLUMN year_statistics.base_area       IS 'Snapshot at row creation; does not track later companies.area_id changes.';
COMMENT ON COLUMN year_statistics.enterprise_type IS 'Snapshot at row creation; does not track later companies.enterprise_type_id changes.';
