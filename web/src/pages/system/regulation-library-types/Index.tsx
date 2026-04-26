import { CrudPage } from '@/components/shared/CrudPage';
import type { CrudPageConfig } from '@/components/shared/CrudPage';

// 法规库类型是个多层级字典 — pid 指向父级，all_name 由后端自动维护
// （父级 all_name + " / " + 自身 title）。列表展示 all_name 直观看出
// 层级，新建/编辑时通过下拉选择父级（不选则为顶级）。
const config: CrudPageConfig = {
    title: '法规库类型',
    endpoint: '/regulation-library-types',
    columns: [
        { key: 'id', header: 'ID', width: 60 },
        { key: 'title', header: '标题', width: 180 },
        { key: 'all_name', header: '完整路径' },
        { key: 'sort', header: '排序', width: 80 },
    ],
    formFields: [
        { key: 'title', label: '标题', type: 'text', required: true },
        {
            key: 'pid',
            label: '上级',
            type: 'select',
            placeholder: '不选则为顶级',
            // optionsEndpoint loads the existing entries; the backend
            // returns rows with `all_name` so the dropdown shows full
            // paths ("安全生产法律 / 行政法规"). When the user keeps it
            // blank, the form posts pid=0 → backend treats as top-level.
            optionsEndpoint: '/regulation-library-types',
            optionsLabelKey: 'all_name',
            optionsValueKey: 'id',
        },
        { key: 'sort', label: '排序', type: 'number' },
    ],
    searchField: 'title',
    searchPlaceholder: '搜索法规库类型...',
    // The dictionary is global — no company_id filter on the listing.
    companyScoped: false,
    permissions: {
        list: 'base-data.view',
        create: 'base-data.manage',
        update: 'base-data.manage',
        delete: 'base-data.manage',
    },
};

export default function RegulationLibraryTypesIndex() {
    return <CrudPage config={config} />;
}
