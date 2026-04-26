import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::index
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:14
* @route '/api/v1/fire-facilities'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/fire-facilities',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::index
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:14
* @route '/api/v1/fire-facilities'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::index
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:14
* @route '/api/v1/fire-facilities'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::index
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:14
* @route '/api/v1/fire-facilities'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::index
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:14
* @route '/api/v1/fire-facilities'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::index
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:14
* @route '/api/v1/fire-facilities'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::index
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:14
* @route '/api/v1/fire-facilities'
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
* @see \App\Http\Controllers\Api\V1\FireFacilityController::show
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:36
* @route '/api/v1/fire-facilities/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/fire-facilities/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::show
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:36
* @route '/api/v1/fire-facilities/{id}'
*/
show.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return show.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::show
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:36
* @route '/api/v1/fire-facilities/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::show
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:36
* @route '/api/v1/fire-facilities/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::show
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:36
* @route '/api/v1/fire-facilities/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::show
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:36
* @route '/api/v1/fire-facilities/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::show
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:36
* @route '/api/v1/fire-facilities/{id}'
*/
showForm.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
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
* @see \App\Http\Controllers\Api\V1\FireFacilityController::store
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:47
* @route '/api/v1/fire-facilities'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/fire-facilities',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::store
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:47
* @route '/api/v1/fire-facilities'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::store
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:47
* @route '/api/v1/fire-facilities'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::store
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:47
* @route '/api/v1/fire-facilities'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::store
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:47
* @route '/api/v1/fire-facilities'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::copy
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:85
* @route '/api/v1/fire-facilities/copy'
*/
export const copy = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: copy.url(options),
    method: 'post',
})

copy.definition = {
    methods: ["post"],
    url: '/api/v1/fire-facilities/copy',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::copy
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:85
* @route '/api/v1/fire-facilities/copy'
*/
copy.url = (options?: RouteQueryOptions) => {
    return copy.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::copy
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:85
* @route '/api/v1/fire-facilities/copy'
*/
copy.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: copy.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::copy
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:85
* @route '/api/v1/fire-facilities/copy'
*/
const copyForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: copy.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::copy
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:85
* @route '/api/v1/fire-facilities/copy'
*/
copyForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: copy.url(options),
    method: 'post',
})

copy.form = copyForm

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::update
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:58
* @route '/api/v1/fire-facilities/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/fire-facilities/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::update
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:58
* @route '/api/v1/fire-facilities/{id}'
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
* @see \App\Http\Controllers\Api\V1\FireFacilityController::update
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:58
* @route '/api/v1/fire-facilities/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::update
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:58
* @route '/api/v1/fire-facilities/{id}'
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
* @see \App\Http\Controllers\Api\V1\FireFacilityController::update
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:58
* @route '/api/v1/fire-facilities/{id}'
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
* @see \App\Http\Controllers\Api\V1\FireFacilityController::destroy
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:72
* @route '/api/v1/fire-facilities/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/fire-facilities/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::destroy
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:72
* @route '/api/v1/fire-facilities/{id}'
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
* @see \App\Http\Controllers\Api\V1\FireFacilityController::destroy
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:72
* @route '/api/v1/fire-facilities/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\FireFacilityController::destroy
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:72
* @route '/api/v1/fire-facilities/{id}'
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
* @see \App\Http\Controllers\Api\V1\FireFacilityController::destroy
* @see app/Http/Controllers/Api/V1/FireFacilityController.php:72
* @route '/api/v1/fire-facilities/{id}'
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

const FireFacilityController = { index, show, store, copy, update, destroy }

export default FireFacilityController