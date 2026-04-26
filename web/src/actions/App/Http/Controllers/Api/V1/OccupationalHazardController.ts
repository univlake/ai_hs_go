import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::index
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:19
* @route '/api/v1/occupational-hazards'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/occupational-hazards',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::index
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:19
* @route '/api/v1/occupational-hazards'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::index
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:19
* @route '/api/v1/occupational-hazards'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::index
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:19
* @route '/api/v1/occupational-hazards'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::index
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:19
* @route '/api/v1/occupational-hazards'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::index
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:19
* @route '/api/v1/occupational-hazards'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::index
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:19
* @route '/api/v1/occupational-hazards'
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
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::show
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:41
* @route '/api/v1/occupational-hazards/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/occupational-hazards/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::show
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:41
* @route '/api/v1/occupational-hazards/{id}'
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
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::show
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:41
* @route '/api/v1/occupational-hazards/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::show
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:41
* @route '/api/v1/occupational-hazards/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::show
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:41
* @route '/api/v1/occupational-hazards/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::show
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:41
* @route '/api/v1/occupational-hazards/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::show
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:41
* @route '/api/v1/occupational-hazards/{id}'
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
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::exportMethod
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:125
* @route '/api/v1/occupational-hazards-export'
*/
export const exportMethod = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: exportMethod.url(options),
    method: 'get',
})

exportMethod.definition = {
    methods: ["get","head"],
    url: '/api/v1/occupational-hazards-export',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::exportMethod
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:125
* @route '/api/v1/occupational-hazards-export'
*/
exportMethod.url = (options?: RouteQueryOptions) => {
    return exportMethod.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::exportMethod
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:125
* @route '/api/v1/occupational-hazards-export'
*/
exportMethod.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: exportMethod.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::exportMethod
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:125
* @route '/api/v1/occupational-hazards-export'
*/
exportMethod.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: exportMethod.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::exportMethod
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:125
* @route '/api/v1/occupational-hazards-export'
*/
const exportMethodForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: exportMethod.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::exportMethod
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:125
* @route '/api/v1/occupational-hazards-export'
*/
exportMethodForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: exportMethod.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::exportMethod
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:125
* @route '/api/v1/occupational-hazards-export'
*/
exportMethodForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: exportMethod.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

exportMethod.form = exportMethodForm

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::template
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:90
* @route '/api/v1/occupational-hazards-template'
*/
export const template = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: template.url(options),
    method: 'get',
})

template.definition = {
    methods: ["get","head"],
    url: '/api/v1/occupational-hazards-template',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::template
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:90
* @route '/api/v1/occupational-hazards-template'
*/
template.url = (options?: RouteQueryOptions) => {
    return template.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::template
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:90
* @route '/api/v1/occupational-hazards-template'
*/
template.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: template.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::template
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:90
* @route '/api/v1/occupational-hazards-template'
*/
template.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: template.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::template
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:90
* @route '/api/v1/occupational-hazards-template'
*/
const templateForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: template.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::template
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:90
* @route '/api/v1/occupational-hazards-template'
*/
templateForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: template.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::template
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:90
* @route '/api/v1/occupational-hazards-template'
*/
templateForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: template.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

template.form = templateForm

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::store
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:52
* @route '/api/v1/occupational-hazards'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/occupational-hazards',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::store
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:52
* @route '/api/v1/occupational-hazards'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::store
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:52
* @route '/api/v1/occupational-hazards'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::store
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:52
* @route '/api/v1/occupational-hazards'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::store
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:52
* @route '/api/v1/occupational-hazards'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::importMethod
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:100
* @route '/api/v1/occupational-hazards-import'
*/
export const importMethod = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: importMethod.url(options),
    method: 'post',
})

importMethod.definition = {
    methods: ["post"],
    url: '/api/v1/occupational-hazards-import',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::importMethod
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:100
* @route '/api/v1/occupational-hazards-import'
*/
importMethod.url = (options?: RouteQueryOptions) => {
    return importMethod.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::importMethod
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:100
* @route '/api/v1/occupational-hazards-import'
*/
importMethod.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: importMethod.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::importMethod
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:100
* @route '/api/v1/occupational-hazards-import'
*/
const importMethodForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: importMethod.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::importMethod
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:100
* @route '/api/v1/occupational-hazards-import'
*/
importMethodForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: importMethod.url(options),
    method: 'post',
})

importMethod.form = importMethodForm

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::update
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:63
* @route '/api/v1/occupational-hazards/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/occupational-hazards/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::update
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:63
* @route '/api/v1/occupational-hazards/{id}'
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
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::update
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:63
* @route '/api/v1/occupational-hazards/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::update
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:63
* @route '/api/v1/occupational-hazards/{id}'
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
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::update
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:63
* @route '/api/v1/occupational-hazards/{id}'
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
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::destroy
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:77
* @route '/api/v1/occupational-hazards/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/occupational-hazards/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::destroy
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:77
* @route '/api/v1/occupational-hazards/{id}'
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
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::destroy
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:77
* @route '/api/v1/occupational-hazards/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::destroy
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:77
* @route '/api/v1/occupational-hazards/{id}'
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
* @see \App\Http\Controllers\Api\V1\OccupationalHazardController::destroy
* @see app/Http/Controllers/Api/V1/OccupationalHazardController.php:77
* @route '/api/v1/occupational-hazards/{id}'
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

const OccupationalHazardController = { index, show, exportMethod, template, store, importMethod, update, destroy, export: exportMethod, import: importMethod }

export default OccupationalHazardController