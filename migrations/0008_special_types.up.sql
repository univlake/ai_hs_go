-- special_types: dictionary of 特种设备类型 (boilers, pressure vessels, lifts, ...)
-- referenced by special_equipment.special_type_id.

CREATE TABLE special_types (
    id          BIGSERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);
CREATE INDEX idx_special_types_title ON special_types (title);

-- 0006 创建 special_equipment 时该表还没建，FK 在此回填。
ALTER TABLE special_equipment
    ADD CONSTRAINT special_equipment_special_type_id_fkey
    FOREIGN KEY (special_type_id) REFERENCES special_types(id) ON DELETE SET NULL;
