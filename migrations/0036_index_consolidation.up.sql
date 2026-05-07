-- 索引整理：统一 partial 索引策略 + 删除冗余 + 补 (company_id, created_at) 复合。
--
-- A. 5 个非 partial 的 idx_*_company 改成 WHERE deleted_at IS NULL — 业务查询
--    全都带 deleted_at IS NULL，partial 让索引体积/扫描行数都小一截。
-- B. idx_users_status 同理改 partial。
-- C. 0025 考核模块的 5 个 (company_id) 单列索引被 (company_id, year) 复合
--    覆盖，删掉省维护开销。
-- D. 给前端高频列表表加 (company_id, created_at DESC) partial 复合索引，
--    让 OFFSET 翻页 / 倒序分页走索引。子表（按 parent_id 查）不在此列表。

-- A. partial 化非 partial idx_*_company ─────────────────────────
DROP INDEX IF EXISTS idx_chemical_ingredients_company;
CREATE INDEX idx_chemical_ingredients_company ON chemical_ingredients (company_id) WHERE deleted_at IS NULL;

DROP INDEX IF EXISTS idx_major_constitutes_company;
CREATE INDEX idx_major_constitutes_company ON major_constitutes (company_id) WHERE deleted_at IS NULL;

DROP INDEX IF EXISTS idx_personnels_company;
CREATE INDEX idx_personnels_company ON personnels (company_id) WHERE deleted_at IS NULL;

DROP INDEX IF EXISTS idx_skill_unit_kids_company;
CREATE INDEX idx_skill_unit_kids_company ON skill_unit_kids (company_id) WHERE deleted_at IS NULL;

DROP INDEX IF EXISTS idx_training_file_logs_company;
CREATE INDEX idx_training_file_logs_company ON training_file_logs (company_id) WHERE deleted_at IS NULL;

DROP INDEX IF EXISTS idx_training_video_logs_company;
CREATE INDEX idx_training_video_logs_company ON training_video_logs (company_id) WHERE deleted_at IS NULL;

-- B. idx_users_status partial ─────────────────────────────────
DROP INDEX IF EXISTS idx_users_status;
CREATE INDEX idx_users_status ON users (status) WHERE deleted_at IS NULL;

-- C. 删 0025 冗余 (company_id) 单列索引 ─────────────────────────
-- (company_id, year) 已存在并能服务 WHERE company_id=? 这种查询。
DROP INDEX IF EXISTS idx_daily_checks_company;
DROP INDEX IF EXISTS idx_accident_checks_company;
DROP INDEX IF EXISTS idx_add_points_company;
DROP INDEX IF EXISTS idx_system_checks_company;
DROP INDEX IF EXISTS idx_norm_checks_company;

-- D. (company_id, created_at DESC) partial 复合索引 ───────────
-- 仅对前端按 created_at 倒序分页的主列表表添加；子表（按 parent_id 查）不在此列表。
CREATE INDEX idx_expert_inspections_co_ct  ON expert_inspections    (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_internal_meetings_co_ct   ON internal_meetings     (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_internal_trainings_co_ct  ON internal_trainings    (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_external_trainings_co_ct  ON external_trainings    (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_other_trainings_co_ct     ON other_trainings       (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_health_files_co_ct        ON health_files          (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_monthly_inv_co_ct         ON monthly_investigations (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_accident_bulletins_co_ct  ON accident_bulletins    (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_meetings_co_ct            ON meetings              (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_office_notifies_co_ct     ON office_notifies       (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_new_notifies_co_ct        ON new_notifies          (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_information_drafts_co_ct  ON information_drafts    (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_change_manages_co_ct      ON change_manages        (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_hidden_danger_co_ct       ON hidden_danger_infos   (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_assess_detects_co_ct      ON assess_detects        (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_three_togethers_co_ct     ON three_togethers       (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_job_management_co_ct      ON job_management        (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_space_works_co_ct         ON space_works           (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_dangerous_equipment_co_ct ON dangerous_equipment   (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_special_equipment_co_ct   ON special_equipment     (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_hazardous_chemicals_co_ct ON hazardous_chemicals   (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_design_units_co_ct        ON design_units          (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_evaluation_units_co_ct    ON evaluation_units      (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_skill_units_co_ct         ON skill_units           (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_fire_safeties_co_ct       ON fire_safeties         (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_health_systems_co_ct      ON health_systems        (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_health_reports_co_ct      ON health_reports        (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_occupational_hazards_co_ct ON occupational_hazards (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_health_checkups_co_ct     ON health_checkups       (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_emergency_drills_co_ct    ON emergency_drills      (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_emergency_supplies_co_ct  ON emergency_supplies    (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_funding_executions_co_ct  ON funding_executions    (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_security_investments_co_ct ON security_investments (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_training_files_co_ct      ON training_files        (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_training_videos_co_ct     ON training_videos       (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_regulation_libraries_co_ct ON regulation_libraries (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_fire_facilities_co_ct     ON fire_facilities       (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_fire_equipment_co_ct      ON fire_equipment        (company_id, created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_elimination_control_co_ct ON elimination_control_staff (company_id, created_at DESC) WHERE deleted_at IS NULL;
