-- 隐患信息 + 隐患附件 — 专家检查模块底层支持。
--
-- hidden_danger_infos    隐患主表（CRUD + rectify / confirm-rectification / archive 三态流转）
-- hidden_upload_annexes  隐患三类附件（is_review = 1 排查证据 / 2 复查证据 / 3 整改证据）
--
-- 跟 0023_expert_module 配套：DetailDrawer 一次拉「受检企业 + 隐患」两份数据，
-- 之前 stub 返回 200 + [] 让列表能渲染；本表上线后真实数据接入。

CREATE TABLE hidden_danger_infos (
    id                    BIGSERIAL PRIMARY KEY,
    expert_inspection_id  BIGINT REFERENCES expert_inspections(id) ON DELETE CASCADE,
    expert_company_id     BIGINT REFERENCES expert_companies(id)   ON DELETE CASCADE,
    check_company_id      BIGINT REFERENCES companies(id)          ON DELETE RESTRICT,
    area                  VARCHAR(255) NOT NULL,
    "session"             VARCHAR(255) NOT NULL DEFAULT '第一轮',
    description           TEXT NOT NULL,
    hidden_type           VARCHAR(255) NOT NULL,
    hidden_level          VARCHAR(255) NOT NULL,
    legalises             TEXT,
    proposal              TEXT,
    correct_deadline      DATE,
    invest_status         SMALLINT NOT NULL DEFAULT 1,
    zg_user_id            BIGINT,
    correct_status        SMALLINT NOT NULL DEFAULT 1,
    correct_date          DATE,
    review_description    TEXT,
    review_user_id        BIGINT,
    review_time           TIMESTAMPTZ,
    archive_user_id       BIGINT,
    "comment"             TEXT,
    company_id            BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id               BIGINT NOT NULL DEFAULT 0,
    last_user_id          BIGINT NOT NULL DEFAULT 0,
    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at            TIMESTAMPTZ
);
CREATE INDEX idx_hidden_danger_infos_inspection      ON hidden_danger_infos (expert_inspection_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_hidden_danger_infos_expert_company  ON hidden_danger_infos (expert_company_id)    WHERE deleted_at IS NULL;
CREATE INDEX idx_hidden_danger_infos_check_company   ON hidden_danger_infos (check_company_id)     WHERE deleted_at IS NULL;
CREATE INDEX idx_hidden_danger_infos_company         ON hidden_danger_infos (company_id)           WHERE deleted_at IS NULL;
CREATE INDEX idx_hidden_danger_infos_correct_status  ON hidden_danger_infos (correct_status)       WHERE deleted_at IS NULL;
CREATE INDEX idx_hidden_danger_infos_invest_status   ON hidden_danger_infos (invest_status)        WHERE deleted_at IS NULL;

CREATE TABLE hidden_upload_annexes (
    id                     BIGSERIAL PRIMARY KEY,
    hidden_danger_info_id  BIGINT NOT NULL REFERENCES hidden_danger_infos(id) ON DELETE CASCADE,
    geolocation            VARCHAR(255) NOT NULL DEFAULT '',
    is_review              SMALLINT NOT NULL DEFAULT 1,
    annex_url              VARCHAR(255) NOT NULL,
    user_id                BIGINT NOT NULL DEFAULT 0,
    created_at             TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at             TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at             TIMESTAMPTZ
);
CREATE INDEX idx_hidden_upload_annexes_parent ON hidden_upload_annexes (hidden_danger_info_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_hidden_upload_annexes_kind   ON hidden_upload_annexes (hidden_danger_info_id, is_review) WHERE deleted_at IS NULL;
