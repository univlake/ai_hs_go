-- 通用 set_updated_at() 触发器，挂到所有有 updated_at 列的业务表上。
--
-- PG 不会自动维护 updated_at；之前完全靠 GORM hook 保证。任何 raw UPDATE
-- （psql 手工修复、软删除 SET deleted_at = NOW()、脚本批处理）都会留下
-- 错误的 updated_at。挂触发器后无论谁写都会被刷新。
--
-- 命名约定 trg_<table>_set_updated_at；新增表的迁移要自己挂触发器。
-- 函数 set_updated_at() 全局共用，下一份迁移可直接 EXECUTE FUNCTION 引用。

CREATE OR REPLACE FUNCTION set_updated_at() RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at := NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DO $$
DECLARE
    rec RECORD;
BEGIN
    FOR rec IN
        SELECT c.table_name
        FROM information_schema.columns c
        WHERE c.table_schema = 'public'
          AND c.column_name = 'updated_at'
        ORDER BY c.table_name
    LOOP
        EXECUTE format(
            'DROP TRIGGER IF EXISTS trg_%I_set_updated_at ON %I; '
            'CREATE TRIGGER trg_%I_set_updated_at '
            'BEFORE UPDATE ON %I '
            'FOR EACH ROW EXECUTE FUNCTION set_updated_at();',
            rec.table_name, rec.table_name, rec.table_name, rec.table_name
        );
    END LOOP;
END;
$$;
