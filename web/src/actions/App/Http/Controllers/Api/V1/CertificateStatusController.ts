import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::index
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:17
* @route '/api/v1/certificate-statuses'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/certificate-statuses',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::index
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:17
* @route '/api/v1/certificate-statuses'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::index
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:17
* @route '/api/v1/certificate-statuses'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::index
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:17
* @route '/api/v1/certificate-statuses'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::index
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:17
* @route '/api/v1/certificate-statuses'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::index
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:17
* @route '/api/v1/certificate-statuses'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::index
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:17
* @route '/api/v1/certificate-statuses'
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
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::show
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:40
* @route '/api/v1/certificate-statuses/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/certificate-statuses/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::show
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:40
* @route '/api/v1/certificate-statuses/{id}'
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
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::show
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:40
* @route '/api/v1/certificate-statuses/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::show
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:40
* @route '/api/v1/certificate-statuses/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::show
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:40
* @route '/api/v1/certificate-statuses/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::show
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:40
* @route '/api/v1/certificate-statuses/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::show
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:40
* @route '/api/v1/certificate-statuses/{id}'
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
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::store
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:54
* @route '/api/v1/certificate-statuses'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/certificate-statuses',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::store
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:54
* @route '/api/v1/certificate-statuses'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::store
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:54
* @route '/api/v1/certificate-statuses'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::store
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:54
* @route '/api/v1/certificate-statuses'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::store
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:54
* @route '/api/v1/certificate-statuses'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::update
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:68
* @route '/api/v1/certificate-statuses/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/certificate-statuses/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::update
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:68
* @route '/api/v1/certificate-statuses/{id}'
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
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::update
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:68
* @route '/api/v1/certificate-statuses/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::update
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:68
* @route '/api/v1/certificate-statuses/{id}'
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
* @see \App\Http\Controllers\Api\V1\CertificateStatusController::update
* @see app/Http/Controllers/Api/V1/CertificateStatusController.php:68
* @route '/api/v1/certificate-statuses/{id}'
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

const CertificateStatusController = { index, show, store, update }

export default CertificateStatusController