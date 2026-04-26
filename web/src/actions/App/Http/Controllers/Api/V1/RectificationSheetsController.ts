import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::generate
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:39
* @route '/api/v1/expert-inspection-companies/{id}/rectification-sheet'
*/
export const generate = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: generate.url(args, options),
    method: 'post',
})

generate.definition = {
    methods: ["post"],
    url: '/api/v1/expert-inspection-companies/{id}/rectification-sheet',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::generate
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:39
* @route '/api/v1/expert-inspection-companies/{id}/rectification-sheet'
*/
generate.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
    if (typeof args === 'string' || typeof args === 'number') {
        args = { id: args }
    }

    if (Array.isArray(args)) {
        args = {
            id: args[0],
        }
    }

    args = applyUrlDefaults(args)

    const parsedArgs = {
        id: args.id,
    }

    return generate.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::generate
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:39
* @route '/api/v1/expert-inspection-companies/{id}/rectification-sheet'
*/
generate.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: generate.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::generate
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:39
* @route '/api/v1/expert-inspection-companies/{id}/rectification-sheet'
*/
const generateForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: generate.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::generate
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:39
* @route '/api/v1/expert-inspection-companies/{id}/rectification-sheet'
*/
generateForm.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: generate.url(args, options),
    method: 'post',
})

generate.form = generateForm

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::download
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:73
* @route '/api/v1/expert-inspection-companies/{id}/rectification-sheet'
*/
export const download = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: download.url(args, options),
    method: 'get',
})

download.definition = {
    methods: ["get","head"],
    url: '/api/v1/expert-inspection-companies/{id}/rectification-sheet',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::download
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:73
* @route '/api/v1/expert-inspection-companies/{id}/rectification-sheet'
*/
download.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
    if (typeof args === 'string' || typeof args === 'number') {
        args = { id: args }
    }

    if (Array.isArray(args)) {
        args = {
            id: args[0],
        }
    }

    args = applyUrlDefaults(args)

    const parsedArgs = {
        id: args.id,
    }

    return download.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::download
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:73
* @route '/api/v1/expert-inspection-companies/{id}/rectification-sheet'
*/
download.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: download.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::download
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:73
* @route '/api/v1/expert-inspection-companies/{id}/rectification-sheet'
*/
download.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: download.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::download
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:73
* @route '/api/v1/expert-inspection-companies/{id}/rectification-sheet'
*/
const downloadForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: download.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::download
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:73
* @route '/api/v1/expert-inspection-companies/{id}/rectification-sheet'
*/
downloadForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: download.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::download
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:73
* @route '/api/v1/expert-inspection-companies/{id}/rectification-sheet'
*/
downloadForm.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: download.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

download.form = downloadForm

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::exportStatistics
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:111
* @route '/api/v1/hazard-statistics'
*/
export const exportStatistics = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: exportStatistics.url(options),
    method: 'get',
})

exportStatistics.definition = {
    methods: ["get","head"],
    url: '/api/v1/hazard-statistics',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::exportStatistics
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:111
* @route '/api/v1/hazard-statistics'
*/
exportStatistics.url = (options?: RouteQueryOptions) => {
    return exportStatistics.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::exportStatistics
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:111
* @route '/api/v1/hazard-statistics'
*/
exportStatistics.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: exportStatistics.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::exportStatistics
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:111
* @route '/api/v1/hazard-statistics'
*/
exportStatistics.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: exportStatistics.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::exportStatistics
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:111
* @route '/api/v1/hazard-statistics'
*/
const exportStatisticsForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: exportStatistics.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::exportStatistics
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:111
* @route '/api/v1/hazard-statistics'
*/
exportStatisticsForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: exportStatistics.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\RectificationSheetsController::exportStatistics
* @see app/Http/Controllers/Api/V1/RectificationSheetsController.php:111
* @route '/api/v1/hazard-statistics'
*/
exportStatisticsForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: exportStatistics.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

exportStatistics.form = exportStatisticsForm

const RectificationSheetsController = { generate, download, exportStatistics }

export default RectificationSheetsController