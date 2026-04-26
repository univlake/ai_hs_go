import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\HomesController::earlyWarning
* @see app/Http/Controllers/Api/V1/HomesController.php:12
* @route '/api/v1/homes-early-warning'
*/
export const earlyWarning = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: earlyWarning.url(options),
    method: 'get',
})

earlyWarning.definition = {
    methods: ["get","head"],
    url: '/api/v1/homes-early-warning',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\HomesController::earlyWarning
* @see app/Http/Controllers/Api/V1/HomesController.php:12
* @route '/api/v1/homes-early-warning'
*/
earlyWarning.url = (options?: RouteQueryOptions) => {
    return earlyWarning.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\HomesController::earlyWarning
* @see app/Http/Controllers/Api/V1/HomesController.php:12
* @route '/api/v1/homes-early-warning'
*/
earlyWarning.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: earlyWarning.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\HomesController::earlyWarning
* @see app/Http/Controllers/Api/V1/HomesController.php:12
* @route '/api/v1/homes-early-warning'
*/
earlyWarning.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: earlyWarning.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\HomesController::earlyWarning
* @see app/Http/Controllers/Api/V1/HomesController.php:12
* @route '/api/v1/homes-early-warning'
*/
const earlyWarningForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: earlyWarning.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\HomesController::earlyWarning
* @see app/Http/Controllers/Api/V1/HomesController.php:12
* @route '/api/v1/homes-early-warning'
*/
earlyWarningForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: earlyWarning.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\HomesController::earlyWarning
* @see app/Http/Controllers/Api/V1/HomesController.php:12
* @route '/api/v1/homes-early-warning'
*/
earlyWarningForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: earlyWarning.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

earlyWarning.form = earlyWarningForm

/**
* @see \App\Http\Controllers\Api\V1\HomesController::companyInfo
* @see app/Http/Controllers/Api/V1/HomesController.php:20
* @route '/api/v1/homes-company-info'
*/
export const companyInfo = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: companyInfo.url(options),
    method: 'get',
})

companyInfo.definition = {
    methods: ["get","head"],
    url: '/api/v1/homes-company-info',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\HomesController::companyInfo
* @see app/Http/Controllers/Api/V1/HomesController.php:20
* @route '/api/v1/homes-company-info'
*/
companyInfo.url = (options?: RouteQueryOptions) => {
    return companyInfo.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\HomesController::companyInfo
* @see app/Http/Controllers/Api/V1/HomesController.php:20
* @route '/api/v1/homes-company-info'
*/
companyInfo.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: companyInfo.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\HomesController::companyInfo
* @see app/Http/Controllers/Api/V1/HomesController.php:20
* @route '/api/v1/homes-company-info'
*/
companyInfo.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: companyInfo.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\HomesController::companyInfo
* @see app/Http/Controllers/Api/V1/HomesController.php:20
* @route '/api/v1/homes-company-info'
*/
const companyInfoForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: companyInfo.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\HomesController::companyInfo
* @see app/Http/Controllers/Api/V1/HomesController.php:20
* @route '/api/v1/homes-company-info'
*/
companyInfoForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: companyInfo.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\HomesController::companyInfo
* @see app/Http/Controllers/Api/V1/HomesController.php:20
* @route '/api/v1/homes-company-info'
*/
companyInfoForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: companyInfo.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

companyInfo.form = companyInfoForm

const HomesController = { earlyWarning, companyInfo }

export default HomesController