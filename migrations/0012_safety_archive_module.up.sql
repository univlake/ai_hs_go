-- 安全档案模块的 4 张新表
-- 5 个模块里 personnels 已经在 0004 建好；其余 4 张这里一并补齐。
--
-- 这 4 张表共享一组「业务列 + company_id 锚点 + audit + 软删除」骨架，
-- 对照 Laravel 各自的迁移文件 (2026_04_*) 一对一搬过来；时间戳列统一
-- TIMESTAMPTZ + DEFAULT NOW() 跟项目里其它表保持一致。

-- 1) 安全架构 — 每家企业 0~1 条；annex_url 是 OSS 路径
CREATE TABLE safety_architectures (
    id            BIGSERIAL PRIMARY KEY,
    thumbnail_url TEXT,
    annex_url     TEXT,
    company_id    BIGINT REFERENCES companies(id) ON DELETE CASCADE,
    user_id       BIGINT NOT NULL DEFAULT 0,
    last_user_id  BIGINT NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ
);
CREATE INDEX idx_safety_architectures_company ON safety_architectures (company_id) WHERE deleted_at IS NULL;

-- 2) 安全资质（备案/许可）— 多对一指向 record_types 字典
CREATE TABLE filing_permissions (
    id                 BIGSERIAL PRIMARY KEY,
    record_type_id     BIGINT REFERENCES record_types(id) ON DELETE SET NULL,
    issuing_department VARCHAR(255),
    filing_range       TEXT,
    filing_time        DATE,
    valid_period       DATE,
    comment            TEXT,
    annex_url          TEXT,
    company_id         BIGINT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    user_id            BIGINT NOT NULL DEFAULT 0,
    last_user_id       BIGINT NOT NULL DEFAULT 0,
    created_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at         TIMESTAMPTZ
);
CREATE INDEX idx_filing_permissions_company  ON filing_permissions (company_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_filing_permissions_record   ON filing_permissions (record_type_id);

-- 3) 建筑物及构筑物 — Laravel 的 structure 列在后续迁移补的，这里一次性写入
CREATE TABLE buildings (
    id           BIGSERIAL PRIMARY KEY,
    title        VARCHAR(255),
    structure    VARCHAR(255),
    area         VARCHAR(64),
    fire_class   VARCHAR(64),
    company_id   BIGINT REFERENCES companies(id) ON DELETE CASCADE,
    user_id      BIGINT NOT NULL DEFAULT 0,
    last_user_id BIGINT NOT NULL DEFAULT 0,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMPTZ
);
CREATE INDEX idx_buildings_company ON buildings (company_id) WHERE deleted_at IS NULL;

-- 4) 安全投入 — 按年份归档 + 类型枚举（前端 select 控制取值）
CREATE TABLE security_investments (
    id                BIGSERIAL PRIMARY KEY,
    year              VARCHAR(8),
    type              VARCHAR(64),
    title             VARCHAR(255),
    content           TEXT,
    budget            NUMERIC(14, 2),
    last_year_income  NUMERIC(14, 2),
    company_id        BIGINT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    user_id           BIGINT NOT NULL DEFAULT 0,
    last_user_id      BIGINT NOT NULL DEFAULT 0,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at        TIMESTAMPTZ
);
CREATE INDEX idx_security_investments_company_year ON security_investments (company_id, year) WHERE deleted_at IS NULL;
