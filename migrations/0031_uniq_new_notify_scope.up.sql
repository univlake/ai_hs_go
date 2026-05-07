-- 0007 originally created new_notify_scopes with column notify_id; 0022 then
-- DROPped that table and rebuilt it with new_notify_id (plus company_group_*,
-- audit fields, deleted_at). This migration enforces (new_notify_id, user_id)
-- uniqueness scoped to live rows so soft-deleting and re-adding the same pair
-- works; a regular UNIQUE would block re-adds forever after a soft delete.
--
-- If duplicate live rows exist, the migration fails; clean them with:
--   DELETE FROM new_notify_scopes a USING new_notify_scopes b
--   WHERE a.id < b.id
--     AND a.deleted_at IS NULL AND b.deleted_at IS NULL
--     AND a.new_notify_id = b.new_notify_id
--     AND a.user_id = b.user_id;

CREATE UNIQUE INDEX uniq_new_notify_scope
    ON new_notify_scopes (new_notify_id, user_id)
    WHERE deleted_at IS NULL;
