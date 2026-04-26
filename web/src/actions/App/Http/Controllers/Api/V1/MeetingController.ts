import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\MeetingController::index
* @see app/Http/Controllers/Api/V1/MeetingController.php:16
* @route '/api/v1/meetings'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/meetings',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::index
* @see app/Http/Controllers/Api/V1/MeetingController.php:16
* @route '/api/v1/meetings'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::index
* @see app/Http/Controllers/Api/V1/MeetingController.php:16
* @route '/api/v1/meetings'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::index
* @see app/Http/Controllers/Api/V1/MeetingController.php:16
* @route '/api/v1/meetings'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::index
* @see app/Http/Controllers/Api/V1/MeetingController.php:16
* @route '/api/v1/meetings'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::index
* @see app/Http/Controllers/Api/V1/MeetingController.php:16
* @route '/api/v1/meetings'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::index
* @see app/Http/Controllers/Api/V1/MeetingController.php:16
* @route '/api/v1/meetings'
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
* @see \App\Http\Controllers\Api\V1\MeetingController::show
* @see app/Http/Controllers/Api/V1/MeetingController.php:54
* @route '/api/v1/meetings/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/meetings/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::show
* @see app/Http/Controllers/Api/V1/MeetingController.php:54
* @route '/api/v1/meetings/{id}'
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
* @see \App\Http\Controllers\Api\V1\MeetingController::show
* @see app/Http/Controllers/Api/V1/MeetingController.php:54
* @route '/api/v1/meetings/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::show
* @see app/Http/Controllers/Api/V1/MeetingController.php:54
* @route '/api/v1/meetings/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::show
* @see app/Http/Controllers/Api/V1/MeetingController.php:54
* @route '/api/v1/meetings/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::show
* @see app/Http/Controllers/Api/V1/MeetingController.php:54
* @route '/api/v1/meetings/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::show
* @see app/Http/Controllers/Api/V1/MeetingController.php:54
* @route '/api/v1/meetings/{id}'
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
* @see \App\Http\Controllers\Api\V1\MeetingController::store
* @see app/Http/Controllers/Api/V1/MeetingController.php:68
* @route '/api/v1/meetings'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/meetings',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::store
* @see app/Http/Controllers/Api/V1/MeetingController.php:68
* @route '/api/v1/meetings'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::store
* @see app/Http/Controllers/Api/V1/MeetingController.php:68
* @route '/api/v1/meetings'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::store
* @see app/Http/Controllers/Api/V1/MeetingController.php:68
* @route '/api/v1/meetings'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::store
* @see app/Http/Controllers/Api/V1/MeetingController.php:68
* @route '/api/v1/meetings'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::update
* @see app/Http/Controllers/Api/V1/MeetingController.php:97
* @route '/api/v1/meetings/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/meetings/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::update
* @see app/Http/Controllers/Api/V1/MeetingController.php:97
* @route '/api/v1/meetings/{id}'
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
* @see \App\Http\Controllers\Api\V1\MeetingController::update
* @see app/Http/Controllers/Api/V1/MeetingController.php:97
* @route '/api/v1/meetings/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::update
* @see app/Http/Controllers/Api/V1/MeetingController.php:97
* @route '/api/v1/meetings/{id}'
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
* @see \App\Http\Controllers\Api\V1\MeetingController::update
* @see app/Http/Controllers/Api/V1/MeetingController.php:97
* @route '/api/v1/meetings/{id}'
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
* @see \App\Http\Controllers\Api\V1\MeetingController::destroy
* @see app/Http/Controllers/Api/V1/MeetingController.php:131
* @route '/api/v1/meetings/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/meetings/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::destroy
* @see app/Http/Controllers/Api/V1/MeetingController.php:131
* @route '/api/v1/meetings/{id}'
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
* @see \App\Http\Controllers\Api\V1\MeetingController::destroy
* @see app/Http/Controllers/Api/V1/MeetingController.php:131
* @route '/api/v1/meetings/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::destroy
* @see app/Http/Controllers/Api/V1/MeetingController.php:131
* @route '/api/v1/meetings/{id}'
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
* @see \App\Http\Controllers\Api\V1\MeetingController::destroy
* @see app/Http/Controllers/Api/V1/MeetingController.php:131
* @route '/api/v1/meetings/{id}'
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
* @see \App\Http\Controllers\Api\V1\MeetingController::notifyNotice
* @see app/Http/Controllers/Api/V1/MeetingController.php:153
* @route '/api/v1/meeting-notices/{id}/notify'
*/
export const notifyNotice = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: notifyNotice.url(args, options),
    method: 'post',
})

notifyNotice.definition = {
    methods: ["post"],
    url: '/api/v1/meeting-notices/{id}/notify',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::notifyNotice
* @see app/Http/Controllers/Api/V1/MeetingController.php:153
* @route '/api/v1/meeting-notices/{id}/notify'
*/
notifyNotice.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return notifyNotice.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::notifyNotice
* @see app/Http/Controllers/Api/V1/MeetingController.php:153
* @route '/api/v1/meeting-notices/{id}/notify'
*/
notifyNotice.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: notifyNotice.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::notifyNotice
* @see app/Http/Controllers/Api/V1/MeetingController.php:153
* @route '/api/v1/meeting-notices/{id}/notify'
*/
const notifyNoticeForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: notifyNotice.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::notifyNotice
* @see app/Http/Controllers/Api/V1/MeetingController.php:153
* @route '/api/v1/meeting-notices/{id}/notify'
*/
notifyNoticeForm.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: notifyNotice.url(args, options),
    method: 'post',
})

notifyNotice.form = notifyNoticeForm

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::confirmNotice
* @see app/Http/Controllers/Api/V1/MeetingController.php:178
* @route '/api/v1/meeting-notices/{id}/confirm'
*/
export const confirmNotice = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: confirmNotice.url(args, options),
    method: 'post',
})

confirmNotice.definition = {
    methods: ["post"],
    url: '/api/v1/meeting-notices/{id}/confirm',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::confirmNotice
* @see app/Http/Controllers/Api/V1/MeetingController.php:178
* @route '/api/v1/meeting-notices/{id}/confirm'
*/
confirmNotice.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return confirmNotice.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::confirmNotice
* @see app/Http/Controllers/Api/V1/MeetingController.php:178
* @route '/api/v1/meeting-notices/{id}/confirm'
*/
confirmNotice.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: confirmNotice.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::confirmNotice
* @see app/Http/Controllers/Api/V1/MeetingController.php:178
* @route '/api/v1/meeting-notices/{id}/confirm'
*/
const confirmNoticeForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: confirmNotice.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::confirmNotice
* @see app/Http/Controllers/Api/V1/MeetingController.php:178
* @route '/api/v1/meeting-notices/{id}/confirm'
*/
confirmNoticeForm.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: confirmNotice.url(args, options),
    method: 'post',
})

confirmNotice.form = confirmNoticeForm

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::pendingNotices
* @see app/Http/Controllers/Api/V1/MeetingController.php:197
* @route '/api/v1/meeting-notifications/pending'
*/
export const pendingNotices = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: pendingNotices.url(options),
    method: 'get',
})

pendingNotices.definition = {
    methods: ["get","head"],
    url: '/api/v1/meeting-notifications/pending',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::pendingNotices
* @see app/Http/Controllers/Api/V1/MeetingController.php:197
* @route '/api/v1/meeting-notifications/pending'
*/
pendingNotices.url = (options?: RouteQueryOptions) => {
    return pendingNotices.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::pendingNotices
* @see app/Http/Controllers/Api/V1/MeetingController.php:197
* @route '/api/v1/meeting-notifications/pending'
*/
pendingNotices.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: pendingNotices.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::pendingNotices
* @see app/Http/Controllers/Api/V1/MeetingController.php:197
* @route '/api/v1/meeting-notifications/pending'
*/
pendingNotices.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: pendingNotices.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::pendingNotices
* @see app/Http/Controllers/Api/V1/MeetingController.php:197
* @route '/api/v1/meeting-notifications/pending'
*/
const pendingNoticesForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: pendingNotices.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::pendingNotices
* @see app/Http/Controllers/Api/V1/MeetingController.php:197
* @route '/api/v1/meeting-notifications/pending'
*/
pendingNoticesForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: pendingNotices.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\MeetingController::pendingNotices
* @see app/Http/Controllers/Api/V1/MeetingController.php:197
* @route '/api/v1/meeting-notifications/pending'
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

const MeetingController = { index, show, store, update, destroy, notifyNotice, confirmNotice, pendingNotices }

export default MeetingController