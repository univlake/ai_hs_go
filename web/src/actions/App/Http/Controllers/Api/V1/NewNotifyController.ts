import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::index
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:16
* @route '/api/v1/new-notifies'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/new-notifies',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::index
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:16
* @route '/api/v1/new-notifies'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::index
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:16
* @route '/api/v1/new-notifies'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::index
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:16
* @route '/api/v1/new-notifies'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::index
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:16
* @route '/api/v1/new-notifies'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::index
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:16
* @route '/api/v1/new-notifies'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::index
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:16
* @route '/api/v1/new-notifies'
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
* @see \App\Http\Controllers\Api\V1\NewNotifyController::show
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:47
* @route '/api/v1/new-notifies/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/new-notifies/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::show
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:47
* @route '/api/v1/new-notifies/{id}'
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
* @see \App\Http\Controllers\Api\V1\NewNotifyController::show
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:47
* @route '/api/v1/new-notifies/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::show
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:47
* @route '/api/v1/new-notifies/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::show
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:47
* @route '/api/v1/new-notifies/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::show
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:47
* @route '/api/v1/new-notifies/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::show
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:47
* @route '/api/v1/new-notifies/{id}'
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
* @see \App\Http\Controllers\Api\V1\NewNotifyController::store
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:58
* @route '/api/v1/new-notifies'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/new-notifies',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::store
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:58
* @route '/api/v1/new-notifies'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::store
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:58
* @route '/api/v1/new-notifies'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::store
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:58
* @route '/api/v1/new-notifies'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::store
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:58
* @route '/api/v1/new-notifies'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::update
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:91
* @route '/api/v1/new-notifies/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/new-notifies/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::update
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:91
* @route '/api/v1/new-notifies/{id}'
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
* @see \App\Http\Controllers\Api\V1\NewNotifyController::update
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:91
* @route '/api/v1/new-notifies/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::update
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:91
* @route '/api/v1/new-notifies/{id}'
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
* @see \App\Http\Controllers\Api\V1\NewNotifyController::update
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:91
* @route '/api/v1/new-notifies/{id}'
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
* @see \App\Http\Controllers\Api\V1\NewNotifyController::destroy
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:129
* @route '/api/v1/new-notifies/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/new-notifies/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::destroy
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:129
* @route '/api/v1/new-notifies/{id}'
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
* @see \App\Http\Controllers\Api\V1\NewNotifyController::destroy
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:129
* @route '/api/v1/new-notifies/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::destroy
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:129
* @route '/api/v1/new-notifies/{id}'
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
* @see \App\Http\Controllers\Api\V1\NewNotifyController::destroy
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:129
* @route '/api/v1/new-notifies/{id}'
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

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::markAsRead
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:168
* @route '/api/v1/new-notify-scopes/{id}/read'
*/
export const markAsRead = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: markAsRead.url(args, options),
    method: 'post',
})

markAsRead.definition = {
    methods: ["post"],
    url: '/api/v1/new-notify-scopes/{id}/read',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::markAsRead
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:168
* @route '/api/v1/new-notify-scopes/{id}/read'
*/
markAsRead.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return markAsRead.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::markAsRead
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:168
* @route '/api/v1/new-notify-scopes/{id}/read'
*/
markAsRead.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: markAsRead.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::markAsRead
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:168
* @route '/api/v1/new-notify-scopes/{id}/read'
*/
const markAsReadForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: markAsRead.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::markAsRead
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:168
* @route '/api/v1/new-notify-scopes/{id}/read'
*/
markAsReadForm.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: markAsRead.url(args, options),
    method: 'post',
})

markAsRead.form = markAsReadForm

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::pendingNotices
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:148
* @route '/api/v1/new-notify-notifications/pending'
*/
export const pendingNotices = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: pendingNotices.url(options),
    method: 'get',
})

pendingNotices.definition = {
    methods: ["get","head"],
    url: '/api/v1/new-notify-notifications/pending',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::pendingNotices
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:148
* @route '/api/v1/new-notify-notifications/pending'
*/
pendingNotices.url = (options?: RouteQueryOptions) => {
    return pendingNotices.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::pendingNotices
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:148
* @route '/api/v1/new-notify-notifications/pending'
*/
pendingNotices.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: pendingNotices.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::pendingNotices
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:148
* @route '/api/v1/new-notify-notifications/pending'
*/
pendingNotices.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: pendingNotices.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::pendingNotices
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:148
* @route '/api/v1/new-notify-notifications/pending'
*/
const pendingNoticesForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: pendingNotices.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::pendingNotices
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:148
* @route '/api/v1/new-notify-notifications/pending'
*/
pendingNoticesForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: pendingNotices.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NewNotifyController::pendingNotices
* @see app/Http/Controllers/Api/V1/NewNotifyController.php:148
* @route '/api/v1/new-notify-notifications/pending'
*/
pendingNoticesForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: pendingNotices.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

pendingNotices.form = pendingNoticesForm

const NewNotifyController = { index, show, store, update, destroy, markAsRead, pendingNotices }

export default NewNotifyController