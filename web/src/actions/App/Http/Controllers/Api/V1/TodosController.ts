import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\TodosController::index
* @see app/Http/Controllers/Api/V1/TodosController.php:11
* @route '/api/v1/todos'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/todos',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\TodosController::index
* @see app/Http/Controllers/Api/V1/TodosController.php:11
* @route '/api/v1/todos'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\TodosController::index
* @see app/Http/Controllers/Api/V1/TodosController.php:11
* @route '/api/v1/todos'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\TodosController::index
* @see app/Http/Controllers/Api/V1/TodosController.php:11
* @route '/api/v1/todos'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\TodosController::index
* @see app/Http/Controllers/Api/V1/TodosController.php:11
* @route '/api/v1/todos'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\TodosController::index
* @see app/Http/Controllers/Api/V1/TodosController.php:11
* @route '/api/v1/todos'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\TodosController::index
* @see app/Http/Controllers/Api/V1/TodosController.php:11
* @route '/api/v1/todos'
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
* @see \App\Http\Controllers\Api\V1\TodosController::listTable
* @see app/Http/Controllers/Api/V1/TodosController.php:22
* @route '/api/v1/todos-list-table'
*/
export const listTable = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: listTable.url(options),
    method: 'get',
})

listTable.definition = {
    methods: ["get","head"],
    url: '/api/v1/todos-list-table',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\TodosController::listTable
* @see app/Http/Controllers/Api/V1/TodosController.php:22
* @route '/api/v1/todos-list-table'
*/
listTable.url = (options?: RouteQueryOptions) => {
    return listTable.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\TodosController::listTable
* @see app/Http/Controllers/Api/V1/TodosController.php:22
* @route '/api/v1/todos-list-table'
*/
listTable.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: listTable.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\TodosController::listTable
* @see app/Http/Controllers/Api/V1/TodosController.php:22
* @route '/api/v1/todos-list-table'
*/
listTable.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: listTable.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\TodosController::listTable
* @see app/Http/Controllers/Api/V1/TodosController.php:22
* @route '/api/v1/todos-list-table'
*/
const listTableForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: listTable.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\TodosController::listTable
* @see app/Http/Controllers/Api/V1/TodosController.php:22
* @route '/api/v1/todos-list-table'
*/
listTableForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: listTable.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\TodosController::listTable
* @see app/Http/Controllers/Api/V1/TodosController.php:22
* @route '/api/v1/todos-list-table'
*/
listTableForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: listTable.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

listTable.form = listTableForm

const TodosController = { index, listTable }

export default TodosController