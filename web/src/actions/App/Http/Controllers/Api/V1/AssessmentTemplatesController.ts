import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\AssessmentTemplatesController::years
* @see app/Http/Controllers/Api/V1/AssessmentTemplatesController.php:32
* @route '/api/v1/assessment-templates/years'
*/
export const years = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: years.url(options),
    method: 'get',
})

years.definition = {
    methods: ["get","head"],
    url: '/api/v1/assessment-templates/years',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\AssessmentTemplatesController::years
* @see app/Http/Controllers/Api/V1/AssessmentTemplatesController.php:32
* @route '/api/v1/assessment-templates/years'
*/
years.url = (options?: RouteQueryOptions) => {
    return years.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AssessmentTemplatesController::years
* @see app/Http/Controllers/Api/V1/AssessmentTemplatesController.php:32
* @route '/api/v1/assessment-templates/years'
*/
years.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: years.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AssessmentTemplatesController::years
* @see app/Http/Controllers/Api/V1/AssessmentTemplatesController.php:32
* @route '/api/v1/assessment-templates/years'
*/
years.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: years.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\AssessmentTemplatesController::years
* @see app/Http/Controllers/Api/V1/AssessmentTemplatesController.php:32
* @route '/api/v1/assessment-templates/years'
*/
const yearsForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: years.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AssessmentTemplatesController::years
* @see app/Http/Controllers/Api/V1/AssessmentTemplatesController.php:32
* @route '/api/v1/assessment-templates/years'
*/
yearsForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: years.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AssessmentTemplatesController::years
* @see app/Http/Controllers/Api/V1/AssessmentTemplatesController.php:32
* @route '/api/v1/assessment-templates/years'
*/
yearsForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: years.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

years.form = yearsForm

/**
* @see \App\Http\Controllers\Api\V1\AssessmentTemplatesController::copy
* @see app/Http/Controllers/Api/V1/AssessmentTemplatesController.php:58
* @route '/api/v1/assessment-templates/copy'
*/
export const copy = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: copy.url(options),
    method: 'post',
})

copy.definition = {
    methods: ["post"],
    url: '/api/v1/assessment-templates/copy',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\AssessmentTemplatesController::copy
* @see app/Http/Controllers/Api/V1/AssessmentTemplatesController.php:58
* @route '/api/v1/assessment-templates/copy'
*/
copy.url = (options?: RouteQueryOptions) => {
    return copy.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AssessmentTemplatesController::copy
* @see app/Http/Controllers/Api/V1/AssessmentTemplatesController.php:58
* @route '/api/v1/assessment-templates/copy'
*/
copy.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: copy.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\AssessmentTemplatesController::copy
* @see app/Http/Controllers/Api/V1/AssessmentTemplatesController.php:58
* @route '/api/v1/assessment-templates/copy'
*/
const copyForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: copy.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\AssessmentTemplatesController::copy
* @see app/Http/Controllers/Api/V1/AssessmentTemplatesController.php:58
* @route '/api/v1/assessment-templates/copy'
*/
copyForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: copy.url(options),
    method: 'post',
})

copy.form = copyForm

const AssessmentTemplatesController = { years, copy }

export default AssessmentTemplatesController