import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::index
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:21
* @route '/api/v1/year-statistics'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/year-statistics',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::index
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:21
* @route '/api/v1/year-statistics'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::index
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:21
* @route '/api/v1/year-statistics'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::index
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:21
* @route '/api/v1/year-statistics'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::index
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:21
* @route '/api/v1/year-statistics'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::index
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:21
* @route '/api/v1/year-statistics'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::index
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:21
* @route '/api/v1/year-statistics'
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
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::stuff
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:163
* @route '/api/v1/year-statistics/stuff'
*/
export const stuff = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: stuff.url(options),
    method: 'get',
})

stuff.definition = {
    methods: ["get","head"],
    url: '/api/v1/year-statistics/stuff',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::stuff
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:163
* @route '/api/v1/year-statistics/stuff'
*/
stuff.url = (options?: RouteQueryOptions) => {
    return stuff.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::stuff
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:163
* @route '/api/v1/year-statistics/stuff'
*/
stuff.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: stuff.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::stuff
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:163
* @route '/api/v1/year-statistics/stuff'
*/
stuff.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: stuff.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::stuff
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:163
* @route '/api/v1/year-statistics/stuff'
*/
const stuffForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: stuff.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::stuff
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:163
* @route '/api/v1/year-statistics/stuff'
*/
stuffForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: stuff.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::stuff
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:163
* @route '/api/v1/year-statistics/stuff'
*/
stuffForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: stuff.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

stuff.form = stuffForm

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::statistics
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:271
* @route '/api/v1/year-statistics/statistics'
*/
export const statistics = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: statistics.url(options),
    method: 'get',
})

statistics.definition = {
    methods: ["get","head"],
    url: '/api/v1/year-statistics/statistics',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::statistics
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:271
* @route '/api/v1/year-statistics/statistics'
*/
statistics.url = (options?: RouteQueryOptions) => {
    return statistics.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::statistics
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:271
* @route '/api/v1/year-statistics/statistics'
*/
statistics.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: statistics.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::statistics
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:271
* @route '/api/v1/year-statistics/statistics'
*/
statistics.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: statistics.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::statistics
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:271
* @route '/api/v1/year-statistics/statistics'
*/
const statisticsForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: statistics.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::statistics
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:271
* @route '/api/v1/year-statistics/statistics'
*/
statisticsForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: statistics.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::statistics
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:271
* @route '/api/v1/year-statistics/statistics'
*/
statisticsForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: statistics.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

statistics.form = statisticsForm

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::show
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:54
* @route '/api/v1/year-statistics/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/year-statistics/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::show
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:54
* @route '/api/v1/year-statistics/{id}'
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
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::show
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:54
* @route '/api/v1/year-statistics/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::show
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:54
* @route '/api/v1/year-statistics/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::show
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:54
* @route '/api/v1/year-statistics/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::show
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:54
* @route '/api/v1/year-statistics/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::show
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:54
* @route '/api/v1/year-statistics/{id}'
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
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::store
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:67
* @route '/api/v1/year-statistics'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/year-statistics',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::store
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:67
* @route '/api/v1/year-statistics'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::store
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:67
* @route '/api/v1/year-statistics'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::store
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:67
* @route '/api/v1/year-statistics'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::store
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:67
* @route '/api/v1/year-statistics'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::update
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:107
* @route '/api/v1/year-statistics/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/year-statistics/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::update
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:107
* @route '/api/v1/year-statistics/{id}'
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
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::update
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:107
* @route '/api/v1/year-statistics/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::update
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:107
* @route '/api/v1/year-statistics/{id}'
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
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::update
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:107
* @route '/api/v1/year-statistics/{id}'
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
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::destroy
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:144
* @route '/api/v1/year-statistics/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/year-statistics/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::destroy
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:144
* @route '/api/v1/year-statistics/{id}'
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
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::destroy
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:144
* @route '/api/v1/year-statistics/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::destroy
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:144
* @route '/api/v1/year-statistics/{id}'
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
* @see \App\Http\Controllers\Api\V1\YearStatisticsController::destroy
* @see app/Http/Controllers/Api/V1/YearStatisticsController.php:144
* @route '/api/v1/year-statistics/{id}'
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

const YearStatisticsController = { index, stuff, statistics, show, store, update, destroy }

export default YearStatisticsController