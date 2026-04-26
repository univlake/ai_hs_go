-- 工艺信息（Process Information）模块 — 三张表
--
-- major_hazardous_processes  重点监管危险工艺：name + 工艺类型 + SIS 字段
-- hazardous_processes        较大危险工艺：process_type + 4 个布尔标记
-- space_works                有限空间辨识：name/position + space_work_type
--
-- 风格沿用 0006_safety_ledger.up.sql / 0014_major_hazards.up.sql：
-- company_id NOT NULL + audit + 软删除；外键全部 ON DELETE SET NULL，
-- 这样字典或物料/设备被清掉也不会拖死工艺记录。

CREATE TABLE major_hazardous_processes (
    id                       BIGSERIAL PRIMARY KEY,
    "name"                   VARCHAR(255),
    dangerous_craft_type_id  BIGINT REFERENCES dangerous_craft_types(id) ON DELETE SET NULL,
    "position"               VARCHAR(255),
    hazardous_chemical_id    BIGINT REFERENCES hazardous_chemicals(id) ON DELETE SET NULL,
    dangerous_equipment_id   BIGINT REFERENCES dangerous_equipment(id) ON DELETE SET NULL,
    operating_people_num     VARCHAR(64),
    temperature              VARCHAR(64),
    pressure                 VARCHAR(64),
    pressure_unit            VARCHAR(32),
    is_sis                   SMALLINT NOT NULL DEFAULT 0,
    sil_level                VARCHAR(32),
    detect_parameter         VARCHAR(255),
    company_id               BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id                  BIGINT NOT NULL DEFAULT 0,
    last_user_id             BIGINT NOT NULL DEFAULT 0,
    created_at               TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at               TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at               TIMESTAMPTZ
);
CREATE INDEX idx_major_hazardous_processes_company   ON major_hazardous_processes (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_major_hazardous_processes_craft     ON major_hazardous_processes (dangerous_craft_type_id);
CREATE INDEX idx_major_hazardous_processes_chemical  ON major_hazardous_processes (hazardous_chemical_id);
CREATE INDEX idx_major_hazardous_processes_equipment ON major_hazardous_processes (dangerous_equipment_id);

CREATE TABLE hazardous_processes (
    id                      BIGSERIAL PRIMARY KEY,
    process_type_id         BIGINT REFERENCES process_types(id) ON DELETE SET NULL,
    "position"              VARCHAR(255),
    hazardous_chemical_id   BIGINT REFERENCES hazardous_chemicals(id) ON DELETE SET NULL,
    dangerous_equipment_id  BIGINT REFERENCES dangerous_equipment(id) ON DELETE SET NULL,
    other_equipment_name    VARCHAR(255),
    operating_people_num    VARCHAR(64),
    is_online_test          SMALLINT NOT NULL DEFAULT 0,
    is_automatic            SMALLINT NOT NULL DEFAULT 0,
    is_control              SMALLINT NOT NULL DEFAULT 0,
    is_dust                 SMALLINT NOT NULL DEFAULT 0,
    company_id              BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id                 BIGINT NOT NULL DEFAULT 0,
    last_user_id            BIGINT NOT NULL DEFAULT 0,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at              TIMESTAMPTZ
);
CREATE INDEX idx_hazardous_processes_company     ON hazardous_processes (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_hazardous_processes_type        ON hazardous_processes (process_type_id);
CREATE INDEX idx_hazardous_processes_chemical    ON hazardous_processes (hazardous_chemical_id);
CREATE INDEX idx_hazardous_processes_equipment   ON hazardous_processes (dangerous_equipment_id);

-- space_works.hazard_id 是逗号分隔的 hazards.id 列表（前端 toggleHazard 拼装），
-- 不建外键 — 多对多以 VARCHAR 表达，遵循 Laravel 原结构。
CREATE TABLE space_works (
    id                  BIGSERIAL PRIMARY KEY,
    "name"              VARCHAR(255),
    "position"          VARCHAR(255),
    space_work_type_id  BIGINT REFERENCES space_work_types(id) ON DELETE SET NULL,
    "size"              VARCHAR(64),
    people_num          INTEGER,
    hazard_id           VARCHAR(255),
    word_content        TEXT,
    company_id          BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id             BIGINT NOT NULL DEFAULT 0,
    last_user_id        BIGINT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);
CREATE INDEX idx_space_works_company ON space_works (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_space_works_type    ON space_works (space_work_type_id);
