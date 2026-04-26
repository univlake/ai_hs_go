import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getInherentRiskNum
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:20
* @route '/api/v1/bloc-risks-inherent'
*/
export const getInherentRiskNum = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: getInherentRiskNum.url(options),
    method: 'get',
})

getInherentRiskNum.definition = {
    methods: ["get","head"],
    url: '/api/v1/bloc-risks-inherent',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getInherentRiskNum
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:20
* @route '/api/v1/bloc-risks-inherent'
*/
getInherentRiskNum.url = (options?: RouteQueryOptions) => {
    return getInherentRiskNum.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getInherentRiskNum
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:20
* @route '/api/v1/bloc-risks-inherent'
*/
getInherentRiskNum.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: getInherentRiskNum.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getInherentRiskNum
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:20
* @route '/api/v1/bloc-risks-inherent'
*/
getInherentRiskNum.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: getInherentRiskNum.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getInherentRiskNum
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:20
* @route '/api/v1/bloc-risks-inherent'
*/
const getInherentRiskNumForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: getInherentRiskNum.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getInherentRiskNum
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:20
* @route '/api/v1/bloc-risks-inherent'
*/
getInherentRiskNumForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: getInherentRiskNum.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getInherentRiskNum
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:20
* @route '/api/v1/bloc-risks-inherent'
*/
getInherentRiskNumForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: getInherentRiskNum.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

getInherentRiskNum.form = getInherentRiskNumForm

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getFloatRiskNum
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:51
* @route '/api/v1/bloc-risks-float'
*/
export const getFloatRiskNum = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: getFloatRiskNum.url(options),
    method: 'get',
})

getFloatRiskNum.definition = {
    methods: ["get","head"],
    url: '/api/v1/bloc-risks-float',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getFloatRiskNum
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:51
* @route '/api/v1/bloc-risks-float'
*/
getFloatRiskNum.url = (options?: RouteQueryOptions) => {
    return getFloatRiskNum.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getFloatRiskNum
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:51
* @route '/api/v1/bloc-risks-float'
*/
getFloatRiskNum.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: getFloatRiskNum.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getFloatRiskNum
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:51
* @route '/api/v1/bloc-risks-float'
*/
getFloatRiskNum.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: getFloatRiskNum.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getFloatRiskNum
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:51
* @route '/api/v1/bloc-risks-float'
*/
const getFloatRiskNumForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: getFloatRiskNum.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getFloatRiskNum
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:51
* @route '/api/v1/bloc-risks-float'
*/
getFloatRiskNumForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: getFloatRiskNum.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getFloatRiskNum
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:51
* @route '/api/v1/bloc-risks-float'
*/
getFloatRiskNumForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: getFloatRiskNum.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

getFloatRiskNum.form = getFloatRiskNumForm

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::floatRiskStatistics
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:83
* @route '/api/v1/bloc-risks-float-statistics'
*/
export const floatRiskStatistics = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: floatRiskStatistics.url(options),
    method: 'get',
})

floatRiskStatistics.definition = {
    methods: ["get","head"],
    url: '/api/v1/bloc-risks-float-statistics',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::floatRiskStatistics
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:83
* @route '/api/v1/bloc-risks-float-statistics'
*/
floatRiskStatistics.url = (options?: RouteQueryOptions) => {
    return floatRiskStatistics.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::floatRiskStatistics
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:83
* @route '/api/v1/bloc-risks-float-statistics'
*/
floatRiskStatistics.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: floatRiskStatistics.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::floatRiskStatistics
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:83
* @route '/api/v1/bloc-risks-float-statistics'
*/
floatRiskStatistics.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: floatRiskStatistics.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::floatRiskStatistics
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:83
* @route '/api/v1/bloc-risks-float-statistics'
*/
const floatRiskStatisticsForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: floatRiskStatistics.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::floatRiskStatistics
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:83
* @route '/api/v1/bloc-risks-float-statistics'
*/
floatRiskStatisticsForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: floatRiskStatistics.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::floatRiskStatistics
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:83
* @route '/api/v1/bloc-risks-float-statistics'
*/
floatRiskStatisticsForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: floatRiskStatistics.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

floatRiskStatistics.form = floatRiskStatisticsForm

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::inherentRiskStatistics
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:98
* @route '/api/v1/bloc-risks-inherent-statistics'
*/
export const inherentRiskStatistics = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: inherentRiskStatistics.url(options),
    method: 'get',
})

inherentRiskStatistics.definition = {
    methods: ["get","head"],
    url: '/api/v1/bloc-risks-inherent-statistics',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::inherentRiskStatistics
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:98
* @route '/api/v1/bloc-risks-inherent-statistics'
*/
inherentRiskStatistics.url = (options?: RouteQueryOptions) => {
    return inherentRiskStatistics.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::inherentRiskStatistics
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:98
* @route '/api/v1/bloc-risks-inherent-statistics'
*/
inherentRiskStatistics.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: inherentRiskStatistics.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::inherentRiskStatistics
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:98
* @route '/api/v1/bloc-risks-inherent-statistics'
*/
inherentRiskStatistics.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: inherentRiskStatistics.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::inherentRiskStatistics
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:98
* @route '/api/v1/bloc-risks-inherent-statistics'
*/
const inherentRiskStatisticsForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: inherentRiskStatistics.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::inherentRiskStatistics
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:98
* @route '/api/v1/bloc-risks-inherent-statistics'
*/
inherentRiskStatisticsForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: inherentRiskStatistics.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::inherentRiskStatistics
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:98
* @route '/api/v1/bloc-risks-inherent-statistics'
*/
inherentRiskStatisticsForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: inherentRiskStatistics.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

inherentRiskStatistics.form = inherentRiskStatisticsForm

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getRiskLog
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:113
* @route '/api/v1/bloc-risks-log'
*/
export const getRiskLog = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: getRiskLog.url(options),
    method: 'get',
})

getRiskLog.definition = {
    methods: ["get","head"],
    url: '/api/v1/bloc-risks-log',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getRiskLog
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:113
* @route '/api/v1/bloc-risks-log'
*/
getRiskLog.url = (options?: RouteQueryOptions) => {
    return getRiskLog.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getRiskLog
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:113
* @route '/api/v1/bloc-risks-log'
*/
getRiskLog.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: getRiskLog.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getRiskLog
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:113
* @route '/api/v1/bloc-risks-log'
*/
getRiskLog.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: getRiskLog.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getRiskLog
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:113
* @route '/api/v1/bloc-risks-log'
*/
const getRiskLogForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: getRiskLog.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getRiskLog
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:113
* @route '/api/v1/bloc-risks-log'
*/
getRiskLogForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: getRiskLog.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::getRiskLog
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:113
* @route '/api/v1/bloc-risks-log'
*/
getRiskLogForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: getRiskLog.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

getRiskLog.form = getRiskLogForm

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::riskSort
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:193
* @route '/api/v1/bloc-risks-sort'
*/
export const riskSort = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: riskSort.url(options),
    method: 'get',
})

riskSort.definition = {
    methods: ["get","head"],
    url: '/api/v1/bloc-risks-sort',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::riskSort
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:193
* @route '/api/v1/bloc-risks-sort'
*/
riskSort.url = (options?: RouteQueryOptions) => {
    return riskSort.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::riskSort
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:193
* @route '/api/v1/bloc-risks-sort'
*/
riskSort.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: riskSort.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::riskSort
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:193
* @route '/api/v1/bloc-risks-sort'
*/
riskSort.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: riskSort.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::riskSort
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:193
* @route '/api/v1/bloc-risks-sort'
*/
const riskSortForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: riskSort.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::riskSort
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:193
* @route '/api/v1/bloc-risks-sort'
*/
riskSortForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: riskSort.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::riskSort
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:193
* @route '/api/v1/bloc-risks-sort'
*/
riskSortForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: riskSort.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

riskSort.form = riskSortForm

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::hiddenDangerSort
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:219
* @route '/api/v1/bloc-hidden-danger-sort'
*/
export const hiddenDangerSort = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: hiddenDangerSort.url(options),
    method: 'get',
})

hiddenDangerSort.definition = {
    methods: ["get","head"],
    url: '/api/v1/bloc-hidden-danger-sort',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::hiddenDangerSort
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:219
* @route '/api/v1/bloc-hidden-danger-sort'
*/
hiddenDangerSort.url = (options?: RouteQueryOptions) => {
    return hiddenDangerSort.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::hiddenDangerSort
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:219
* @route '/api/v1/bloc-hidden-danger-sort'
*/
hiddenDangerSort.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: hiddenDangerSort.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::hiddenDangerSort
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:219
* @route '/api/v1/bloc-hidden-danger-sort'
*/
hiddenDangerSort.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: hiddenDangerSort.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::hiddenDangerSort
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:219
* @route '/api/v1/bloc-hidden-danger-sort'
*/
const hiddenDangerSortForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: hiddenDangerSort.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::hiddenDangerSort
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:219
* @route '/api/v1/bloc-hidden-danger-sort'
*/
hiddenDangerSortForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: hiddenDangerSort.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocRisksController::hiddenDangerSort
* @see app/Http/Controllers/Api/V1/BlocRisksController.php:219
* @route '/api/v1/bloc-hidden-danger-sort'
*/
hiddenDangerSortForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: hiddenDangerSort.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

hiddenDangerSort.form = hiddenDangerSortForm

const BlocRisksController = { getInherentRiskNum, getFloatRiskNum, floatRiskStatistics, inherentRiskStatistics, getRiskLog, riskSort, hiddenDangerSort }

export default BlocRisksController