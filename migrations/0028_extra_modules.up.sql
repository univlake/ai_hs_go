-- 剩余模块迁移：变更管理 / 法律法规库 / 培训资料(文件+视频) / 风险画像
--
-- change_manages          变更管理（前端 /change-management）
-- regulation_libraries    法律法规库（前端 /regulation/libraries），依赖 regulation_library_types
-- training_files          培训资料-文件
-- training_file_logs      下载行为流水
-- training_videos         培训资料-视频
-- training_video_logs     观看行为流水
-- inherent_risk_logs      固有风险日志（21 维度评分快照）
-- float_risk_logs         浮动风险日志（17 维度评分快照）
--
-- risk_logs 已经在 0027_risk_monitoring_module 里建过，本迁移不再 CREATE。
-- /risks/portrait 用 inherent_risk_logs + float_risk_logs + risk_logs 组合返回；
-- /risks/log 是 risk_logs 的 per-company 列表。

CREATE TABLE change_manages (
    id                          BIGSERIAL PRIMARY KEY,
    title                       VARCHAR(255),
    department_name             VARCHAR(255),
    apply_date                  DATE,
    reason                      TEXT,
    apply_leader_opinion        TEXT,
    apply_leader_opinion_date   DATE,
    risk_analyze                TEXT,
    manage_leader_opinion       TEXT,
    manage_leader_opinion_date  DATE,
    approve                     TEXT,
    approve_date                DATE,
    acceptance_info             TEXT,
    acceptance_info_date        DATE,
    company_id                  BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id                     BIGINT NOT NULL DEFAULT 0,
    last_user_id                BIGINT NOT NULL DEFAULT 0,
    created_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at                  TIMESTAMPTZ
);
CREATE INDEX idx_change_manages_company ON change_manages (company_id) WHERE deleted_at IS NULL;

CREATE TABLE regulation_libraries (
    id                           BIGSERIAL PRIMARY KEY,
    title                        VARCHAR(255) NOT NULL,
    regulation_library_type_id   BIGINT,
    established_date             DATE,
    effective_date               DATE,
    promulgating_department      VARCHAR(255),
    standard_encoding            VARCHAR(255),
    identify_date                DATE,
    status                       VARCHAR(64),
    annex_url                    TEXT,
    company_id                   BIGINT NOT NULL DEFAULT 0,
    user_id                      BIGINT NOT NULL DEFAULT 0,
    last_user_id                 BIGINT NOT NULL DEFAULT 0,
    created_at                   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at                   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at                   TIMESTAMPTZ
);
CREATE INDEX idx_regulation_libraries_type    ON regulation_libraries (regulation_library_type_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_regulation_libraries_company ON regulation_libraries (company_id) WHERE deleted_at IS NULL;

CREATE TABLE training_files (
    id              BIGSERIAL PRIMARY KEY,
    title           VARCHAR(255),
    file_url        TEXT,
    times_download  INTEGER NOT NULL DEFAULT 0,
    company_id      BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id         BIGINT NOT NULL DEFAULT 0,
    last_user_id    BIGINT NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);
CREATE INDEX idx_training_files_company ON training_files (company_id) WHERE deleted_at IS NULL;

CREATE TABLE training_file_logs (
    id                BIGSERIAL PRIMARY KEY,
    training_file_id  BIGINT NOT NULL,
    company_id        BIGINT NOT NULL,
    user_id           BIGINT NOT NULL DEFAULT 0,
    last_user_id      BIGINT NOT NULL DEFAULT 0,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at        TIMESTAMPTZ
);
CREATE INDEX idx_training_file_logs_file    ON training_file_logs (training_file_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_training_file_logs_company ON training_file_logs (company_id);

CREATE TABLE training_videos (
    id             BIGSERIAL PRIMARY KEY,
    title          VARCHAR(255),
    introduce      TEXT,
    thumb_url      TEXT,
    video_url      TEXT,
    times_watched  INTEGER NOT NULL DEFAULT 0,
    duration       VARCHAR(64),
    lector         VARCHAR(255),
    company_id     BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id        BIGINT NOT NULL DEFAULT 0,
    last_user_id   BIGINT NOT NULL DEFAULT 0,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at     TIMESTAMPTZ
);
CREATE INDEX idx_training_videos_company ON training_videos (company_id) WHERE deleted_at IS NULL;

CREATE TABLE training_video_logs (
    id                 BIGSERIAL PRIMARY KEY,
    training_video_id  BIGINT NOT NULL,
    company_id         BIGINT NOT NULL,
    user_id            BIGINT NOT NULL DEFAULT 0,
    last_user_id       BIGINT NOT NULL DEFAULT 0,
    created_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at         TIMESTAMPTZ
);
CREATE INDEX idx_training_video_logs_video   ON training_video_logs (training_video_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_training_video_logs_company ON training_video_logs (company_id);

-- 固有风险评分（21 维度），每企业每次评估写一行。
CREATE TABLE inherent_risk_logs (
    id                                  BIGSERIAL PRIMARY KEY,
    safety_architecture_value           NUMERIC(10, 4) NOT NULL DEFAULT 0,
    security_officer_scale              NUMERIC(10, 4) NOT NULL DEFAULT 0,
    already_obtain_people_scale         NUMERIC(10, 4) NOT NULL DEFAULT 0,
    last_year_income_scale              NUMERIC(10, 4) NOT NULL DEFAULT 0,
    funding_execution_scale             NUMERIC(10, 4) NOT NULL DEFAULT 0,
    production_standardization_scale    NUMERIC(10, 4) NOT NULL DEFAULT 0,
    health_system_scale                 NUMERIC(10, 4) NOT NULL DEFAULT 0,
    assess_detect_scale                 NUMERIC(10, 4) NOT NULL DEFAULT 0,
    high_school_diploma_scale           NUMERIC(10, 4) NOT NULL DEFAULT 0,
    acreage_scale                       NUMERIC(10, 4) NOT NULL DEFAULT 0,
    staff_scale                         NUMERIC(10, 4) NOT NULL DEFAULT 0,
    safety_engineering_major_scale      NUMERIC(10, 4) NOT NULL DEFAULT 0,
    hazardous_processes_scale           NUMERIC(10, 4) NOT NULL DEFAULT 0,
    hazardous_process_scale             NUMERIC(10, 4) NOT NULL DEFAULT 0,
    hazardous_process_more_scale        NUMERIC(10, 4) NOT NULL DEFAULT 0,
    process_is_automatic_scale          NUMERIC(10, 4) NOT NULL DEFAULT 0,
    special_equipment_number_scale      NUMERIC(10, 4) NOT NULL DEFAULT 0,
    equipment_more_number_scale         NUMERIC(10, 4) NOT NULL DEFAULT 0,
    hazardous_chemical_scale            NUMERIC(10, 4) NOT NULL DEFAULT 0,
    hazardous_chemicals_number_scale    NUMERIC(10, 4) NOT NULL DEFAULT 0,
    major_hazard_scale                  NUMERIC(10, 4) NOT NULL DEFAULT 0,
    company_id                          BIGINT NOT NULL DEFAULT 0,
    user_id                             BIGINT NOT NULL DEFAULT 0,
    last_user_id                        BIGINT NOT NULL DEFAULT 0,
    created_at                          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at                          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at                          TIMESTAMPTZ
);
CREATE INDEX idx_inherent_risk_logs_company ON inherent_risk_logs (company_id) WHERE deleted_at IS NULL;

-- 浮动风险评分（17 维度），每企业每次评估写一行。
CREATE TABLE float_risk_logs (
    id                                BIGSERIAL PRIMARY KEY,
    internal_meeting_scale            NUMERIC(10, 4) NOT NULL DEFAULT 0,
    internal_training_scale           NUMERIC(10, 4) NOT NULL DEFAULT 0,
    accident_bulletin_scale           NUMERIC(10, 4) NOT NULL DEFAULT 0,
    accident_bulletin_severe_scale    NUMERIC(10, 4) NOT NULL DEFAULT 0,
    accident_report_url_scale         NUMERIC(10, 4) NOT NULL DEFAULT 0,
    emergency_team_scale              NUMERIC(10, 4) NOT NULL DEFAULT 0,
    emergency_drill_scale             NUMERIC(10, 4) NOT NULL DEFAULT 0,
    job_management_scale              NUMERIC(10, 4) NOT NULL DEFAULT 0,
    risk_source_scale                 NUMERIC(10, 4) NOT NULL DEFAULT 0,
    monthly_investigation_scale       NUMERIC(10, 4) NOT NULL DEFAULT 0,
    three_together_over_scale         NUMERIC(10, 4) NOT NULL DEFAULT 0,
    unit_sum_number_scale             NUMERIC(10, 4) NOT NULL DEFAULT 0,
    evaluation_unit_scale             NUMERIC(10, 4) NOT NULL DEFAULT 0,
    evaluation_unit_school_scale      NUMERIC(10, 4) NOT NULL DEFAULT 0,
    health_checkup_people_scale       NUMERIC(10, 4) NOT NULL DEFAULT 0,
    health_checkup_scale              NUMERIC(10, 4) NOT NULL DEFAULT 0,
    health_report_scale               NUMERIC(10, 4) NOT NULL DEFAULT 0,
    company_id                        BIGINT NOT NULL DEFAULT 0,
    user_id                           BIGINT NOT NULL DEFAULT 0,
    last_user_id                      BIGINT NOT NULL DEFAULT 0,
    created_at                        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at                        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at                        TIMESTAMPTZ
);
CREATE INDEX idx_float_risk_logs_company ON float_risk_logs (company_id) WHERE deleted_at IS NULL;
