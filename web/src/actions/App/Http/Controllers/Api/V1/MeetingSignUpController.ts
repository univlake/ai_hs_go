import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::index
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:14
* @route '/api/v1/meeting-sign-ups'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/meeting-sign-ups',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::index
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:14
* @route '/api/v1/meeting-sign-ups'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::index
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:14
* @route '/api/v1/meeting-sign-ups'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::index
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:14
* @route '/api/v1/meeting-sign-ups'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::index
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:14
* @route '/api/v1/meeting-sign-ups'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::index
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:14
* @route '/api/v1/meeting-sign-ups'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::index
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:14
* @route '/api/v1/meeting-sign-ups'
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
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::show
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:32
* @route '/api/v1/meeting-sign-ups/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/meeting-sign-ups/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::show
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:32
* @route '/api/v1/meeting-sign-ups/{id}'
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
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::show
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:32
* @route '/api/v1/meeting-sign-ups/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::show
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:32
* @route '/api/v1/meeting-sign-ups/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::show
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:32
* @route '/api/v1/meeting-sign-ups/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::show
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:32
* @route '/api/v1/meeting-sign-ups/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::show
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:32
* @route '/api/v1/meeting-sign-ups/{id}'
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
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::store
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:43
* @route '/api/v1/meeting-sign-ups'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/meeting-sign-ups',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::store
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:43
* @route '/api/v1/meeting-sign-ups'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::store
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:43
* @route '/api/v1/meeting-sign-ups'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::store
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:43
* @route '/api/v1/meeting-sign-ups'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::store
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:43
* @route '/api/v1/meeting-sign-ups'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::update
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:54
* @route '/api/v1/meeting-sign-ups/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/meeting-sign-ups/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::update
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:54
* @route '/api/v1/meeting-sign-ups/{id}'
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
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::update
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:54
* @route '/api/v1/meeting-sign-ups/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::update
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:54
* @route '/api/v1/meeting-sign-ups/{id}'
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
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::update
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:54
* @route '/api/v1/meeting-sign-ups/{id}'
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
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::destroy
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:68
* @route '/api/v1/meeting-sign-ups/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/meeting-sign-ups/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::destroy
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:68
* @route '/api/v1/meeting-sign-ups/{id}'
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
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::destroy
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:68
* @route '/api/v1/meeting-sign-ups/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::destroy
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:68
* @route '/api/v1/meeting-sign-ups/{id}'
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
* @see \App\Http\Controllers\Api\V1\MeetingSignUpController::destroy
* @see app/Http/Controllers/Api/V1/MeetingSignUpController.php:68
* @route '/api/v1/meeting-sign-ups/{id}'
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

const MeetingSignUpController = { index, show, store, update, destroy }

export default MeetingSignUpController