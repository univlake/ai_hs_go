import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::index
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:15
* @route '/api/v1/information-drafts'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/information-drafts',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::index
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:15
* @route '/api/v1/information-drafts'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::index
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:15
* @route '/api/v1/information-drafts'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::index
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:15
* @route '/api/v1/information-drafts'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::index
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:15
* @route '/api/v1/information-drafts'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::index
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:15
* @route '/api/v1/information-drafts'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::index
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:15
* @route '/api/v1/information-drafts'
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
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::dashboardReadings
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:82
* @route '/api/v1/information-drafts/dashboard/readings'
*/
export const dashboardReadings = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: dashboardReadings.url(options),
    method: 'get',
})

dashboardReadings.definition = {
    methods: ["get","head"],
    url: '/api/v1/information-drafts/dashboard/readings',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::dashboardReadings
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:82
* @route '/api/v1/information-drafts/dashboard/readings'
*/
dashboardReadings.url = (options?: RouteQueryOptions) => {
    return dashboardReadings.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::dashboardReadings
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:82
* @route '/api/v1/information-drafts/dashboard/readings'
*/
dashboardReadings.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: dashboardReadings.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::dashboardReadings
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:82
* @route '/api/v1/information-drafts/dashboard/readings'
*/
dashboardReadings.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: dashboardReadings.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::dashboardReadings
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:82
* @route '/api/v1/information-drafts/dashboard/readings'
*/
const dashboardReadingsForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: dashboardReadings.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::dashboardReadings
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:82
* @route '/api/v1/information-drafts/dashboard/readings'
*/
dashboardReadingsForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: dashboardReadings.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::dashboardReadings
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:82
* @route '/api/v1/information-drafts/dashboard/readings'
*/
dashboardReadingsForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: dashboardReadings.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

dashboardReadings.form = dashboardReadingsForm

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::dashboardPending
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:102
* @route '/api/v1/information-drafts/dashboard/pending'
*/
export const dashboardPending = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: dashboardPending.url(options),
    method: 'get',
})

dashboardPending.definition = {
    methods: ["get","head"],
    url: '/api/v1/information-drafts/dashboard/pending',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::dashboardPending
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:102
* @route '/api/v1/information-drafts/dashboard/pending'
*/
dashboardPending.url = (options?: RouteQueryOptions) => {
    return dashboardPending.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::dashboardPending
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:102
* @route '/api/v1/information-drafts/dashboard/pending'
*/
dashboardPending.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: dashboardPending.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::dashboardPending
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:102
* @route '/api/v1/information-drafts/dashboard/pending'
*/
dashboardPending.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: dashboardPending.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::dashboardPending
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:102
* @route '/api/v1/information-drafts/dashboard/pending'
*/
const dashboardPendingForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: dashboardPending.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::dashboardPending
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:102
* @route '/api/v1/information-drafts/dashboard/pending'
*/
dashboardPendingForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: dashboardPending.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::dashboardPending
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:102
* @route '/api/v1/information-drafts/dashboard/pending'
*/
dashboardPendingForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: dashboardPending.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

dashboardPending.form = dashboardPendingForm

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::show
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:136
* @route '/api/v1/information-drafts/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/information-drafts/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::show
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:136
* @route '/api/v1/information-drafts/{id}'
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
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::show
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:136
* @route '/api/v1/information-drafts/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::show
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:136
* @route '/api/v1/information-drafts/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::show
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:136
* @route '/api/v1/information-drafts/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::show
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:136
* @route '/api/v1/information-drafts/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::show
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:136
* @route '/api/v1/information-drafts/{id}'
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
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::store
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:147
* @route '/api/v1/information-drafts'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/information-drafts',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::store
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:147
* @route '/api/v1/information-drafts'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::store
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:147
* @route '/api/v1/information-drafts'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::store
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:147
* @route '/api/v1/information-drafts'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::store
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:147
* @route '/api/v1/information-drafts'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::update
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:158
* @route '/api/v1/information-drafts/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/information-drafts/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::update
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:158
* @route '/api/v1/information-drafts/{id}'
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
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::update
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:158
* @route '/api/v1/information-drafts/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::update
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:158
* @route '/api/v1/information-drafts/{id}'
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
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::update
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:158
* @route '/api/v1/information-drafts/{id}'
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
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::destroy
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:172
* @route '/api/v1/information-drafts/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/information-drafts/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::destroy
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:172
* @route '/api/v1/information-drafts/{id}'
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
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::destroy
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:172
* @route '/api/v1/information-drafts/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::destroy
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:172
* @route '/api/v1/information-drafts/{id}'
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
* @see \App\Http\Controllers\Api\V1\InformationDraftsController::destroy
* @see app/Http/Controllers/Api/V1/InformationDraftsController.php:172
* @route '/api/v1/information-drafts/{id}'
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

const InformationDraftsController = { index, dashboardReadings, dashboardPending, show, store, update, destroy }

export default InformationDraftsController