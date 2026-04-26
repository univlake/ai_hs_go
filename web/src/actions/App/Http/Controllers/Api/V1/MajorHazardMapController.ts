import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\MajorHazardMapController::index
* @see app/Http/Controllers/Api/V1/MajorHazardMapController.php:15
* @route '/api/v1/major-hazard-map'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/major-hazard-map',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\MajorHazardMapController::index
* @see app/Http/Controllers/Api/V1/MajorHazardMapController.php:15
* @route '/api/v1/major-hazard-map'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\MajorHazardMapController::index
* @see app/Http/Controllers/Api/V1/MajorHazardMapController.php:15
* @route '/api/v1/major-hazard-map'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MajorHazardMapController::index
* @see app/Http/Controllers/Api/V1/MajorHazardMapController.php:15
* @route '/api/v1/major-hazard-map'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\MajorHazardMapController::index
* @see app/Http/Controllers/Api/V1/MajorHazardMapController.php:15
* @route '/api/v1/major-hazard-map'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MajorHazardMapController::index
* @see app/Http/Controllers/Api/V1/MajorHazardMapController.php:15
* @route '/api/v1/major-hazard-map'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MajorHazardMapController::index
* @see app/Http/Controllers/Api/V1/MajorHazardMapController.php:15
* @route '/api/v1/major-hazard-map'
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

const MajorHazardMapController = { index }

export default MajorHazardMapController