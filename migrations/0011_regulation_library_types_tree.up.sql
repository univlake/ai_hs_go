-- 法规库类型升级为多层级（树）
--
-- 0010 把它当成普通字典建了 {id, title, sort, ts...}; 现在补齐 Laravel
-- 项目里实际的列：pid (父级 id, 0=顶级)、all_name (父级 / 自身的全路径
-- 字符串，便于展示层级)、company_id (留作租户隔离的占位，目前都填 0
-- 因为字典是全局的)、user_id / last_user_id (审计)。
--
-- pid 默认 0，意味着已有 26 条种子记录都默认成顶级 — 配合 seeder 中
-- 同步把 all_name 写成自己的 title 即可。

ALTER TABLE regulation_library_types
    ADD COLUMN pid          BIGINT  NOT NULL DEFAULT 0,
    ADD COLUMN all_name     VARCHAR(255),
    ADD COLUMN company_id   BIGINT  NOT NULL DEFAULT 0,
    ADD COLUMN user_id      BIGINT  NOT NULL DEFAULT 0,
    ADD COLUMN last_user_id BIGINT  NOT NULL DEFAULT 0;

CREATE INDEX idx_regulation_library_types_pid ON regulation_library_types (pid);

-- Backfill all_name for any rows already inserted by the previous seed run
-- (single-level only — they're all top-level so all_name = title is correct).
UPDATE regulation_library_types
SET all_name = title
WHERE all_name IS NULL AND deleted_at IS NULL;
