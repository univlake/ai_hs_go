-- 清理 seeder 已删除的死权限（前端无对应 UI 或不需要独立页面）。
-- 关联表 role_has_permissions / model_has_permissions 通过 ON DELETE CASCADE 自动清理；
-- casbin_rule 需要手工删除对应 policy。

DELETE FROM casbin_rule
WHERE ptype = 'p' AND (
    (v1 = 'bloc-accident-bulletins' AND v2 = 'list')
 OR (v1 = 'todos'                   AND v2 = 'list')
 OR (v1 = 'hidden-danger-infos'     AND v2 IN ('list', 'create', 'update', 'delete'))
 OR (v1 = 'focus-companies'         AND v2 IN ('list', 'create', 'delete'))
 OR (v1 = 'month-statistics'        AND v2 = 'delete')
 OR (v1 = 'year-statistics'         AND v2 = 'delete')
);

DELETE FROM permissions
WHERE guard_name = 'api' AND name IN (
    'bloc-accident-bulletins.list',
    'todos.list',
    'hidden-danger-infos.list',
    'hidden-danger-infos.create',
    'hidden-danger-infos.update',
    'hidden-danger-infos.delete',
    'focus-companies.list',
    'focus-companies.create',
    'focus-companies.delete',
    'month-statistics.delete',
    'year-statistics.delete'
);
