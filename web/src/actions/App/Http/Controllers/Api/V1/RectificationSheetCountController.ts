import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetCountController::getStatistics
* @see app/Http/Controllers/Api/V1/RectificationSheetCountController.php:15
* @route '/api/v1/rectification-sheet-count'
*/
export const getStatistics = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: getStatistics.url(options),
    method: 'get',
})

getStatistics.definition = {
    methods: ["get","head"],
    url: '/api/v1/rectification-sheet-count',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetCountController::getStatistics
* @see app/Http/Controllers/Api/V1/RectificationSheetCountController.php:15
* @route '/api/v1/rectification-sheet-count'
*/
getStatistics.url = (options?: RouteQueryOptions) => {
    return getStatistics.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetCountController::getStatistics
* @see app/Http/Controllers/Api/V1/RectificationSheetCountController.php:15
* @route '/api/v1/rectification-sheet-count'
*/
getStatistics.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: getStatistics.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetCountController::getStatistics
* @see app/Http/Controllers/Api/V1/RectificationSheetCountController.php:15
* @route '/api/v1/rectification-sheet-count'
*/
getStatistics.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: getStatistics.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetCountController::getStatistics
* @see app/Http/Controllers/Api/V1/RectificationSheetCountController.php:15
* @route '/api/v1/rectification-sheet-count'
*/
const getStatisticsForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: getStatistics.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetCountController::getStatistics
* @see app/Http/Controllers/Api/V1/RectificationSheetCountController.php:15
* @route '/api/v1/rectification-sheet-count'
*/
getStatisticsForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: getStatistics.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetCountController::getStatistics
* @see app/Http/Controllers/Api/V1/RectificationSheetCountController.php:15
* @route '/api/v1/rectification-sheet-count'
*/
getStatisticsForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: getStatistics.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

getStatistics.form = getStatisticsForm

const RectificationSheetCountController = { getStatistics }

export default RectificationSheetCountController