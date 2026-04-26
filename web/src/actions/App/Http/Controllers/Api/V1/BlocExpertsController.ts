import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::index
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:14
* @route '/api/v1/bloc-experts'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/bloc-experts',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::index
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:14
* @route '/api/v1/bloc-experts'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::index
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:14
* @route '/api/v1/bloc-experts'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::index
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:14
* @route '/api/v1/bloc-experts'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::index
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:14
* @route '/api/v1/bloc-experts'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::index
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:14
* @route '/api/v1/bloc-experts'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::index
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:14
* @route '/api/v1/bloc-experts'
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
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::show
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:28
* @route '/api/v1/bloc-experts/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/bloc-experts/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::show
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:28
* @route '/api/v1/bloc-experts/{id}'
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
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::show
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:28
* @route '/api/v1/bloc-experts/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::show
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:28
* @route '/api/v1/bloc-experts/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::show
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:28
* @route '/api/v1/bloc-experts/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::show
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:28
* @route '/api/v1/bloc-experts/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::show
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:28
* @route '/api/v1/bloc-experts/{id}'
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
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::list
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:63
* @route '/api/v1/bloc-experts-list'
*/
export const list = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: list.url(options),
    method: 'get',
})

list.definition = {
    methods: ["get","head"],
    url: '/api/v1/bloc-experts-list',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::list
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:63
* @route '/api/v1/bloc-experts-list'
*/
list.url = (options?: RouteQueryOptions) => {
    return list.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::list
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:63
* @route '/api/v1/bloc-experts-list'
*/
list.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: list.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::list
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:63
* @route '/api/v1/bloc-experts-list'
*/
list.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: list.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::list
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:63
* @route '/api/v1/bloc-experts-list'
*/
const listForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: list.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::list
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:63
* @route '/api/v1/bloc-experts-list'
*/
listForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: list.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::list
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:63
* @route '/api/v1/bloc-experts-list'
*/
listForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: list.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

list.form = listForm

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::store
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:39
* @route '/api/v1/bloc-experts'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/bloc-experts',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::store
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:39
* @route '/api/v1/bloc-experts'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::store
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:39
* @route '/api/v1/bloc-experts'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::store
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:39
* @route '/api/v1/bloc-experts'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::store
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:39
* @route '/api/v1/bloc-experts'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::destroy
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:50
* @route '/api/v1/bloc-experts/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/bloc-experts/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::destroy
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:50
* @route '/api/v1/bloc-experts/{id}'
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
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::destroy
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:50
* @route '/api/v1/bloc-experts/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::destroy
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:50
* @route '/api/v1/bloc-experts/{id}'
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
* @see \App\Http\Controllers\Api\V1\BlocExpertsController::destroy
* @see app/Http/Controllers/Api/V1/BlocExpertsController.php:50
* @route '/api/v1/bloc-experts/{id}'
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

const BlocExpertsController = { index, show, list, store, destroy }

export default BlocExpertsController