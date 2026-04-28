-- 0012 declared safety_architectures.company_id and buildings.company_id as
-- nullable references to companies. Every other table in safety_archive_module
-- uses NOT NULL — the relaxed declarations were an oversight that lets rows
-- bypass per-company filters (`WHERE company_id = ?` silently drops NULL rows).
--
-- This migration tightens both columns. It will fail loudly if NULL rows
-- exist; the operator should clean them up first (orphaned rows have no
-- meaningful tenant and should be deleted or backfilled).

ALTER TABLE safety_architectures ALTER COLUMN company_id SET NOT NULL;
ALTER TABLE buildings            ALTER COLUMN company_id SET NOT NULL;
