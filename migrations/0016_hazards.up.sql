-- 危险有害因素（hazards）— 基础数据字典补漏
--
-- 0010_basedata_dictionaries 当时漏掉这一张：confined-space 页面的多选框、
-- 以及未来 风险源 / 隐患排查 类模块都从这里选项。
-- 形状和其它 base-data 字典一致：id + title + sort + timestamps + 软删除，
-- 走 DictService 共用的 CRUD。

CREATE TABLE hazards (
    id          BIGSERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    sort        INTEGER NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);
CREATE INDEX idx_hazards_title ON hazards (title);
