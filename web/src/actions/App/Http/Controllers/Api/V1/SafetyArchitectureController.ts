import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\SafetyArchitectureController::index
* @see app/Http/Controllers/Api/V1/SafetyArchitectureController.php:17
* @route '/api/v1/safety-architectures'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/safety-architectures',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\SafetyArchitectureController::index
* @see app/Http/Controllers/Api/V1/SafetyArchitectureController.php:17
* @route '/api/v1/safety-architectures'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\SafetyArchitectureController::index
* @see app/Http/Controllers/Api/V1/SafetyArchitectureController.php:17
* @route '/api/v1/safety-architectures'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyArchitectureController::index
* @see app/Http/Controllers/Api/V1/SafetyArchitectureController.php:17
* @route '/api/v1/safety-architectures'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyArchitectureController::index
* @see app/Http/Controllers/Api/V1/SafetyArchitectureController.php:17
* @route '/api/v1/safety-architectures'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyArchitectureController::index
* @see app/Http/Controllers/Api/V1/SafetyArchitectureController.php:17
* @route '/api/v1/safety-architectures'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyArchitectureController::index
* @see app/Http/Controllers/Api/V1/SafetyArchitectureController.php:17
* @route '/api/v1/safety-architectures'
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
* @see \App\Http\Controllers\Api\V1\SafetyArchitectureController::store
* @see app/Http/Controllers/Api/V1/SafetyArchitectureController.php:33
* @route '/api/v1/safety-architectures'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/safety-architectures',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\SafetyArchitectureController::store
* @see app/Http/Controllers/Api/V1/SafetyArchitectureController.php:33
* @route '/api/v1/safety-architectures'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\SafetyArchitectureController::store
* @see app/Http/Controllers/Api/V1/SafetyArchitectureController.php:33
* @route '/api/v1/safety-architectures'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyArchitectureController::store
* @see app/Http/Controllers/Api/V1/SafetyArchitectureController.php:33
* @route '/api/v1/safety-architectures'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyArchitectureController::store
* @see app/Http/Controllers/Api/V1/SafetyArchitectureController.php:33
* @route '/api/v1/safety-architectures'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\SafetyArchitectureController::update
* @see app/Http/Controllers/Api/V1/SafetyArchitectureController.php:47
* @route '/api/v1/safety-architectures/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/safety-architectures/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\SafetyArchitectureController::update
* @see app/Http/Controllers/Api/V1/SafetyArchitectureController.php:47
* @route '/api/v1/safety-architectures/{id}'
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
* @see \App\Http\Controllers\Api\V1\SafetyArchitectureController::update
* @see app/Http/Controllers/Api/V1/SafetyArchitectureController.php:47
* @route '/api/v1/safety-architectures/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyArchitectureController::update
* @see app/Http/Controllers/Api/V1/SafetyArchitectureController.php:47
* @route '/api/v1/safety-architectures/{id}'
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
* @see \App\Http\Controllers\Api\V1\SafetyArchitectureController::update
* @see app/Http/Controllers/Api/V1/SafetyArchitectureController.php:47
* @route '/api/v1/safety-architectures/{id}'
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

const SafetyArchitectureController = { index, store, update }

export default SafetyArchitectureController