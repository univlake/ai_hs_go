DROP TRIGGER IF EXISTS trg_companies_root_id_after  ON companies;
DROP TRIGGER IF EXISTS trg_companies_root_id_insert ON companies;
DROP FUNCTION IF EXISTS companies_finalize_root_id();
DROP FUNCTION IF EXISTS companies_set_root_id();
DROP TABLE IF EXISTS companies;
DROP TABLE IF EXISTS operational_types;
DROP TABLE IF EXISTS enterprise_types;
DROP TABLE IF EXISTS base_areas;
