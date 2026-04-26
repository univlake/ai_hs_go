import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\AssessmentSummaryController::summary
* @see app/Http/Controllers/Api/V1/AssessmentSummaryController.php:36
* @route '/api/v1/assessment-summary'
*/
export const summary = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: summary.url(options),
    method: 'get',
})

summary.definition = {
    methods: ["get","head"],
    url: '/api/v1/assessment-summary',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\AssessmentSummaryController::summary
* @see app/Http/Controllers/Api/V1/AssessmentSummaryController.php:36
* @route '/api/v1/assessment-summary'
*/
summary.url = (options?: RouteQueryOptions) => {
    return summary.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AssessmentSummaryController::summary
* @see app/Http/Controllers/Api/V1/AssessmentSummaryController.php:36
* @route '/api/v1/assessment-summary'
*/
summary.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: summary.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AssessmentSummaryController::summary
* @see app/Http/Controllers/Api/V1/AssessmentSummaryController.php:36
* @route '/api/v1/assessment-summary'
*/
summary.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: summary.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\AssessmentSummaryController::summary
* @see app/Http/Controllers/Api/V1/AssessmentSummaryController.php:36
* @route '/api/v1/assessment-summary'
*/
const summaryForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: summary.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AssessmentSummaryController::summary
* @see app/Http/Controllers/Api/V1/AssessmentSummaryController.php:36
* @route '/api/v1/assessment-summary'
*/
summaryForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: summary.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AssessmentSummaryController::summary
* @see app/Http/Controllers/Api/V1/AssessmentSummaryController.php:36
* @route '/api/v1/assessment-summary'
*/
summaryForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: summary.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

summary.form = summaryForm

/**
* @see \App\Http\Controllers\Api\V1\AssessmentSummaryController::exportMethod
* @see app/Http/Controllers/Api/V1/AssessmentSummaryController.php:65
* @route '/api/v1/assessment-summary/export'
*/
export const exportMethod = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: exportMethod.url(options),
    method: 'get',
})

exportMethod.definition = {
    methods: ["get","head"],
    url: '/api/v1/assessment-summary/export',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\AssessmentSummaryController::exportMethod
* @see app/Http/Controllers/Api/V1/AssessmentSummaryController.php:65
* @route '/api/v1/assessment-summary/export'
*/
exportMethod.url = (options?: RouteQueryOptions) => {
    return exportMethod.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AssessmentSummaryController::exportMethod
* @see app/Http/Controllers/Api/V1/AssessmentSummaryController.php:65
* @route '/api/v1/assessment-summary/export'
*/
exportMethod.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: exportMethod.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AssessmentSummaryController::exportMethod
* @see app/Http/Controllers/Api/V1/AssessmentSummaryController.php:65
* @route '/api/v1/assessment-summary/export'
*/
exportMethod.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: exportMethod.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\AssessmentSummaryController::exportMethod
* @see app/Http/Controllers/Api/V1/AssessmentSummaryController.php:65
* @route '/api/v1/assessment-summary/export'
*/
const exportMethodForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: exportMethod.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AssessmentSummaryController::exportMethod
* @see app/Http/Controllers/Api/V1/AssessmentSummaryController.php:65
* @route '/api/v1/assessment-summary/export'
*/
exportMethodForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: exportMethod.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AssessmentSummaryController::exportMethod
* @see app/Http/Controllers/Api/V1/AssessmentSummaryController.php:65
* @route '/api/v1/assessment-summary/export'
*/
exportMethodForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: exportMethod.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

exportMethod.form = exportMethodForm

const AssessmentSummaryController = { summary, exportMethod, export: exportMethod }

export default AssessmentSummaryController