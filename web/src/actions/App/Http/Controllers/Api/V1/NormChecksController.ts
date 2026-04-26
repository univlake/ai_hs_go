import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::index
* @see app/Http/Controllers/Api/V1/NormChecksController.php:15
* @route '/api/v1/norm-checks'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/norm-checks',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::index
* @see app/Http/Controllers/Api/V1/NormChecksController.php:15
* @route '/api/v1/norm-checks'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::index
* @see app/Http/Controllers/Api/V1/NormChecksController.php:15
* @route '/api/v1/norm-checks'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::index
* @see app/Http/Controllers/Api/V1/NormChecksController.php:15
* @route '/api/v1/norm-checks'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::index
* @see app/Http/Controllers/Api/V1/NormChecksController.php:15
* @route '/api/v1/norm-checks'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::index
* @see app/Http/Controllers/Api/V1/NormChecksController.php:15
* @route '/api/v1/norm-checks'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::index
* @see app/Http/Controllers/Api/V1/NormChecksController.php:15
* @route '/api/v1/norm-checks'
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
* @see \App\Http\Controllers\Api\V1\NormChecksController::submissions
* @see app/Http/Controllers/Api/V1/NormChecksController.php:91
* @route '/api/v1/norm-checks/submissions'
*/
export const submissions = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: submissions.url(options),
    method: 'get',
})

submissions.definition = {
    methods: ["get","head"],
    url: '/api/v1/norm-checks/submissions',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::submissions
* @see app/Http/Controllers/Api/V1/NormChecksController.php:91
* @route '/api/v1/norm-checks/submissions'
*/
submissions.url = (options?: RouteQueryOptions) => {
    return submissions.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::submissions
* @see app/Http/Controllers/Api/V1/NormChecksController.php:91
* @route '/api/v1/norm-checks/submissions'
*/
submissions.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: submissions.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::submissions
* @see app/Http/Controllers/Api/V1/NormChecksController.php:91
* @route '/api/v1/norm-checks/submissions'
*/
submissions.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: submissions.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::submissions
* @see app/Http/Controllers/Api/V1/NormChecksController.php:91
* @route '/api/v1/norm-checks/submissions'
*/
const submissionsForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: submissions.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::submissions
* @see app/Http/Controllers/Api/V1/NormChecksController.php:91
* @route '/api/v1/norm-checks/submissions'
*/
submissionsForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: submissions.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::submissions
* @see app/Http/Controllers/Api/V1/NormChecksController.php:91
* @route '/api/v1/norm-checks/submissions'
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
* @see \App\Http\Controllers\Api\V1\NormChecksController::templateLink
* @see app/Http/Controllers/Api/V1/NormChecksController.php:161
* @route '/api/v1/norm-checks/template-link'
*/
export const templateLink = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: templateLink.url(options),
    method: 'get',
})

templateLink.definition = {
    methods: ["get","head"],
    url: '/api/v1/norm-checks/template-link',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::templateLink
* @see app/Http/Controllers/Api/V1/NormChecksController.php:161
* @route '/api/v1/norm-checks/template-link'
*/
templateLink.url = (options?: RouteQueryOptions) => {
    return templateLink.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::templateLink
* @see app/Http/Controllers/Api/V1/NormChecksController.php:161
* @route '/api/v1/norm-checks/template-link'
*/
templateLink.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: templateLink.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::templateLink
* @see app/Http/Controllers/Api/V1/NormChecksController.php:161
* @route '/api/v1/norm-checks/template-link'
*/
templateLink.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: templateLink.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::templateLink
* @see app/Http/Controllers/Api/V1/NormChecksController.php:161
* @route '/api/v1/norm-checks/template-link'
*/
const templateLinkForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: templateLink.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::templateLink
* @see app/Http/Controllers/Api/V1/NormChecksController.php:161
* @route '/api/v1/norm-checks/template-link'
*/
templateLinkForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: templateLink.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::templateLink
* @see app/Http/Controllers/Api/V1/NormChecksController.php:161
* @route '/api/v1/norm-checks/template-link'
*/
templateLinkForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: templateLink.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

templateLink.form = templateLinkForm

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::show
* @see app/Http/Controllers/Api/V1/NormChecksController.php:38
* @route '/api/v1/norm-checks/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/norm-checks/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::show
* @see app/Http/Controllers/Api/V1/NormChecksController.php:38
* @route '/api/v1/norm-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\NormChecksController::show
* @see app/Http/Controllers/Api/V1/NormChecksController.php:38
* @route '/api/v1/norm-checks/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::show
* @see app/Http/Controllers/Api/V1/NormChecksController.php:38
* @route '/api/v1/norm-checks/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::show
* @see app/Http/Controllers/Api/V1/NormChecksController.php:38
* @route '/api/v1/norm-checks/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::show
* @see app/Http/Controllers/Api/V1/NormChecksController.php:38
* @route '/api/v1/norm-checks/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::show
* @see app/Http/Controllers/Api/V1/NormChecksController.php:38
* @route '/api/v1/norm-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\NormChecksController::store
* @see app/Http/Controllers/Api/V1/NormChecksController.php:49
* @route '/api/v1/norm-checks'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/norm-checks',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::store
* @see app/Http/Controllers/Api/V1/NormChecksController.php:49
* @route '/api/v1/norm-checks'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::store
* @see app/Http/Controllers/Api/V1/NormChecksController.php:49
* @route '/api/v1/norm-checks'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::store
* @see app/Http/Controllers/Api/V1/NormChecksController.php:49
* @route '/api/v1/norm-checks'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::store
* @see app/Http/Controllers/Api/V1/NormChecksController.php:49
* @route '/api/v1/norm-checks'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::saveScore
* @see app/Http/Controllers/Api/V1/NormChecksController.php:128
* @route '/api/v1/norm-checks/save-score'
*/
export const saveScore = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: saveScore.url(options),
    method: 'post',
})

saveScore.definition = {
    methods: ["post"],
    url: '/api/v1/norm-checks/save-score',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::saveScore
* @see app/Http/Controllers/Api/V1/NormChecksController.php:128
* @route '/api/v1/norm-checks/save-score'
*/
saveScore.url = (options?: RouteQueryOptions) => {
    return saveScore.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::saveScore
* @see app/Http/Controllers/Api/V1/NormChecksController.php:128
* @route '/api/v1/norm-checks/save-score'
*/
saveScore.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: saveScore.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::saveScore
* @see app/Http/Controllers/Api/V1/NormChecksController.php:128
* @route '/api/v1/norm-checks/save-score'
*/
const saveScoreForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: saveScore.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::saveScore
* @see app/Http/Controllers/Api/V1/NormChecksController.php:128
* @route '/api/v1/norm-checks/save-score'
*/
saveScoreForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: saveScore.url(options),
    method: 'post',
})

saveScore.form = saveScoreForm

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::update
* @see app/Http/Controllers/Api/V1/NormChecksController.php:60
* @route '/api/v1/norm-checks/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/norm-checks/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::update
* @see app/Http/Controllers/Api/V1/NormChecksController.php:60
* @route '/api/v1/norm-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\NormChecksController::update
* @see app/Http/Controllers/Api/V1/NormChecksController.php:60
* @route '/api/v1/norm-checks/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::update
* @see app/Http/Controllers/Api/V1/NormChecksController.php:60
* @route '/api/v1/norm-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\NormChecksController::update
* @see app/Http/Controllers/Api/V1/NormChecksController.php:60
* @route '/api/v1/norm-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\NormChecksController::destroy
* @see app/Http/Controllers/Api/V1/NormChecksController.php:74
* @route '/api/v1/norm-checks/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/norm-checks/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::destroy
* @see app/Http/Controllers/Api/V1/NormChecksController.php:74
* @route '/api/v1/norm-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\NormChecksController::destroy
* @see app/Http/Controllers/Api/V1/NormChecksController.php:74
* @route '/api/v1/norm-checks/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\NormChecksController::destroy
* @see app/Http/Controllers/Api/V1/NormChecksController.php:74
* @route '/api/v1/norm-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\NormChecksController::destroy
* @see app/Http/Controllers/Api/V1/NormChecksController.php:74
* @route '/api/v1/norm-checks/{id}'
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

const NormChecksController = { index, submissions, templateLink, show, store, saveScore, update, destroy }

export default NormChecksController