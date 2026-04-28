-- 回滚仅恢复 permissions 表记录；role/casbin 关联无法重建（信息已丢失），
-- 需要重新跑 seed 才能恢复角色绑定。
INSERT INTO permissions (name, guard_name) VALUES
    ('bloc-accident-bulletins.list', 'api'),
    ('todos.list',                    'api'),
    ('hidden-danger-infos.list',      'api'),
    ('hidden-danger-infos.create',    'api'),
    ('hidden-danger-infos.update',    'api'),
    ('hidden-danger-infos.delete',    'api'),
    ('focus-companies.list',          'api'),
    ('focus-companies.create',        'api'),
    ('focus-companies.delete',        'api'),
    ('month-statistics.delete',       'api'),
    ('year-statistics.delete',        'api')
ON CONFLICT (name, guard_name) DO NOTHING;
