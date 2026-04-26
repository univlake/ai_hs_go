import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\CompanyController::publicList
* @see app/Http/Controllers/Api/V1/CompanyController.php:27
* @route '/api/v1/companies-public'
*/
export const publicList = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: publicList.url(options),
    method: 'get',
})

publicList.definition = {
    methods: ["get","head"],
    url: '/api/v1/companies-public',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::publicList
* @see app/Http/Controllers/Api/V1/CompanyController.php:27
* @route '/api/v1/companies-public'
*/
publicList.url = (options?: RouteQueryOptions) => {
    return publicList.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::publicList
* @see app/Http/Controllers/Api/V1/CompanyController.php:27
* @route '/api/v1/companies-public'
*/
publicList.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: publicList.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::publicList
* @see app/Http/Controllers/Api/V1/CompanyController.php:27
* @route '/api/v1/companies-public'
*/
publicList.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: publicList.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::publicList
* @see app/Http/Controllers/Api/V1/CompanyController.php:27
* @route '/api/v1/companies-public'
*/
const publicListForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: publicList.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::publicList
* @see app/Http/Controllers/Api/V1/CompanyController.php:27
* @route '/api/v1/companies-public'
*/
publicListForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: publicList.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::publicList
* @see app/Http/Controllers/Api/V1/CompanyController.php:27
* @route '/api/v1/companies-public'
*/
publicListForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: publicList.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

publicList.form = publicListForm

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::index
* @see app/Http/Controllers/Api/V1/CompanyController.php:40
* @route '/api/v1/companies'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/companies',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::index
* @see app/Http/Controllers/Api/V1/CompanyController.php:40
* @route '/api/v1/companies'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::index
* @see app/Http/Controllers/Api/V1/CompanyController.php:40
* @route '/api/v1/companies'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::index
* @see app/Http/Controllers/Api/V1/CompanyController.php:40
* @route '/api/v1/companies'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::index
* @see app/Http/Controllers/Api/V1/CompanyController.php:40
* @route '/api/v1/companies'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::index
* @see app/Http/Controllers/Api/V1/CompanyController.php:40
* @route '/api/v1/companies'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::index
* @see app/Http/Controllers/Api/V1/CompanyController.php:40
* @route '/api/v1/companies'
*/
indexForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

index.form = indexForm

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::show
* @see app/Http/Controllers/Api/V1/CompanyController.php:76
* @route '/api/v1/companies/{company}'
*/
export const show = (args: { company: string | number } | [company: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/companies/{company}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::show
* @see app/Http/Controllers/Api/V1/CompanyController.php:76
* @route '/api/v1/companies/{company}'
*/
show.url = (args: { company: string | number } | [company: string | number ] | string | number, options?: RouteQueryOptions) => {
    if (typeof args === 'string' || typeof args === 'number') {
        args = { company: args }
    }

    if (Array.isArray(args)) {
        args = {
            company: args[0],
        }
    }

    args = applyUrlDefaults(args)

    const parsedArgs = {
        company: args.company,
    }

    return show.definition.url
            .replace('{company}', parsedArgs.company.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::show
* @see app/Http/Controllers/Api/V1/CompanyController.php:76
* @route '/api/v1/companies/{company}'
*/
show.get = (args: { company: string | number } | [company: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::show
* @see app/Http/Controllers/Api/V1/CompanyController.php:76
* @route '/api/v1/companies/{company}'
*/
show.head = (args: { company: string | number } | [company: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::show
* @see app/Http/Controllers/Api/V1/CompanyController.php:76
* @route '/api/v1/companies/{company}'
*/
const showForm = (args: { company: string | number } | [company: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::show
* @see app/Http/Controllers/Api/V1/CompanyController.php:76
* @route '/api/v1/companies/{company}'
*/
showForm.get = (args: { company: string | number } | [company: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::show
* @see app/Http/Controllers/Api/V1/CompanyController.php:76
* @route '/api/v1/companies/{company}'
*/
showForm.head = (args: { company: string | number } | [company: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

show.form = showForm

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::tree
* @see app/Http/Controllers/Api/V1/CompanyController.php:181
* @route '/api/v1/companies-tree'
*/
export const tree = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: tree.url(options),
    method: 'get',
})

tree.definition = {
    methods: ["get","head"],
    url: '/api/v1/companies-tree',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::tree
* @see app/Http/Controllers/Api/V1/CompanyController.php:181
* @route '/api/v1/companies-tree'
*/
tree.url = (options?: RouteQueryOptions) => {
    return tree.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::tree
* @see app/Http/Controllers/Api/V1/CompanyController.php:181
* @route '/api/v1/companies-tree'
*/
tree.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: tree.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::tree
* @see app/Http/Controllers/Api/V1/CompanyController.php:181
* @route '/api/v1/companies-tree'
*/
tree.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: tree.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::tree
* @see app/Http/Controllers/Api/V1/CompanyController.php:181
* @route '/api/v1/companies-tree'
*/
const treeForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: tree.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::tree
* @see app/Http/Controllers/Api/V1/CompanyController.php:181
* @route '/api/v1/companies-tree'
*/
treeForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: tree.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::tree
* @see app/Http/Controllers/Api/V1/CompanyController.php:181
* @route '/api/v1/companies-tree'
*/
treeForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: tree.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

tree.form = treeForm

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::store
* @see app/Http/Controllers/Api/V1/CompanyController.php:95
* @route '/api/v1/companies'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/companies',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::store
* @see app/Http/Controllers/Api/V1/CompanyController.php:95
* @route '/api/v1/companies'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::store
* @see app/Http/Controllers/Api/V1/CompanyController.php:95
* @route '/api/v1/companies'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::store
* @see app/Http/Controllers/Api/V1/CompanyController.php:95
* @route '/api/v1/companies'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::store
* @see app/Http/Controllers/Api/V1/CompanyController.php:95
* @route '/api/v1/companies'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::update
* @see app/Http/Controllers/Api/V1/CompanyController.php:132
* @route '/api/v1/companies/{company}'
*/
export const update = (args: { company: string | number } | [company: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/companies/{company}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::update
* @see app/Http/Controllers/Api/V1/CompanyController.php:132
* @route '/api/v1/companies/{company}'
*/
update.url = (args: { company: string | number } | [company: string | number ] | string | number, options?: RouteQueryOptions) => {
    if (typeof args === 'string' || typeof args === 'number') {
        args = { company: args }
    }

    if (Array.isArray(args)) {
        args = {
            company: args[0],
        }
    }

    args = applyUrlDefaults(args)

    const parsedArgs = {
        company: args.company,
    }

    return update.definition.url
            .replace('{company}', parsedArgs.company.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::update
* @see app/Http/Controllers/Api/V1/CompanyController.php:132
* @route '/api/v1/companies/{company}'
*/
update.put = (args: { company: string | number } | [company: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::update
* @see app/Http/Controllers/Api/V1/CompanyController.php:132
* @route '/api/v1/companies/{company}'
*/
const updateForm = (args: { company: string | number } | [company: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: update.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'PUT',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::update
* @see app/Http/Controllers/Api/V1/CompanyController.php:132
* @route '/api/v1/companies/{company}'
*/
updateForm.put = (args: { company: string | number } | [company: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: update.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'PUT',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

update.form = updateForm

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::toggleStatus
* @see app/Http/Controllers/Api/V1/CompanyController.php:257
* @route '/api/v1/companies/{id}/toggle-status'
*/
export const toggleStatus = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'patch'> => ({
    url: toggleStatus.url(args, options),
    method: 'patch',
})

toggleStatus.definition = {
    methods: ["patch"],
    url: '/api/v1/companies/{id}/toggle-status',
} satisfies RouteDefinition<["patch"]>

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::toggleStatus
* @see app/Http/Controllers/Api/V1/CompanyController.php:257
* @route '/api/v1/companies/{id}/toggle-status'
*/
toggleStatus.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
    if (typeof args === 'string' || typeof args === 'number') {
        args = { id: args }
    }

    if (Array.isArray(args)) {
        args = {
            id: args[0],
        }
    }

    args = applyUrlDefaults(args)

    const parsedArgs = {
        id: args.id,
    }

    return toggleStatus.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::toggleStatus
* @see app/Http/Controllers/Api/V1/CompanyController.php:257
* @route '/api/v1/companies/{id}/toggle-status'
*/
toggleStatus.patch = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'patch'> => ({
    url: toggleStatus.url(args, options),
    method: 'patch',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::toggleStatus
* @see app/Http/Controllers/Api/V1/CompanyController.php:257
* @route '/api/v1/companies/{id}/toggle-status'
*/
const toggleStatusForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: toggleStatus.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'PATCH',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::toggleStatus
* @see app/Http/Controllers/Api/V1/CompanyController.php:257
* @route '/api/v1/companies/{id}/toggle-status'
*/
toggleStatusForm.patch = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: toggleStatus.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'PATCH',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

toggleStatus.form = toggleStatusForm

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::destroy
* @see app/Http/Controllers/Api/V1/CompanyController.php:158
* @route '/api/v1/companies/{company}'
*/
export const destroy = (args: { company: string | number } | [company: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/companies/{company}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::destroy
* @see app/Http/Controllers/Api/V1/CompanyController.php:158
* @route '/api/v1/companies/{company}'
*/
destroy.url = (args: { company: string | number } | [company: string | number ] | string | number, options?: RouteQueryOptions) => {
    if (typeof args === 'string' || typeof args === 'number') {
        args = { company: args }
    }

    if (Array.isArray(args)) {
        args = {
            company: args[0],
        }
    }

    args = applyUrlDefaults(args)

    const parsedArgs = {
        company: args.company,
    }

    return destroy.definition.url
            .replace('{company}', parsedArgs.company.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::destroy
* @see app/Http/Controllers/Api/V1/CompanyController.php:158
* @route '/api/v1/companies/{company}'
*/
destroy.delete = (args: { company: string | number } | [company: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::destroy
* @see app/Http/Controllers/Api/V1/CompanyController.php:158
* @route '/api/v1/companies/{company}'
*/
const destroyForm = (args: { company: string | number } | [company: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: destroy.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'DELETE',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::destroy
* @see app/Http/Controllers/Api/V1/CompanyController.php:158
* @route '/api/v1/companies/{company}'
*/
destroyForm.delete = (args: { company: string | number } | [company: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: destroy.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'DELETE',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

destroy.form = destroyForm

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::reviewList
* @see app/Http/Controllers/Api/V1/CompanyController.php:203
* @route '/api/v1/companies-review-list'
*/
export const reviewList = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: reviewList.url(options),
    method: 'get',
})

reviewList.definition = {
    methods: ["get","head"],
    url: '/api/v1/companies-review-list',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::reviewList
* @see app/Http/Controllers/Api/V1/CompanyController.php:203
* @route '/api/v1/companies-review-list'
*/
reviewList.url = (options?: RouteQueryOptions) => {
    return reviewList.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::reviewList
* @see app/Http/Controllers/Api/V1/CompanyController.php:203
* @route '/api/v1/companies-review-list'
*/
reviewList.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: reviewList.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::reviewList
* @see app/Http/Controllers/Api/V1/CompanyController.php:203
* @route '/api/v1/companies-review-list'
*/
reviewList.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: reviewList.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::reviewList
* @see app/Http/Controllers/Api/V1/CompanyController.php:203
* @route '/api/v1/companies-review-list'
*/
const reviewListForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: reviewList.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::reviewList
* @see app/Http/Controllers/Api/V1/CompanyController.php:203
* @route '/api/v1/companies-review-list'
*/
reviewListForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: reviewList.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::reviewList
* @see app/Http/Controllers/Api/V1/CompanyController.php:203
* @route '/api/v1/companies-review-list'
*/
reviewListForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: reviewList.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

reviewList.form = reviewListForm

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::review
* @see app/Http/Controllers/Api/V1/CompanyController.php:234
* @route '/api/v1/companies/{id}/review'
*/
export const review = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: review.url(args, options),
    method: 'post',
})

review.definition = {
    methods: ["post"],
    url: '/api/v1/companies/{id}/review',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::review
* @see app/Http/Controllers/Api/V1/CompanyController.php:234
* @route '/api/v1/companies/{id}/review'
*/
review.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
    if (typeof args === 'string' || typeof args === 'number') {
        args = { id: args }
    }

    if (Array.isArray(args)) {
        args = {
            id: args[0],
        }
    }

    args = applyUrlDefaults(args)

    const parsedArgs = {
        id: args.id,
    }

    return review.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::review
* @see app/Http/Controllers/Api/V1/CompanyController.php:234
* @route '/api/v1/companies/{id}/review'
*/
review.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: review.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::review
* @see app/Http/Controllers/Api/V1/CompanyController.php:234
* @route '/api/v1/companies/{id}/review'
*/
const reviewForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: review.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyController::review
* @see app/Http/Controllers/Api/V1/CompanyController.php:234
* @route '/api/v1/companies/{id}/review'
*/
reviewForm.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: review.url(args, options),
    method: 'post',
})

review.form = reviewForm

const CompanyController = { publicList, index, show, tree, store, update, toggleStatus, destroy, reviewList, review }

export default CompanyController