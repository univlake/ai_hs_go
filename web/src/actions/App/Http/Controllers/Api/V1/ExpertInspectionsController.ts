import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition, applyUrlDefaults } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::index
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:23
* @route '/api/v1/expert-inspections'
*/
export const index = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

index.definition = {
    methods: ["get","head"],
    url: '/api/v1/expert-inspections',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::index
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:23
* @route '/api/v1/expert-inspections'
*/
index.url = (options?: RouteQueryOptions) => {
    return index.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::index
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:23
* @route '/api/v1/expert-inspections'
*/
index.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::index
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:23
* @route '/api/v1/expert-inspections'
*/
index.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: index.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::index
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:23
* @route '/api/v1/expert-inspections'
*/
const indexForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::index
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:23
* @route '/api/v1/expert-inspections'
*/
indexForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: index.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::index
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:23
* @route '/api/v1/expert-inspections'
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
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::show
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:64
* @route '/api/v1/expert-inspections/{id}'
*/
export const show = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

show.definition = {
    methods: ["get","head"],
    url: '/api/v1/expert-inspections/{id}',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::show
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:64
* @route '/api/v1/expert-inspections/{id}'
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
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::show
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:64
* @route '/api/v1/expert-inspections/{id}'
*/
show.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::show
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:64
* @route '/api/v1/expert-inspections/{id}'
*/
show.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: show.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::show
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:64
* @route '/api/v1/expert-inspections/{id}'
*/
const showForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::show
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:64
* @route '/api/v1/expert-inspections/{id}'
*/
showForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: show.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::show
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:64
* @route '/api/v1/expert-inspections/{id}'
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
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::companies
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:120
* @route '/api/v1/expert-inspection-companies'
*/
export const companies = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: companies.url(options),
    method: 'get',
})

companies.definition = {
    methods: ["get","head"],
    url: '/api/v1/expert-inspection-companies',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::companies
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:120
* @route '/api/v1/expert-inspection-companies'
*/
companies.url = (options?: RouteQueryOptions) => {
    return companies.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::companies
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:120
* @route '/api/v1/expert-inspection-companies'
*/
companies.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: companies.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::companies
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:120
* @route '/api/v1/expert-inspection-companies'
*/
companies.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: companies.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::companies
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:120
* @route '/api/v1/expert-inspection-companies'
*/
const companiesForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: companies.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::companies
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:120
* @route '/api/v1/expert-inspection-companies'
*/
companiesForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: companies.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::companies
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:120
* @route '/api/v1/expert-inspection-companies'
*/
companiesForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: companies.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

companies.form = companiesForm

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::users
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:280
* @route '/api/v1/expert-inspection-users'
*/
export const users = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: users.url(options),
    method: 'get',
})

users.definition = {
    methods: ["get","head"],
    url: '/api/v1/expert-inspection-users',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::users
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:280
* @route '/api/v1/expert-inspection-users'
*/
users.url = (options?: RouteQueryOptions) => {
    return users.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::users
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:280
* @route '/api/v1/expert-inspection-users'
*/
users.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: users.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::users
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:280
* @route '/api/v1/expert-inspection-users'
*/
users.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: users.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::users
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:280
* @route '/api/v1/expert-inspection-users'
*/
const usersForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: users.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::users
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:280
* @route '/api/v1/expert-inspection-users'
*/
usersForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: users.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::users
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:280
* @route '/api/v1/expert-inspection-users'
*/
usersForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: users.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

users.form = usersForm

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::store
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:80
* @route '/api/v1/expert-inspections'
*/
export const store = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

store.definition = {
    methods: ["post"],
    url: '/api/v1/expert-inspections',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::store
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:80
* @route '/api/v1/expert-inspections'
*/
store.url = (options?: RouteQueryOptions) => {
    return store.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::store
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:80
* @route '/api/v1/expert-inspections'
*/
store.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::store
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:80
* @route '/api/v1/expert-inspections'
*/
const storeForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::store
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:80
* @route '/api/v1/expert-inspections'
*/
storeForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: store.url(options),
    method: 'post',
})

store.form = storeForm

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::storeCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:134
* @route '/api/v1/expert-inspection-companies'
*/
export const storeCompany = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: storeCompany.url(options),
    method: 'post',
})

storeCompany.definition = {
    methods: ["post"],
    url: '/api/v1/expert-inspection-companies',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::storeCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:134
* @route '/api/v1/expert-inspection-companies'
*/
storeCompany.url = (options?: RouteQueryOptions) => {
    return storeCompany.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::storeCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:134
* @route '/api/v1/expert-inspection-companies'
*/
storeCompany.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: storeCompany.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::storeCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:134
* @route '/api/v1/expert-inspection-companies'
*/
const storeCompanyForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: storeCompany.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::storeCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:134
* @route '/api/v1/expert-inspection-companies'
*/
storeCompanyForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: storeCompany.url(options),
    method: 'post',
})

storeCompany.form = storeCompanyForm

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::storeUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:293
* @route '/api/v1/expert-inspection-users'
*/
export const storeUser = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: storeUser.url(options),
    method: 'post',
})

storeUser.definition = {
    methods: ["post"],
    url: '/api/v1/expert-inspection-users',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::storeUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:293
* @route '/api/v1/expert-inspection-users'
*/
storeUser.url = (options?: RouteQueryOptions) => {
    return storeUser.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::storeUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:293
* @route '/api/v1/expert-inspection-users'
*/
storeUser.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: storeUser.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::storeUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:293
* @route '/api/v1/expert-inspection-users'
*/
const storeUserForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: storeUser.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::storeUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:293
* @route '/api/v1/expert-inspection-users'
*/
storeUserForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: storeUser.url(options),
    method: 'post',
})

storeUser.form = storeUserForm

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::update
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:91
* @route '/api/v1/expert-inspections/{id}'
*/
export const update = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

update.definition = {
    methods: ["put"],
    url: '/api/v1/expert-inspections/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::update
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:91
* @route '/api/v1/expert-inspections/{id}'
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
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::update
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:91
* @route '/api/v1/expert-inspections/{id}'
*/
update.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: update.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::update
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:91
* @route '/api/v1/expert-inspections/{id}'
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
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::update
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:91
* @route '/api/v1/expert-inspections/{id}'
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
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::updateCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:146
* @route '/api/v1/expert-inspection-companies/{id}'
*/
export const updateCompany = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: updateCompany.url(args, options),
    method: 'put',
})

updateCompany.definition = {
    methods: ["put"],
    url: '/api/v1/expert-inspection-companies/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::updateCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:146
* @route '/api/v1/expert-inspection-companies/{id}'
*/
updateCompany.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return updateCompany.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::updateCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:146
* @route '/api/v1/expert-inspection-companies/{id}'
*/
updateCompany.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: updateCompany.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::updateCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:146
* @route '/api/v1/expert-inspection-companies/{id}'
*/
const updateCompanyForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: updateCompany.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'PUT',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::updateCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:146
* @route '/api/v1/expert-inspection-companies/{id}'
*/
updateCompanyForm.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: updateCompany.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'PUT',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

updateCompany.form = updateCompanyForm

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::updateUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:305
* @route '/api/v1/expert-inspection-users/{id}'
*/
export const updateUser = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: updateUser.url(args, options),
    method: 'put',
})

updateUser.definition = {
    methods: ["put"],
    url: '/api/v1/expert-inspection-users/{id}',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::updateUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:305
* @route '/api/v1/expert-inspection-users/{id}'
*/
updateUser.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return updateUser.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::updateUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:305
* @route '/api/v1/expert-inspection-users/{id}'
*/
updateUser.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: updateUser.url(args, options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::updateUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:305
* @route '/api/v1/expert-inspection-users/{id}'
*/
const updateUserForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: updateUser.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'PUT',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::updateUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:305
* @route '/api/v1/expert-inspection-users/{id}'
*/
updateUserForm.put = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: updateUser.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'PUT',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

updateUser.form = updateUserForm

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::destroy
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:105
* @route '/api/v1/expert-inspections/{id}'
*/
export const destroy = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

destroy.definition = {
    methods: ["delete"],
    url: '/api/v1/expert-inspections/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::destroy
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:105
* @route '/api/v1/expert-inspections/{id}'
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
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::destroy
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:105
* @route '/api/v1/expert-inspections/{id}'
*/
destroy.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroy.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::destroy
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:105
* @route '/api/v1/expert-inspections/{id}'
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
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::destroy
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:105
* @route '/api/v1/expert-inspections/{id}'
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
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::destroyCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:162
* @route '/api/v1/expert-inspection-companies/{id}'
*/
export const destroyCompany = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroyCompany.url(args, options),
    method: 'delete',
})

destroyCompany.definition = {
    methods: ["delete"],
    url: '/api/v1/expert-inspection-companies/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::destroyCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:162
* @route '/api/v1/expert-inspection-companies/{id}'
*/
destroyCompany.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return destroyCompany.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::destroyCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:162
* @route '/api/v1/expert-inspection-companies/{id}'
*/
destroyCompany.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroyCompany.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::destroyCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:162
* @route '/api/v1/expert-inspection-companies/{id}'
*/
const destroyCompanyForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: destroyCompany.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'DELETE',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::destroyCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:162
* @route '/api/v1/expert-inspection-companies/{id}'
*/
destroyCompanyForm.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: destroyCompany.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'DELETE',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

destroyCompany.form = destroyCompanyForm

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::destroyUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:321
* @route '/api/v1/expert-inspection-users/{id}'
*/
export const destroyUser = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroyUser.url(args, options),
    method: 'delete',
})

destroyUser.definition = {
    methods: ["delete"],
    url: '/api/v1/expert-inspection-users/{id}',
} satisfies RouteDefinition<["delete"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::destroyUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:321
* @route '/api/v1/expert-inspection-users/{id}'
*/
destroyUser.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return destroyUser.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::destroyUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:321
* @route '/api/v1/expert-inspection-users/{id}'
*/
destroyUser.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'delete'> => ({
    url: destroyUser.url(args, options),
    method: 'delete',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::destroyUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:321
* @route '/api/v1/expert-inspection-users/{id}'
*/
const destroyUserForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: destroyUser.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'DELETE',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::destroyUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:321
* @route '/api/v1/expert-inspection-users/{id}'
*/
destroyUserForm.delete = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: destroyUser.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'DELETE',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

destroyUser.form = destroyUserForm

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::companyNotifications
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:198
* @route '/api/v1/expert-inspection-notifications/company'
*/
export const companyNotifications = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: companyNotifications.url(options),
    method: 'get',
})

companyNotifications.definition = {
    methods: ["get","head"],
    url: '/api/v1/expert-inspection-notifications/company',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::companyNotifications
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:198
* @route '/api/v1/expert-inspection-notifications/company'
*/
companyNotifications.url = (options?: RouteQueryOptions) => {
    return companyNotifications.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::companyNotifications
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:198
* @route '/api/v1/expert-inspection-notifications/company'
*/
companyNotifications.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: companyNotifications.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::companyNotifications
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:198
* @route '/api/v1/expert-inspection-notifications/company'
*/
companyNotifications.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: companyNotifications.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::companyNotifications
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:198
* @route '/api/v1/expert-inspection-notifications/company'
*/
const companyNotificationsForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: companyNotifications.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::companyNotifications
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:198
* @route '/api/v1/expert-inspection-notifications/company'
*/
companyNotificationsForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: companyNotifications.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::companyNotifications
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:198
* @route '/api/v1/expert-inspection-notifications/company'
*/
companyNotificationsForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: companyNotifications.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

companyNotifications.form = companyNotificationsForm

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::expertNotifications
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:354
* @route '/api/v1/expert-inspection-notifications/expert'
*/
export const expertNotifications = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: expertNotifications.url(options),
    method: 'get',
})

expertNotifications.definition = {
    methods: ["get","head"],
    url: '/api/v1/expert-inspection-notifications/expert',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::expertNotifications
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:354
* @route '/api/v1/expert-inspection-notifications/expert'
*/
expertNotifications.url = (options?: RouteQueryOptions) => {
    return expertNotifications.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::expertNotifications
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:354
* @route '/api/v1/expert-inspection-notifications/expert'
*/
expertNotifications.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: expertNotifications.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::expertNotifications
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:354
* @route '/api/v1/expert-inspection-notifications/expert'
*/
expertNotifications.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: expertNotifications.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::expertNotifications
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:354
* @route '/api/v1/expert-inspection-notifications/expert'
*/
const expertNotificationsForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: expertNotifications.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::expertNotifications
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:354
* @route '/api/v1/expert-inspection-notifications/expert'
*/
expertNotificationsForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: expertNotifications.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::expertNotifications
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:354
* @route '/api/v1/expert-inspection-notifications/expert'
*/
expertNotificationsForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: expertNotifications.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

expertNotifications.form = expertNotificationsForm

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::myCheckInspections
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:223
* @route '/api/v1/my-check-inspections'
*/
export const myCheckInspections = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: myCheckInspections.url(options),
    method: 'get',
})

myCheckInspections.definition = {
    methods: ["get","head"],
    url: '/api/v1/my-check-inspections',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::myCheckInspections
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:223
* @route '/api/v1/my-check-inspections'
*/
myCheckInspections.url = (options?: RouteQueryOptions) => {
    return myCheckInspections.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::myCheckInspections
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:223
* @route '/api/v1/my-check-inspections'
*/
myCheckInspections.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: myCheckInspections.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::myCheckInspections
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:223
* @route '/api/v1/my-check-inspections'
*/
myCheckInspections.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: myCheckInspections.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::myCheckInspections
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:223
* @route '/api/v1/my-check-inspections'
*/
const myCheckInspectionsForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: myCheckInspections.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::myCheckInspections
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:223
* @route '/api/v1/my-check-inspections'
*/
myCheckInspectionsForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: myCheckInspections.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::myCheckInspections
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:223
* @route '/api/v1/my-check-inspections'
*/
myCheckInspectionsForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: myCheckInspections.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

myCheckInspections.form = myCheckInspectionsForm

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::myExpertAssignments
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:258
* @route '/api/v1/my-expert-assignments'
*/
export const myExpertAssignments = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: myExpertAssignments.url(options),
    method: 'get',
})

myExpertAssignments.definition = {
    methods: ["get","head"],
    url: '/api/v1/my-expert-assignments',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::myExpertAssignments
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:258
* @route '/api/v1/my-expert-assignments'
*/
myExpertAssignments.url = (options?: RouteQueryOptions) => {
    return myExpertAssignments.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::myExpertAssignments
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:258
* @route '/api/v1/my-expert-assignments'
*/
myExpertAssignments.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: myExpertAssignments.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::myExpertAssignments
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:258
* @route '/api/v1/my-expert-assignments'
*/
myExpertAssignments.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: myExpertAssignments.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::myExpertAssignments
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:258
* @route '/api/v1/my-expert-assignments'
*/
const myExpertAssignmentsForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: myExpertAssignments.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::myExpertAssignments
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:258
* @route '/api/v1/my-expert-assignments'
*/
myExpertAssignmentsForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: myExpertAssignments.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::myExpertAssignments
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:258
* @route '/api/v1/my-expert-assignments'
*/
myExpertAssignmentsForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: myExpertAssignments.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

myExpertAssignments.form = myExpertAssignmentsForm

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::confirmCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:175
* @route '/api/v1/expert-inspection-companies/{id}/confirm'
*/
export const confirmCompany = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: confirmCompany.url(args, options),
    method: 'post',
})

confirmCompany.definition = {
    methods: ["post"],
    url: '/api/v1/expert-inspection-companies/{id}/confirm',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::confirmCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:175
* @route '/api/v1/expert-inspection-companies/{id}/confirm'
*/
confirmCompany.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return confirmCompany.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::confirmCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:175
* @route '/api/v1/expert-inspection-companies/{id}/confirm'
*/
confirmCompany.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: confirmCompany.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::confirmCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:175
* @route '/api/v1/expert-inspection-companies/{id}/confirm'
*/
const confirmCompanyForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: confirmCompany.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::confirmCompany
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:175
* @route '/api/v1/expert-inspection-companies/{id}/confirm'
*/
confirmCompanyForm.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: confirmCompany.url(args, options),
    method: 'post',
})

confirmCompany.form = confirmCompanyForm

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::confirmUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:334
* @route '/api/v1/expert-inspection-users/{id}/confirm'
*/
export const confirmUser = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: confirmUser.url(args, options),
    method: 'post',
})

confirmUser.definition = {
    methods: ["post"],
    url: '/api/v1/expert-inspection-users/{id}/confirm',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::confirmUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:334
* @route '/api/v1/expert-inspection-users/{id}/confirm'
*/
confirmUser.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return confirmUser.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::confirmUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:334
* @route '/api/v1/expert-inspection-users/{id}/confirm'
*/
confirmUser.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: confirmUser.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::confirmUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:334
* @route '/api/v1/expert-inspection-users/{id}/confirm'
*/
const confirmUserForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: confirmUser.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::confirmUser
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:334
* @route '/api/v1/expert-inspection-users/{id}/confirm'
*/
confirmUserForm.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: confirmUser.url(args, options),
    method: 'post',
})

confirmUser.form = confirmUserForm

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::generateInspectionNotice
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:383
* @route '/api/v1/expert-inspection-companies/{id}/inspection-notice'
*/
export const generateInspectionNotice = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: generateInspectionNotice.url(args, options),
    method: 'post',
})

generateInspectionNotice.definition = {
    methods: ["post"],
    url: '/api/v1/expert-inspection-companies/{id}/inspection-notice',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::generateInspectionNotice
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:383
* @route '/api/v1/expert-inspection-companies/{id}/inspection-notice'
*/
generateInspectionNotice.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return generateInspectionNotice.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::generateInspectionNotice
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:383
* @route '/api/v1/expert-inspection-companies/{id}/inspection-notice'
*/
generateInspectionNotice.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: generateInspectionNotice.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::generateInspectionNotice
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:383
* @route '/api/v1/expert-inspection-companies/{id}/inspection-notice'
*/
const generateInspectionNoticeForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: generateInspectionNotice.url(args, options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::generateInspectionNotice
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:383
* @route '/api/v1/expert-inspection-companies/{id}/inspection-notice'
*/
generateInspectionNoticeForm.post = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: generateInspectionNotice.url(args, options),
    method: 'post',
})

generateInspectionNotice.form = generateInspectionNoticeForm

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::downloadInspectionNotice
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:444
* @route '/api/v1/expert-inspection-companies/{id}/inspection-notice'
*/
export const downloadInspectionNotice = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: downloadInspectionNotice.url(args, options),
    method: 'get',
})

downloadInspectionNotice.definition = {
    methods: ["get","head"],
    url: '/api/v1/expert-inspection-companies/{id}/inspection-notice',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::downloadInspectionNotice
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:444
* @route '/api/v1/expert-inspection-companies/{id}/inspection-notice'
*/
downloadInspectionNotice.url = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions) => {
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

    return downloadInspectionNotice.definition.url
            .replace('{id}', parsedArgs.id.toString())
            .replace(/\/+$/, '') + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::downloadInspectionNotice
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:444
* @route '/api/v1/expert-inspection-companies/{id}/inspection-notice'
*/
downloadInspectionNotice.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: downloadInspectionNotice.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::downloadInspectionNotice
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:444
* @route '/api/v1/expert-inspection-companies/{id}/inspection-notice'
*/
downloadInspectionNotice.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: downloadInspectionNotice.url(args, options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::downloadInspectionNotice
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:444
* @route '/api/v1/expert-inspection-companies/{id}/inspection-notice'
*/
const downloadInspectionNoticeForm = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: downloadInspectionNotice.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::downloadInspectionNotice
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:444
* @route '/api/v1/expert-inspection-companies/{id}/inspection-notice'
*/
downloadInspectionNoticeForm.get = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: downloadInspectionNotice.url(args, options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\ExpertInspectionsController::downloadInspectionNotice
* @see app/Http/Controllers/Api/V1/ExpertInspectionsController.php:444
* @route '/api/v1/expert-inspection-companies/{id}/inspection-notice'
*/
downloadInspectionNoticeForm.head = (args: { id: string | number } | [id: string | number ] | string | number, options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: downloadInspectionNotice.url(args, {
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

downloadInspectionNotice.form = downloadInspectionNoticeForm

const ExpertInspectionsController = { index, show, companies, users, store, storeCompany, storeUser, update, updateCompany, updateUser, destroy, destroyCompany, destroyUser, companyNotifications, expertNotifications, myCheckInspections, myExpertAssignments, confirmCompany, confirmUser, generateInspectionNotice, downloadInspectionNotice }

export default ExpertInspectionsController