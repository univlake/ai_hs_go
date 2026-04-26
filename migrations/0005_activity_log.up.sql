-- Spatie-compatible activity_log schema. attribute_changes / properties as
-- JSONB to support GIN indexing instead of Laravel's TEXT.

CREATE TABLE activity_log (
    id                  BIGSERIAL PRIMARY KEY,
    log_name            VARCHAR(255),
    description         TEXT NOT NULL,
    subject_type        VARCHAR(255),
    subject_id          BIGINT,
    causer_type         VARCHAR(255),
    causer_id           BIGINT,
    event               VARCHAR(64),
    properties          JSONB NOT NULL DEFAULT '{}'::jsonb,
    attribute_changes   JSONB NOT NULL DEFAULT '{}'::jsonb,
    batch_uuid          UUID,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_activity_log_subject  ON activity_log (subject_type, subject_id);
CREATE INDEX idx_activity_log_causer   ON activity_log (causer_type, causer_id);
CREATE INDEX idx_activity_log_log_name ON activity_log (log_name);
CREATE INDEX idx_activity_log_event    ON activity_log (event);
CREATE INDEX idx_activity_log_props    ON activity_log USING GIN (properties);
