import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\BlocNoticesController::index
* @see app/Http/Controllers/Api/V1/BlocNoticesController.php:11
* @route '/api/v1/bloc-notices'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/bloc-notices',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\BlocNoticesController::index
* @see app/Http/Controllers/Api/V1/BlocNoticesController.php:11
* @route '/api/v1/bloc-notices'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\BlocNoticesController::index
* @see app/Http/Controllers/Api/V1/BlocNoticesController.php:11
* @route '/api/v1/bloc-notices'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocNoticesController::index
* @see app/Http/Controllers/Api/V1/BlocNoticesController.php:11
* @route '/api/v1/bloc-notices'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocNoticesController::index
* @see app/Http/Controllers/Api/V1/BlocNoticesController.php:11
* @route '/api/v1/bloc-notices'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocNoticesController::index
* @see app/Http/Controllers/Api/V1/BlocNoticesController.php:11
* @route '/api/v1/bloc-notices'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocNoticesController::index
* @see app/Http/Controllers/Api/V1/BlocNoticesController.php:11
* @route '/api/v1/bloc-notices'
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

const BlocNoticesController = { index }

export default BlocNoticesController