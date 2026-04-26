import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::template
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:21
* @route '/api/v1/special-equipment-staff/template'
*/
export const template = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: template.url(options),
    method: 'get',
})

template.definition = {
    methods: ["get","head"],
    url: '/api/v1/special-equipment-staff/template',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::template
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:21
* @route '/api/v1/special-equipment-staff/template'
*/
template.url = (options?: RouteQueryOptions) => {
    return template.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::template
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:21
* @route '/api/v1/special-equipment-staff/template'
*/
template.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: template.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::template
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:21
* @route '/api/v1/special-equipment-staff/template'
*/
template.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: template.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::template
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:21
* @route '/api/v1/special-equipment-staff/template'
*/
const templateForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: template.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::template
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:21
* @route '/api/v1/special-equipment-staff/template'
*/
templateForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: template.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::template
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:21
* @route '/api/v1/special-equipment-staff/template'
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
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::importMethod
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:34
* @route '/api/v1/special-equipment-staff/import'
*/
export const importMethod = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: importMethod.url(options),
    method: 'post',
})

importMethod.definition = {
    methods: ["post"],
    url: '/api/v1/special-equipment-staff/import',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::importMethod
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:34
* @route '/api/v1/special-equipment-staff/import'
*/
importMethod.url = (options?: RouteQueryOptions) => {
    return importMethod.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::importMethod
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:34
* @route '/api/v1/special-equipment-staff/import'
*/
importMethod.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: importMethod.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::importMethod
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:34
* @route '/api/v1/special-equipment-staff/import'
*/
const importMethodForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: importMethod.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::importMethod
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:34
* @route '/api/v1/special-equipment-staff/import'
*/
importMethodForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: importMethod.url(options),
    method: 'post',
})

importMethod.form = importMethodForm

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::index
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:51
* @route '/api/v1/special-equipment-staff'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/special-equipment-staff',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::index
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:51
* @route '/api/v1/special-equipment-staff'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::index
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:51
* @route '/api/v1/special-equipment-staff'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::index
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:51
* @route '/api/v1/special-equipment-staff'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::index
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:51
* @route '/api/v1/special-equipment-staff'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::index
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:51
* @route '/api/v1/special-equipment-staff'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::index
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:51
* @route '/api/v1/special-equipment-staff'
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
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::show
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:69
* @route '/api/v1/special-equipment-staff/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/special-equipment-staff/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::show
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:69
* @route '/api/v1/special-equipment-staff/{id}'
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
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::show
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:69
* @route '/api/v1/special-equipment-staff/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::show
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:69
* @route '/api/v1/special-equipment-staff/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::show
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:69
* @route '/api/v1/special-equipment-staff/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::show
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:69
* @route '/api/v1/special-equipment-staff/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::show
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:69
* @route '/api/v1/special-equipment-staff/{id}'
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
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::store
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:80
* @route '/api/v1/special-equipment-staff'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/special-equipment-staff',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::store
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:80
* @route '/api/v1/special-equipment-staff'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::store
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:80
* @route '/api/v1/special-equipment-staff'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::store
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:80
* @route '/api/v1/special-equipment-staff'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::store
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:80
* @route '/api/v1/special-equipment-staff'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::update
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:91
* @route '/api/v1/special-equipment-staff/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/special-equipment-staff/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::update
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:91
* @route '/api/v1/special-equipment-staff/{id}'
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
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::update
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:91
* @route '/api/v1/special-equipment-staff/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::update
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:91
* @route '/api/v1/special-equipment-staff/{id}'
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
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::update
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:91
* @route '/api/v1/special-equipment-staff/{id}'
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
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::destroy
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:105
* @route '/api/v1/special-equipment-staff/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/special-equipment-staff/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::destroy
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:105
* @route '/api/v1/special-equipment-staff/{id}'
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
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::destroy
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:105
* @route '/api/v1/special-equipment-staff/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::destroy
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:105
* @route '/api/v1/special-equipment-staff/{id}'
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
* @see \App\Http\Controllers\Api\V1\SpecialEquipmentStaffController::destroy
* @see app/Http/Controllers/Api/V1/SpecialEquipmentStaffController.php:105
* @route '/api/v1/special-equipment-staff/{id}'
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

const SpecialEquipmentStaffController = { template, importMethod, index, show, store, update, destroy, import: importMethod }

export default SpecialEquipmentStaffController