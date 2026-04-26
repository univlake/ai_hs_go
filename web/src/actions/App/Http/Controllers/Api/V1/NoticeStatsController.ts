import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\NoticeStatsController::summary
* @see app/Http/Controllers/Api/V1/NoticeStatsController.php:17
* @route '/api/v1/notice-stats/summary'
*/
export const summary = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: summary.url(options),
    method: 'get',
})

summary.definition = {
    methods: ["get","head"],
    url: '/api/v1/notice-stats/summary',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\NoticeStatsController::summary
* @see app/Http/Controllers/Api/V1/NoticeStatsController.php:17
* @route '/api/v1/notice-stats/summary'
*/
summary.url = (options?: RouteQueryOptions) => {
    return summary.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\NoticeStatsController::summary
* @see app/Http/Controllers/Api/V1/NoticeStatsController.php:17
* @route '/api/v1/notice-stats/summary'
*/
summary.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: summary.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NoticeStatsController::summary
* @see app/Http/Controllers/Api/V1/NoticeStatsController.php:17
* @route '/api/v1/notice-stats/summary'
*/
summary.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: summary.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\NoticeStatsController::summary
* @see app/Http/Controllers/Api/V1/NoticeStatsController.php:17
* @route '/api/v1/notice-stats/summary'
*/
const summaryForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: summary.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NoticeStatsController::summary
* @see app/Http/Controllers/Api/V1/NoticeStatsController.php:17
* @route '/api/v1/notice-stats/summary'
*/
summaryForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: summary.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NoticeStatsController::summary
* @see app/Http/Controllers/Api/V1/NoticeStatsController.php:17
* @route '/api/v1/notice-stats/summary'
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
* @see \App\Http\Controllers\Api\V1\NoticeStatsController::pending
* @see app/Http/Controllers/Api/V1/NoticeStatsController.php:75
* @route '/api/v1/notice-stats/pending'
*/
export const pending = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: pending.url(options),
    method: 'get',
})

pending.definition = {
    methods: ["get","head"],
    url: '/api/v1/notice-stats/pending',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\NoticeStatsController::pending
* @see app/Http/Controllers/Api/V1/NoticeStatsController.php:75
* @route '/api/v1/notice-stats/pending'
*/
pending.url = (options?: RouteQueryOptions) => {
    return pending.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\NoticeStatsController::pending
* @see app/Http/Controllers/Api/V1/NoticeStatsController.php:75
* @route '/api/v1/notice-stats/pending'
*/
pending.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: pending.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NoticeStatsController::pending
* @see app/Http/Controllers/Api/V1/NoticeStatsController.php:75
* @route '/api/v1/notice-stats/pending'
*/
pending.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: pending.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\NoticeStatsController::pending
* @see app/Http/Controllers/Api/V1/NoticeStatsController.php:75
* @route '/api/v1/notice-stats/pending'
*/
const pendingForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: pending.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NoticeStatsController::pending
* @see app/Http/Controllers/Api/V1/NoticeStatsController.php:75
* @route '/api/v1/notice-stats/pending'
*/
pendingForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: pending.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NoticeStatsController::pending
* @see app/Http/Controllers/Api/V1/NoticeStatsController.php:75
* @route '/api/v1/notice-stats/pending'
*/
pendingForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: pending.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

pending.form = pendingForm

const NoticeStatsController = { summary, pending }

export default NoticeStatsController