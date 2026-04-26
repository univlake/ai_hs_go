import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::index
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:18
* @route '/api/v1/system-checks'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/system-checks',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::index
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:18
* @route '/api/v1/system-checks'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::index
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:18
* @route '/api/v1/system-checks'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::index
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:18
* @route '/api/v1/system-checks'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::index
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:18
* @route '/api/v1/system-checks'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::index
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:18
* @route '/api/v1/system-checks'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::index
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:18
* @route '/api/v1/system-checks'
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
* @see \App\Http\Controllers\Api\V1\SystemChecksController::submissions
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:89
* @route '/api/v1/system-checks/submissions'
*/
export const submissions = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: submissions.url(options),
    method: 'get',
})

submissions.definition = {
    methods: ["get","head"],
    url: '/api/v1/system-checks/submissions',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::submissions
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:89
* @route '/api/v1/system-checks/submissions'
*/
submissions.url = (options?: RouteQueryOptions) => {
    return submissions.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::submissions
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:89
* @route '/api/v1/system-checks/submissions'
*/
submissions.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: submissions.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::submissions
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:89
* @route '/api/v1/system-checks/submissions'
*/
submissions.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: submissions.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::submissions
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:89
* @route '/api/v1/system-checks/submissions'
*/
const submissionsForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: submissions.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::submissions
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:89
* @route '/api/v1/system-checks/submissions'
*/
submissionsForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: submissions.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::submissions
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:89
* @route '/api/v1/system-checks/submissions'
*/
submissionsForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: submissions.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

submissions.form = submissionsForm

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::show
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:37
* @route '/api/v1/system-checks/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/system-checks/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::show
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:37
* @route '/api/v1/system-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\SystemChecksController::show
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:37
* @route '/api/v1/system-checks/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::show
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:37
* @route '/api/v1/system-checks/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::show
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:37
* @route '/api/v1/system-checks/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::show
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:37
* @route '/api/v1/system-checks/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::show
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:37
* @route '/api/v1/system-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\SystemChecksController::store
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:48
* @route '/api/v1/system-checks'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/system-checks',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::store
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:48
* @route '/api/v1/system-checks'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::store
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:48
* @route '/api/v1/system-checks'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::store
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:48
* @route '/api/v1/system-checks'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::store
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:48
* @route '/api/v1/system-checks'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::computeScore
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:129
* @route '/api/v1/system-checks/compute'
*/
export const computeScore = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: computeScore.url(options),
    method: 'post',
})

computeScore.definition = {
    methods: ["post"],
    url: '/api/v1/system-checks/compute',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::computeScore
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:129
* @route '/api/v1/system-checks/compute'
*/
computeScore.url = (options?: RouteQueryOptions) => {
    return computeScore.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::computeScore
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:129
* @route '/api/v1/system-checks/compute'
*/
computeScore.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: computeScore.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::computeScore
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:129
* @route '/api/v1/system-checks/compute'
*/
const computeScoreForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: computeScore.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::computeScore
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:129
* @route '/api/v1/system-checks/compute'
*/
computeScoreForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: computeScore.url(options),
    method: 'post',
})

computeScore.form = computeScoreForm

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::update
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:59
* @route '/api/v1/system-checks/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/system-checks/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::update
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:59
* @route '/api/v1/system-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\SystemChecksController::update
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:59
* @route '/api/v1/system-checks/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::update
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:59
* @route '/api/v1/system-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\SystemChecksController::update
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:59
* @route '/api/v1/system-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\SystemChecksController::destroy
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:73
* @route '/api/v1/system-checks/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/system-checks/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::destroy
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:73
* @route '/api/v1/system-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\SystemChecksController::destroy
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:73
* @route '/api/v1/system-checks/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\SystemChecksController::destroy
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:73
* @route '/api/v1/system-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\SystemChecksController::destroy
* @see app/Http/Controllers/Api/V1/SystemChecksController.php:73
* @route '/api/v1/system-checks/{id}'
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

const SystemChecksController = { index, submissions, show, store, computeScore, update, destroy }

export default SystemChecksController