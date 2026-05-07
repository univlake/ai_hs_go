-- 基础数据 / 字典模块
--
-- 9 个字典表共享同一形状（id + title + sort + timestamps + soft-delete）；
-- 6 张需要新建，3 张已经存在但缺 sort 字段，统一补齐使前端 CrudPage 能
-- 用同一份配置驱动。
-- 另外还有 checklist_libraries —— 检查表库，结构略复杂带附件。
--
-- 未引入 Laravel 端 safety_management_types.is_cert /
-- regulation_library_types.{pid, all_name, company_id, ...} 等"特例字段"，
-- 因为前端 system/*-types 页面统一只暴露 title + sort。后续如需要这些
-- 维度，单独迁移升级而不是塞进这个公共骨架。

CREATE TABLE record_types (
    id          BIGSERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    sort        INTEGER NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);
CREATE INDEX idx_record_types_title ON record_types (title);

CREATE TABLE process_types (
    id          BIGSERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    sort        INTEGER NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);
CREATE INDEX idx_process_types_title ON process_types (title);

CREATE TABLE space_work_types (
    id          BIGSERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    sort        INTEGER NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);
CREATE INDEX idx_space_work_types_title ON space_work_types (title);

CREATE TABLE safety_management_types (
    id          BIGSERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    sort        INTEGER NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);
CREATE INDEX idx_safety_management_types_title ON safety_management_types (title);

CREATE TABLE dangerous_craft_types (
    id          BIGSERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    sort        INTEGER NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);
CREATE INDEX idx_dangerous_craft_types_title ON dangerous_craft_types (title);

CREATE TABLE regulation_library_types (
    id          BIGSERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    sort        INTEGER NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);
CREATE INDEX idx_regulation_library_types_title ON regulation_library_types (title);

-- Existing dict tables — backfill the sort column the frontend CrudPage
-- expects on every list. Default 0 keeps current ordering (by id) intact.
ALTER TABLE enterprise_types  ADD COLUMN sort INTEGER NOT NULL DEFAULT 0;
ALTER TABLE operational_types ADD COLUMN sort INTEGER NOT NULL DEFAULT 0;
ALTER TABLE special_types     ADD COLUMN sort INTEGER NOT NULL DEFAULT 0;

-- 检查表库 — title + 附件 URL 对(预览 + 实际)。company_id 锚定到企业，
-- 软删除字段保留 Laravel 行为。
CREATE TABLE checklist_libraries (
    id                BIGSERIAL PRIMARY KEY,
    title             VARCHAR(255) NOT NULL,
    status            SMALLINT NOT NULL DEFAULT 1,
    preview_annex_url TEXT,
    annex_url         TEXT,
    company_id        BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id           BIGINT NOT NULL DEFAULT 0,
    last_user_id      BIGINT NOT NULL DEFAULT 0,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at        TIMESTAMPTZ
);
CREATE INDEX idx_checklist_libraries_company ON checklist_libraries (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_checklist_libraries_title   ON checklist_libraries (title);
