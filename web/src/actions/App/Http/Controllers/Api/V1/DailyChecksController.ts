import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::index
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:319
* @route '/api/v1/daily-checks'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/daily-checks',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::index
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:319
* @route '/api/v1/daily-checks'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::index
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:319
* @route '/api/v1/daily-checks'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::index
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:319
* @route '/api/v1/daily-checks'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::index
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:319
* @route '/api/v1/daily-checks'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::index
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:319
* @route '/api/v1/daily-checks'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::index
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:319
* @route '/api/v1/daily-checks'
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
* @see \App\Http\Controllers\Api\V1\DailyChecksController::submissions
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:27
* @route '/api/v1/daily-checks/submissions'
*/
export const submissions = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: submissions.url(options),
    method: 'get',
})

submissions.definition = {
    methods: ["get","head"],
    url: '/api/v1/daily-checks/submissions',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::submissions
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:27
* @route '/api/v1/daily-checks/submissions'
*/
submissions.url = (options?: RouteQueryOptions) => {
    return submissions.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::submissions
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:27
* @route '/api/v1/daily-checks/submissions'
*/
submissions.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: submissions.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::submissions
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:27
* @route '/api/v1/daily-checks/submissions'
*/
submissions.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: submissions.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::submissions
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:27
* @route '/api/v1/daily-checks/submissions'
*/
const submissionsForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: submissions.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::submissions
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:27
* @route '/api/v1/daily-checks/submissions'
*/
submissionsForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: submissions.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::submissions
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:27
* @route '/api/v1/daily-checks/submissions'
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
* @see \App\Http\Controllers\Api\V1\DailyChecksController::startForm
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:95
* @route '/api/v1/daily-checks/start-form'
*/
export const startForm = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: startForm.url(options),
    method: 'get',
})

startForm.definition = {
    methods: ["get","head"],
    url: '/api/v1/daily-checks/start-form',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::startForm
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:95
* @route '/api/v1/daily-checks/start-form'
*/
startForm.url = (options?: RouteQueryOptions) => {
    return startForm.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::startForm
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:95
* @route '/api/v1/daily-checks/start-form'
*/
startForm.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: startForm.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::startForm
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:95
* @route '/api/v1/daily-checks/start-form'
*/
startForm.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: startForm.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::startForm
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:95
* @route '/api/v1/daily-checks/start-form'
*/
const startFormForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: startForm.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::startForm
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:95
* @route '/api/v1/daily-checks/start-form'
*/
startFormForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: startForm.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::startForm
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:95
* @route '/api/v1/daily-checks/start-form'
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
* @see \App\Http\Controllers\Api\V1\DailyChecksController::showSubmission
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:248
* @route '/api/v1/daily-checks/submission'
*/
export const showSubmission = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: showSubmission.url(options),
    method: 'get',
})

showSubmission.definition = {
    methods: ["get","head"],
    url: '/api/v1/daily-checks/submission',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::showSubmission
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:248
* @route '/api/v1/daily-checks/submission'
*/
showSubmission.url = (options?: RouteQueryOptions) => {
    return showSubmission.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::showSubmission
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:248
* @route '/api/v1/daily-checks/submission'
*/
showSubmission.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: showSubmission.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::showSubmission
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:248
* @route '/api/v1/daily-checks/submission'
*/
showSubmission.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: showSubmission.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::showSubmission
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:248
* @route '/api/v1/daily-checks/submission'
*/
const showSubmissionForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: showSubmission.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::showSubmission
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:248
* @route '/api/v1/daily-checks/submission'
*/
showSubmissionForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: showSubmission.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::showSubmission
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:248
* @route '/api/v1/daily-checks/submission'
*/
showSubmissionForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: showSubmission.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

showSubmission.form = showSubmissionForm

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::show
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:338
* @route '/api/v1/daily-checks/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/daily-checks/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::show
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:338
* @route '/api/v1/daily-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\DailyChecksController::show
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:338
* @route '/api/v1/daily-checks/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::show
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:338
* @route '/api/v1/daily-checks/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::show
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:338
* @route '/api/v1/daily-checks/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::show
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:338
* @route '/api/v1/daily-checks/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::show
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:338
* @route '/api/v1/daily-checks/{id}'
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
* @see \App\Http\Controllers\Api\V1\DailyChecksController::saveDraft
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:135
* @route '/api/v1/daily-checks/save-draft'
*/
export const saveDraft = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: saveDraft.url(options),
    method: 'post',
})

saveDraft.definition = {
    methods: ["post"],
    url: '/api/v1/daily-checks/save-draft',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::saveDraft
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:135
* @route '/api/v1/daily-checks/save-draft'
*/
saveDraft.url = (options?: RouteQueryOptions) => {
    return saveDraft.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::saveDraft
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:135
* @route '/api/v1/daily-checks/save-draft'
*/
saveDraft.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: saveDraft.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::saveDraft
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:135
* @route '/api/v1/daily-checks/save-draft'
*/
const saveDraftForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: saveDraft.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::saveDraft
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:135
* @route '/api/v1/daily-checks/save-draft'
*/
saveDraftForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: saveDraft.url(options),
    method: 'post',
})

saveDraft.form = saveDraftForm

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::complete
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:183
* @route '/api/v1/daily-checks/complete'
*/
export const complete = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: complete.url(options),
    method: 'post',
})

complete.definition = {
    methods: ["post"],
    url: '/api/v1/daily-checks/complete',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::complete
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:183
* @route '/api/v1/daily-checks/complete'
*/
complete.url = (options?: RouteQueryOptions) => {
    return complete.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::complete
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:183
* @route '/api/v1/daily-checks/complete'
*/
complete.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: complete.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::complete
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:183
* @route '/api/v1/daily-checks/complete'
*/
const completeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: complete.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::complete
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:183
* @route '/api/v1/daily-checks/complete'
*/
completeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: complete.url(options),
    method: 'post',
})

complete.form = completeForm

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::score
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:212
* @route '/api/v1/daily-checks/score'
*/
export const score = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: score.url(options),
    method: 'post',
})

score.definition = {
    methods: ["post"],
    url: '/api/v1/daily-checks/score',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::score
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:212
* @route '/api/v1/daily-checks/score'
*/
score.url = (options?: RouteQueryOptions) => {
    return score.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::score
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:212
* @route '/api/v1/daily-checks/score'
*/
score.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: score.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::score
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:212
* @route '/api/v1/daily-checks/score'
*/
const scoreForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: score.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::score
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:212
* @route '/api/v1/daily-checks/score'
*/
scoreForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: score.url(options),
    method: 'post',
})

score.form = scoreForm

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::destroyForm
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:289
* @route '/api/v1/daily-checks/destroy-form'
*/
export const destroyForm = (options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroyForm.url(options),
    method: 'delete',
})

destroyForm.definition = {
    methods: ["delete"],
    url: '/api/v1/daily-checks/destroy-form',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::destroyForm
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:289
* @route '/api/v1/daily-checks/destroy-form'
*/
destroyForm.url = (options?: RouteQueryOptions) => {
    return destroyForm.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::destroyForm
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:289
* @route '/api/v1/daily-checks/destroy-form'
*/
destroyForm.delete = (options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroyForm.url(options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::destroyForm
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:289
* @route '/api/v1/daily-checks/destroy-form'
*/
const destroyFormForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: destroyForm.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'DELETE',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\DailyChecksController::destroyForm
* @see app/Http/Controllers/Api/V1/DailyChecksController.php:289
* @route '/api/v1/daily-checks/destroy-form'
*/
destroyFormForm.delete = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: destroyForm.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'DELETE',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

destroyForm.form = destroyFormForm

const DailyChecksController = { index, submissions, startForm, showSubmission, show, saveDraft, complete, score, destroyForm }

export default DailyChecksController