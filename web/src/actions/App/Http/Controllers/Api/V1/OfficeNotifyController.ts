import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::index
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:16
* @route '/api/v1/office-notifies'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/office-notifies',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::index
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:16
* @route '/api/v1/office-notifies'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::index
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:16
* @route '/api/v1/office-notifies'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::index
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:16
* @route '/api/v1/office-notifies'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::index
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:16
* @route '/api/v1/office-notifies'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::index
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:16
* @route '/api/v1/office-notifies'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::index
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:16
* @route '/api/v1/office-notifies'
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
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::show
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:47
* @route '/api/v1/office-notifies/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/office-notifies/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::show
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:47
* @route '/api/v1/office-notifies/{id}'
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
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::show
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:47
* @route '/api/v1/office-notifies/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::show
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:47
* @route '/api/v1/office-notifies/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::show
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:47
* @route '/api/v1/office-notifies/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::show
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:47
* @route '/api/v1/office-notifies/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::show
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:47
* @route '/api/v1/office-notifies/{id}'
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
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::store
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:58
* @route '/api/v1/office-notifies'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/office-notifies',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::store
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:58
* @route '/api/v1/office-notifies'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::store
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:58
* @route '/api/v1/office-notifies'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::store
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:58
* @route '/api/v1/office-notifies'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::store
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:58
* @route '/api/v1/office-notifies'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::update
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:91
* @route '/api/v1/office-notifies/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/office-notifies/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::update
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:91
* @route '/api/v1/office-notifies/{id}'
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
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::update
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:91
* @route '/api/v1/office-notifies/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::update
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:91
* @route '/api/v1/office-notifies/{id}'
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
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::update
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:91
* @route '/api/v1/office-notifies/{id}'
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
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::destroy
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:129
* @route '/api/v1/office-notifies/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/office-notifies/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::destroy
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:129
* @route '/api/v1/office-notifies/{id}'
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
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::destroy
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:129
* @route '/api/v1/office-notifies/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::destroy
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:129
* @route '/api/v1/office-notifies/{id}'
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
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::destroy
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:129
* @route '/api/v1/office-notifies/{id}'
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
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::submitFeedback
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:168
* @route '/api/v1/office-notify-scopes/{id}/feedback'
*/
export const submitFeedback = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: submitFeedback.url(args, options),
    method: 'post',
})

submitFeedback.definition = {
    methods: ["post"],
    url: '/api/v1/office-notify-scopes/{id}/feedback',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::submitFeedback
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:168
* @route '/api/v1/office-notify-scopes/{id}/feedback'
*/
submitFeedback.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return submitFeedback.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::submitFeedback
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:168
* @route '/api/v1/office-notify-scopes/{id}/feedback'
*/
submitFeedback.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: submitFeedback.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::submitFeedback
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:168
* @route '/api/v1/office-notify-scopes/{id}/feedback'
*/
const submitFeedbackForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: submitFeedback.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::submitFeedback
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:168
* @route '/api/v1/office-notify-scopes/{id}/feedback'
*/
submitFeedbackForm.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: submitFeedback.url(args, options),
    method: 'post',
})

submitFeedback.form = submitFeedbackForm

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::pendingNotices
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:148
* @route '/api/v1/office-notify-notifications/pending'
*/
export const pendingNotices = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: pendingNotices.url(options),
    method: 'get',
})

pendingNotices.definition = {
    methods: ["get","head"],
    url: '/api/v1/office-notify-notifications/pending',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::pendingNotices
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:148
* @route '/api/v1/office-notify-notifications/pending'
*/
pendingNotices.url = (options?: RouteQueryOptions) => {
    return pendingNotices.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::pendingNotices
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:148
* @route '/api/v1/office-notify-notifications/pending'
*/
pendingNotices.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: pendingNotices.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::pendingNotices
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:148
* @route '/api/v1/office-notify-notifications/pending'
*/
pendingNotices.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: pendingNotices.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::pendingNotices
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:148
* @route '/api/v1/office-notify-notifications/pending'
*/
const pendingNoticesForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: pendingNotices.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::pendingNotices
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:148
* @route '/api/v1/office-notify-notifications/pending'
*/
pendingNoticesForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: pendingNotices.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OfficeNotifyController::pendingNotices
* @see app/Http/Controllers/Api/V1/OfficeNotifyController.php:148
* @route '/api/v1/office-notify-notifications/pending'
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

const OfficeNotifyController = { index, show, store, update, destroy, submitFeedback, pendingNotices }

export default OfficeNotifyController