DO $$
DECLARE
    rec RECORD;
BEGIN
    FOR rec IN
        SELECT c.table_name
        FROM information_schema.columns c
        WHERE c.table_schema = 'public'
          AND c.column_name = 'updated_at'
    LOOP
        EXECUTE format('DROP TRIGGER IF EXISTS trg_%I_set_updated_at ON %I;',
                       rec.table_name, rec.table_name);
    END LOOP;
END;
$$;

DROP FUNCTION IF EXISTS set_updated_at();
