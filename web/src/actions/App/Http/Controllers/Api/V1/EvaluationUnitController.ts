import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::index
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:17
* @route '/api/v1/evaluation-units'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/evaluation-units',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::index
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:17
* @route '/api/v1/evaluation-units'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::index
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:17
* @route '/api/v1/evaluation-units'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::index
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:17
* @route '/api/v1/evaluation-units'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::index
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:17
* @route '/api/v1/evaluation-units'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::index
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:17
* @route '/api/v1/evaluation-units'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::index
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:17
* @route '/api/v1/evaluation-units'
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
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::show
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:38
* @route '/api/v1/evaluation-units/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/evaluation-units/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::show
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:38
* @route '/api/v1/evaluation-units/{id}'
*/
show.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return show.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::show
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:38
* @route '/api/v1/evaluation-units/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::show
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:38
* @route '/api/v1/evaluation-units/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::show
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:38
* @route '/api/v1/evaluation-units/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::show
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:38
* @route '/api/v1/evaluation-units/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::show
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:38
* @route '/api/v1/evaluation-units/{id}'
*/
showForm.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

show.form = showForm

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::store
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:52
* @route '/api/v1/evaluation-units'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/evaluation-units',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::store
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:52
* @route '/api/v1/evaluation-units'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::store
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:52
* @route '/api/v1/evaluation-units'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::store
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:52
* @route '/api/v1/evaluation-units'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::store
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:52
* @route '/api/v1/evaluation-units'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::update
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:66
* @route '/api/v1/evaluation-units/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/evaluation-units/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::update
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:66
* @route '/api/v1/evaluation-units/{id}'
*/
update.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return update.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::update
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:66
* @route '/api/v1/evaluation-units/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::update
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:66
* @route '/api/v1/evaluation-units/{id}'
*/
const updateForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: update.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'PUT',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::update
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:66
* @route '/api/v1/evaluation-units/{id}'
*/
updateForm.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: update.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'PUT',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

update.form = updateForm

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::destroy
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:83
* @route '/api/v1/evaluation-units/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/evaluation-units/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::destroy
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:83
* @route '/api/v1/evaluation-units/{id}'
*/
destroy.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return destroy.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::destroy
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:83
* @route '/api/v1/evaluation-units/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::destroy
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:83
* @route '/api/v1/evaluation-units/{id}'
*/
const destroyForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: destroy.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'DELETE',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\EvaluationUnitController::destroy
* @see app/Http/Controllers/Api/V1/EvaluationUnitController.php:83
* @route '/api/v1/evaluation-units/{id}'
*/
destroyForm.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: destroy.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'DELETE',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

destroy.form = destroyForm

const EvaluationUnitController = { index, show, store, update, destroy }

export default EvaluationUnitController