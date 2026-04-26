import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\EnterpriseReportController::index
* @see app/Http/Controllers/Api/V1/EnterpriseReportController.php:11
* @route '/api/v1/enterprise-report'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/enterprise-report',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\EnterpriseReportController::index
* @see app/Http/Controllers/Api/V1/EnterpriseReportController.php:11
* @route '/api/v1/enterprise-report'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\EnterpriseReportController::index
* @see app/Http/Controllers/Api/V1/EnterpriseReportController.php:11
* @route '/api/v1/enterprise-report'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\EnterpriseReportController::index
* @see app/Http/Controllers/Api/V1/EnterpriseReportController.php:11
* @route '/api/v1/enterprise-report'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\EnterpriseReportController::index
* @see app/Http/Controllers/Api/V1/EnterpriseReportController.php:11
* @route '/api/v1/enterprise-report'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\EnterpriseReportController::index
* @see app/Http/Controllers/Api/V1/EnterpriseReportController.php:11
* @route '/api/v1/enterprise-report'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\EnterpriseReportController::index
* @see app/Http/Controllers/Api/V1/EnterpriseReportController.php:11
* @route '/api/v1/enterprise-report'
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

const EnterpriseReportController = { index }

export default EnterpriseReportController