import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::template
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:100
* @route '/api/v1/safety-management/template'
*/
export const template = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: template.url(options),
    method: 'get',
})

template.definition = {
    methods: ["get","head"],
    url: '/api/v1/safety-management/template',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::template
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:100
* @route '/api/v1/safety-management/template'
*/
template.url = (options?: RouteQueryOptions) => {
    return template.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::template
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:100
* @route '/api/v1/safety-management/template'
*/
template.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: template.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::template
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:100
* @route '/api/v1/safety-management/template'
*/
template.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: template.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::template
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:100
* @route '/api/v1/safety-management/template'
*/
const templateForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: template.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::template
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:100
* @route '/api/v1/safety-management/template'
*/
templateForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: template.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::template
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:100
* @route '/api/v1/safety-management/template'
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
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::importMethod
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:113
* @route '/api/v1/safety-management/import'
*/
export const importMethod = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: importMethod.url(options),
    method: 'post',
})

importMethod.definition = {
    methods: ["post"],
    url: '/api/v1/safety-management/import',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::importMethod
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:113
* @route '/api/v1/safety-management/import'
*/
importMethod.url = (options?: RouteQueryOptions) => {
    return importMethod.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::importMethod
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:113
* @route '/api/v1/safety-management/import'
*/
importMethod.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: importMethod.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::importMethod
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:113
* @route '/api/v1/safety-management/import'
*/
const importMethodForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: importMethod.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::importMethod
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:113
* @route '/api/v1/safety-management/import'
*/
importMethodForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: importMethod.url(options),
    method: 'post',
})

importMethod.form = importMethodForm

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::index
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:21
* @route '/api/v1/safety-management'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/safety-management',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::index
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:21
* @route '/api/v1/safety-management'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::index
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:21
* @route '/api/v1/safety-management'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::index
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:21
* @route '/api/v1/safety-management'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::index
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:21
* @route '/api/v1/safety-management'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::index
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:21
* @route '/api/v1/safety-management'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::index
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:21
* @route '/api/v1/safety-management'
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
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::show
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:39
* @route '/api/v1/safety-management/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/safety-management/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::show
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:39
* @route '/api/v1/safety-management/{id}'
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
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::show
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:39
* @route '/api/v1/safety-management/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::show
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:39
* @route '/api/v1/safety-management/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::show
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:39
* @route '/api/v1/safety-management/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::show
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:39
* @route '/api/v1/safety-management/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::show
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:39
* @route '/api/v1/safety-management/{id}'
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
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::store
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:53
* @route '/api/v1/safety-management'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/safety-management',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::store
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:53
* @route '/api/v1/safety-management'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::store
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:53
* @route '/api/v1/safety-management'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::store
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:53
* @route '/api/v1/safety-management'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::store
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:53
* @route '/api/v1/safety-management'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::update
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:67
* @route '/api/v1/safety-management/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/safety-management/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::update
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:67
* @route '/api/v1/safety-management/{id}'
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
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::update
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:67
* @route '/api/v1/safety-management/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::update
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:67
* @route '/api/v1/safety-management/{id}'
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
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::update
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:67
* @route '/api/v1/safety-management/{id}'
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
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::destroy
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:84
* @route '/api/v1/safety-management/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/safety-management/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::destroy
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:84
* @route '/api/v1/safety-management/{id}'
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
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::destroy
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:84
* @route '/api/v1/safety-management/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::destroy
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:84
* @route '/api/v1/safety-management/{id}'
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
* @see \App\Http\Controllers\Api\V1\SafetyManagementController::destroy
* @see app/Http/Controllers/Api/V1/SafetyManagementController.php:84
* @route '/api/v1/safety-management/{id}'
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

const SafetyManagementController = { template, importMethod, index, show, store, update, destroy, import: importMethod }

export default SafetyManagementController