import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\BlocAccidentBulletinsController::index
* @see app/Http/Controllers/Api/V1/BlocAccidentBulletinsController.php:17
* @route '/api/v1/bloc-accident-bulletins'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/bloc-accident-bulletins',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\BlocAccidentBulletinsController::index
* @see app/Http/Controllers/Api/V1/BlocAccidentBulletinsController.php:17
* @route '/api/v1/bloc-accident-bulletins'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\BlocAccidentBulletinsController::index
* @see app/Http/Controllers/Api/V1/BlocAccidentBulletinsController.php:17
* @route '/api/v1/bloc-accident-bulletins'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocAccidentBulletinsController::index
* @see app/Http/Controllers/Api/V1/BlocAccidentBulletinsController.php:17
* @route '/api/v1/bloc-accident-bulletins'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocAccidentBulletinsController::index
* @see app/Http/Controllers/Api/V1/BlocAccidentBulletinsController.php:17
* @route '/api/v1/bloc-accident-bulletins'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocAccidentBulletinsController::index
* @see app/Http/Controllers/Api/V1/BlocAccidentBulletinsController.php:17
* @route '/api/v1/bloc-accident-bulletins'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocAccidentBulletinsController::index
* @see app/Http/Controllers/Api/V1/BlocAccidentBulletinsController.php:17
* @route '/api/v1/bloc-accident-bulletins'
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

/**
* @see \App\Http\Controllers\Api\V1\BlocAccidentBulletinsController::annualStats
* @see app/Http/Controllers/Api/V1/BlocAccidentBulletinsController.php:42
* @route '/api/v1/bloc-accident-bulletins/annual-stats'
*/
export const annualStats = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: annualStats.url(options),
    method: 'get',
})

annualStats.definition = {
    methods: ["get","head"],
    url: '/api/v1/bloc-accident-bulletins/annual-stats',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\BlocAccidentBulletinsController::annualStats
* @see app/Http/Controllers/Api/V1/BlocAccidentBulletinsController.php:42
* @route '/api/v1/bloc-accident-bulletins/annual-stats'
*/
annualStats.url = (options?: RouteQueryOptions) => {
    return annualStats.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\BlocAccidentBulletinsController::annualStats
* @see app/Http/Controllers/Api/V1/BlocAccidentBulletinsController.php:42
* @route '/api/v1/bloc-accident-bulletins/annual-stats'
*/
annualStats.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: annualStats.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocAccidentBulletinsController::annualStats
* @see app/Http/Controllers/Api/V1/BlocAccidentBulletinsController.php:42
* @route '/api/v1/bloc-accident-bulletins/annual-stats'
*/
annualStats.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: annualStats.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocAccidentBulletinsController::annualStats
* @see app/Http/Controllers/Api/V1/BlocAccidentBulletinsController.php:42
* @route '/api/v1/bloc-accident-bulletins/annual-stats'
*/
const annualStatsForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: annualStats.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocAccidentBulletinsController::annualStats
* @see app/Http/Controllers/Api/V1/BlocAccidentBulletinsController.php:42
* @route '/api/v1/bloc-accident-bulletins/annual-stats'
*/
annualStatsForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: annualStats.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocAccidentBulletinsController::annualStats
* @see app/Http/Controllers/Api/V1/BlocAccidentBulletinsController.php:42
* @route '/api/v1/bloc-accident-bulletins/annual-stats'
*/
annualStatsForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: annualStats.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

annualStats.form = annualStatsForm

const BlocAccidentBulletinsController = { index, annualStats }

export default BlocAccidentBulletinsController