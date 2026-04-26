import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\CompanyMapController::index
* @see app/Http/Controllers/Api/V1/CompanyMapController.php:12
* @route '/api/v1/company-map'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/company-map',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\CompanyMapController::index
* @see app/Http/Controllers/Api/V1/CompanyMapController.php:12
* @route '/api/v1/company-map'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CompanyMapController::index
* @see app/Http/Controllers/Api/V1/CompanyMapController.php:12
* @route '/api/v1/company-map'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyMapController::index
* @see app/Http/Controllers/Api/V1/CompanyMapController.php:12
* @route '/api/v1/company-map'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyMapController::index
* @see app/Http/Controllers/Api/V1/CompanyMapController.php:12
* @route '/api/v1/company-map'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyMapController::index
* @see app/Http/Controllers/Api/V1/CompanyMapController.php:12
* @route '/api/v1/company-map'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CompanyMapController::index
* @see app/Http/Controllers/Api/V1/CompanyMapController.php:12
* @route '/api/v1/company-map'
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

const CompanyMapController = { index }

export default CompanyMapController