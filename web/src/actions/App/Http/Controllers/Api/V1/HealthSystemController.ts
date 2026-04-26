import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\HealthSystemController::index
* @see app/Http/Controllers/Api/V1/HealthSystemController.php:14
* @route '/api/v1/health-systems'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/health-systems',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\HealthSystemController::index
* @see app/Http/Controllers/Api/V1/HealthSystemController.php:14
* @route '/api/v1/health-systems'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\HealthSystemController::index
* @see app/Http/Controllers/Api/V1/HealthSystemController.php:14
* @route '/api/v1/health-systems'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\HealthSystemController::index
* @see app/Http/Controllers/Api/V1/HealthSystemController.php:14
* @route '/api/v1/health-systems'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\HealthSystemController::index
* @see app/Http/Controllers/Api/V1/HealthSystemController.php:14
* @route '/api/v1/health-systems'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\HealthSystemController::index
* @see app/Http/Controllers/Api/V1/HealthSystemController.php:14
* @route '/api/v1/health-systems'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\HealthSystemController::index
* @see app/Http/Controllers/Api/V1/HealthSystemController.php:14
* @route '/api/v1/health-systems'
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
* @see \App\Http\Controllers\Api\V1\HealthSystemController::store
* @see app/Http/Controllers/Api/V1/HealthSystemController.php:27
* @route '/api/v1/health-systems'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/health-systems',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\HealthSystemController::store
* @see app/Http/Controllers/Api/V1/HealthSystemController.php:27
* @route '/api/v1/health-systems'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\HealthSystemController::store
* @see app/Http/Controllers/Api/V1/HealthSystemController.php:27
* @route '/api/v1/health-systems'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\HealthSystemController::store
* @see app/Http/Controllers/Api/V1/HealthSystemController.php:27
* @route '/api/v1/health-systems'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\HealthSystemController::store
* @see app/Http/Controllers/Api/V1/HealthSystemController.php:27
* @route '/api/v1/health-systems'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\HealthSystemController::update
* @see app/Http/Controllers/Api/V1/HealthSystemController.php:38
* @route '/api/v1/health-systems/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/health-systems/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\HealthSystemController::update
* @see app/Http/Controllers/Api/V1/HealthSystemController.php:38
* @route '/api/v1/health-systems/{id}'
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
* @see \App\Http\Controllers\Api\V1\HealthSystemController::update
* @see app/Http/Controllers/Api/V1/HealthSystemController.php:38
* @route '/api/v1/health-systems/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\HealthSystemController::update
* @see app/Http/Controllers/Api/V1/HealthSystemController.php:38
* @route '/api/v1/health-systems/{id}'
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
* @see \App\Http\Controllers\Api\V1\HealthSystemController::update
* @see app/Http/Controllers/Api/V1/HealthSystemController.php:38
* @route '/api/v1/health-systems/{id}'
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

const HealthSystemController = { index, store, update }

export default HealthSystemController