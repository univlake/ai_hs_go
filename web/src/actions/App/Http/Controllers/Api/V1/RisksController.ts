import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\RisksController::portrait
* @see app/Http/Controllers/Api/V1/RisksController.php:18
* @route '/api/v1/risks/portrait'
*/
export const portrait = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: portrait.url(options),
    method: 'get',
})

portrait.definition = {
    methods: ["get","head"],
    url: '/api/v1/risks/portrait',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\RisksController::portrait
* @see app/Http/Controllers/Api/V1/RisksController.php:18
* @route '/api/v1/risks/portrait'
*/
portrait.url = (options?: RouteQueryOptions) => {
    return portrait.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\RisksController::portrait
* @see app/Http/Controllers/Api/V1/RisksController.php:18
* @route '/api/v1/risks/portrait'
*/
portrait.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: portrait.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\RisksController::portrait
* @see app/Http/Controllers/Api/V1/RisksController.php:18
* @route '/api/v1/risks/portrait'
*/
portrait.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: portrait.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\RisksController::portrait
* @see app/Http/Controllers/Api/V1/RisksController.php:18
* @route '/api/v1/risks/portrait'
*/
const portraitForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: portrait.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\RisksController::portrait
* @see app/Http/Controllers/Api/V1/RisksController.php:18
* @route '/api/v1/risks/portrait'
*/
portraitForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: portrait.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\RisksController::portrait
* @see app/Http/Controllers/Api/V1/RisksController.php:18
* @route '/api/v1/risks/portrait'
*/
portraitForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: portrait.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

portrait.form = portraitForm

/**
* @see \App\Http\Controllers\Api\V1\RisksController::riskLog
* @see app/Http/Controllers/Api/V1/RisksController.php:55
* @route '/api/v1/risks/log'
*/
export const riskLog = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: riskLog.url(options),
    method: 'get',
})

riskLog.definition = {
    methods: ["get","head"],
    url: '/api/v1/risks/log',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\RisksController::riskLog
* @see app/Http/Controllers/Api/V1/RisksController.php:55
* @route '/api/v1/risks/log'
*/
riskLog.url = (options?: RouteQueryOptions) => {
    return riskLog.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\RisksController::riskLog
* @see app/Http/Controllers/Api/V1/RisksController.php:55
* @route '/api/v1/risks/log'
*/
riskLog.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: riskLog.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\RisksController::riskLog
* @see app/Http/Controllers/Api/V1/RisksController.php:55
* @route '/api/v1/risks/log'
*/
riskLog.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: riskLog.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\RisksController::riskLog
* @see app/Http/Controllers/Api/V1/RisksController.php:55
* @route '/api/v1/risks/log'
*/
const riskLogForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: riskLog.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\RisksController::riskLog
* @see app/Http/Controllers/Api/V1/RisksController.php:55
* @route '/api/v1/risks/log'
*/
riskLogForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: riskLog.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\RisksController::riskLog
* @see app/Http/Controllers/Api/V1/RisksController.php:55
* @route '/api/v1/risks/log'
*/
riskLogForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: riskLog.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

riskLog.form = riskLogForm

const RisksController = { portrait, riskLog }

export default RisksController