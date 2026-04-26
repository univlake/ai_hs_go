import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::index
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:20
* @route '/api/v1/checklist-libraries'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/checklist-libraries',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::index
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:20
* @route '/api/v1/checklist-libraries'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::index
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:20
* @route '/api/v1/checklist-libraries'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::index
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:20
* @route '/api/v1/checklist-libraries'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::index
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:20
* @route '/api/v1/checklist-libraries'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::index
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:20
* @route '/api/v1/checklist-libraries'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::index
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:20
* @route '/api/v1/checklist-libraries'
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
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::exportMethod
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:38
* @route '/api/v1/checklist-libraries/export'
*/
export const exportMethod = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: exportMethod.url(options),
    method: 'get',
})

exportMethod.definition = {
    methods: ["get","head"],
    url: '/api/v1/checklist-libraries/export',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::exportMethod
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:38
* @route '/api/v1/checklist-libraries/export'
*/
exportMethod.url = (options?: RouteQueryOptions) => {
    return exportMethod.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::exportMethod
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:38
* @route '/api/v1/checklist-libraries/export'
*/
exportMethod.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: exportMethod.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::exportMethod
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:38
* @route '/api/v1/checklist-libraries/export'
*/
exportMethod.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: exportMethod.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::exportMethod
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:38
* @route '/api/v1/checklist-libraries/export'
*/
const exportMethodForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: exportMethod.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::exportMethod
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:38
* @route '/api/v1/checklist-libraries/export'
*/
exportMethodForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: exportMethod.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::exportMethod
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:38
* @route '/api/v1/checklist-libraries/export'
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
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::store
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:49
* @route '/api/v1/checklist-libraries'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/checklist-libraries',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::store
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:49
* @route '/api/v1/checklist-libraries'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::store
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:49
* @route '/api/v1/checklist-libraries'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::store
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:49
* @route '/api/v1/checklist-libraries'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::store
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:49
* @route '/api/v1/checklist-libraries'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::update
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:61
* @route '/api/v1/checklist-libraries/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/checklist-libraries/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::update
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:61
* @route '/api/v1/checklist-libraries/{id}'
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
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::update
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:61
* @route '/api/v1/checklist-libraries/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::update
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:61
* @route '/api/v1/checklist-libraries/{id}'
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
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::update
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:61
* @route '/api/v1/checklist-libraries/{id}'
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
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::destroy
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:78
* @route '/api/v1/checklist-libraries/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/checklist-libraries/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::destroy
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:78
* @route '/api/v1/checklist-libraries/{id}'
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
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::destroy
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:78
* @route '/api/v1/checklist-libraries/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::destroy
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:78
* @route '/api/v1/checklist-libraries/{id}'
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
* @see \App\Http\Controllers\Api\V1\ChecklistLibraryController::destroy
* @see app/Http/Controllers/Api/V1/ChecklistLibraryController.php:78
* @route '/api/v1/checklist-libraries/{id}'
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

const ChecklistLibraryController = { index, exportMethod, store, update, destroy, export: exportMethod }

export default ChecklistLibraryController