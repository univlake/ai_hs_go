import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::index
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:17
* @route '/api/v1/accident-checks'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/accident-checks',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::index
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:17
* @route '/api/v1/accident-checks'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::index
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:17
* @route '/api/v1/accident-checks'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::index
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:17
* @route '/api/v1/accident-checks'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::index
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:17
* @route '/api/v1/accident-checks'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::index
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:17
* @route '/api/v1/accident-checks'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::index
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:17
* @route '/api/v1/accident-checks'
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
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::submissions
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:107
* @route '/api/v1/accident-checks/submissions'
*/
export const submissions = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: submissions.url(options),
    method: 'get',
})

submissions.definition = {
    methods: ["get","head"],
    url: '/api/v1/accident-checks/submissions',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::submissions
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:107
* @route '/api/v1/accident-checks/submissions'
*/
submissions.url = (options?: RouteQueryOptions) => {
    return submissions.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::submissions
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:107
* @route '/api/v1/accident-checks/submissions'
*/
submissions.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: submissions.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::submissions
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:107
* @route '/api/v1/accident-checks/submissions'
*/
submissions.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: submissions.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::submissions
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:107
* @route '/api/v1/accident-checks/submissions'
*/
const submissionsForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: submissions.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::submissions
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:107
* @route '/api/v1/accident-checks/submissions'
*/
submissionsForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: submissions.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::submissions
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:107
* @route '/api/v1/accident-checks/submissions'
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
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::startForm
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:39
* @route '/api/v1/accident-checks/start-form'
*/
export const startForm = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: startForm.url(options),
    method: 'get',
})

startForm.definition = {
    methods: ["get","head"],
    url: '/api/v1/accident-checks/start-form',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::startForm
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:39
* @route '/api/v1/accident-checks/start-form'
*/
startForm.url = (options?: RouteQueryOptions) => {
    return startForm.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::startForm
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:39
* @route '/api/v1/accident-checks/start-form'
*/
startForm.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: startForm.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::startForm
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:39
* @route '/api/v1/accident-checks/start-form'
*/
startForm.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: startForm.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::startForm
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:39
* @route '/api/v1/accident-checks/start-form'
*/
const startFormForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: startForm.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::startForm
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:39
* @route '/api/v1/accident-checks/start-form'
*/
startFormForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: startForm.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::startForm
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:39
* @route '/api/v1/accident-checks/start-form'
*/
startFormForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: startForm.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

startForm.form = startFormForm

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::show
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:146
* @route '/api/v1/accident-checks/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/accident-checks/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::show
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:146
* @route '/api/v1/accident-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::show
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:146
* @route '/api/v1/accident-checks/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::show
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:146
* @route '/api/v1/accident-checks/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::show
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:146
* @route '/api/v1/accident-checks/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::show
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:146
* @route '/api/v1/accident-checks/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::show
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:146
* @route '/api/v1/accident-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::store
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:157
* @route '/api/v1/accident-checks'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/accident-checks',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::store
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:157
* @route '/api/v1/accident-checks'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::store
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:157
* @route '/api/v1/accident-checks'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::store
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:157
* @route '/api/v1/accident-checks'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::store
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:157
* @route '/api/v1/accident-checks'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::saveScore
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:68
* @route '/api/v1/accident-checks/save-score'
*/
export const saveScore = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: saveScore.url(options),
    method: 'post',
})

saveScore.definition = {
    methods: ["post"],
    url: '/api/v1/accident-checks/save-score',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::saveScore
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:68
* @route '/api/v1/accident-checks/save-score'
*/
saveScore.url = (options?: RouteQueryOptions) => {
    return saveScore.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::saveScore
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:68
* @route '/api/v1/accident-checks/save-score'
*/
saveScore.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: saveScore.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::saveScore
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:68
* @route '/api/v1/accident-checks/save-score'
*/
const saveScoreForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: saveScore.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::saveScore
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:68
* @route '/api/v1/accident-checks/save-score'
*/
saveScoreForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: saveScore.url(options),
    method: 'post',
})

saveScore.form = saveScoreForm

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::update
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:168
* @route '/api/v1/accident-checks/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/accident-checks/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::update
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:168
* @route '/api/v1/accident-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::update
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:168
* @route '/api/v1/accident-checks/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::update
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:168
* @route '/api/v1/accident-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::update
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:168
* @route '/api/v1/accident-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::destroy
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:182
* @route '/api/v1/accident-checks/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/accident-checks/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::destroy
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:182
* @route '/api/v1/accident-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::destroy
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:182
* @route '/api/v1/accident-checks/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::destroy
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:182
* @route '/api/v1/accident-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\AccidentChecksController::destroy
* @see app/Http/Controllers/Api/V1/AccidentChecksController.php:182
* @route '/api/v1/accident-checks/{id}'
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

const AccidentChecksController = { index, submissions, startForm, show, store, saveScore, update, destroy }

export default AccidentChecksController