-- Constraint policy consistency + denormalized-snapshot documentation.
--
-- M9: ON DELETE CASCADE on companies(id) → RESTRICT.
--   Day-to-day paths soft-delete companies (deleted_at), so CASCADE never
--   fires; only a raw hard delete (operator script) triggers it. CASCADE
--   then silently cleans 15+ business tables, which is the wrong default —
--   operators should see RESTRICT and decide what to archive first. The
--   one CASCADE that's intentional (focus_companies.focus_id, the
--   "关注关系" link) stays — focus relations carry no business data.
--
-- M5: Snapshot columns get COMMENT ON COLUMN so future readers know the
--   string is captured at insert time and won't track later renames.

-- M9 ─ flip company_id ON DELETE from CASCADE to RESTRICT ────────
-- Pattern: drop, re-add. PG doesn't have ALTER CONSTRAINT … ON DELETE.

ALTER TABLE buildings
    DROP CONSTRAINT buildings_company_id_fkey,
    ADD CONSTRAINT buildings_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE certificate_statuses
    DROP CONSTRAINT certificate_statuses_company_id_fkey,
    ADD CONSTRAINT certificate_statuses_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE checklist_libraries
    DROP CONSTRAINT checklist_libraries_company_id_fkey,
    ADD CONSTRAINT checklist_libraries_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE company_groups
    DROP CONSTRAINT company_groups_company_id_fkey,
    ADD CONSTRAINT company_groups_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE company_group_kids
    DROP CONSTRAINT company_group_kids_company_id_fkey,
    ADD CONSTRAINT company_group_kids_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE elimination_control_staff
    DROP CONSTRAINT elimination_control_staff_company_id_fkey,
    ADD CONSTRAINT elimination_control_staff_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE filing_permissions
    DROP CONSTRAINT filing_permissions_company_id_fkey,
    ADD CONSTRAINT filing_permissions_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE funding_executions
    DROP CONSTRAINT funding_executions_company_id_fkey,
    ADD CONSTRAINT funding_executions_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE occupational_health_safeties
    DROP CONSTRAINT occupational_health_safeties_company_id_fkey,
    ADD CONSTRAINT occupational_health_safeties_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE personnels
    DROP CONSTRAINT personnels_company_id_fkey,
    ADD CONSTRAINT personnels_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE production_standardizations
    DROP CONSTRAINT production_standardizations_company_id_fkey,
    ADD CONSTRAINT production_standardizations_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE safety_architectures
    DROP CONSTRAINT safety_architectures_company_id_fkey,
    ADD CONSTRAINT safety_architectures_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE safety_management_staff
    DROP CONSTRAINT safety_management_staff_company_id_fkey,
    ADD CONSTRAINT safety_management_staff_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE security_investments
    DROP CONSTRAINT security_investments_company_id_fkey,
    ADD CONSTRAINT security_investments_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

-- M5 ─ snapshot column comments ─────────────────────────────────
COMMENT ON COLUMN risk_logs.company_name IS
    'Snapshot at row creation; does not track later companies.title renames. Join via company_id for current names.';

COMMENT ON COLUMN year_statistics.company_name IS
    'Snapshot at row creation; does not track later companies.title renames.';
COMMENT ON COLUMN year_statistics.base_area IS
    'Snapshot at row creation; does not track later companies.area_id changes.';
COMMENT ON COLUMN year_statistics.enterprise_type IS
    'Snapshot at row creation; does not track later companies.enterprise_type_id changes.';

COMMENT ON COLUMN information_drafts.company_name IS
    'Snapshot at row creation; does not track later companies.title renames.';

COMMENT ON COLUMN office_notify_scopes.company_name IS
    'Snapshot at row creation; for display when notification was issued.';
COMMENT ON COLUMN office_notify_scopes.company_group_name IS
    'Snapshot at row creation; for display when notification was issued.';

COMMENT ON COLUMN new_notify_scopes.company_name IS
    'Snapshot at row creation; for display when notification was issued.';
COMMENT ON COLUMN new_notify_scopes.company_group_name IS
    'Snapshot at row creation; for display when notification was issued.';

COMMENT ON COLUMN meeting_notices.company_name IS
    'Snapshot at row creation; for display in meeting recipient list.';
COMMENT ON COLUMN meeting_notices.company_group_name IS
    'Snapshot at row creation; for display in meeting recipient list.';

COMMENT ON COLUMN meeting_sign_ups.company_name IS
    'Snapshot at sign-up; for historical sign-up sheet display.';
