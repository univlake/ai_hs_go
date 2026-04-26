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
