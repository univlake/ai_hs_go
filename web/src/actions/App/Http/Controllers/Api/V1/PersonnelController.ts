import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::index
* @see app/Http/Controllers/Api/V1/PersonnelController.php:18
* @route '/api/v1/personnels'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/personnels',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::index
* @see app/Http/Controllers/Api/V1/PersonnelController.php:18
* @route '/api/v1/personnels'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::index
* @see app/Http/Controllers/Api/V1/PersonnelController.php:18
* @route '/api/v1/personnels'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::index
* @see app/Http/Controllers/Api/V1/PersonnelController.php:18
* @route '/api/v1/personnels'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::index
* @see app/Http/Controllers/Api/V1/PersonnelController.php:18
* @route '/api/v1/personnels'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::index
* @see app/Http/Controllers/Api/V1/PersonnelController.php:18
* @route '/api/v1/personnels'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::index
* @see app/Http/Controllers/Api/V1/PersonnelController.php:18
* @route '/api/v1/personnels'
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
* @see \App\Http\Controllers\Api\V1\PersonnelController::show
* @see app/Http/Controllers/Api/V1/PersonnelController.php:36
* @route '/api/v1/personnels/{personnel}'
*/
export const show = (args: { personnel: string | number } | [personnel: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/personnels/{personnel}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::show
* @see app/Http/Controllers/Api/V1/PersonnelController.php:36
* @route '/api/v1/personnels/{personnel}'
*/
show.url = (args: { personnel: string | number } | [personnel: string | number ] | string | number, options?: RouteQueryOptions) => {
    if (typeof args === 'string' || typeof args === 'number') {
        args = { personnel: args }
    }

    if (Array.isArray(args)) {
        args = {
            personnel: args[0],
        }
    }

    args = applyUrlDefaults(args)

    const parsedArgs = {
        personnel: args.personnel,
    }

    return show.definition.url
            .replace('{personnel}', parsedArgs.personnel.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::show
* @see app/Http/Controllers/Api/V1/PersonnelController.php:36
* @route '/api/v1/personnels/{personnel}'
*/
show.get = (args: { personnel: string | number } | [personnel: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::show
* @see app/Http/Controllers/Api/V1/PersonnelController.php:36
* @route '/api/v1/personnels/{personnel}'
*/
show.head = (args: { personnel: string | number } | [personnel: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::show
* @see app/Http/Controllers/Api/V1/PersonnelController.php:36
* @route '/api/v1/personnels/{personnel}'
*/
const showForm = (args: { personnel: string | number } | [personnel: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::show
* @see app/Http/Controllers/Api/V1/PersonnelController.php:36
* @route '/api/v1/personnels/{personnel}'
*/
showForm.get = (args: { personnel: string | number } | [personnel: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::show
* @see app/Http/Controllers/Api/V1/PersonnelController.php:36
* @route '/api/v1/personnels/{personnel}'
*/
showForm.head = (args: { personnel: string | number } | [personnel: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
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
* @see \App\Http\Controllers\Api\V1\PersonnelController::store
* @see app/Http/Controllers/Api/V1/PersonnelController.php:52
* @route '/api/v1/personnels'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/personnels',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::store
* @see app/Http/Controllers/Api/V1/PersonnelController.php:52
* @route '/api/v1/personnels'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::store
* @see app/Http/Controllers/Api/V1/PersonnelController.php:52
* @route '/api/v1/personnels'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::store
* @see app/Http/Controllers/Api/V1/PersonnelController.php:52
* @route '/api/v1/personnels'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::store
* @see app/Http/Controllers/Api/V1/PersonnelController.php:52
* @route '/api/v1/personnels'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::update
* @see app/Http/Controllers/Api/V1/PersonnelController.php:64
* @route '/api/v1/personnels/{personnel}'
*/
export const update = (args: { personnel: string | number } | [personnel: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/personnels/{personnel}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::update
* @see app/Http/Controllers/Api/V1/PersonnelController.php:64
* @route '/api/v1/personnels/{personnel}'
*/
update.url = (args: { personnel: string | number } | [personnel: string | number ] | string | number, options?: RouteQueryOptions) => {
    if (typeof args === 'string' || typeof args === 'number') {
        args = { personnel: args }
    }

    if (Array.isArray(args)) {
        args = {
            personnel: args[0],
        }
    }

    args = applyUrlDefaults(args)

    const parsedArgs = {
        personnel: args.personnel,
    }

    return update.definition.url
            .replace('{personnel}', parsedArgs.personnel.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::update
* @see app/Http/Controllers/Api/V1/PersonnelController.php:64
* @route '/api/v1/personnels/{personnel}'
*/
update.put = (args: { personnel: string | number } | [personnel: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::update
* @see app/Http/Controllers/Api/V1/PersonnelController.php:64
* @route '/api/v1/personnels/{personnel}'
*/
const updateForm = (args: { personnel: string | number } | [personnel: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: update.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'PUT',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\PersonnelController::update
* @see app/Http/Controllers/Api/V1/PersonnelController.php:64
* @route '/api/v1/personnels/{personnel}'
*/
updateForm.put = (args: { personnel: string | number } | [personnel: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: update.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'PUT',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

update.form = updateForm

const PersonnelController = { index, show, store, update }

export default PersonnelController