-- 给 SMALLINT/INTEGER 枚举字段加 CHECK 约束。只覆盖代码里有显式常量定义、
-- 或迁移注释里写明取值的列；含义不明的保持开放，避免误伤后续未知用法。
--
-- 命名约定 ck_<table>_<column>。所有约束 NOT VALID 兜底兼容（生产可能有
-- 早期遗留行）；本地预检通过。运维确认后可执行 VALIDATE CONSTRAINT。

-- 1. 二值布尔 (SMALLINT 0/1) ─────────────────────────────────────
ALTER TABLE users
    ADD CONSTRAINT ck_users_is_expert        CHECK (is_expert IN (0, 1)) NOT VALID;

ALTER TABLE companies
    ADD CONSTRAINT ck_companies_is_produce   CHECK (is_produce IN (0, 1)) NOT VALID,
    ADD CONSTRAINT ck_companies_is_bloc      CHECK (is_bloc IN (0, 1)) NOT VALID;

ALTER TABLE hazardous_chemicals
    ADD CONSTRAINT ck_hazchem_supervision    CHECK (is_supervision IN (0, 1)) NOT VALID,
    ADD CONSTRAINT ck_hazchem_toxin          CHECK (is_toxin IN (0, 1)) NOT VALID,
    ADD CONSTRAINT ck_hazchem_explosion      CHECK (is_explosion IN (0, 1)) NOT VALID,
    ADD CONSTRAINT ck_hazchem_easy_explode   CHECK (is_easy_explode IN (0, 1)) NOT VALID,
    ADD CONSTRAINT ck_hazchem_narcotics      CHECK (is_narcotics IN (0, 1)) NOT VALID,
    ADD CONSTRAINT ck_hazchem_monitor        CHECK (is_monitor IN (0, 1)) NOT VALID,
    ADD CONSTRAINT ck_hazchem_mixing         CHECK (is_mixing IN (0, 1)) NOT VALID;

ALTER TABLE hazardous_processes
    ADD CONSTRAINT ck_hazproc_automatic      CHECK (is_automatic IN (0, 1)) NOT VALID,
    ADD CONSTRAINT ck_hazproc_control        CHECK (is_control IN (0, 1)) NOT VALID,
    ADD CONSTRAINT ck_hazproc_dust           CHECK (is_dust IN (0, 1)) NOT VALID,
    ADD CONSTRAINT ck_hazproc_online_test    CHECK (is_online_test IN (0, 1)) NOT VALID;

ALTER TABLE major_hazards
    ADD CONSTRAINT ck_major_hazards_is_sis   CHECK (is_sis IN (0, 1)) NOT VALID;

ALTER TABLE major_hazardous_processes
    ADD CONSTRAINT ck_majorhazproc_is_sis    CHECK (is_sis IN (0, 1)) NOT VALID;

ALTER TABLE monthly_investigations
    ADD CONSTRAINT ck_minv_is_fire_fighting  CHECK (is_fire_fighting IN (0, 1)) NOT VALID,
    ADD CONSTRAINT ck_minv_is_emergency      CHECK (is_emergency IN (0, 1)) NOT VALID,
    ADD CONSTRAINT ck_minv_is_govern         CHECK (is_govern IN (0, 1)) NOT VALID;

ALTER TABLE job_management
    ADD CONSTRAINT ck_job_is_outsourcing     CHECK (is_outsourcing IN (0, 1)) NOT VALID;

ALTER TABLE health_reports
    ADD CONSTRAINT ck_health_reports_is_true CHECK (is_true IN (0, 1)) NOT VALID;

ALTER TABLE hidden_upload_annexes
    ADD CONSTRAINT ck_hda_is_review          CHECK (is_review IN (0, 1)) NOT VALID;

-- 2. 三值资质标记 (1=有, 2=未填) ─────────────────────────────────
-- 代码里 defaultIfZero(IsQualification, 2) 把 0 也归到 2，所以合法值是 1 / 2。
ALTER TABLE evaluation_units
    ADD CONSTRAINT ck_eval_units_is_qualification CHECK (is_qualification IN (1, 2)) NOT VALID,
    ADD CONSTRAINT ck_eval_units_is_agreement     CHECK (is_agreement     IN (1, 2)) NOT VALID;

ALTER TABLE skill_units
    ADD CONSTRAINT ck_skill_units_is_qualification CHECK (is_qualification IN (1, 2)) NOT VALID;

-- 3. 业务状态枚举 ───────────────────────────────────────────────
-- users:    UserStatusEnable=1, UserStatusDisable=2 (auth_service.go)
ALTER TABLE users
    ADD CONSTRAINT ck_users_status           CHECK (status IN (1, 2)) NOT VALID;

-- companies: Pending=0, Enable=1, Disabled=2 (company_service.go)
ALTER TABLE companies
    ADD CONSTRAINT ck_companies_status       CHECK (status IN (0, 1, 2)) NOT VALID;

-- daily_checks: Draft=0, Completed=1, Scored=2 (assessment_service.go)
ALTER TABLE daily_checks
    ADD CONSTRAINT ck_daily_checks_status    CHECK (status IN (0, 1, 2)) NOT VALID;

-- hidden_danger_infos: invest_status 1=进行中, 2=归档；correct_status 1=未整改, 2=已整改, 3=已确认
ALTER TABLE hidden_danger_infos
    ADD CONSTRAINT ck_hidden_invest_status   CHECK (invest_status   IN (1, 2)) NOT VALID,
    ADD CONSTRAINT ck_hidden_correct_status  CHECK (correct_status  IN (1, 2, 3)) NOT VALID;

-- risk_logs.type: 1=固有, 2=浮动 (0027 迁移注释)
ALTER TABLE risk_logs
    ADD CONSTRAINT ck_risk_logs_type         CHECK (type IN (1, 2)) NOT VALID;
