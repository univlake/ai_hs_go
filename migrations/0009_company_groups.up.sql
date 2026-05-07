-- 企业分组：一级企业可以把下级公司打成自定义集合（如"华北片区"），
-- 配合 bloc 视图做风险/统计聚合。
--
-- Laravel 对照：app/Models/CompanyGroup, CompanyGroupKid + 迁移
-- 2026_04_13_070005 / 070006。
--
-- 设计要点：
--   * 每个分组属于某个 owner（创建该分组的企业 / 一级企业），通过
--     company_id 锚定；无跨租户分组。
--   * 成员 (kids) 表存 group_id × company_id，含 (group_id, company_id)
--     唯一索引避免重复添加。
--   * 软删除复用统一的 deleted_at 模式。

CREATE TABLE company_groups (
    id           BIGSERIAL PRIMARY KEY,
    name         VARCHAR(255) NOT NULL,
    company_id   BIGINT NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
    user_id      BIGINT NOT NULL DEFAULT 0,
    last_user_id BIGINT NOT NULL DEFAULT 0,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMPTZ
);
CREATE INDEX idx_company_groups_company ON company_groups (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_company_groups_name    ON company_groups (name);

CREATE TABLE company_group_kids (
    id               BIGSERIAL PRIMARY KEY,
    company_group_id BIGINT NOT NULL REFERENCES company_groups(id) ON DELETE CASCADE,
    company_id       BIGINT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at       TIMESTAMPTZ
);
CREATE UNIQUE INDEX uq_company_group_kids_pair
    ON company_group_kids (company_group_id, company_id)
    WHERE deleted_at IS NULL;
CREATE INDEX idx_company_group_kids_company
    ON company_group_kids (company_id) WHERE deleted_at IS NULL;
