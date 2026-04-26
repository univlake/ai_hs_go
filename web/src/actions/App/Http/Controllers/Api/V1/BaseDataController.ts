import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\BaseDataController::baseAreas
* @see app/Http/Controllers/Api/V1/BaseDataController.php:43
* @route '/api/v1/base-areas'
*/
export const baseAreas = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: baseAreas.url(options),
    method: 'get',
})

baseAreas.definition = {
    methods: ["get","head"],
    url: '/api/v1/base-areas',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\BaseDataController::baseAreas
* @see app/Http/Controllers/Api/V1/BaseDataController.php:43
* @route '/api/v1/base-areas'
*/
baseAreas.url = (options?: RouteQueryOptions) => {
    return baseAreas.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\BaseDataController::baseAreas
* @see app/Http/Controllers/Api/V1/BaseDataController.php:43
* @route '/api/v1/base-areas'
*/
baseAreas.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: baseAreas.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BaseDataController::baseAreas
* @see app/Http/Controllers/Api/V1/BaseDataController.php:43
* @route '/api/v1/base-areas'
*/
baseAreas.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: baseAreas.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\BaseDataController::baseAreas
* @see app/Http/Controllers/Api/V1/BaseDataController.php:43
* @route '/api/v1/base-areas'
*/
const baseAreasForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: baseAreas.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BaseDataController::baseAreas
* @see app/Http/Controllers/Api/V1/BaseDataController.php:43
* @route '/api/v1/base-areas'
*/
baseAreasForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: baseAreas.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BaseDataController::baseAreas
* @see app/Http/Controllers/Api/V1/BaseDataController.php:43
* @route '/api/v1/base-areas'
*/
baseAreasForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: baseAreas.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

baseAreas.form = baseAreasForm

const BaseDataController = { baseAreas }

export default BaseDataController