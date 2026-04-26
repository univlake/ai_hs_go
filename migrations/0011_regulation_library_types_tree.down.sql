DROP INDEX IF EXISTS idx_regulation_library_types_pid;
ALTER TABLE regulation_library_types
    DROP COLUMN IF EXISTS last_user_id,
    DROP COLUMN IF EXISTS user_id,
    DROP COLUMN IF EXISTS company_id,
    DROP COLUMN IF EXISTS all_name,
    DROP COLUMN IF EXISTS pid;
