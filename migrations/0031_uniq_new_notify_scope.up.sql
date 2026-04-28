-- 0007 added new_notify_scopes without a uniqueness constraint on
-- (notify_id, user_id), so the same user can be appended to the same notify
-- multiple times if the application code forgets to dedup. Make duplicates
-- impossible at the database layer.
--
-- If duplicates already exist, this migration will fail; clean them with:
--   DELETE FROM new_notify_scopes a
--   USING new_notify_scopes b
--   WHERE a.id < b.id
--     AND a.notify_id = b.notify_id
--     AND a.user_id   = b.user_id;

ALTER TABLE new_notify_scopes
    ADD CONSTRAINT uniq_new_notify_scope UNIQUE (notify_id, user_id);
