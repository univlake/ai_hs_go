-- Reverse: drop comments, restore CASCADE policy.

COMMENT ON COLUMN meeting_sign_ups.company_name      IS NULL;
COMMENT ON COLUMN meeting_notices.company_group_name IS NULL;
COMMENT ON COLUMN meeting_notices.company_name       IS NULL;
COMMENT ON COLUMN new_notify_scopes.company_group_name    IS NULL;
COMMENT ON COLUMN new_notify_scopes.company_name          IS NULL;
COMMENT ON COLUMN office_notify_scopes.company_group_name IS NULL;
COMMENT ON COLUMN office_notify_scopes.company_name       IS NULL;
COMMENT ON COLUMN information_drafts.company_name    IS NULL;
COMMENT ON COLUMN year_statistics.enterprise_type    IS NULL;
COMMENT ON COLUMN year_statistics.base_area          IS NULL;
COMMENT ON COLUMN year_statistics.company_name       IS NULL;
COMMENT ON COLUMN risk_logs.company_name             IS NULL;

ALTER TABLE security_investments
    DROP CONSTRAINT security_investments_company_id_fkey,
    ADD CONSTRAINT security_investments_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;

ALTER TABLE safety_management_staff
    DROP CONSTRAINT safety_management_staff_company_id_fkey,
    ADD CONSTRAINT safety_management_staff_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;

ALTER TABLE safety_architectures
    DROP CONSTRAINT safety_architectures_company_id_fkey,
    ADD CONSTRAINT safety_architectures_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;

ALTER TABLE production_standardizations
    DROP CONSTRAINT production_standardizations_company_id_fkey,
    ADD CONSTRAINT production_standardizations_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;

ALTER TABLE personnels
    DROP CONSTRAINT personnels_company_id_fkey,
    ADD CONSTRAINT personnels_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;

ALTER TABLE occupational_health_safeties
    DROP CONSTRAINT occupational_health_safeties_company_id_fkey,
    ADD CONSTRAINT occupational_health_safeties_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;

ALTER TABLE funding_executions
    DROP CONSTRAINT funding_executions_company_id_fkey,
    ADD CONSTRAINT funding_executions_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;

ALTER TABLE filing_permissions
    DROP CONSTRAINT filing_permissions_company_id_fkey,
    ADD CONSTRAINT filing_permissions_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;

ALTER TABLE elimination_control_staff
    DROP CONSTRAINT elimination_control_staff_company_id_fkey,
    ADD CONSTRAINT elimination_control_staff_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;

ALTER TABLE company_group_kids
    DROP CONSTRAINT company_group_kids_company_id_fkey,
    ADD CONSTRAINT company_group_kids_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;

ALTER TABLE company_groups
    DROP CONSTRAINT company_groups_company_id_fkey,
    ADD CONSTRAINT company_groups_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;

ALTER TABLE checklist_libraries
    DROP CONSTRAINT checklist_libraries_company_id_fkey,
    ADD CONSTRAINT checklist_libraries_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;

ALTER TABLE certificate_statuses
    DROP CONSTRAINT certificate_statuses_company_id_fkey,
    ADD CONSTRAINT certificate_statuses_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;

ALTER TABLE buildings
    DROP CONSTRAINT buildings_company_id_fkey,
    ADD CONSTRAINT buildings_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;
