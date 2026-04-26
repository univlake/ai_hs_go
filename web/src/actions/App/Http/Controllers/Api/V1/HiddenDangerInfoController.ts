import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::index
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:16
* @route '/api/v1/hidden-danger-infos'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/hidden-danger-infos',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::index
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:16
* @route '/api/v1/hidden-danger-infos'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::index
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:16
* @route '/api/v1/hidden-danger-infos'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::index
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:16
* @route '/api/v1/hidden-danger-infos'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::index
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:16
* @route '/api/v1/hidden-danger-infos'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::index
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:16
* @route '/api/v1/hidden-danger-infos'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::index
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:16
* @route '/api/v1/hidden-danger-infos'
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
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::show
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:41
* @route '/api/v1/hidden-danger-infos/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/hidden-danger-infos/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::show
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:41
* @route '/api/v1/hidden-danger-infos/{id}'
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
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::show
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:41
* @route '/api/v1/hidden-danger-infos/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::show
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:41
* @route '/api/v1/hidden-danger-infos/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::show
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:41
* @route '/api/v1/hidden-danger-infos/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::show
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:41
* @route '/api/v1/hidden-danger-infos/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::show
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:41
* @route '/api/v1/hidden-danger-infos/{id}'
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
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::store
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:59
* @route '/api/v1/hidden-danger-infos'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/hidden-danger-infos',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::store
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:59
* @route '/api/v1/hidden-danger-infos'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::store
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:59
* @route '/api/v1/hidden-danger-infos'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::store
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:59
* @route '/api/v1/hidden-danger-infos'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::store
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:59
* @route '/api/v1/hidden-danger-infos'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::update
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:83
* @route '/api/v1/hidden-danger-infos/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/hidden-danger-infos/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::update
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:83
* @route '/api/v1/hidden-danger-infos/{id}'
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
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::update
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:83
* @route '/api/v1/hidden-danger-infos/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::update
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:83
* @route '/api/v1/hidden-danger-infos/{id}'
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
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::update
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:83
* @route '/api/v1/hidden-danger-infos/{id}'
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
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::destroy
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:125
* @route '/api/v1/hidden-danger-infos/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/hidden-danger-infos/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::destroy
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:125
* @route '/api/v1/hidden-danger-infos/{id}'
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
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::destroy
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:125
* @route '/api/v1/hidden-danger-infos/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::destroy
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:125
* @route '/api/v1/hidden-danger-infos/{id}'
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
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::destroy
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:125
* @route '/api/v1/hidden-danger-infos/{id}'
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
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::rectify
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:147
* @route '/api/v1/hidden-danger-infos/{id}/rectify'
*/
export const rectify = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: rectify.url(args, options),
    method: 'post',
})

rectify.definition = {
    methods: ["post"],
    url: '/api/v1/hidden-danger-infos/{id}/rectify',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::rectify
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:147
* @route '/api/v1/hidden-danger-infos/{id}/rectify'
*/
rectify.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return rectify.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::rectify
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:147
* @route '/api/v1/hidden-danger-infos/{id}/rectify'
*/
rectify.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: rectify.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::rectify
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:147
* @route '/api/v1/hidden-danger-infos/{id}/rectify'
*/
const rectifyForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: rectify.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::rectify
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:147
* @route '/api/v1/hidden-danger-infos/{id}/rectify'
*/
rectifyForm.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: rectify.url(args, options),
    method: 'post',
})

rectify.form = rectifyForm

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::confirmRectification
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:185
* @route '/api/v1/hidden-danger-infos/{id}/confirm-rectification'
*/
export const confirmRectification = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: confirmRectification.url(args, options),
    method: 'post',
})

confirmRectification.definition = {
    methods: ["post"],
    url: '/api/v1/hidden-danger-infos/{id}/confirm-rectification',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::confirmRectification
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:185
* @route '/api/v1/hidden-danger-infos/{id}/confirm-rectification'
*/
confirmRectification.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return confirmRectification.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::confirmRectification
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:185
* @route '/api/v1/hidden-danger-infos/{id}/confirm-rectification'
*/
confirmRectification.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: confirmRectification.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::confirmRectification
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:185
* @route '/api/v1/hidden-danger-infos/{id}/confirm-rectification'
*/
const confirmRectificationForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: confirmRectification.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::confirmRectification
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:185
* @route '/api/v1/hidden-danger-infos/{id}/confirm-rectification'
*/
confirmRectificationForm.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: confirmRectification.url(args, options),
    method: 'post',
})

confirmRectification.form = confirmRectificationForm

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::archive
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:219
* @route '/api/v1/hidden-danger-infos/{id}/archive'
*/
export const archive = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: archive.url(args, options),
    method: 'post',
})

archive.definition = {
    methods: ["post"],
    url: '/api/v1/hidden-danger-infos/{id}/archive',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::archive
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:219
* @route '/api/v1/hidden-danger-infos/{id}/archive'
*/
archive.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return archive.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::archive
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:219
* @route '/api/v1/hidden-danger-infos/{id}/archive'
*/
archive.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: archive.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::archive
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:219
* @route '/api/v1/hidden-danger-infos/{id}/archive'
*/
const archiveForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: archive.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\HiddenDangerInfoController::archive
* @see app/Http/Controllers/Api/V1/HiddenDangerInfoController.php:219
* @route '/api/v1/hidden-danger-infos/{id}/archive'
*/
archiveForm.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: archive.url(args, options),
    method: 'post',
})

archive.form = archiveForm

const HiddenDangerInfoController = { index, show, store, update, destroy, rectify, confirmRectification, archive }

export default HiddenDangerInfoController