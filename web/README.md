# ai_hs_go web (SPA 前端)

从源 Laravel 项目 (`/Users/zero/Herd/aifix_hs/resources/js`) 抽离出来的 React 19 + Tailwind v4 前端，已切换为**纯 SPA**（去 Inertia.js）。

## 启动

```bash
cp .env.example .env
npm install         # or pnpm/yarn
npm run dev         # http://localhost:5173 — vite proxy 把 /api/* 转给 Go :8080
```

## 生产部署

```bash
npm run build       # 输出到 dist/
# Go server 会自动托管 dist 目录（见 internal/api/spa.go）
```

## 目录差异（vs Laravel resources/js）

- 原 `app.tsx`（Inertia 入口）→ 已重命名为 `_app.legacy.tsx`，仅作参考。新入口 `src/main.tsx` + `src/routes/index.tsx`。
- 新增 `src/lib/api-routes.ts`：集中管理 URL，**逐步替代** `@/actions/*` 与 `@/routes/*`（wayfinder 生成代码）。
- `src/lib/api.ts` 改用 `import.meta.env.VITE_API_BASE_URL`。

## 未完成的迁移工作（Phase 8 后续）

172 个页面文件已原样复制，但每个页面内的以下调用仍指向 Inertia/Wayfinder，需在使用对应模块时**逐页替换**：

| 调用 | 替换为 |
|---|---|
| `import { Link } from '@inertiajs/react'` | `import { Link } from 'react-router-dom'`，`href` → `to` |
| `router.visit(...)` / `router.replace(...)` | `useNavigate()` |
| `useForm()` from `@inertiajs/react` | `useForm` from `react-hook-form` 或自定义 hook |
| `import { ... } from '@/actions/...'` 或 `@/routes/...` | 改用 `@/lib/api-routes.ts` 中的常量 + axios |
| Inertia `<Head>`、`usePage()`、`useLayoutProps()` | 移除；标题用 `document.title`，状态用 React Context |

可以在 PR 中按模块批量重构（先 dashboard / auth / safety-ledger 等 MVP 路径），其余在后续迁移阶段处理。

## 主要依赖

| 包 | 作用 |
|---|---|
| react / react-dom 19 | 框架 |
| react-router-dom 7 | SPA 路由 |
| @tailwindcss/vite + tailwindcss 4 | 样式 |
| @tanstack/react-table | 数据表 |
| axios | HTTP |
| sonner | toast |
| jsencrypt | 登录密码 RSA 加密 |
| leaflet + react-leaflet + @geoman-io/leaflet-geoman-free | 地图 |

## 注意

- 主题切换在 `ThemeContext`，localStorage key `app-theme`，支持 `light / dark / auto`（auto 在 6/18 点切换）。
- 多租户切换在 `ScopeContext`，localStorage key `scope_state`。
- 平台超管登录后视野是「多个一级企业」，需要 `ScopeContext` 暴露 `availableTenants` + 顶部租户切换器（**详见 Phase 11.3 计划，尚待 UI 实现**）。
