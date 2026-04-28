-- 0006 declared special_equipment.special_type_id BIGINT and added an index,
-- but skipped the foreign key to special_types(id). Other lookup references
-- in this module (e.g. hazardous_chemical_id, company_id) all carry a FK.
-- Match the convention so orphan rows can't accumulate.
--
-- ON DELETE SET NULL keeps equipment rows visible if a special_type is later
-- removed — the deletion of a lookup value should not cascade-delete equipment.

ALTER TABLE special_equipment
    ADD CONSTRAINT special_equipment_special_type_id_fkey
    FOREIGN KEY (special_type_id) REFERENCES special_types(id) ON DELETE SET NULL;
