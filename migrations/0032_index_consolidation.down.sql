-- D. drop composite indexes
DROP INDEX IF EXISTS idx_elimination_control_co_ct;
DROP INDEX IF EXISTS idx_fire_equipment_co_ct;
DROP INDEX IF EXISTS idx_fire_facilities_co_ct;
DROP INDEX IF EXISTS idx_regulation_libraries_co_ct;
DROP INDEX IF EXISTS idx_training_videos_co_ct;
DROP INDEX IF EXISTS idx_training_files_co_ct;
DROP INDEX IF EXISTS idx_security_investments_co_ct;
DROP INDEX IF EXISTS idx_funding_executions_co_ct;
DROP INDEX IF EXISTS idx_emergency_supplies_co_ct;
DROP INDEX IF EXISTS idx_emergency_drills_co_ct;
DROP INDEX IF EXISTS idx_health_checkups_co_ct;
DROP INDEX IF EXISTS idx_occupational_hazards_co_ct;
DROP INDEX IF EXISTS idx_health_reports_co_ct;
DROP INDEX IF EXISTS idx_health_systems_co_ct;
DROP INDEX IF EXISTS idx_fire_safeties_co_ct;
DROP INDEX IF EXISTS idx_skill_units_co_ct;
DROP INDEX IF EXISTS idx_evaluation_units_co_ct;
DROP INDEX IF EXISTS idx_design_units_co_ct;
DROP INDEX IF EXISTS idx_hazardous_chemicals_co_ct;
DROP INDEX IF EXISTS idx_special_equipment_co_ct;
DROP INDEX IF EXISTS idx_dangerous_equipment_co_ct;
DROP INDEX IF EXISTS idx_space_works_co_ct;
DROP INDEX IF EXISTS idx_job_management_co_ct;
DROP INDEX IF EXISTS idx_three_togethers_co_ct;
DROP INDEX IF EXISTS idx_assess_detects_co_ct;
DROP INDEX IF EXISTS idx_hidden_danger_co_ct;
DROP INDEX IF EXISTS idx_change_manages_co_ct;
DROP INDEX IF EXISTS idx_information_drafts_co_ct;
DROP INDEX IF EXISTS idx_new_notifies_co_ct;
DROP INDEX IF EXISTS idx_office_notifies_co_ct;
DROP INDEX IF EXISTS idx_meetings_co_ct;
DROP INDEX IF EXISTS idx_accident_bulletins_co_ct;
DROP INDEX IF EXISTS idx_monthly_inv_co_ct;
DROP INDEX IF EXISTS idx_health_files_co_ct;
DROP INDEX IF EXISTS idx_other_trainings_co_ct;
DROP INDEX IF EXISTS idx_external_trainings_co_ct;
DROP INDEX IF EXISTS idx_internal_trainings_co_ct;
DROP INDEX IF EXISTS idx_internal_meetings_co_ct;
DROP INDEX IF EXISTS idx_expert_inspections_co_ct;

-- C. restore 0025 single-column (company_id) partial indexes
CREATE INDEX idx_norm_checks_company     ON norm_checks     (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_system_checks_company   ON system_checks   (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_add_points_company      ON add_points      (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_accident_checks_company ON accident_checks (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_daily_checks_company    ON daily_checks    (company_id) WHERE deleted_at IS NULL;

-- B. restore idx_users_status non-partial
DROP INDEX IF EXISTS idx_users_status;
CREATE INDEX idx_users_status ON users (status);

-- A. restore non-partial idx_*_company
DROP INDEX IF EXISTS idx_training_video_logs_company;
CREATE INDEX idx_training_video_logs_company ON training_video_logs (company_id);

DROP INDEX IF EXISTS idx_training_file_logs_company;
CREATE INDEX idx_training_file_logs_company ON training_file_logs (company_id);

DROP INDEX IF EXISTS idx_skill_unit_kids_company;
CREATE INDEX idx_skill_unit_kids_company ON skill_unit_kids (company_id);

DROP INDEX IF EXISTS idx_personnels_company;
CREATE INDEX idx_personnels_company ON personnels (company_id);

DROP INDEX IF EXISTS idx_major_constitutes_company;
CREATE INDEX idx_major_constitutes_company ON major_constitutes (company_id);

DROP INDEX IF EXISTS idx_chemical_ingredients_company;
CREATE INDEX idx_chemical_ingredients_company ON chemical_ingredients (company_id);
