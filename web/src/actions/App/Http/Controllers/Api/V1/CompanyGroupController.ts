import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::index
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:18
* @route '/api/v1/company-groups'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/company-groups',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::index
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:18
* @route '/api/v1/company-groups'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::index
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:18
* @route '/api/v1/company-groups'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::index
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:18
* @route '/api/v1/company-groups'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::index
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:18
* @route '/api/v1/company-groups'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::index
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:18
* @route '/api/v1/company-groups'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::index
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:18
* @route '/api/v1/company-groups'
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
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::store
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:34
* @route '/api/v1/company-groups'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/company-groups',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::store
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:34
* @route '/api/v1/company-groups'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::store
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:34
* @route '/api/v1/company-groups'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::store
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:34
* @route '/api/v1/company-groups'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::store
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:34
* @route '/api/v1/company-groups'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::update
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:49
* @route '/api/v1/company-groups/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/company-groups/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::update
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:49
* @route '/api/v1/company-groups/{id}'
*/
update.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return update.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::update
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:49
* @route '/api/v1/company-groups/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::update
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:49
* @route '/api/v1/company-groups/{id}'
*/
const updateForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: update.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'PUT',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::update
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:49
* @route '/api/v1/company-groups/{id}'
*/
updateForm.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
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
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::syncKids
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:85
* @route '/api/v1/company-groups/{id}/sync-kids'
*/
export const syncKids = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: syncKids.url(args, options),
    method: 'post',
})

syncKids.definition = {
    methods: ["post"],
    url: '/api/v1/company-groups/{id}/sync-kids',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::syncKids
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:85
* @route '/api/v1/company-groups/{id}/sync-kids'
*/
syncKids.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return syncKids.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::syncKids
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:85
* @route '/api/v1/company-groups/{id}/sync-kids'
*/
syncKids.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: syncKids.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::syncKids
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:85
* @route '/api/v1/company-groups/{id}/sync-kids'
*/
const syncKidsForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: syncKids.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::syncKids
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:85
* @route '/api/v1/company-groups/{id}/sync-kids'
*/
syncKidsForm.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: syncKids.url(args, options),
    method: 'post',
})

syncKids.form = syncKidsForm

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::destroy
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:68
* @route '/api/v1/company-groups/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/company-groups/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::destroy
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:68
* @route '/api/v1/company-groups/{id}'
*/
destroy.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return destroy.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::destroy
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:68
* @route '/api/v1/company-groups/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::destroy
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:68
* @route '/api/v1/company-groups/{id}'
*/
const destroyForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: destroy.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'DELETE',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyGroupController::destroy
* @see app/Http/Controllers/Api/V1/CompanyGroupController.php:68
* @route '/api/v1/company-groups/{id}'
*/
destroyForm.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: destroy.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'DELETE',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

destroy.form = destroyForm

const CompanyGroupController = { index, store, update, syncKids, destroy }

export default CompanyGroupController