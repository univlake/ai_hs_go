import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\BlocMonthlyInvestigationsController::index
* @see app/Http/Controllers/Api/V1/BlocMonthlyInvestigationsController.php:16
* @route '/api/v1/bloc-monthly-investigations'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/bloc-monthly-investigations',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\BlocMonthlyInvestigationsController::index
* @see app/Http/Controllers/Api/V1/BlocMonthlyInvestigationsController.php:16
* @route '/api/v1/bloc-monthly-investigations'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\BlocMonthlyInvestigationsController::index
* @see app/Http/Controllers/Api/V1/BlocMonthlyInvestigationsController.php:16
* @route '/api/v1/bloc-monthly-investigations'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocMonthlyInvestigationsController::index
* @see app/Http/Controllers/Api/V1/BlocMonthlyInvestigationsController.php:16
* @route '/api/v1/bloc-monthly-investigations'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocMonthlyInvestigationsController::index
* @see app/Http/Controllers/Api/V1/BlocMonthlyInvestigationsController.php:16
* @route '/api/v1/bloc-monthly-investigations'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocMonthlyInvestigationsController::index
* @see app/Http/Controllers/Api/V1/BlocMonthlyInvestigationsController.php:16
* @route '/api/v1/bloc-monthly-investigations'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocMonthlyInvestigationsController::index
* @see app/Http/Controllers/Api/V1/BlocMonthlyInvestigationsController.php:16
* @route '/api/v1/bloc-monthly-investigations'
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

const BlocMonthlyInvestigationsController = { index }

export default BlocMonthlyInvestigationsController