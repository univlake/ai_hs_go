import { CrudPage } from '@/components/shared/CrudPage';
import type { CrudPageConfig } from '@/components/shared/CrudPage';

const config: CrudPageConfig = {
    title: '安全管理人员类型',
    endpoint: '/safety-management-types',
    columns: [
        { key: 'title', header: '标题' },
        { key: 'sort', header: '排序', width: 80 },
    ],
    formFields: [
        { key: 'title', label: '标题', type: 'text', required: true },
        { key: 'sort', label: '排序', type: 'number' },
    ],
    searchField: 'title',
    searchPlaceholder: '搜索安全管理人员类型...',
    permissions: {
        list: 'base-data.view',
        create: 'base-data.manage',
        update: 'base-data.manage',
        delete: 'base-data.manage',
    },
};

export default function SafetyManagementTypesIndex() {
    return <CrudPage config={config} />;
}
