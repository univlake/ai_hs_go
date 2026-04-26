import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::index
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:14
* @route '/api/v1/emergency-drills'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/emergency-drills',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::index
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:14
* @route '/api/v1/emergency-drills'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::index
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:14
* @route '/api/v1/emergency-drills'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::index
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:14
* @route '/api/v1/emergency-drills'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::index
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:14
* @route '/api/v1/emergency-drills'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::index
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:14
* @route '/api/v1/emergency-drills'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::index
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:14
* @route '/api/v1/emergency-drills'
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
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::show
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:32
* @route '/api/v1/emergency-drills/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/emergency-drills/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::show
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:32
* @route '/api/v1/emergency-drills/{id}'
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
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::show
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:32
* @route '/api/v1/emergency-drills/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::show
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:32
* @route '/api/v1/emergency-drills/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::show
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:32
* @route '/api/v1/emergency-drills/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::show
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:32
* @route '/api/v1/emergency-drills/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::show
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:32
* @route '/api/v1/emergency-drills/{id}'
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
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::store
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:43
* @route '/api/v1/emergency-drills'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/emergency-drills',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::store
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:43
* @route '/api/v1/emergency-drills'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::store
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:43
* @route '/api/v1/emergency-drills'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::store
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:43
* @route '/api/v1/emergency-drills'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::store
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:43
* @route '/api/v1/emergency-drills'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::update
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:54
* @route '/api/v1/emergency-drills/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/emergency-drills/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::update
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:54
* @route '/api/v1/emergency-drills/{id}'
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
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::update
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:54
* @route '/api/v1/emergency-drills/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::update
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:54
* @route '/api/v1/emergency-drills/{id}'
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
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::update
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:54
* @route '/api/v1/emergency-drills/{id}'
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
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::destroy
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:68
* @route '/api/v1/emergency-drills/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/emergency-drills/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::destroy
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:68
* @route '/api/v1/emergency-drills/{id}'
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
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::destroy
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:68
* @route '/api/v1/emergency-drills/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::destroy
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:68
* @route '/api/v1/emergency-drills/{id}'
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
* @see \App\Http\Controllers\Api\V1\EmergencyDrillController::destroy
* @see app/Http/Controllers/Api/V1/EmergencyDrillController.php:68
* @route '/api/v1/emergency-drills/{id}'
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

const EmergencyDrillController = { index, show, store, update, destroy }

export default EmergencyDrillController