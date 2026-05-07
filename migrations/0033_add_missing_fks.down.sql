-- Reverse 0033: drop all FKs and restore DEFAULT 0 sentinels.

ALTER TABLE float_risk_logs
    DROP CONSTRAINT IF EXISTS float_risk_logs_company_id_fkey,
    ALTER COLUMN company_id SET DEFAULT 0;

ALTER TABLE inherent_risk_logs
    DROP CONSTRAINT IF EXISTS inherent_risk_logs_company_id_fkey,
    ALTER COLUMN company_id SET DEFAULT 0;

ALTER TABLE regulation_libraries
    DROP CONSTRAINT IF EXISTS regulation_libraries_company_id_fkey,
    DROP CONSTRAINT IF EXISTS regulation_libraries_regulation_library_type_id_fkey,
    ALTER COLUMN company_id SET DEFAULT 0;

ALTER TABLE risk_logs
    DROP CONSTRAINT IF EXISTS risk_logs_company_id_fkey,
    ALTER COLUMN company_id SET DEFAULT 0;

ALTER TABLE training_video_logs
    DROP CONSTRAINT IF EXISTS training_video_logs_training_video_id_fkey,
    DROP CONSTRAINT IF EXISTS training_video_logs_company_id_fkey;

ALTER TABLE training_file_logs
    DROP CONSTRAINT IF EXISTS training_file_logs_training_file_id_fkey,
    DROP CONSTRAINT IF EXISTS training_file_logs_company_id_fkey;

ALTER TABLE norm_checks
    DROP CONSTRAINT IF EXISTS norm_checks_company_id_fkey;

ALTER TABLE system_checks
    DROP CONSTRAINT IF EXISTS system_checks_company_id_fkey;

ALTER TABLE add_points
    DROP CONSTRAINT IF EXISTS add_points_company_id_fkey,
    DROP CONSTRAINT IF EXISTS add_points_template_id_fkey;

ALTER TABLE accident_checks
    DROP CONSTRAINT IF EXISTS accident_checks_company_id_fkey,
    DROP CONSTRAINT IF EXISTS accident_checks_template_id_fkey;

ALTER TABLE daily_checks
    DROP CONSTRAINT IF EXISTS daily_checks_company_id_fkey,
    DROP CONSTRAINT IF EXISTS daily_checks_template_id_fkey;

ALTER TABLE focus_companies
    DROP CONSTRAINT IF EXISTS focus_companies_company_id_fkey,
    DROP CONSTRAINT IF EXISTS focus_companies_focus_id_fkey;

ALTER TABLE information_drafts
    DROP CONSTRAINT IF EXISTS information_drafts_next_reviewer_fkey;

ALTER TABLE meeting_notices
    DROP CONSTRAINT IF EXISTS meeting_notices_company_group_id_fkey;

ALTER TABLE new_notify_scopes
    DROP CONSTRAINT IF EXISTS new_notify_scopes_company_group_id_fkey;

ALTER TABLE office_notify_scopes
    DROP CONSTRAINT IF EXISTS office_notify_scopes_company_group_id_fkey;

ALTER TABLE internal_meeting_kids
    DROP CONSTRAINT IF EXISTS internal_meeting_kids_internal_meeting_id_fkey;

ALTER TABLE skill_unit_kids
    DROP CONSTRAINT IF EXISTS skill_unit_kids_skill_unit_id_fkey;

ALTER TABLE companies
    DROP CONSTRAINT IF EXISTS companies_area_id_fkey,
    DROP CONSTRAINT IF EXISTS companies_operational_type_id_fkey,
    DROP CONSTRAINT IF EXISTS companies_enterprise_type_id_fkey;
