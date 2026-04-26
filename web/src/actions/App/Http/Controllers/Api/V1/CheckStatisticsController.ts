import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\CheckStatisticsController::alerts
* @see app/Http/Controllers/Api/V1/CheckStatisticsController.php:95
* @route '/api/v1/inspection-alerts'
*/
export const alerts = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: alerts.url(options),
    method: 'get',
})

alerts.definition = {
    methods: ["get","head"],
    url: '/api/v1/inspection-alerts',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\CheckStatisticsController::alerts
* @see app/Http/Controllers/Api/V1/CheckStatisticsController.php:95
* @route '/api/v1/inspection-alerts'
*/
alerts.url = (options?: RouteQueryOptions) => {
    return alerts.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CheckStatisticsController::alerts
* @see app/Http/Controllers/Api/V1/CheckStatisticsController.php:95
* @route '/api/v1/inspection-alerts'
*/
alerts.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: alerts.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CheckStatisticsController::alerts
* @see app/Http/Controllers/Api/V1/CheckStatisticsController.php:95
* @route '/api/v1/inspection-alerts'
*/
alerts.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: alerts.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\CheckStatisticsController::alerts
* @see app/Http/Controllers/Api/V1/CheckStatisticsController.php:95
* @route '/api/v1/inspection-alerts'
*/
const alertsForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: alerts.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CheckStatisticsController::alerts
* @see app/Http/Controllers/Api/V1/CheckStatisticsController.php:95
* @route '/api/v1/inspection-alerts'
*/
alertsForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: alerts.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CheckStatisticsController::alerts
* @see app/Http/Controllers/Api/V1/CheckStatisticsController.php:95
* @route '/api/v1/inspection-alerts'
*/
alertsForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: alerts.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

alerts.form = alertsForm

const CheckStatisticsController = { alerts }

export default CheckStatisticsController