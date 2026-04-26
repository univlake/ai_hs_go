import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStatisticsController::index
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStatisticsController.php:19
* @route '/api/v1/special-equipment-statistics'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/special-equipment-statistics',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStatisticsController::index
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStatisticsController.php:19
* @route '/api/v1/special-equipment-statistics'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStatisticsController::index
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStatisticsController.php:19
* @route '/api/v1/special-equipment-statistics'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStatisticsController::index
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStatisticsController.php:19
* @route '/api/v1/special-equipment-statistics'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStatisticsController::index
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStatisticsController.php:19
* @route '/api/v1/special-equipment-statistics'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStatisticsController::index
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStatisticsController.php:19
* @route '/api/v1/special-equipment-statistics'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStatisticsController::index
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStatisticsController.php:19
* @route '/api/v1/special-equipment-statistics'
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

const SpecialEquipmentStatisticsController = { index }

export default SpecialEquipmentStatisticsController