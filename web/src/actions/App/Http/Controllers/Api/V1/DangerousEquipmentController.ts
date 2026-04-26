import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::template
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:89
* @route '/api/v1/dangerous-equipment/template'
*/
export const template = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: template.url(options),
    method: 'get',
})

template.definition = {
    methods: ["get","head"],
    url: '/api/v1/dangerous-equipment/template',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::template
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:89
* @route '/api/v1/dangerous-equipment/template'
*/
template.url = (options?: RouteQueryOptions) => {
    return template.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::template
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:89
* @route '/api/v1/dangerous-equipment/template'
*/
template.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: template.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::template
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:89
* @route '/api/v1/dangerous-equipment/template'
*/
template.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: template.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::template
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:89
* @route '/api/v1/dangerous-equipment/template'
*/
const templateForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: template.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::template
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:89
* @route '/api/v1/dangerous-equipment/template'
*/
templateForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: template.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::template
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:89
* @route '/api/v1/dangerous-equipment/template'
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
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::importMethod
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:108
* @route '/api/v1/dangerous-equipment/import'
*/
export const importMethod = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: importMethod.url(options),
    method: 'post',
})

importMethod.definition = {
    methods: ["post"],
    url: '/api/v1/dangerous-equipment/import',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::importMethod
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:108
* @route '/api/v1/dangerous-equipment/import'
*/
importMethod.url = (options?: RouteQueryOptions) => {
    return importMethod.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::importMethod
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:108
* @route '/api/v1/dangerous-equipment/import'
*/
importMethod.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: importMethod.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::importMethod
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:108
* @route '/api/v1/dangerous-equipment/import'
*/
const importMethodForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: importMethod.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::importMethod
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:108
* @route '/api/v1/dangerous-equipment/import'
*/
importMethodForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: importMethod.url(options),
    method: 'post',
})

importMethod.form = importMethodForm

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::copy
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:139
* @route '/api/v1/dangerous-equipment/copy'
*/
export const copy = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: copy.url(options),
    method: 'post',
})

copy.definition = {
    methods: ["post"],
    url: '/api/v1/dangerous-equipment/copy',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::copy
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:139
* @route '/api/v1/dangerous-equipment/copy'
*/
copy.url = (options?: RouteQueryOptions) => {
    return copy.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::copy
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:139
* @route '/api/v1/dangerous-equipment/copy'
*/
copy.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: copy.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::copy
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:139
* @route '/api/v1/dangerous-equipment/copy'
*/
const copyForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: copy.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::copy
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:139
* @route '/api/v1/dangerous-equipment/copy'
*/
copyForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: copy.url(options),
    method: 'post',
})

copy.form = copyForm

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::exportMethod
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:126
* @route '/api/v1/dangerous-equipment/export'
*/
export const exportMethod = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: exportMethod.url(options),
    method: 'get',
})

exportMethod.definition = {
    methods: ["get","head"],
    url: '/api/v1/dangerous-equipment/export',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::exportMethod
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:126
* @route '/api/v1/dangerous-equipment/export'
*/
exportMethod.url = (options?: RouteQueryOptions) => {
    return exportMethod.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::exportMethod
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:126
* @route '/api/v1/dangerous-equipment/export'
*/
exportMethod.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: exportMethod.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::exportMethod
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:126
* @route '/api/v1/dangerous-equipment/export'
*/
exportMethod.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: exportMethod.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::exportMethod
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:126
* @route '/api/v1/dangerous-equipment/export'
*/
const exportMethodForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: exportMethod.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::exportMethod
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:126
* @route '/api/v1/dangerous-equipment/export'
*/
exportMethodForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: exportMethod.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::exportMethod
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:126
* @route '/api/v1/dangerous-equipment/export'
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
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::index
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:19
* @route '/api/v1/dangerous-equipment'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/dangerous-equipment',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::index
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:19
* @route '/api/v1/dangerous-equipment'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::index
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:19
* @route '/api/v1/dangerous-equipment'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::index
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:19
* @route '/api/v1/dangerous-equipment'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::index
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:19
* @route '/api/v1/dangerous-equipment'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::index
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:19
* @route '/api/v1/dangerous-equipment'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::index
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:19
* @route '/api/v1/dangerous-equipment'
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
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::show
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:37
* @route '/api/v1/dangerous-equipment/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/dangerous-equipment/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::show
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:37
* @route '/api/v1/dangerous-equipment/{id}'
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
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::show
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:37
* @route '/api/v1/dangerous-equipment/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::show
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:37
* @route '/api/v1/dangerous-equipment/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::show
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:37
* @route '/api/v1/dangerous-equipment/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::show
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:37
* @route '/api/v1/dangerous-equipment/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::show
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:37
* @route '/api/v1/dangerous-equipment/{id}'
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
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::store
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:48
* @route '/api/v1/dangerous-equipment'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/dangerous-equipment',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::store
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:48
* @route '/api/v1/dangerous-equipment'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::store
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:48
* @route '/api/v1/dangerous-equipment'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::store
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:48
* @route '/api/v1/dangerous-equipment'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::store
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:48
* @route '/api/v1/dangerous-equipment'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::update
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:59
* @route '/api/v1/dangerous-equipment/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/dangerous-equipment/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::update
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:59
* @route '/api/v1/dangerous-equipment/{id}'
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
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::update
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:59
* @route '/api/v1/dangerous-equipment/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::update
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:59
* @route '/api/v1/dangerous-equipment/{id}'
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
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::update
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:59
* @route '/api/v1/dangerous-equipment/{id}'
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
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::destroy
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:73
* @route '/api/v1/dangerous-equipment/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/dangerous-equipment/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::destroy
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:73
* @route '/api/v1/dangerous-equipment/{id}'
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
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::destroy
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:73
* @route '/api/v1/dangerous-equipment/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::destroy
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:73
* @route '/api/v1/dangerous-equipment/{id}'
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
* @see \App\Http\Controllers\Api\V1\DangerousEquipmentController::destroy
* @see app/Http/Controllers/Api/V1/DangerousEquipmentController.php:73
* @route '/api/v1/dangerous-equipment/{id}'
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

const DangerousEquipmentController = { template, importMethod, copy, exportMethod, index, show, store, update, destroy, import: importMethod, export: exportMethod }

export default DangerousEquipmentController