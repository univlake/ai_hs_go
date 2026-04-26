import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../wayfinder'
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/'
*/
const Controller980bb49ee7ae63891f1d891d2fbcf1c9 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller980bb49ee7ae63891f1d891d2fbcf1c9.url(options),
    method: 'get',
})

Controller980bb49ee7ae63891f1d891d2fbcf1c9.definition = {
    methods: ["get","head"],
    url: '/',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/'
*/
Controller980bb49ee7ae63891f1d891d2fbcf1c9.url = (options?: RouteQueryOptions) => {
    return Controller980bb49ee7ae63891f1d891d2fbcf1c9.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/'
*/
Controller980bb49ee7ae63891f1d891d2fbcf1c9.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller980bb49ee7ae63891f1d891d2fbcf1c9.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/'
*/
Controller980bb49ee7ae63891f1d891d2fbcf1c9.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller980bb49ee7ae63891f1d891d2fbcf1c9.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/'
*/
const Controller980bb49ee7ae63891f1d891d2fbcf1c9Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller980bb49ee7ae63891f1d891d2fbcf1c9.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/'
*/
Controller980bb49ee7ae63891f1d891d2fbcf1c9Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller980bb49ee7ae63891f1d891d2fbcf1c9.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/'
*/
Controller980bb49ee7ae63891f1d891d2fbcf1c9Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller980bb49ee7ae63891f1d891d2fbcf1c9.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller980bb49ee7ae63891f1d891d2fbcf1c9.form = Controller980bb49ee7ae63891f1d891d2fbcf1c9Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/login'
*/
const Controllerb6041c76e8e1cd791f8f89d035d48611 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerb6041c76e8e1cd791f8f89d035d48611.url(options),
    method: 'get',
})

Controllerb6041c76e8e1cd791f8f89d035d48611.definition = {
    methods: ["get","head"],
    url: '/login',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/login'
*/
Controllerb6041c76e8e1cd791f8f89d035d48611.url = (options?: RouteQueryOptions) => {
    return Controllerb6041c76e8e1cd791f8f89d035d48611.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/login'
*/
Controllerb6041c76e8e1cd791f8f89d035d48611.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerb6041c76e8e1cd791f8f89d035d48611.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/login'
*/
Controllerb6041c76e8e1cd791f8f89d035d48611.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerb6041c76e8e1cd791f8f89d035d48611.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/login'
*/
const Controllerb6041c76e8e1cd791f8f89d035d48611Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb6041c76e8e1cd791f8f89d035d48611.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/login'
*/
Controllerb6041c76e8e1cd791f8f89d035d48611Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb6041c76e8e1cd791f8f89d035d48611.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/login'
*/
Controllerb6041c76e8e1cd791f8f89d035d48611Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb6041c76e8e1cd791f8f89d035d48611.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerb6041c76e8e1cd791f8f89d035d48611.form = Controllerb6041c76e8e1cd791f8f89d035d48611Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/register'
*/
const Controllere9819db9819a1d19b38dd89a0c4218c4 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllere9819db9819a1d19b38dd89a0c4218c4.url(options),
    method: 'get',
})

Controllere9819db9819a1d19b38dd89a0c4218c4.definition = {
    methods: ["get","head"],
    url: '/register',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/register'
*/
Controllere9819db9819a1d19b38dd89a0c4218c4.url = (options?: RouteQueryOptions) => {
    return Controllere9819db9819a1d19b38dd89a0c4218c4.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/register'
*/
Controllere9819db9819a1d19b38dd89a0c4218c4.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllere9819db9819a1d19b38dd89a0c4218c4.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/register'
*/
Controllere9819db9819a1d19b38dd89a0c4218c4.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllere9819db9819a1d19b38dd89a0c4218c4.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/register'
*/
const Controllere9819db9819a1d19b38dd89a0c4218c4Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllere9819db9819a1d19b38dd89a0c4218c4.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/register'
*/
Controllere9819db9819a1d19b38dd89a0c4218c4Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllere9819db9819a1d19b38dd89a0c4218c4.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/register'
*/
Controllere9819db9819a1d19b38dd89a0c4218c4Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllere9819db9819a1d19b38dd89a0c4218c4.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllere9819db9819a1d19b38dd89a0c4218c4.form = Controllere9819db9819a1d19b38dd89a0c4218c4Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/first-login-setup'
*/
const Controller8faf738c13d56f0246b14ec9e7e25ed8 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller8faf738c13d56f0246b14ec9e7e25ed8.url(options),
    method: 'get',
})

Controller8faf738c13d56f0246b14ec9e7e25ed8.definition = {
    methods: ["get","head"],
    url: '/first-login-setup',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/first-login-setup'
*/
Controller8faf738c13d56f0246b14ec9e7e25ed8.url = (options?: RouteQueryOptions) => {
    return Controller8faf738c13d56f0246b14ec9e7e25ed8.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/first-login-setup'
*/
Controller8faf738c13d56f0246b14ec9e7e25ed8.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller8faf738c13d56f0246b14ec9e7e25ed8.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/first-login-setup'
*/
Controller8faf738c13d56f0246b14ec9e7e25ed8.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller8faf738c13d56f0246b14ec9e7e25ed8.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/first-login-setup'
*/
const Controller8faf738c13d56f0246b14ec9e7e25ed8Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller8faf738c13d56f0246b14ec9e7e25ed8.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/first-login-setup'
*/
Controller8faf738c13d56f0246b14ec9e7e25ed8Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller8faf738c13d56f0246b14ec9e7e25ed8.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/first-login-setup'
*/
Controller8faf738c13d56f0246b14ec9e7e25ed8Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller8faf738c13d56f0246b14ec9e7e25ed8.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller8faf738c13d56f0246b14ec9e7e25ed8.form = Controller8faf738c13d56f0246b14ec9e7e25ed8Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/dashboard'
*/
const Controller42a740574ecbfbac32f8cc353fc32db9 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller42a740574ecbfbac32f8cc353fc32db9.url(options),
    method: 'get',
})

Controller42a740574ecbfbac32f8cc353fc32db9.definition = {
    methods: ["get","head"],
    url: '/dashboard',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/dashboard'
*/
Controller42a740574ecbfbac32f8cc353fc32db9.url = (options?: RouteQueryOptions) => {
    return Controller42a740574ecbfbac32f8cc353fc32db9.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/dashboard'
*/
Controller42a740574ecbfbac32f8cc353fc32db9.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller42a740574ecbfbac32f8cc353fc32db9.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/dashboard'
*/
Controller42a740574ecbfbac32f8cc353fc32db9.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller42a740574ecbfbac32f8cc353fc32db9.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/dashboard'
*/
const Controller42a740574ecbfbac32f8cc353fc32db9Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller42a740574ecbfbac32f8cc353fc32db9.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/dashboard'
*/
Controller42a740574ecbfbac32f8cc353fc32db9Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller42a740574ecbfbac32f8cc353fc32db9.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/dashboard'
*/
Controller42a740574ecbfbac32f8cc353fc32db9Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller42a740574ecbfbac32f8cc353fc32db9.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller42a740574ecbfbac32f8cc353fc32db9.form = Controller42a740574ecbfbac32f8cc353fc32db9Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/profile'
*/
const Controlleraed0d18b6699640aa4d85cb0451e9ce2 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controlleraed0d18b6699640aa4d85cb0451e9ce2.url(options),
    method: 'get',
})

Controlleraed0d18b6699640aa4d85cb0451e9ce2.definition = {
    methods: ["get","head"],
    url: '/profile',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/profile'
*/
Controlleraed0d18b6699640aa4d85cb0451e9ce2.url = (options?: RouteQueryOptions) => {
    return Controlleraed0d18b6699640aa4d85cb0451e9ce2.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/profile'
*/
Controlleraed0d18b6699640aa4d85cb0451e9ce2.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controlleraed0d18b6699640aa4d85cb0451e9ce2.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/profile'
*/
Controlleraed0d18b6699640aa4d85cb0451e9ce2.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controlleraed0d18b6699640aa4d85cb0451e9ce2.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/profile'
*/
const Controlleraed0d18b6699640aa4d85cb0451e9ce2Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controlleraed0d18b6699640aa4d85cb0451e9ce2.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/profile'
*/
Controlleraed0d18b6699640aa4d85cb0451e9ce2Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controlleraed0d18b6699640aa4d85cb0451e9ce2.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/profile'
*/
Controlleraed0d18b6699640aa4d85cb0451e9ce2Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controlleraed0d18b6699640aa4d85cb0451e9ce2.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controlleraed0d18b6699640aa4d85cb0451e9ce2.form = Controlleraed0d18b6699640aa4d85cb0451e9ce2Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/change-password'
*/
const Controllerfdd726fe943ba5011c80134906e72889 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerfdd726fe943ba5011c80134906e72889.url(options),
    method: 'get',
})

Controllerfdd726fe943ba5011c80134906e72889.definition = {
    methods: ["get","head"],
    url: '/change-password',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/change-password'
*/
Controllerfdd726fe943ba5011c80134906e72889.url = (options?: RouteQueryOptions) => {
    return Controllerfdd726fe943ba5011c80134906e72889.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/change-password'
*/
Controllerfdd726fe943ba5011c80134906e72889.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerfdd726fe943ba5011c80134906e72889.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/change-password'
*/
Controllerfdd726fe943ba5011c80134906e72889.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerfdd726fe943ba5011c80134906e72889.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/change-password'
*/
const Controllerfdd726fe943ba5011c80134906e72889Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerfdd726fe943ba5011c80134906e72889.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/change-password'
*/
Controllerfdd726fe943ba5011c80134906e72889Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerfdd726fe943ba5011c80134906e72889.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/change-password'
*/
Controllerfdd726fe943ba5011c80134906e72889Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerfdd726fe943ba5011c80134906e72889.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerfdd726fe943ba5011c80134906e72889.form = Controllerfdd726fe943ba5011c80134906e72889Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/current'
*/
const Controllerd1dbe9162bf030ca8d815a2896726ea8 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerd1dbe9162bf030ca8d815a2896726ea8.url(options),
    method: 'get',
})

Controllerd1dbe9162bf030ca8d815a2896726ea8.definition = {
    methods: ["get","head"],
    url: '/enterprise/current',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/current'
*/
Controllerd1dbe9162bf030ca8d815a2896726ea8.url = (options?: RouteQueryOptions) => {
    return Controllerd1dbe9162bf030ca8d815a2896726ea8.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/current'
*/
Controllerd1dbe9162bf030ca8d815a2896726ea8.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerd1dbe9162bf030ca8d815a2896726ea8.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/current'
*/
Controllerd1dbe9162bf030ca8d815a2896726ea8.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerd1dbe9162bf030ca8d815a2896726ea8.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/current'
*/
const Controllerd1dbe9162bf030ca8d815a2896726ea8Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerd1dbe9162bf030ca8d815a2896726ea8.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/current'
*/
Controllerd1dbe9162bf030ca8d815a2896726ea8Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerd1dbe9162bf030ca8d815a2896726ea8.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/current'
*/
Controllerd1dbe9162bf030ca8d815a2896726ea8Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerd1dbe9162bf030ca8d815a2896726ea8.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerd1dbe9162bf030ca8d815a2896726ea8.form = Controllerd1dbe9162bf030ca8d815a2896726ea8Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/subsidiaries'
*/
const Controllerdc63d3ff2633f051c4e6fafefcc65b36 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerdc63d3ff2633f051c4e6fafefcc65b36.url(options),
    method: 'get',
})

Controllerdc63d3ff2633f051c4e6fafefcc65b36.definition = {
    methods: ["get","head"],
    url: '/enterprise/subsidiaries',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/subsidiaries'
*/
Controllerdc63d3ff2633f051c4e6fafefcc65b36.url = (options?: RouteQueryOptions) => {
    return Controllerdc63d3ff2633f051c4e6fafefcc65b36.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/subsidiaries'
*/
Controllerdc63d3ff2633f051c4e6fafefcc65b36.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerdc63d3ff2633f051c4e6fafefcc65b36.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/subsidiaries'
*/
Controllerdc63d3ff2633f051c4e6fafefcc65b36.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerdc63d3ff2633f051c4e6fafefcc65b36.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/subsidiaries'
*/
const Controllerdc63d3ff2633f051c4e6fafefcc65b36Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerdc63d3ff2633f051c4e6fafefcc65b36.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/subsidiaries'
*/
Controllerdc63d3ff2633f051c4e6fafefcc65b36Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerdc63d3ff2633f051c4e6fafefcc65b36.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/subsidiaries'
*/
Controllerdc63d3ff2633f051c4e6fafefcc65b36Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerdc63d3ff2633f051c4e6fafefcc65b36.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerdc63d3ff2633f051c4e6fafefcc65b36.form = Controllerdc63d3ff2633f051c4e6fafefcc65b36Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/approval'
*/
const Controllere8686b6a496a704bc4ba429ba6e98199 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllere8686b6a496a704bc4ba429ba6e98199.url(options),
    method: 'get',
})

Controllere8686b6a496a704bc4ba429ba6e98199.definition = {
    methods: ["get","head"],
    url: '/enterprise/approval',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/approval'
*/
Controllere8686b6a496a704bc4ba429ba6e98199.url = (options?: RouteQueryOptions) => {
    return Controllere8686b6a496a704bc4ba429ba6e98199.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/approval'
*/
Controllere8686b6a496a704bc4ba429ba6e98199.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllere8686b6a496a704bc4ba429ba6e98199.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/approval'
*/
Controllere8686b6a496a704bc4ba429ba6e98199.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllere8686b6a496a704bc4ba429ba6e98199.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/approval'
*/
const Controllere8686b6a496a704bc4ba429ba6e98199Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllere8686b6a496a704bc4ba429ba6e98199.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/approval'
*/
Controllere8686b6a496a704bc4ba429ba6e98199Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllere8686b6a496a704bc4ba429ba6e98199.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/approval'
*/
Controllere8686b6a496a704bc4ba429ba6e98199Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllere8686b6a496a704bc4ba429ba6e98199.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllere8686b6a496a704bc4ba429ba6e98199.form = Controllere8686b6a496a704bc4ba429ba6e98199Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/grouping'
*/
const Controllerc35191ff34467f1a013ce9a4c98651af = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerc35191ff34467f1a013ce9a4c98651af.url(options),
    method: 'get',
})

Controllerc35191ff34467f1a013ce9a4c98651af.definition = {
    methods: ["get","head"],
    url: '/enterprise/grouping',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/grouping'
*/
Controllerc35191ff34467f1a013ce9a4c98651af.url = (options?: RouteQueryOptions) => {
    return Controllerc35191ff34467f1a013ce9a4c98651af.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/grouping'
*/
Controllerc35191ff34467f1a013ce9a4c98651af.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerc35191ff34467f1a013ce9a4c98651af.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/grouping'
*/
Controllerc35191ff34467f1a013ce9a4c98651af.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerc35191ff34467f1a013ce9a4c98651af.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/grouping'
*/
const Controllerc35191ff34467f1a013ce9a4c98651afForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerc35191ff34467f1a013ce9a4c98651af.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/grouping'
*/
Controllerc35191ff34467f1a013ce9a4c98651afForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerc35191ff34467f1a013ce9a4c98651af.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/enterprise/grouping'
*/
Controllerc35191ff34467f1a013ce9a4c98651afForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerc35191ff34467f1a013ce9a4c98651af.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerc35191ff34467f1a013ce9a4c98651af.form = Controllerc35191ff34467f1a013ce9a4c98651afForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/enterprise-map'
*/
const Controller05aab3f771482e0f736ba647c68e1796 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller05aab3f771482e0f736ba647c68e1796.url(options),
    method: 'get',
})

Controller05aab3f771482e0f736ba647c68e1796.definition = {
    methods: ["get","head"],
    url: '/risk-monitoring/enterprise-map',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/enterprise-map'
*/
Controller05aab3f771482e0f736ba647c68e1796.url = (options?: RouteQueryOptions) => {
    return Controller05aab3f771482e0f736ba647c68e1796.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/enterprise-map'
*/
Controller05aab3f771482e0f736ba647c68e1796.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller05aab3f771482e0f736ba647c68e1796.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/enterprise-map'
*/
Controller05aab3f771482e0f736ba647c68e1796.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller05aab3f771482e0f736ba647c68e1796.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/enterprise-map'
*/
const Controller05aab3f771482e0f736ba647c68e1796Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller05aab3f771482e0f736ba647c68e1796.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/enterprise-map'
*/
Controller05aab3f771482e0f736ba647c68e1796Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller05aab3f771482e0f736ba647c68e1796.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/enterprise-map'
*/
Controller05aab3f771482e0f736ba647c68e1796Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller05aab3f771482e0f736ba647c68e1796.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller05aab3f771482e0f736ba647c68e1796.form = Controller05aab3f771482e0f736ba647c68e1796Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/hazard-map'
*/
const Controller41d0c8523849475797401e3bca6562af = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller41d0c8523849475797401e3bca6562af.url(options),
    method: 'get',
})

Controller41d0c8523849475797401e3bca6562af.definition = {
    methods: ["get","head"],
    url: '/risk-monitoring/hazard-map',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/hazard-map'
*/
Controller41d0c8523849475797401e3bca6562af.url = (options?: RouteQueryOptions) => {
    return Controller41d0c8523849475797401e3bca6562af.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/hazard-map'
*/
Controller41d0c8523849475797401e3bca6562af.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller41d0c8523849475797401e3bca6562af.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/hazard-map'
*/
Controller41d0c8523849475797401e3bca6562af.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller41d0c8523849475797401e3bca6562af.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/hazard-map'
*/
const Controller41d0c8523849475797401e3bca6562afForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller41d0c8523849475797401e3bca6562af.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/hazard-map'
*/
Controller41d0c8523849475797401e3bca6562afForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller41d0c8523849475797401e3bca6562af.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/hazard-map'
*/
Controller41d0c8523849475797401e3bca6562afForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller41d0c8523849475797401e3bca6562af.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller41d0c8523849475797401e3bca6562af.form = Controller41d0c8523849475797401e3bca6562afForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/online-inspection'
*/
const Controller1ff7759eea7e5ed022e6f9862a7742b6 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller1ff7759eea7e5ed022e6f9862a7742b6.url(options),
    method: 'get',
})

Controller1ff7759eea7e5ed022e6f9862a7742b6.definition = {
    methods: ["get","head"],
    url: '/risk-monitoring/online-inspection',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/online-inspection'
*/
Controller1ff7759eea7e5ed022e6f9862a7742b6.url = (options?: RouteQueryOptions) => {
    return Controller1ff7759eea7e5ed022e6f9862a7742b6.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/online-inspection'
*/
Controller1ff7759eea7e5ed022e6f9862a7742b6.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller1ff7759eea7e5ed022e6f9862a7742b6.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/online-inspection'
*/
Controller1ff7759eea7e5ed022e6f9862a7742b6.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller1ff7759eea7e5ed022e6f9862a7742b6.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/online-inspection'
*/
const Controller1ff7759eea7e5ed022e6f9862a7742b6Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller1ff7759eea7e5ed022e6f9862a7742b6.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/online-inspection'
*/
Controller1ff7759eea7e5ed022e6f9862a7742b6Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller1ff7759eea7e5ed022e6f9862a7742b6.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/online-inspection'
*/
Controller1ff7759eea7e5ed022e6f9862a7742b6Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller1ff7759eea7e5ed022e6f9862a7742b6.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller1ff7759eea7e5ed022e6f9862a7742b6.form = Controller1ff7759eea7e5ed022e6f9862a7742b6Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/risk-monitor'
*/
const Controller37f5f03d82b3e02f8541822ce4bb2798 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller37f5f03d82b3e02f8541822ce4bb2798.url(options),
    method: 'get',
})

Controller37f5f03d82b3e02f8541822ce4bb2798.definition = {
    methods: ["get","head"],
    url: '/risk-monitoring/risk-monitor',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/risk-monitor'
*/
Controller37f5f03d82b3e02f8541822ce4bb2798.url = (options?: RouteQueryOptions) => {
    return Controller37f5f03d82b3e02f8541822ce4bb2798.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/risk-monitor'
*/
Controller37f5f03d82b3e02f8541822ce4bb2798.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller37f5f03d82b3e02f8541822ce4bb2798.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/risk-monitor'
*/
Controller37f5f03d82b3e02f8541822ce4bb2798.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller37f5f03d82b3e02f8541822ce4bb2798.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/risk-monitor'
*/
const Controller37f5f03d82b3e02f8541822ce4bb2798Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller37f5f03d82b3e02f8541822ce4bb2798.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/risk-monitor'
*/
Controller37f5f03d82b3e02f8541822ce4bb2798Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller37f5f03d82b3e02f8541822ce4bb2798.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/risk-monitor'
*/
Controller37f5f03d82b3e02f8541822ce4bb2798Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller37f5f03d82b3e02f8541822ce4bb2798.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller37f5f03d82b3e02f8541822ce4bb2798.form = Controller37f5f03d82b3e02f8541822ce4bb2798Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/monthly-inspection'
*/
const Controlleraeeb3a833b5a24551356fc46b501899c = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controlleraeeb3a833b5a24551356fc46b501899c.url(options),
    method: 'get',
})

Controlleraeeb3a833b5a24551356fc46b501899c.definition = {
    methods: ["get","head"],
    url: '/risk-control/monthly-inspection',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/monthly-inspection'
*/
Controlleraeeb3a833b5a24551356fc46b501899c.url = (options?: RouteQueryOptions) => {
    return Controlleraeeb3a833b5a24551356fc46b501899c.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/monthly-inspection'
*/
Controlleraeeb3a833b5a24551356fc46b501899c.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controlleraeeb3a833b5a24551356fc46b501899c.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/monthly-inspection'
*/
Controlleraeeb3a833b5a24551356fc46b501899c.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controlleraeeb3a833b5a24551356fc46b501899c.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/monthly-inspection'
*/
const Controlleraeeb3a833b5a24551356fc46b501899cForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controlleraeeb3a833b5a24551356fc46b501899c.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/monthly-inspection'
*/
Controlleraeeb3a833b5a24551356fc46b501899cForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controlleraeeb3a833b5a24551356fc46b501899c.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/monthly-inspection'
*/
Controlleraeeb3a833b5a24551356fc46b501899cForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controlleraeeb3a833b5a24551356fc46b501899c.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controlleraeeb3a833b5a24551356fc46b501899c.form = Controlleraeeb3a833b5a24551356fc46b501899cForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/expert-inspection'
*/
const Controller7172acbc864e98d79046caf7595a5c87 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller7172acbc864e98d79046caf7595a5c87.url(options),
    method: 'get',
})

Controller7172acbc864e98d79046caf7595a5c87.definition = {
    methods: ["get","head"],
    url: '/risk-control/expert-inspection',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/expert-inspection'
*/
Controller7172acbc864e98d79046caf7595a5c87.url = (options?: RouteQueryOptions) => {
    return Controller7172acbc864e98d79046caf7595a5c87.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/expert-inspection'
*/
Controller7172acbc864e98d79046caf7595a5c87.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller7172acbc864e98d79046caf7595a5c87.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/expert-inspection'
*/
Controller7172acbc864e98d79046caf7595a5c87.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller7172acbc864e98d79046caf7595a5c87.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/expert-inspection'
*/
const Controller7172acbc864e98d79046caf7595a5c87Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller7172acbc864e98d79046caf7595a5c87.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/expert-inspection'
*/
Controller7172acbc864e98d79046caf7595a5c87Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller7172acbc864e98d79046caf7595a5c87.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/expert-inspection'
*/
Controller7172acbc864e98d79046caf7595a5c87Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller7172acbc864e98d79046caf7595a5c87.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller7172acbc864e98d79046caf7595a5c87.form = Controller7172acbc864e98d79046caf7595a5c87Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/companies'
*/
const Controller22317c48b944ba855e883e487b7b892e = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller22317c48b944ba855e883e487b7b892e.url(options),
    method: 'get',
})

Controller22317c48b944ba855e883e487b7b892e.definition = {
    methods: ["get","head"],
    url: '/system/companies',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/companies'
*/
Controller22317c48b944ba855e883e487b7b892e.url = (options?: RouteQueryOptions) => {
    return Controller22317c48b944ba855e883e487b7b892e.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/companies'
*/
Controller22317c48b944ba855e883e487b7b892e.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller22317c48b944ba855e883e487b7b892e.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/companies'
*/
Controller22317c48b944ba855e883e487b7b892e.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller22317c48b944ba855e883e487b7b892e.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/companies'
*/
const Controller22317c48b944ba855e883e487b7b892eForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller22317c48b944ba855e883e487b7b892e.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/companies'
*/
Controller22317c48b944ba855e883e487b7b892eForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller22317c48b944ba855e883e487b7b892e.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/companies'
*/
Controller22317c48b944ba855e883e487b7b892eForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller22317c48b944ba855e883e487b7b892e.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller22317c48b944ba855e883e487b7b892e.form = Controller22317c48b944ba855e883e487b7b892eForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/users'
*/
const Controller418f97c1428914f8efc179504e0cb069 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller418f97c1428914f8efc179504e0cb069.url(options),
    method: 'get',
})

Controller418f97c1428914f8efc179504e0cb069.definition = {
    methods: ["get","head"],
    url: '/system/users',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/users'
*/
Controller418f97c1428914f8efc179504e0cb069.url = (options?: RouteQueryOptions) => {
    return Controller418f97c1428914f8efc179504e0cb069.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/users'
*/
Controller418f97c1428914f8efc179504e0cb069.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller418f97c1428914f8efc179504e0cb069.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/users'
*/
Controller418f97c1428914f8efc179504e0cb069.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller418f97c1428914f8efc179504e0cb069.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/users'
*/
const Controller418f97c1428914f8efc179504e0cb069Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller418f97c1428914f8efc179504e0cb069.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/users'
*/
Controller418f97c1428914f8efc179504e0cb069Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller418f97c1428914f8efc179504e0cb069.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/users'
*/
Controller418f97c1428914f8efc179504e0cb069Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller418f97c1428914f8efc179504e0cb069.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller418f97c1428914f8efc179504e0cb069.form = Controller418f97c1428914f8efc179504e0cb069Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/roles'
*/
const Controller8e4cd5574fbacbae03adf3fbb312404b = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller8e4cd5574fbacbae03adf3fbb312404b.url(options),
    method: 'get',
})

Controller8e4cd5574fbacbae03adf3fbb312404b.definition = {
    methods: ["get","head"],
    url: '/system/roles',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/roles'
*/
Controller8e4cd5574fbacbae03adf3fbb312404b.url = (options?: RouteQueryOptions) => {
    return Controller8e4cd5574fbacbae03adf3fbb312404b.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/roles'
*/
Controller8e4cd5574fbacbae03adf3fbb312404b.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller8e4cd5574fbacbae03adf3fbb312404b.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/roles'
*/
Controller8e4cd5574fbacbae03adf3fbb312404b.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller8e4cd5574fbacbae03adf3fbb312404b.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/roles'
*/
const Controller8e4cd5574fbacbae03adf3fbb312404bForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller8e4cd5574fbacbae03adf3fbb312404b.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/roles'
*/
Controller8e4cd5574fbacbae03adf3fbb312404bForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller8e4cd5574fbacbae03adf3fbb312404b.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/roles'
*/
Controller8e4cd5574fbacbae03adf3fbb312404bForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller8e4cd5574fbacbae03adf3fbb312404b.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller8e4cd5574fbacbae03adf3fbb312404b.form = Controller8e4cd5574fbacbae03adf3fbb312404bForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/permissions'
*/
const Controllera91f2d42251113e168610b33abb3d6a3 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllera91f2d42251113e168610b33abb3d6a3.url(options),
    method: 'get',
})

Controllera91f2d42251113e168610b33abb3d6a3.definition = {
    methods: ["get","head"],
    url: '/system/permissions',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/permissions'
*/
Controllera91f2d42251113e168610b33abb3d6a3.url = (options?: RouteQueryOptions) => {
    return Controllera91f2d42251113e168610b33abb3d6a3.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/permissions'
*/
Controllera91f2d42251113e168610b33abb3d6a3.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllera91f2d42251113e168610b33abb3d6a3.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/permissions'
*/
Controllera91f2d42251113e168610b33abb3d6a3.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllera91f2d42251113e168610b33abb3d6a3.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/permissions'
*/
const Controllera91f2d42251113e168610b33abb3d6a3Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera91f2d42251113e168610b33abb3d6a3.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/permissions'
*/
Controllera91f2d42251113e168610b33abb3d6a3Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera91f2d42251113e168610b33abb3d6a3.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/permissions'
*/
Controllera91f2d42251113e168610b33abb3d6a3Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera91f2d42251113e168610b33abb3d6a3.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllera91f2d42251113e168610b33abb3d6a3.form = Controllera91f2d42251113e168610b33abb3d6a3Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/enterprise-types'
*/
const Controller044adaaf6c398030da2a3d4b1ded58b3 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller044adaaf6c398030da2a3d4b1ded58b3.url(options),
    method: 'get',
})

Controller044adaaf6c398030da2a3d4b1ded58b3.definition = {
    methods: ["get","head"],
    url: '/system/enterprise-types',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/enterprise-types'
*/
Controller044adaaf6c398030da2a3d4b1ded58b3.url = (options?: RouteQueryOptions) => {
    return Controller044adaaf6c398030da2a3d4b1ded58b3.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/enterprise-types'
*/
Controller044adaaf6c398030da2a3d4b1ded58b3.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller044adaaf6c398030da2a3d4b1ded58b3.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/enterprise-types'
*/
Controller044adaaf6c398030da2a3d4b1ded58b3.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller044adaaf6c398030da2a3d4b1ded58b3.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/enterprise-types'
*/
const Controller044adaaf6c398030da2a3d4b1ded58b3Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller044adaaf6c398030da2a3d4b1ded58b3.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/enterprise-types'
*/
Controller044adaaf6c398030da2a3d4b1ded58b3Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller044adaaf6c398030da2a3d4b1ded58b3.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/enterprise-types'
*/
Controller044adaaf6c398030da2a3d4b1ded58b3Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller044adaaf6c398030da2a3d4b1ded58b3.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller044adaaf6c398030da2a3d4b1ded58b3.form = Controller044adaaf6c398030da2a3d4b1ded58b3Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/operational-types'
*/
const Controllere4b2217aeb9c91d75d471461a9e4519a = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllere4b2217aeb9c91d75d471461a9e4519a.url(options),
    method: 'get',
})

Controllere4b2217aeb9c91d75d471461a9e4519a.definition = {
    methods: ["get","head"],
    url: '/system/operational-types',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/operational-types'
*/
Controllere4b2217aeb9c91d75d471461a9e4519a.url = (options?: RouteQueryOptions) => {
    return Controllere4b2217aeb9c91d75d471461a9e4519a.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/operational-types'
*/
Controllere4b2217aeb9c91d75d471461a9e4519a.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllere4b2217aeb9c91d75d471461a9e4519a.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/operational-types'
*/
Controllere4b2217aeb9c91d75d471461a9e4519a.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllere4b2217aeb9c91d75d471461a9e4519a.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/operational-types'
*/
const Controllere4b2217aeb9c91d75d471461a9e4519aForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllere4b2217aeb9c91d75d471461a9e4519a.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/operational-types'
*/
Controllere4b2217aeb9c91d75d471461a9e4519aForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllere4b2217aeb9c91d75d471461a9e4519a.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/operational-types'
*/
Controllere4b2217aeb9c91d75d471461a9e4519aForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllere4b2217aeb9c91d75d471461a9e4519a.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllere4b2217aeb9c91d75d471461a9e4519a.form = Controllere4b2217aeb9c91d75d471461a9e4519aForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/dangerous-craft-types'
*/
const Controllercfa7b1d6544bfa7299202292496bd119 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllercfa7b1d6544bfa7299202292496bd119.url(options),
    method: 'get',
})

Controllercfa7b1d6544bfa7299202292496bd119.definition = {
    methods: ["get","head"],
    url: '/system/dangerous-craft-types',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/dangerous-craft-types'
*/
Controllercfa7b1d6544bfa7299202292496bd119.url = (options?: RouteQueryOptions) => {
    return Controllercfa7b1d6544bfa7299202292496bd119.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/dangerous-craft-types'
*/
Controllercfa7b1d6544bfa7299202292496bd119.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllercfa7b1d6544bfa7299202292496bd119.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/dangerous-craft-types'
*/
Controllercfa7b1d6544bfa7299202292496bd119.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllercfa7b1d6544bfa7299202292496bd119.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/dangerous-craft-types'
*/
const Controllercfa7b1d6544bfa7299202292496bd119Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercfa7b1d6544bfa7299202292496bd119.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/dangerous-craft-types'
*/
Controllercfa7b1d6544bfa7299202292496bd119Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercfa7b1d6544bfa7299202292496bd119.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/dangerous-craft-types'
*/
Controllercfa7b1d6544bfa7299202292496bd119Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercfa7b1d6544bfa7299202292496bd119.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllercfa7b1d6544bfa7299202292496bd119.form = Controllercfa7b1d6544bfa7299202292496bd119Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/special-types'
*/
const Controllerbc273d698d993e09426e4f7b4dced41b = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerbc273d698d993e09426e4f7b4dced41b.url(options),
    method: 'get',
})

Controllerbc273d698d993e09426e4f7b4dced41b.definition = {
    methods: ["get","head"],
    url: '/system/special-types',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/special-types'
*/
Controllerbc273d698d993e09426e4f7b4dced41b.url = (options?: RouteQueryOptions) => {
    return Controllerbc273d698d993e09426e4f7b4dced41b.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/special-types'
*/
Controllerbc273d698d993e09426e4f7b4dced41b.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerbc273d698d993e09426e4f7b4dced41b.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/special-types'
*/
Controllerbc273d698d993e09426e4f7b4dced41b.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerbc273d698d993e09426e4f7b4dced41b.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/special-types'
*/
const Controllerbc273d698d993e09426e4f7b4dced41bForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerbc273d698d993e09426e4f7b4dced41b.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/special-types'
*/
Controllerbc273d698d993e09426e4f7b4dced41bForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerbc273d698d993e09426e4f7b4dced41b.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/special-types'
*/
Controllerbc273d698d993e09426e4f7b4dced41bForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerbc273d698d993e09426e4f7b4dced41b.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerbc273d698d993e09426e4f7b4dced41b.form = Controllerbc273d698d993e09426e4f7b4dced41bForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/process-types'
*/
const Controllercc5be3c6fdf69815d72033596df0ee33 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllercc5be3c6fdf69815d72033596df0ee33.url(options),
    method: 'get',
})

Controllercc5be3c6fdf69815d72033596df0ee33.definition = {
    methods: ["get","head"],
    url: '/system/process-types',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/process-types'
*/
Controllercc5be3c6fdf69815d72033596df0ee33.url = (options?: RouteQueryOptions) => {
    return Controllercc5be3c6fdf69815d72033596df0ee33.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/process-types'
*/
Controllercc5be3c6fdf69815d72033596df0ee33.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllercc5be3c6fdf69815d72033596df0ee33.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/process-types'
*/
Controllercc5be3c6fdf69815d72033596df0ee33.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllercc5be3c6fdf69815d72033596df0ee33.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/process-types'
*/
const Controllercc5be3c6fdf69815d72033596df0ee33Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercc5be3c6fdf69815d72033596df0ee33.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/process-types'
*/
Controllercc5be3c6fdf69815d72033596df0ee33Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercc5be3c6fdf69815d72033596df0ee33.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/process-types'
*/
Controllercc5be3c6fdf69815d72033596df0ee33Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercc5be3c6fdf69815d72033596df0ee33.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllercc5be3c6fdf69815d72033596df0ee33.form = Controllercc5be3c6fdf69815d72033596df0ee33Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/space-work-types'
*/
const Controllere4caab3e61c1fa82f402eba68b119b58 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllere4caab3e61c1fa82f402eba68b119b58.url(options),
    method: 'get',
})

Controllere4caab3e61c1fa82f402eba68b119b58.definition = {
    methods: ["get","head"],
    url: '/system/space-work-types',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/space-work-types'
*/
Controllere4caab3e61c1fa82f402eba68b119b58.url = (options?: RouteQueryOptions) => {
    return Controllere4caab3e61c1fa82f402eba68b119b58.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/space-work-types'
*/
Controllere4caab3e61c1fa82f402eba68b119b58.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllere4caab3e61c1fa82f402eba68b119b58.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/space-work-types'
*/
Controllere4caab3e61c1fa82f402eba68b119b58.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllere4caab3e61c1fa82f402eba68b119b58.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/space-work-types'
*/
const Controllere4caab3e61c1fa82f402eba68b119b58Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllere4caab3e61c1fa82f402eba68b119b58.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/space-work-types'
*/
Controllere4caab3e61c1fa82f402eba68b119b58Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllere4caab3e61c1fa82f402eba68b119b58.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/space-work-types'
*/
Controllere4caab3e61c1fa82f402eba68b119b58Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllere4caab3e61c1fa82f402eba68b119b58.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllere4caab3e61c1fa82f402eba68b119b58.form = Controllere4caab3e61c1fa82f402eba68b119b58Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/safety-management-types'
*/
const Controllerbc930121bc0a562e140d40c3767083fd = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerbc930121bc0a562e140d40c3767083fd.url(options),
    method: 'get',
})

Controllerbc930121bc0a562e140d40c3767083fd.definition = {
    methods: ["get","head"],
    url: '/system/safety-management-types',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/safety-management-types'
*/
Controllerbc930121bc0a562e140d40c3767083fd.url = (options?: RouteQueryOptions) => {
    return Controllerbc930121bc0a562e140d40c3767083fd.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/safety-management-types'
*/
Controllerbc930121bc0a562e140d40c3767083fd.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerbc930121bc0a562e140d40c3767083fd.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/safety-management-types'
*/
Controllerbc930121bc0a562e140d40c3767083fd.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerbc930121bc0a562e140d40c3767083fd.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/safety-management-types'
*/
const Controllerbc930121bc0a562e140d40c3767083fdForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerbc930121bc0a562e140d40c3767083fd.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/safety-management-types'
*/
Controllerbc930121bc0a562e140d40c3767083fdForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerbc930121bc0a562e140d40c3767083fd.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/safety-management-types'
*/
Controllerbc930121bc0a562e140d40c3767083fdForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerbc930121bc0a562e140d40c3767083fd.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerbc930121bc0a562e140d40c3767083fd.form = Controllerbc930121bc0a562e140d40c3767083fdForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/record-types'
*/
const Controller942b9a21421566f7b3f8dd5be36f9fed = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller942b9a21421566f7b3f8dd5be36f9fed.url(options),
    method: 'get',
})

Controller942b9a21421566f7b3f8dd5be36f9fed.definition = {
    methods: ["get","head"],
    url: '/system/record-types',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/record-types'
*/
Controller942b9a21421566f7b3f8dd5be36f9fed.url = (options?: RouteQueryOptions) => {
    return Controller942b9a21421566f7b3f8dd5be36f9fed.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/record-types'
*/
Controller942b9a21421566f7b3f8dd5be36f9fed.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller942b9a21421566f7b3f8dd5be36f9fed.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/record-types'
*/
Controller942b9a21421566f7b3f8dd5be36f9fed.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller942b9a21421566f7b3f8dd5be36f9fed.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/record-types'
*/
const Controller942b9a21421566f7b3f8dd5be36f9fedForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller942b9a21421566f7b3f8dd5be36f9fed.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/record-types'
*/
Controller942b9a21421566f7b3f8dd5be36f9fedForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller942b9a21421566f7b3f8dd5be36f9fed.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/record-types'
*/
Controller942b9a21421566f7b3f8dd5be36f9fedForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller942b9a21421566f7b3f8dd5be36f9fed.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller942b9a21421566f7b3f8dd5be36f9fed.form = Controller942b9a21421566f7b3f8dd5be36f9fedForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/regulation-library-types'
*/
const Controller6923a40b06398712bb2cc26d54b91ba5 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller6923a40b06398712bb2cc26d54b91ba5.url(options),
    method: 'get',
})

Controller6923a40b06398712bb2cc26d54b91ba5.definition = {
    methods: ["get","head"],
    url: '/system/regulation-library-types',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/regulation-library-types'
*/
Controller6923a40b06398712bb2cc26d54b91ba5.url = (options?: RouteQueryOptions) => {
    return Controller6923a40b06398712bb2cc26d54b91ba5.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/regulation-library-types'
*/
Controller6923a40b06398712bb2cc26d54b91ba5.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller6923a40b06398712bb2cc26d54b91ba5.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/regulation-library-types'
*/
Controller6923a40b06398712bb2cc26d54b91ba5.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller6923a40b06398712bb2cc26d54b91ba5.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/regulation-library-types'
*/
const Controller6923a40b06398712bb2cc26d54b91ba5Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller6923a40b06398712bb2cc26d54b91ba5.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/regulation-library-types'
*/
Controller6923a40b06398712bb2cc26d54b91ba5Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller6923a40b06398712bb2cc26d54b91ba5.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/regulation-library-types'
*/
Controller6923a40b06398712bb2cc26d54b91ba5Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller6923a40b06398712bb2cc26d54b91ba5.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller6923a40b06398712bb2cc26d54b91ba5.form = Controller6923a40b06398712bb2cc26d54b91ba5Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/assessment-templates'
*/
const Controller2490d96901b856cc50b47539eeb24dc8 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller2490d96901b856cc50b47539eeb24dc8.url(options),
    method: 'get',
})

Controller2490d96901b856cc50b47539eeb24dc8.definition = {
    methods: ["get","head"],
    url: '/system/assessment-templates',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/assessment-templates'
*/
Controller2490d96901b856cc50b47539eeb24dc8.url = (options?: RouteQueryOptions) => {
    return Controller2490d96901b856cc50b47539eeb24dc8.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/assessment-templates'
*/
Controller2490d96901b856cc50b47539eeb24dc8.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller2490d96901b856cc50b47539eeb24dc8.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/assessment-templates'
*/
Controller2490d96901b856cc50b47539eeb24dc8.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller2490d96901b856cc50b47539eeb24dc8.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/assessment-templates'
*/
const Controller2490d96901b856cc50b47539eeb24dc8Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller2490d96901b856cc50b47539eeb24dc8.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/assessment-templates'
*/
Controller2490d96901b856cc50b47539eeb24dc8Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller2490d96901b856cc50b47539eeb24dc8.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/system/assessment-templates'
*/
Controller2490d96901b856cc50b47539eeb24dc8Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller2490d96901b856cc50b47539eeb24dc8.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller2490d96901b856cc50b47539eeb24dc8.form = Controller2490d96901b856cc50b47539eeb24dc8Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/buildings'
*/
const Controller35c989abf185d02a8446cf8abc38b6ca = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller35c989abf185d02a8446cf8abc38b6ca.url(options),
    method: 'get',
})

Controller35c989abf185d02a8446cf8abc38b6ca.definition = {
    methods: ["get","head"],
    url: '/safety/buildings',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/buildings'
*/
Controller35c989abf185d02a8446cf8abc38b6ca.url = (options?: RouteQueryOptions) => {
    return Controller35c989abf185d02a8446cf8abc38b6ca.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/buildings'
*/
Controller35c989abf185d02a8446cf8abc38b6ca.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller35c989abf185d02a8446cf8abc38b6ca.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/buildings'
*/
Controller35c989abf185d02a8446cf8abc38b6ca.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller35c989abf185d02a8446cf8abc38b6ca.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/buildings'
*/
const Controller35c989abf185d02a8446cf8abc38b6caForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller35c989abf185d02a8446cf8abc38b6ca.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/buildings'
*/
Controller35c989abf185d02a8446cf8abc38b6caForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller35c989abf185d02a8446cf8abc38b6ca.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/buildings'
*/
Controller35c989abf185d02a8446cf8abc38b6caForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller35c989abf185d02a8446cf8abc38b6ca.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller35c989abf185d02a8446cf8abc38b6ca.form = Controller35c989abf185d02a8446cf8abc38b6caForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/job-management'
*/
const Controller715956940b0d8010718e06152a395138 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller715956940b0d8010718e06152a395138.url(options),
    method: 'get',
})

Controller715956940b0d8010718e06152a395138.definition = {
    methods: ["get","head"],
    url: '/safety/job-management',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/job-management'
*/
Controller715956940b0d8010718e06152a395138.url = (options?: RouteQueryOptions) => {
    return Controller715956940b0d8010718e06152a395138.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/job-management'
*/
Controller715956940b0d8010718e06152a395138.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller715956940b0d8010718e06152a395138.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/job-management'
*/
Controller715956940b0d8010718e06152a395138.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller715956940b0d8010718e06152a395138.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/job-management'
*/
const Controller715956940b0d8010718e06152a395138Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller715956940b0d8010718e06152a395138.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/job-management'
*/
Controller715956940b0d8010718e06152a395138Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller715956940b0d8010718e06152a395138.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/job-management'
*/
Controller715956940b0d8010718e06152a395138Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller715956940b0d8010718e06152a395138.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller715956940b0d8010718e06152a395138.form = Controller715956940b0d8010718e06152a395138Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/personnels'
*/
const Controllerec095dcfcb6f8fb5a6563900689f141a = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerec095dcfcb6f8fb5a6563900689f141a.url(options),
    method: 'get',
})

Controllerec095dcfcb6f8fb5a6563900689f141a.definition = {
    methods: ["get","head"],
    url: '/safety/personnels',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/personnels'
*/
Controllerec095dcfcb6f8fb5a6563900689f141a.url = (options?: RouteQueryOptions) => {
    return Controllerec095dcfcb6f8fb5a6563900689f141a.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/personnels'
*/
Controllerec095dcfcb6f8fb5a6563900689f141a.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerec095dcfcb6f8fb5a6563900689f141a.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/personnels'
*/
Controllerec095dcfcb6f8fb5a6563900689f141a.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerec095dcfcb6f8fb5a6563900689f141a.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/personnels'
*/
const Controllerec095dcfcb6f8fb5a6563900689f141aForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerec095dcfcb6f8fb5a6563900689f141a.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/personnels'
*/
Controllerec095dcfcb6f8fb5a6563900689f141aForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerec095dcfcb6f8fb5a6563900689f141a.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/personnels'
*/
Controllerec095dcfcb6f8fb5a6563900689f141aForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerec095dcfcb6f8fb5a6563900689f141a.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerec095dcfcb6f8fb5a6563900689f141a.form = Controllerec095dcfcb6f8fb5a6563900689f141aForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/special-equipment'
*/
const Controller99d56c37532773e1512d1bba444c5af9 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller99d56c37532773e1512d1bba444c5af9.url(options),
    method: 'get',
})

Controller99d56c37532773e1512d1bba444c5af9.definition = {
    methods: ["get","head"],
    url: '/safety/special-equipment',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/special-equipment'
*/
Controller99d56c37532773e1512d1bba444c5af9.url = (options?: RouteQueryOptions) => {
    return Controller99d56c37532773e1512d1bba444c5af9.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/special-equipment'
*/
Controller99d56c37532773e1512d1bba444c5af9.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller99d56c37532773e1512d1bba444c5af9.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/special-equipment'
*/
Controller99d56c37532773e1512d1bba444c5af9.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller99d56c37532773e1512d1bba444c5af9.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/special-equipment'
*/
const Controller99d56c37532773e1512d1bba444c5af9Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller99d56c37532773e1512d1bba444c5af9.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/special-equipment'
*/
Controller99d56c37532773e1512d1bba444c5af9Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller99d56c37532773e1512d1bba444c5af9.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/special-equipment'
*/
Controller99d56c37532773e1512d1bba444c5af9Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller99d56c37532773e1512d1bba444c5af9.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller99d56c37532773e1512d1bba444c5af9.form = Controller99d56c37532773e1512d1bba444c5af9Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/hazardous-chemicals'
*/
const Controller703b76c5ec2c428d8d5d9020cc82dafa = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller703b76c5ec2c428d8d5d9020cc82dafa.url(options),
    method: 'get',
})

Controller703b76c5ec2c428d8d5d9020cc82dafa.definition = {
    methods: ["get","head"],
    url: '/safety/hazardous-chemicals',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/hazardous-chemicals'
*/
Controller703b76c5ec2c428d8d5d9020cc82dafa.url = (options?: RouteQueryOptions) => {
    return Controller703b76c5ec2c428d8d5d9020cc82dafa.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/hazardous-chemicals'
*/
Controller703b76c5ec2c428d8d5d9020cc82dafa.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller703b76c5ec2c428d8d5d9020cc82dafa.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/hazardous-chemicals'
*/
Controller703b76c5ec2c428d8d5d9020cc82dafa.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller703b76c5ec2c428d8d5d9020cc82dafa.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/hazardous-chemicals'
*/
const Controller703b76c5ec2c428d8d5d9020cc82dafaForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller703b76c5ec2c428d8d5d9020cc82dafa.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/hazardous-chemicals'
*/
Controller703b76c5ec2c428d8d5d9020cc82dafaForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller703b76c5ec2c428d8d5d9020cc82dafa.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/hazardous-chemicals'
*/
Controller703b76c5ec2c428d8d5d9020cc82dafaForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller703b76c5ec2c428d8d5d9020cc82dafa.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller703b76c5ec2c428d8d5d9020cc82dafa.form = Controller703b76c5ec2c428d8d5d9020cc82dafaForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/major-hazards'
*/
const Controller403dfddf0437cb6cf052cb1b0209e94d = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller403dfddf0437cb6cf052cb1b0209e94d.url(options),
    method: 'get',
})

Controller403dfddf0437cb6cf052cb1b0209e94d.definition = {
    methods: ["get","head"],
    url: '/safety/major-hazards',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/major-hazards'
*/
Controller403dfddf0437cb6cf052cb1b0209e94d.url = (options?: RouteQueryOptions) => {
    return Controller403dfddf0437cb6cf052cb1b0209e94d.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/major-hazards'
*/
Controller403dfddf0437cb6cf052cb1b0209e94d.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller403dfddf0437cb6cf052cb1b0209e94d.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/major-hazards'
*/
Controller403dfddf0437cb6cf052cb1b0209e94d.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller403dfddf0437cb6cf052cb1b0209e94d.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/major-hazards'
*/
const Controller403dfddf0437cb6cf052cb1b0209e94dForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller403dfddf0437cb6cf052cb1b0209e94d.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/major-hazards'
*/
Controller403dfddf0437cb6cf052cb1b0209e94dForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller403dfddf0437cb6cf052cb1b0209e94d.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/major-hazards'
*/
Controller403dfddf0437cb6cf052cb1b0209e94dForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller403dfddf0437cb6cf052cb1b0209e94d.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller403dfddf0437cb6cf052cb1b0209e94d.form = Controller403dfddf0437cb6cf052cb1b0209e94dForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/fire-safeties'
*/
const Controllera1dc36a9de9ac5e855e8a9dd7cf132dd = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllera1dc36a9de9ac5e855e8a9dd7cf132dd.url(options),
    method: 'get',
})

Controllera1dc36a9de9ac5e855e8a9dd7cf132dd.definition = {
    methods: ["get","head"],
    url: '/safety/fire-safeties',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/fire-safeties'
*/
Controllera1dc36a9de9ac5e855e8a9dd7cf132dd.url = (options?: RouteQueryOptions) => {
    return Controllera1dc36a9de9ac5e855e8a9dd7cf132dd.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/fire-safeties'
*/
Controllera1dc36a9de9ac5e855e8a9dd7cf132dd.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllera1dc36a9de9ac5e855e8a9dd7cf132dd.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/fire-safeties'
*/
Controllera1dc36a9de9ac5e855e8a9dd7cf132dd.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllera1dc36a9de9ac5e855e8a9dd7cf132dd.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/fire-safeties'
*/
const Controllera1dc36a9de9ac5e855e8a9dd7cf132ddForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera1dc36a9de9ac5e855e8a9dd7cf132dd.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/fire-safeties'
*/
Controllera1dc36a9de9ac5e855e8a9dd7cf132ddForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera1dc36a9de9ac5e855e8a9dd7cf132dd.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety/fire-safeties'
*/
Controllera1dc36a9de9ac5e855e8a9dd7cf132ddForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera1dc36a9de9ac5e855e8a9dd7cf132dd.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllera1dc36a9de9ac5e855e8a9dd7cf132dd.form = Controllera1dc36a9de9ac5e855e8a9dd7cf132ddForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/hazards'
*/
const Controller32cf2db60b9a6fb418346afc40b53a78 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller32cf2db60b9a6fb418346afc40b53a78.url(options),
    method: 'get',
})

Controller32cf2db60b9a6fb418346afc40b53a78.definition = {
    methods: ["get","head"],
    url: '/risk/hazards',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/hazards'
*/
Controller32cf2db60b9a6fb418346afc40b53a78.url = (options?: RouteQueryOptions) => {
    return Controller32cf2db60b9a6fb418346afc40b53a78.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/hazards'
*/
Controller32cf2db60b9a6fb418346afc40b53a78.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller32cf2db60b9a6fb418346afc40b53a78.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/hazards'
*/
Controller32cf2db60b9a6fb418346afc40b53a78.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller32cf2db60b9a6fb418346afc40b53a78.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/hazards'
*/
const Controller32cf2db60b9a6fb418346afc40b53a78Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller32cf2db60b9a6fb418346afc40b53a78.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/hazards'
*/
Controller32cf2db60b9a6fb418346afc40b53a78Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller32cf2db60b9a6fb418346afc40b53a78.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/hazards'
*/
Controller32cf2db60b9a6fb418346afc40b53a78Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller32cf2db60b9a6fb418346afc40b53a78.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller32cf2db60b9a6fb418346afc40b53a78.form = Controller32cf2db60b9a6fb418346afc40b53a78Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/risk-logs'
*/
const Controller931dd43bc86e8b5355757ba1bddccc4d = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller931dd43bc86e8b5355757ba1bddccc4d.url(options),
    method: 'get',
})

Controller931dd43bc86e8b5355757ba1bddccc4d.definition = {
    methods: ["get","head"],
    url: '/risk/risk-logs',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/risk-logs'
*/
Controller931dd43bc86e8b5355757ba1bddccc4d.url = (options?: RouteQueryOptions) => {
    return Controller931dd43bc86e8b5355757ba1bddccc4d.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/risk-logs'
*/
Controller931dd43bc86e8b5355757ba1bddccc4d.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller931dd43bc86e8b5355757ba1bddccc4d.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/risk-logs'
*/
Controller931dd43bc86e8b5355757ba1bddccc4d.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller931dd43bc86e8b5355757ba1bddccc4d.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/risk-logs'
*/
const Controller931dd43bc86e8b5355757ba1bddccc4dForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller931dd43bc86e8b5355757ba1bddccc4d.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/risk-logs'
*/
Controller931dd43bc86e8b5355757ba1bddccc4dForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller931dd43bc86e8b5355757ba1bddccc4d.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/risk-logs'
*/
Controller931dd43bc86e8b5355757ba1bddccc4dForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller931dd43bc86e8b5355757ba1bddccc4d.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller931dd43bc86e8b5355757ba1bddccc4d.form = Controller931dd43bc86e8b5355757ba1bddccc4dForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/monthly-investigation'
*/
const Controllera611b44ef0b191664a928e05a42a8ce4 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllera611b44ef0b191664a928e05a42a8ce4.url(options),
    method: 'get',
})

Controllera611b44ef0b191664a928e05a42a8ce4.definition = {
    methods: ["get","head"],
    url: '/risk/monthly-investigation',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/monthly-investigation'
*/
Controllera611b44ef0b191664a928e05a42a8ce4.url = (options?: RouteQueryOptions) => {
    return Controllera611b44ef0b191664a928e05a42a8ce4.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/monthly-investigation'
*/
Controllera611b44ef0b191664a928e05a42a8ce4.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllera611b44ef0b191664a928e05a42a8ce4.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/monthly-investigation'
*/
Controllera611b44ef0b191664a928e05a42a8ce4.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllera611b44ef0b191664a928e05a42a8ce4.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/monthly-investigation'
*/
const Controllera611b44ef0b191664a928e05a42a8ce4Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera611b44ef0b191664a928e05a42a8ce4.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/monthly-investigation'
*/
Controllera611b44ef0b191664a928e05a42a8ce4Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera611b44ef0b191664a928e05a42a8ce4.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/monthly-investigation'
*/
Controllera611b44ef0b191664a928e05a42a8ce4Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera611b44ef0b191664a928e05a42a8ce4.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllera611b44ef0b191664a928e05a42a8ce4.form = Controllera611b44ef0b191664a928e05a42a8ce4Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/rectification-sheets'
*/
const Controllerd566f42701af8ddf67cda28d9ac9db37 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerd566f42701af8ddf67cda28d9ac9db37.url(options),
    method: 'get',
})

Controllerd566f42701af8ddf67cda28d9ac9db37.definition = {
    methods: ["get","head"],
    url: '/risk/rectification-sheets',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/rectification-sheets'
*/
Controllerd566f42701af8ddf67cda28d9ac9db37.url = (options?: RouteQueryOptions) => {
    return Controllerd566f42701af8ddf67cda28d9ac9db37.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/rectification-sheets'
*/
Controllerd566f42701af8ddf67cda28d9ac9db37.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerd566f42701af8ddf67cda28d9ac9db37.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/rectification-sheets'
*/
Controllerd566f42701af8ddf67cda28d9ac9db37.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerd566f42701af8ddf67cda28d9ac9db37.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/rectification-sheets'
*/
const Controllerd566f42701af8ddf67cda28d9ac9db37Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerd566f42701af8ddf67cda28d9ac9db37.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/rectification-sheets'
*/
Controllerd566f42701af8ddf67cda28d9ac9db37Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerd566f42701af8ddf67cda28d9ac9db37.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/rectification-sheets'
*/
Controllerd566f42701af8ddf67cda28d9ac9db37Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerd566f42701af8ddf67cda28d9ac9db37.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerd566f42701af8ddf67cda28d9ac9db37.form = Controllerd566f42701af8ddf67cda28d9ac9db37Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/daily-checks'
*/
const Controller848e8b566bd81867beab2986e6102b1d = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller848e8b566bd81867beab2986e6102b1d.url(options),
    method: 'get',
})

Controller848e8b566bd81867beab2986e6102b1d.definition = {
    methods: ["get","head"],
    url: '/inspection/daily-checks',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/daily-checks'
*/
Controller848e8b566bd81867beab2986e6102b1d.url = (options?: RouteQueryOptions) => {
    return Controller848e8b566bd81867beab2986e6102b1d.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/daily-checks'
*/
Controller848e8b566bd81867beab2986e6102b1d.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller848e8b566bd81867beab2986e6102b1d.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/daily-checks'
*/
Controller848e8b566bd81867beab2986e6102b1d.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller848e8b566bd81867beab2986e6102b1d.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/daily-checks'
*/
const Controller848e8b566bd81867beab2986e6102b1dForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller848e8b566bd81867beab2986e6102b1d.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/daily-checks'
*/
Controller848e8b566bd81867beab2986e6102b1dForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller848e8b566bd81867beab2986e6102b1d.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/daily-checks'
*/
Controller848e8b566bd81867beab2986e6102b1dForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller848e8b566bd81867beab2986e6102b1d.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller848e8b566bd81867beab2986e6102b1d.form = Controller848e8b566bd81867beab2986e6102b1dForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/system-checks'
*/
const Controllera36f412fb20084ab7c9404fa6d8a5f46 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllera36f412fb20084ab7c9404fa6d8a5f46.url(options),
    method: 'get',
})

Controllera36f412fb20084ab7c9404fa6d8a5f46.definition = {
    methods: ["get","head"],
    url: '/inspection/system-checks',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/system-checks'
*/
Controllera36f412fb20084ab7c9404fa6d8a5f46.url = (options?: RouteQueryOptions) => {
    return Controllera36f412fb20084ab7c9404fa6d8a5f46.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/system-checks'
*/
Controllera36f412fb20084ab7c9404fa6d8a5f46.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllera36f412fb20084ab7c9404fa6d8a5f46.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/system-checks'
*/
Controllera36f412fb20084ab7c9404fa6d8a5f46.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllera36f412fb20084ab7c9404fa6d8a5f46.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/system-checks'
*/
const Controllera36f412fb20084ab7c9404fa6d8a5f46Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera36f412fb20084ab7c9404fa6d8a5f46.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/system-checks'
*/
Controllera36f412fb20084ab7c9404fa6d8a5f46Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera36f412fb20084ab7c9404fa6d8a5f46.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/system-checks'
*/
Controllera36f412fb20084ab7c9404fa6d8a5f46Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera36f412fb20084ab7c9404fa6d8a5f46.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllera36f412fb20084ab7c9404fa6d8a5f46.form = Controllera36f412fb20084ab7c9404fa6d8a5f46Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/norm-checks'
*/
const Controllercf195f51fc323080aba270ea267a22a5 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllercf195f51fc323080aba270ea267a22a5.url(options),
    method: 'get',
})

Controllercf195f51fc323080aba270ea267a22a5.definition = {
    methods: ["get","head"],
    url: '/inspection/norm-checks',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/norm-checks'
*/
Controllercf195f51fc323080aba270ea267a22a5.url = (options?: RouteQueryOptions) => {
    return Controllercf195f51fc323080aba270ea267a22a5.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/norm-checks'
*/
Controllercf195f51fc323080aba270ea267a22a5.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllercf195f51fc323080aba270ea267a22a5.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/norm-checks'
*/
Controllercf195f51fc323080aba270ea267a22a5.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllercf195f51fc323080aba270ea267a22a5.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/norm-checks'
*/
const Controllercf195f51fc323080aba270ea267a22a5Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercf195f51fc323080aba270ea267a22a5.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/norm-checks'
*/
Controllercf195f51fc323080aba270ea267a22a5Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercf195f51fc323080aba270ea267a22a5.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/norm-checks'
*/
Controllercf195f51fc323080aba270ea267a22a5Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercf195f51fc323080aba270ea267a22a5.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllercf195f51fc323080aba270ea267a22a5.form = Controllercf195f51fc323080aba270ea267a22a5Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/accident-checks'
*/
const Controllere706152e0dcff487ab0e79404ee6f9c3 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllere706152e0dcff487ab0e79404ee6f9c3.url(options),
    method: 'get',
})

Controllere706152e0dcff487ab0e79404ee6f9c3.definition = {
    methods: ["get","head"],
    url: '/inspection/accident-checks',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/accident-checks'
*/
Controllere706152e0dcff487ab0e79404ee6f9c3.url = (options?: RouteQueryOptions) => {
    return Controllere706152e0dcff487ab0e79404ee6f9c3.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/accident-checks'
*/
Controllere706152e0dcff487ab0e79404ee6f9c3.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllere706152e0dcff487ab0e79404ee6f9c3.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/accident-checks'
*/
Controllere706152e0dcff487ab0e79404ee6f9c3.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllere706152e0dcff487ab0e79404ee6f9c3.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/accident-checks'
*/
const Controllere706152e0dcff487ab0e79404ee6f9c3Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllere706152e0dcff487ab0e79404ee6f9c3.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/accident-checks'
*/
Controllere706152e0dcff487ab0e79404ee6f9c3Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllere706152e0dcff487ab0e79404ee6f9c3.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/inspection/accident-checks'
*/
Controllere706152e0dcff487ab0e79404ee6f9c3Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllere706152e0dcff487ab0e79404ee6f9c3.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllere706152e0dcff487ab0e79404ee6f9c3.form = Controllere706152e0dcff487ab0e79404ee6f9c3Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/occupational-hazards'
*/
const Controllercf17a5673925405b2682957c0642ef97 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllercf17a5673925405b2682957c0642ef97.url(options),
    method: 'get',
})

Controllercf17a5673925405b2682957c0642ef97.definition = {
    methods: ["get","head"],
    url: '/health/occupational-hazards',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/occupational-hazards'
*/
Controllercf17a5673925405b2682957c0642ef97.url = (options?: RouteQueryOptions) => {
    return Controllercf17a5673925405b2682957c0642ef97.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/occupational-hazards'
*/
Controllercf17a5673925405b2682957c0642ef97.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllercf17a5673925405b2682957c0642ef97.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/occupational-hazards'
*/
Controllercf17a5673925405b2682957c0642ef97.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllercf17a5673925405b2682957c0642ef97.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/occupational-hazards'
*/
const Controllercf17a5673925405b2682957c0642ef97Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercf17a5673925405b2682957c0642ef97.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/occupational-hazards'
*/
Controllercf17a5673925405b2682957c0642ef97Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercf17a5673925405b2682957c0642ef97.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/occupational-hazards'
*/
Controllercf17a5673925405b2682957c0642ef97Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercf17a5673925405b2682957c0642ef97.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllercf17a5673925405b2682957c0642ef97.form = Controllercf17a5673925405b2682957c0642ef97Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/health-files'
*/
const Controller00b00e6ebd69280348b8be7b977ff5e2 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller00b00e6ebd69280348b8be7b977ff5e2.url(options),
    method: 'get',
})

Controller00b00e6ebd69280348b8be7b977ff5e2.definition = {
    methods: ["get","head"],
    url: '/health/health-files',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/health-files'
*/
Controller00b00e6ebd69280348b8be7b977ff5e2.url = (options?: RouteQueryOptions) => {
    return Controller00b00e6ebd69280348b8be7b977ff5e2.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/health-files'
*/
Controller00b00e6ebd69280348b8be7b977ff5e2.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller00b00e6ebd69280348b8be7b977ff5e2.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/health-files'
*/
Controller00b00e6ebd69280348b8be7b977ff5e2.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller00b00e6ebd69280348b8be7b977ff5e2.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/health-files'
*/
const Controller00b00e6ebd69280348b8be7b977ff5e2Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller00b00e6ebd69280348b8be7b977ff5e2.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/health-files'
*/
Controller00b00e6ebd69280348b8be7b977ff5e2Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller00b00e6ebd69280348b8be7b977ff5e2.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/health-files'
*/
Controller00b00e6ebd69280348b8be7b977ff5e2Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller00b00e6ebd69280348b8be7b977ff5e2.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller00b00e6ebd69280348b8be7b977ff5e2.form = Controller00b00e6ebd69280348b8be7b977ff5e2Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/health-checkups'
*/
const Controller7a6036087be56a47df63d7013588f45c = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller7a6036087be56a47df63d7013588f45c.url(options),
    method: 'get',
})

Controller7a6036087be56a47df63d7013588f45c.definition = {
    methods: ["get","head"],
    url: '/health/health-checkups',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/health-checkups'
*/
Controller7a6036087be56a47df63d7013588f45c.url = (options?: RouteQueryOptions) => {
    return Controller7a6036087be56a47df63d7013588f45c.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/health-checkups'
*/
Controller7a6036087be56a47df63d7013588f45c.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller7a6036087be56a47df63d7013588f45c.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/health-checkups'
*/
Controller7a6036087be56a47df63d7013588f45c.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller7a6036087be56a47df63d7013588f45c.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/health-checkups'
*/
const Controller7a6036087be56a47df63d7013588f45cForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller7a6036087be56a47df63d7013588f45c.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/health-checkups'
*/
Controller7a6036087be56a47df63d7013588f45cForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller7a6036087be56a47df63d7013588f45c.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/health-checkups'
*/
Controller7a6036087be56a47df63d7013588f45cForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller7a6036087be56a47df63d7013588f45c.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller7a6036087be56a47df63d7013588f45c.form = Controller7a6036087be56a47df63d7013588f45cForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/hygiene-evaluations'
*/
const Controlleraa1f88d6fc5920f38d9e2eeb08ca6216 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controlleraa1f88d6fc5920f38d9e2eeb08ca6216.url(options),
    method: 'get',
})

Controlleraa1f88d6fc5920f38d9e2eeb08ca6216.definition = {
    methods: ["get","head"],
    url: '/health/hygiene-evaluations',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/hygiene-evaluations'
*/
Controlleraa1f88d6fc5920f38d9e2eeb08ca6216.url = (options?: RouteQueryOptions) => {
    return Controlleraa1f88d6fc5920f38d9e2eeb08ca6216.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/hygiene-evaluations'
*/
Controlleraa1f88d6fc5920f38d9e2eeb08ca6216.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controlleraa1f88d6fc5920f38d9e2eeb08ca6216.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/hygiene-evaluations'
*/
Controlleraa1f88d6fc5920f38d9e2eeb08ca6216.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controlleraa1f88d6fc5920f38d9e2eeb08ca6216.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/hygiene-evaluations'
*/
const Controlleraa1f88d6fc5920f38d9e2eeb08ca6216Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controlleraa1f88d6fc5920f38d9e2eeb08ca6216.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/hygiene-evaluations'
*/
Controlleraa1f88d6fc5920f38d9e2eeb08ca6216Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controlleraa1f88d6fc5920f38d9e2eeb08ca6216.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/hygiene-evaluations'
*/
Controlleraa1f88d6fc5920f38d9e2eeb08ca6216Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controlleraa1f88d6fc5920f38d9e2eeb08ca6216.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controlleraa1f88d6fc5920f38d9e2eeb08ca6216.form = Controlleraa1f88d6fc5920f38d9e2eeb08ca6216Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/fire-safety'
*/
const Controllerd40c2199ac32bc35ff00293ec9a60196 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerd40c2199ac32bc35ff00293ec9a60196.url(options),
    method: 'get',
})

Controllerd40c2199ac32bc35ff00293ec9a60196.definition = {
    methods: ["get","head"],
    url: '/emergency/fire-safety',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/fire-safety'
*/
Controllerd40c2199ac32bc35ff00293ec9a60196.url = (options?: RouteQueryOptions) => {
    return Controllerd40c2199ac32bc35ff00293ec9a60196.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/fire-safety'
*/
Controllerd40c2199ac32bc35ff00293ec9a60196.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerd40c2199ac32bc35ff00293ec9a60196.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/fire-safety'
*/
Controllerd40c2199ac32bc35ff00293ec9a60196.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerd40c2199ac32bc35ff00293ec9a60196.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/fire-safety'
*/
const Controllerd40c2199ac32bc35ff00293ec9a60196Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerd40c2199ac32bc35ff00293ec9a60196.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/fire-safety'
*/
Controllerd40c2199ac32bc35ff00293ec9a60196Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerd40c2199ac32bc35ff00293ec9a60196.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/fire-safety'
*/
Controllerd40c2199ac32bc35ff00293ec9a60196Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerd40c2199ac32bc35ff00293ec9a60196.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerd40c2199ac32bc35ff00293ec9a60196.form = Controllerd40c2199ac32bc35ff00293ec9a60196Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-drills'
*/
const Controller96c1bb13a3ad3f9b5edb57c77c96c216 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller96c1bb13a3ad3f9b5edb57c77c96c216.url(options),
    method: 'get',
})

Controller96c1bb13a3ad3f9b5edb57c77c96c216.definition = {
    methods: ["get","head"],
    url: '/emergency/emergency-drills',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-drills'
*/
Controller96c1bb13a3ad3f9b5edb57c77c96c216.url = (options?: RouteQueryOptions) => {
    return Controller96c1bb13a3ad3f9b5edb57c77c96c216.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-drills'
*/
Controller96c1bb13a3ad3f9b5edb57c77c96c216.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller96c1bb13a3ad3f9b5edb57c77c96c216.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-drills'
*/
Controller96c1bb13a3ad3f9b5edb57c77c96c216.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller96c1bb13a3ad3f9b5edb57c77c96c216.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-drills'
*/
const Controller96c1bb13a3ad3f9b5edb57c77c96c216Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller96c1bb13a3ad3f9b5edb57c77c96c216.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-drills'
*/
Controller96c1bb13a3ad3f9b5edb57c77c96c216Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller96c1bb13a3ad3f9b5edb57c77c96c216.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-drills'
*/
Controller96c1bb13a3ad3f9b5edb57c77c96c216Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller96c1bb13a3ad3f9b5edb57c77c96c216.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller96c1bb13a3ad3f9b5edb57c77c96c216.form = Controller96c1bb13a3ad3f9b5edb57c77c96c216Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-supplies'
*/
const Controllerf907fde89fb2af26dda0618330d90e33 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerf907fde89fb2af26dda0618330d90e33.url(options),
    method: 'get',
})

Controllerf907fde89fb2af26dda0618330d90e33.definition = {
    methods: ["get","head"],
    url: '/emergency/emergency-supplies',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-supplies'
*/
Controllerf907fde89fb2af26dda0618330d90e33.url = (options?: RouteQueryOptions) => {
    return Controllerf907fde89fb2af26dda0618330d90e33.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-supplies'
*/
Controllerf907fde89fb2af26dda0618330d90e33.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerf907fde89fb2af26dda0618330d90e33.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-supplies'
*/
Controllerf907fde89fb2af26dda0618330d90e33.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerf907fde89fb2af26dda0618330d90e33.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-supplies'
*/
const Controllerf907fde89fb2af26dda0618330d90e33Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerf907fde89fb2af26dda0618330d90e33.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-supplies'
*/
Controllerf907fde89fb2af26dda0618330d90e33Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerf907fde89fb2af26dda0618330d90e33.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-supplies'
*/
Controllerf907fde89fb2af26dda0618330d90e33Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerf907fde89fb2af26dda0618330d90e33.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerf907fde89fb2af26dda0618330d90e33.form = Controllerf907fde89fb2af26dda0618330d90e33Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-teams'
*/
const Controllerb5fc8dd61ced1b57e86be40c73008108 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerb5fc8dd61ced1b57e86be40c73008108.url(options),
    method: 'get',
})

Controllerb5fc8dd61ced1b57e86be40c73008108.definition = {
    methods: ["get","head"],
    url: '/emergency/emergency-teams',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-teams'
*/
Controllerb5fc8dd61ced1b57e86be40c73008108.url = (options?: RouteQueryOptions) => {
    return Controllerb5fc8dd61ced1b57e86be40c73008108.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-teams'
*/
Controllerb5fc8dd61ced1b57e86be40c73008108.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerb5fc8dd61ced1b57e86be40c73008108.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-teams'
*/
Controllerb5fc8dd61ced1b57e86be40c73008108.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerb5fc8dd61ced1b57e86be40c73008108.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-teams'
*/
const Controllerb5fc8dd61ced1b57e86be40c73008108Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb5fc8dd61ced1b57e86be40c73008108.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-teams'
*/
Controllerb5fc8dd61ced1b57e86be40c73008108Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb5fc8dd61ced1b57e86be40c73008108.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/emergency/emergency-teams'
*/
Controllerb5fc8dd61ced1b57e86be40c73008108Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb5fc8dd61ced1b57e86be40c73008108.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerb5fc8dd61ced1b57e86be40c73008108.form = Controllerb5fc8dd61ced1b57e86be40c73008108Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/training-files'
*/
const Controllera32e7ab4ac8fc9480b96291a84b03264 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllera32e7ab4ac8fc9480b96291a84b03264.url(options),
    method: 'get',
})

Controllera32e7ab4ac8fc9480b96291a84b03264.definition = {
    methods: ["get","head"],
    url: '/training/training-files',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/training-files'
*/
Controllera32e7ab4ac8fc9480b96291a84b03264.url = (options?: RouteQueryOptions) => {
    return Controllera32e7ab4ac8fc9480b96291a84b03264.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/training-files'
*/
Controllera32e7ab4ac8fc9480b96291a84b03264.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllera32e7ab4ac8fc9480b96291a84b03264.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/training-files'
*/
Controllera32e7ab4ac8fc9480b96291a84b03264.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllera32e7ab4ac8fc9480b96291a84b03264.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/training-files'
*/
const Controllera32e7ab4ac8fc9480b96291a84b03264Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera32e7ab4ac8fc9480b96291a84b03264.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/training-files'
*/
Controllera32e7ab4ac8fc9480b96291a84b03264Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera32e7ab4ac8fc9480b96291a84b03264.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/training-files'
*/
Controllera32e7ab4ac8fc9480b96291a84b03264Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera32e7ab4ac8fc9480b96291a84b03264.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllera32e7ab4ac8fc9480b96291a84b03264.form = Controllera32e7ab4ac8fc9480b96291a84b03264Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/training-videos'
*/
const Controller86903e311a24a371314611c417a8e28d = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller86903e311a24a371314611c417a8e28d.url(options),
    method: 'get',
})

Controller86903e311a24a371314611c417a8e28d.definition = {
    methods: ["get","head"],
    url: '/training/training-videos',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/training-videos'
*/
Controller86903e311a24a371314611c417a8e28d.url = (options?: RouteQueryOptions) => {
    return Controller86903e311a24a371314611c417a8e28d.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/training-videos'
*/
Controller86903e311a24a371314611c417a8e28d.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller86903e311a24a371314611c417a8e28d.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/training-videos'
*/
Controller86903e311a24a371314611c417a8e28d.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller86903e311a24a371314611c417a8e28d.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/training-videos'
*/
const Controller86903e311a24a371314611c417a8e28dForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller86903e311a24a371314611c417a8e28d.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/training-videos'
*/
Controller86903e311a24a371314611c417a8e28dForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller86903e311a24a371314611c417a8e28d.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/training-videos'
*/
Controller86903e311a24a371314611c417a8e28dForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller86903e311a24a371314611c417a8e28d.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller86903e311a24a371314611c417a8e28d.form = Controller86903e311a24a371314611c417a8e28dForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/internal-trainings'
*/
const Controllerb0d5689953570e04c957bb3ebe226451 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerb0d5689953570e04c957bb3ebe226451.url(options),
    method: 'get',
})

Controllerb0d5689953570e04c957bb3ebe226451.definition = {
    methods: ["get","head"],
    url: '/training/internal-trainings',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/internal-trainings'
*/
Controllerb0d5689953570e04c957bb3ebe226451.url = (options?: RouteQueryOptions) => {
    return Controllerb0d5689953570e04c957bb3ebe226451.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/internal-trainings'
*/
Controllerb0d5689953570e04c957bb3ebe226451.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerb0d5689953570e04c957bb3ebe226451.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/internal-trainings'
*/
Controllerb0d5689953570e04c957bb3ebe226451.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerb0d5689953570e04c957bb3ebe226451.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/internal-trainings'
*/
const Controllerb0d5689953570e04c957bb3ebe226451Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb0d5689953570e04c957bb3ebe226451.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/internal-trainings'
*/
Controllerb0d5689953570e04c957bb3ebe226451Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb0d5689953570e04c957bb3ebe226451.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/internal-trainings'
*/
Controllerb0d5689953570e04c957bb3ebe226451Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb0d5689953570e04c957bb3ebe226451.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerb0d5689953570e04c957bb3ebe226451.form = Controllerb0d5689953570e04c957bb3ebe226451Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/external-trainings'
*/
const Controllercc02a54a8a581f4b995a2da3573af252 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllercc02a54a8a581f4b995a2da3573af252.url(options),
    method: 'get',
})

Controllercc02a54a8a581f4b995a2da3573af252.definition = {
    methods: ["get","head"],
    url: '/training/external-trainings',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/external-trainings'
*/
Controllercc02a54a8a581f4b995a2da3573af252.url = (options?: RouteQueryOptions) => {
    return Controllercc02a54a8a581f4b995a2da3573af252.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/external-trainings'
*/
Controllercc02a54a8a581f4b995a2da3573af252.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllercc02a54a8a581f4b995a2da3573af252.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/external-trainings'
*/
Controllercc02a54a8a581f4b995a2da3573af252.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllercc02a54a8a581f4b995a2da3573af252.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/external-trainings'
*/
const Controllercc02a54a8a581f4b995a2da3573af252Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercc02a54a8a581f4b995a2da3573af252.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/external-trainings'
*/
Controllercc02a54a8a581f4b995a2da3573af252Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercc02a54a8a581f4b995a2da3573af252.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/external-trainings'
*/
Controllercc02a54a8a581f4b995a2da3573af252Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercc02a54a8a581f4b995a2da3573af252.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllercc02a54a8a581f4b995a2da3573af252.form = Controllercc02a54a8a581f4b995a2da3573af252Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/other-trainings'
*/
const Controller5e6f8aff4f45ec9e5a3d2073e62940fa = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller5e6f8aff4f45ec9e5a3d2073e62940fa.url(options),
    method: 'get',
})

Controller5e6f8aff4f45ec9e5a3d2073e62940fa.definition = {
    methods: ["get","head"],
    url: '/training/other-trainings',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/other-trainings'
*/
Controller5e6f8aff4f45ec9e5a3d2073e62940fa.url = (options?: RouteQueryOptions) => {
    return Controller5e6f8aff4f45ec9e5a3d2073e62940fa.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/other-trainings'
*/
Controller5e6f8aff4f45ec9e5a3d2073e62940fa.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller5e6f8aff4f45ec9e5a3d2073e62940fa.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/other-trainings'
*/
Controller5e6f8aff4f45ec9e5a3d2073e62940fa.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller5e6f8aff4f45ec9e5a3d2073e62940fa.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/other-trainings'
*/
const Controller5e6f8aff4f45ec9e5a3d2073e62940faForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller5e6f8aff4f45ec9e5a3d2073e62940fa.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/other-trainings'
*/
Controller5e6f8aff4f45ec9e5a3d2073e62940faForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller5e6f8aff4f45ec9e5a3d2073e62940fa.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/other-trainings'
*/
Controller5e6f8aff4f45ec9e5a3d2073e62940faForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller5e6f8aff4f45ec9e5a3d2073e62940fa.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller5e6f8aff4f45ec9e5a3d2073e62940fa.form = Controller5e6f8aff4f45ec9e5a3d2073e62940faForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/education'
*/
const Controller6f7985dd84e48575c462555a011345b1 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller6f7985dd84e48575c462555a011345b1.url(options),
    method: 'get',
})

Controller6f7985dd84e48575c462555a011345b1.definition = {
    methods: ["get","head"],
    url: '/training/education',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/education'
*/
Controller6f7985dd84e48575c462555a011345b1.url = (options?: RouteQueryOptions) => {
    return Controller6f7985dd84e48575c462555a011345b1.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/education'
*/
Controller6f7985dd84e48575c462555a011345b1.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller6f7985dd84e48575c462555a011345b1.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/education'
*/
Controller6f7985dd84e48575c462555a011345b1.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller6f7985dd84e48575c462555a011345b1.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/education'
*/
const Controller6f7985dd84e48575c462555a011345b1Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller6f7985dd84e48575c462555a011345b1.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/education'
*/
Controller6f7985dd84e48575c462555a011345b1Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller6f7985dd84e48575c462555a011345b1.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/training/education'
*/
Controller6f7985dd84e48575c462555a011345b1Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller6f7985dd84e48575c462555a011345b1.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller6f7985dd84e48575c462555a011345b1.form = Controller6f7985dd84e48575c462555a011345b1Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/meetings'
*/
const Controller7f834e4898e9bd3488b8c65a8277d5da = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller7f834e4898e9bd3488b8c65a8277d5da.url(options),
    method: 'get',
})

Controller7f834e4898e9bd3488b8c65a8277d5da.definition = {
    methods: ["get","head"],
    url: '/meeting/meetings',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/meetings'
*/
Controller7f834e4898e9bd3488b8c65a8277d5da.url = (options?: RouteQueryOptions) => {
    return Controller7f834e4898e9bd3488b8c65a8277d5da.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/meetings'
*/
Controller7f834e4898e9bd3488b8c65a8277d5da.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller7f834e4898e9bd3488b8c65a8277d5da.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/meetings'
*/
Controller7f834e4898e9bd3488b8c65a8277d5da.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller7f834e4898e9bd3488b8c65a8277d5da.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/meetings'
*/
const Controller7f834e4898e9bd3488b8c65a8277d5daForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller7f834e4898e9bd3488b8c65a8277d5da.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/meetings'
*/
Controller7f834e4898e9bd3488b8c65a8277d5daForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller7f834e4898e9bd3488b8c65a8277d5da.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/meetings'
*/
Controller7f834e4898e9bd3488b8c65a8277d5daForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller7f834e4898e9bd3488b8c65a8277d5da.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller7f834e4898e9bd3488b8c65a8277d5da.form = Controller7f834e4898e9bd3488b8c65a8277d5daForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/meeting-sign-ups'
*/
const Controller09216ea80d8b90623529eb88c253494c = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller09216ea80d8b90623529eb88c253494c.url(options),
    method: 'get',
})

Controller09216ea80d8b90623529eb88c253494c.definition = {
    methods: ["get","head"],
    url: '/meeting/meeting-sign-ups',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/meeting-sign-ups'
*/
Controller09216ea80d8b90623529eb88c253494c.url = (options?: RouteQueryOptions) => {
    return Controller09216ea80d8b90623529eb88c253494c.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/meeting-sign-ups'
*/
Controller09216ea80d8b90623529eb88c253494c.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller09216ea80d8b90623529eb88c253494c.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/meeting-sign-ups'
*/
Controller09216ea80d8b90623529eb88c253494c.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller09216ea80d8b90623529eb88c253494c.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/meeting-sign-ups'
*/
const Controller09216ea80d8b90623529eb88c253494cForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller09216ea80d8b90623529eb88c253494c.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/meeting-sign-ups'
*/
Controller09216ea80d8b90623529eb88c253494cForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller09216ea80d8b90623529eb88c253494c.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/meeting-sign-ups'
*/
Controller09216ea80d8b90623529eb88c253494cForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller09216ea80d8b90623529eb88c253494c.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller09216ea80d8b90623529eb88c253494c.form = Controller09216ea80d8b90623529eb88c253494cForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/internal-meetings'
*/
const Controller909bb78633799ac7f9687ea722d32ac2 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller909bb78633799ac7f9687ea722d32ac2.url(options),
    method: 'get',
})

Controller909bb78633799ac7f9687ea722d32ac2.definition = {
    methods: ["get","head"],
    url: '/meeting/internal-meetings',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/internal-meetings'
*/
Controller909bb78633799ac7f9687ea722d32ac2.url = (options?: RouteQueryOptions) => {
    return Controller909bb78633799ac7f9687ea722d32ac2.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/internal-meetings'
*/
Controller909bb78633799ac7f9687ea722d32ac2.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller909bb78633799ac7f9687ea722d32ac2.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/internal-meetings'
*/
Controller909bb78633799ac7f9687ea722d32ac2.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller909bb78633799ac7f9687ea722d32ac2.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/internal-meetings'
*/
const Controller909bb78633799ac7f9687ea722d32ac2Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller909bb78633799ac7f9687ea722d32ac2.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/internal-meetings'
*/
Controller909bb78633799ac7f9687ea722d32ac2Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller909bb78633799ac7f9687ea722d32ac2.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/meeting/internal-meetings'
*/
Controller909bb78633799ac7f9687ea722d32ac2Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller909bb78633799ac7f9687ea722d32ac2.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller909bb78633799ac7f9687ea722d32ac2.form = Controller909bb78633799ac7f9687ea722d32ac2Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/month-statistics'
*/
const Controllerf32fbe8105905cbebc888eeaca6761b4 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerf32fbe8105905cbebc888eeaca6761b4.url(options),
    method: 'get',
})

Controllerf32fbe8105905cbebc888eeaca6761b4.definition = {
    methods: ["get","head"],
    url: '/statistics/month-statistics',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/month-statistics'
*/
Controllerf32fbe8105905cbebc888eeaca6761b4.url = (options?: RouteQueryOptions) => {
    return Controllerf32fbe8105905cbebc888eeaca6761b4.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/month-statistics'
*/
Controllerf32fbe8105905cbebc888eeaca6761b4.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerf32fbe8105905cbebc888eeaca6761b4.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/month-statistics'
*/
Controllerf32fbe8105905cbebc888eeaca6761b4.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerf32fbe8105905cbebc888eeaca6761b4.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/month-statistics'
*/
const Controllerf32fbe8105905cbebc888eeaca6761b4Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerf32fbe8105905cbebc888eeaca6761b4.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/month-statistics'
*/
Controllerf32fbe8105905cbebc888eeaca6761b4Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerf32fbe8105905cbebc888eeaca6761b4.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/month-statistics'
*/
Controllerf32fbe8105905cbebc888eeaca6761b4Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerf32fbe8105905cbebc888eeaca6761b4.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerf32fbe8105905cbebc888eeaca6761b4.form = Controllerf32fbe8105905cbebc888eeaca6761b4Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/year-statistics'
*/
const Controller05d5d86ffee2d15dd9fe346c88e51a3e = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller05d5d86ffee2d15dd9fe346c88e51a3e.url(options),
    method: 'get',
})

Controller05d5d86ffee2d15dd9fe346c88e51a3e.definition = {
    methods: ["get","head"],
    url: '/statistics/year-statistics',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/year-statistics'
*/
Controller05d5d86ffee2d15dd9fe346c88e51a3e.url = (options?: RouteQueryOptions) => {
    return Controller05d5d86ffee2d15dd9fe346c88e51a3e.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/year-statistics'
*/
Controller05d5d86ffee2d15dd9fe346c88e51a3e.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller05d5d86ffee2d15dd9fe346c88e51a3e.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/year-statistics'
*/
Controller05d5d86ffee2d15dd9fe346c88e51a3e.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller05d5d86ffee2d15dd9fe346c88e51a3e.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/year-statistics'
*/
const Controller05d5d86ffee2d15dd9fe346c88e51a3eForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller05d5d86ffee2d15dd9fe346c88e51a3e.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/year-statistics'
*/
Controller05d5d86ffee2d15dd9fe346c88e51a3eForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller05d5d86ffee2d15dd9fe346c88e51a3e.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/year-statistics'
*/
Controller05d5d86ffee2d15dd9fe346c88e51a3eForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller05d5d86ffee2d15dd9fe346c88e51a3e.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller05d5d86ffee2d15dd9fe346c88e51a3e.form = Controller05d5d86ffee2d15dd9fe346c88e51a3eForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/expert-inspections'
*/
const Controller3b431e7c6692756fbb94f688636d5142 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller3b431e7c6692756fbb94f688636d5142.url(options),
    method: 'get',
})

Controller3b431e7c6692756fbb94f688636d5142.definition = {
    methods: ["get","head"],
    url: '/statistics/expert-inspections',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/expert-inspections'
*/
Controller3b431e7c6692756fbb94f688636d5142.url = (options?: RouteQueryOptions) => {
    return Controller3b431e7c6692756fbb94f688636d5142.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/expert-inspections'
*/
Controller3b431e7c6692756fbb94f688636d5142.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller3b431e7c6692756fbb94f688636d5142.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/expert-inspections'
*/
Controller3b431e7c6692756fbb94f688636d5142.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller3b431e7c6692756fbb94f688636d5142.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/expert-inspections'
*/
const Controller3b431e7c6692756fbb94f688636d5142Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller3b431e7c6692756fbb94f688636d5142.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/expert-inspections'
*/
Controller3b431e7c6692756fbb94f688636d5142Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller3b431e7c6692756fbb94f688636d5142.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/expert-inspections'
*/
Controller3b431e7c6692756fbb94f688636d5142Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller3b431e7c6692756fbb94f688636d5142.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller3b431e7c6692756fbb94f688636d5142.form = Controller3b431e7c6692756fbb94f688636d5142Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/accident-bulletins'
*/
const Controller3987fe4e864d121bbc5e553a44f3f63e = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller3987fe4e864d121bbc5e553a44f3f63e.url(options),
    method: 'get',
})

Controller3987fe4e864d121bbc5e553a44f3f63e.definition = {
    methods: ["get","head"],
    url: '/statistics/accident-bulletins',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/accident-bulletins'
*/
Controller3987fe4e864d121bbc5e553a44f3f63e.url = (options?: RouteQueryOptions) => {
    return Controller3987fe4e864d121bbc5e553a44f3f63e.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/accident-bulletins'
*/
Controller3987fe4e864d121bbc5e553a44f3f63e.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller3987fe4e864d121bbc5e553a44f3f63e.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/accident-bulletins'
*/
Controller3987fe4e864d121bbc5e553a44f3f63e.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller3987fe4e864d121bbc5e553a44f3f63e.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/accident-bulletins'
*/
const Controller3987fe4e864d121bbc5e553a44f3f63eForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller3987fe4e864d121bbc5e553a44f3f63e.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/accident-bulletins'
*/
Controller3987fe4e864d121bbc5e553a44f3f63eForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller3987fe4e864d121bbc5e553a44f3f63e.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/statistics/accident-bulletins'
*/
Controller3987fe4e864d121bbc5e553a44f3f63eForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller3987fe4e864d121bbc5e553a44f3f63e.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller3987fe4e864d121bbc5e553a44f3f63e.form = Controller3987fe4e864d121bbc5e553a44f3f63eForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/regulation/libraries'
*/
const Controllerb7157957ec9c2ec304b591d2cb37bbee = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerb7157957ec9c2ec304b591d2cb37bbee.url(options),
    method: 'get',
})

Controllerb7157957ec9c2ec304b591d2cb37bbee.definition = {
    methods: ["get","head"],
    url: '/regulation/libraries',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/regulation/libraries'
*/
Controllerb7157957ec9c2ec304b591d2cb37bbee.url = (options?: RouteQueryOptions) => {
    return Controllerb7157957ec9c2ec304b591d2cb37bbee.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/regulation/libraries'
*/
Controllerb7157957ec9c2ec304b591d2cb37bbee.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerb7157957ec9c2ec304b591d2cb37bbee.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/regulation/libraries'
*/
Controllerb7157957ec9c2ec304b591d2cb37bbee.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerb7157957ec9c2ec304b591d2cb37bbee.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/regulation/libraries'
*/
const Controllerb7157957ec9c2ec304b591d2cb37bbeeForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb7157957ec9c2ec304b591d2cb37bbee.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/regulation/libraries'
*/
Controllerb7157957ec9c2ec304b591d2cb37bbeeForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb7157957ec9c2ec304b591d2cb37bbee.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/regulation/libraries'
*/
Controllerb7157957ec9c2ec304b591d2cb37bbeeForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb7157957ec9c2ec304b591d2cb37bbee.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerb7157957ec9c2ec304b591d2cb37bbee.form = Controllerb7157957ec9c2ec304b591d2cb37bbeeForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/space-work/space-works'
*/
const Controller766be9f9f27b6636ae9ccc917b8410b7 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller766be9f9f27b6636ae9ccc917b8410b7.url(options),
    method: 'get',
})

Controller766be9f9f27b6636ae9ccc917b8410b7.definition = {
    methods: ["get","head"],
    url: '/space-work/space-works',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/space-work/space-works'
*/
Controller766be9f9f27b6636ae9ccc917b8410b7.url = (options?: RouteQueryOptions) => {
    return Controller766be9f9f27b6636ae9ccc917b8410b7.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/space-work/space-works'
*/
Controller766be9f9f27b6636ae9ccc917b8410b7.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller766be9f9f27b6636ae9ccc917b8410b7.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/space-work/space-works'
*/
Controller766be9f9f27b6636ae9ccc917b8410b7.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller766be9f9f27b6636ae9ccc917b8410b7.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/space-work/space-works'
*/
const Controller766be9f9f27b6636ae9ccc917b8410b7Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller766be9f9f27b6636ae9ccc917b8410b7.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/space-work/space-works'
*/
Controller766be9f9f27b6636ae9ccc917b8410b7Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller766be9f9f27b6636ae9ccc917b8410b7.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/space-work/space-works'
*/
Controller766be9f9f27b6636ae9ccc917b8410b7Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller766be9f9f27b6636ae9ccc917b8410b7.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller766be9f9f27b6636ae9ccc917b8410b7.form = Controller766be9f9f27b6636ae9ccc917b8410b7Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/space-work/three-togethers'
*/
const Controller8796c98bda408c975eec5ade64700017 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller8796c98bda408c975eec5ade64700017.url(options),
    method: 'get',
})

Controller8796c98bda408c975eec5ade64700017.definition = {
    methods: ["get","head"],
    url: '/space-work/three-togethers',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/space-work/three-togethers'
*/
Controller8796c98bda408c975eec5ade64700017.url = (options?: RouteQueryOptions) => {
    return Controller8796c98bda408c975eec5ade64700017.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/space-work/three-togethers'
*/
Controller8796c98bda408c975eec5ade64700017.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller8796c98bda408c975eec5ade64700017.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/space-work/three-togethers'
*/
Controller8796c98bda408c975eec5ade64700017.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller8796c98bda408c975eec5ade64700017.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/space-work/three-togethers'
*/
const Controller8796c98bda408c975eec5ade64700017Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller8796c98bda408c975eec5ade64700017.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/space-work/three-togethers'
*/
Controller8796c98bda408c975eec5ade64700017Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller8796c98bda408c975eec5ade64700017.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/space-work/three-togethers'
*/
Controller8796c98bda408c975eec5ade64700017Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller8796c98bda408c975eec5ade64700017.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller8796c98bda408c975eec5ade64700017.form = Controller8796c98bda408c975eec5ade64700017Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/org-structure'
*/
const Controllerb7cd76195a4a9cf3f9f9d95da3eb09dc = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerb7cd76195a4a9cf3f9f9d95da3eb09dc.url(options),
    method: 'get',
})

Controllerb7cd76195a4a9cf3f9f9d95da3eb09dc.definition = {
    methods: ["get","head"],
    url: '/safety-archives/org-structure',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/org-structure'
*/
Controllerb7cd76195a4a9cf3f9f9d95da3eb09dc.url = (options?: RouteQueryOptions) => {
    return Controllerb7cd76195a4a9cf3f9f9d95da3eb09dc.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/org-structure'
*/
Controllerb7cd76195a4a9cf3f9f9d95da3eb09dc.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerb7cd76195a4a9cf3f9f9d95da3eb09dc.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/org-structure'
*/
Controllerb7cd76195a4a9cf3f9f9d95da3eb09dc.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerb7cd76195a4a9cf3f9f9d95da3eb09dc.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/org-structure'
*/
const Controllerb7cd76195a4a9cf3f9f9d95da3eb09dcForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb7cd76195a4a9cf3f9f9d95da3eb09dc.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/org-structure'
*/
Controllerb7cd76195a4a9cf3f9f9d95da3eb09dcForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb7cd76195a4a9cf3f9f9d95da3eb09dc.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/org-structure'
*/
Controllerb7cd76195a4a9cf3f9f9d95da3eb09dcForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb7cd76195a4a9cf3f9f9d95da3eb09dc.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerb7cd76195a4a9cf3f9f9d95da3eb09dc.form = Controllerb7cd76195a4a9cf3f9f9d95da3eb09dcForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/qualifications'
*/
const Controllerb33f068b9007af396523a675de9dfd1f = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerb33f068b9007af396523a675de9dfd1f.url(options),
    method: 'get',
})

Controllerb33f068b9007af396523a675de9dfd1f.definition = {
    methods: ["get","head"],
    url: '/safety-archives/qualifications',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/qualifications'
*/
Controllerb33f068b9007af396523a675de9dfd1f.url = (options?: RouteQueryOptions) => {
    return Controllerb33f068b9007af396523a675de9dfd1f.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/qualifications'
*/
Controllerb33f068b9007af396523a675de9dfd1f.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerb33f068b9007af396523a675de9dfd1f.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/qualifications'
*/
Controllerb33f068b9007af396523a675de9dfd1f.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerb33f068b9007af396523a675de9dfd1f.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/qualifications'
*/
const Controllerb33f068b9007af396523a675de9dfd1fForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb33f068b9007af396523a675de9dfd1f.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/qualifications'
*/
Controllerb33f068b9007af396523a675de9dfd1fForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb33f068b9007af396523a675de9dfd1f.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/qualifications'
*/
Controllerb33f068b9007af396523a675de9dfd1fForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerb33f068b9007af396523a675de9dfd1f.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerb33f068b9007af396523a675de9dfd1f.form = Controllerb33f068b9007af396523a675de9dfd1fForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/safety-investment'
*/
const Controller40c606abc4d95e40c8a1bbe7a7b965f9 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller40c606abc4d95e40c8a1bbe7a7b965f9.url(options),
    method: 'get',
})

Controller40c606abc4d95e40c8a1bbe7a7b965f9.definition = {
    methods: ["get","head"],
    url: '/safety-archives/safety-investment',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/safety-investment'
*/
Controller40c606abc4d95e40c8a1bbe7a7b965f9.url = (options?: RouteQueryOptions) => {
    return Controller40c606abc4d95e40c8a1bbe7a7b965f9.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/safety-investment'
*/
Controller40c606abc4d95e40c8a1bbe7a7b965f9.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller40c606abc4d95e40c8a1bbe7a7b965f9.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/safety-investment'
*/
Controller40c606abc4d95e40c8a1bbe7a7b965f9.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller40c606abc4d95e40c8a1bbe7a7b965f9.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/safety-investment'
*/
const Controller40c606abc4d95e40c8a1bbe7a7b965f9Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller40c606abc4d95e40c8a1bbe7a7b965f9.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/safety-investment'
*/
Controller40c606abc4d95e40c8a1bbe7a7b965f9Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller40c606abc4d95e40c8a1bbe7a7b965f9.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/safety-archives/safety-investment'
*/
Controller40c606abc4d95e40c8a1bbe7a7b965f9Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller40c606abc4d95e40c8a1bbe7a7b965f9.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller40c606abc4d95e40c8a1bbe7a7b965f9.form = Controller40c606abc4d95e40c8a1bbe7a7b965f9Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/equipment/major-hazardous'
*/
const Controllerbe8d43bd91af207ce1972a6ea9b07b97 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerbe8d43bd91af207ce1972a6ea9b07b97.url(options),
    method: 'get',
})

Controllerbe8d43bd91af207ce1972a6ea9b07b97.definition = {
    methods: ["get","head"],
    url: '/equipment/major-hazardous',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/equipment/major-hazardous'
*/
Controllerbe8d43bd91af207ce1972a6ea9b07b97.url = (options?: RouteQueryOptions) => {
    return Controllerbe8d43bd91af207ce1972a6ea9b07b97.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/equipment/major-hazardous'
*/
Controllerbe8d43bd91af207ce1972a6ea9b07b97.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerbe8d43bd91af207ce1972a6ea9b07b97.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/equipment/major-hazardous'
*/
Controllerbe8d43bd91af207ce1972a6ea9b07b97.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerbe8d43bd91af207ce1972a6ea9b07b97.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/equipment/major-hazardous'
*/
const Controllerbe8d43bd91af207ce1972a6ea9b07b97Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerbe8d43bd91af207ce1972a6ea9b07b97.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/equipment/major-hazardous'
*/
Controllerbe8d43bd91af207ce1972a6ea9b07b97Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerbe8d43bd91af207ce1972a6ea9b07b97.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/equipment/major-hazardous'
*/
Controllerbe8d43bd91af207ce1972a6ea9b07b97Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerbe8d43bd91af207ce1972a6ea9b07b97.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerbe8d43bd91af207ce1972a6ea9b07b97.form = Controllerbe8d43bd91af207ce1972a6ea9b07b97Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/process/key-hazardous'
*/
const Controller108e5a2f69ee121654ee34809a401fb1 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller108e5a2f69ee121654ee34809a401fb1.url(options),
    method: 'get',
})

Controller108e5a2f69ee121654ee34809a401fb1.definition = {
    methods: ["get","head"],
    url: '/process/key-hazardous',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/process/key-hazardous'
*/
Controller108e5a2f69ee121654ee34809a401fb1.url = (options?: RouteQueryOptions) => {
    return Controller108e5a2f69ee121654ee34809a401fb1.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/process/key-hazardous'
*/
Controller108e5a2f69ee121654ee34809a401fb1.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller108e5a2f69ee121654ee34809a401fb1.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/process/key-hazardous'
*/
Controller108e5a2f69ee121654ee34809a401fb1.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller108e5a2f69ee121654ee34809a401fb1.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/process/key-hazardous'
*/
const Controller108e5a2f69ee121654ee34809a401fb1Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller108e5a2f69ee121654ee34809a401fb1.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/process/key-hazardous'
*/
Controller108e5a2f69ee121654ee34809a401fb1Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller108e5a2f69ee121654ee34809a401fb1.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/process/key-hazardous'
*/
Controller108e5a2f69ee121654ee34809a401fb1Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller108e5a2f69ee121654ee34809a401fb1.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller108e5a2f69ee121654ee34809a401fb1.form = Controller108e5a2f69ee121654ee34809a401fb1Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/process/hazardous'
*/
const Controller86338a2b8fa3273fd00718be7bfc535e = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller86338a2b8fa3273fd00718be7bfc535e.url(options),
    method: 'get',
})

Controller86338a2b8fa3273fd00718be7bfc535e.definition = {
    methods: ["get","head"],
    url: '/process/hazardous',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/process/hazardous'
*/
Controller86338a2b8fa3273fd00718be7bfc535e.url = (options?: RouteQueryOptions) => {
    return Controller86338a2b8fa3273fd00718be7bfc535e.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/process/hazardous'
*/
Controller86338a2b8fa3273fd00718be7bfc535e.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller86338a2b8fa3273fd00718be7bfc535e.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/process/hazardous'
*/
Controller86338a2b8fa3273fd00718be7bfc535e.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller86338a2b8fa3273fd00718be7bfc535e.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/process/hazardous'
*/
const Controller86338a2b8fa3273fd00718be7bfc535eForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller86338a2b8fa3273fd00718be7bfc535e.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/process/hazardous'
*/
Controller86338a2b8fa3273fd00718be7bfc535eForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller86338a2b8fa3273fd00718be7bfc535e.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/process/hazardous'
*/
Controller86338a2b8fa3273fd00718be7bfc535eForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller86338a2b8fa3273fd00718be7bfc535e.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller86338a2b8fa3273fd00718be7bfc535e.form = Controller86338a2b8fa3273fd00718be7bfc535eForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/safety-assessment'
*/
const Controllercb5a69541b20685d1ed664d21d4eec41 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllercb5a69541b20685d1ed664d21d4eec41.url(options),
    method: 'get',
})

Controllercb5a69541b20685d1ed664d21d4eec41.definition = {
    methods: ["get","head"],
    url: '/risk/safety-assessment',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/safety-assessment'
*/
Controllercb5a69541b20685d1ed664d21d4eec41.url = (options?: RouteQueryOptions) => {
    return Controllercb5a69541b20685d1ed664d21d4eec41.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/safety-assessment'
*/
Controllercb5a69541b20685d1ed664d21d4eec41.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllercb5a69541b20685d1ed664d21d4eec41.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/safety-assessment'
*/
Controllercb5a69541b20685d1ed664d21d4eec41.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllercb5a69541b20685d1ed664d21d4eec41.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/safety-assessment'
*/
const Controllercb5a69541b20685d1ed664d21d4eec41Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercb5a69541b20685d1ed664d21d4eec41.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/safety-assessment'
*/
Controllercb5a69541b20685d1ed664d21d4eec41Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercb5a69541b20685d1ed664d21d4eec41.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/safety-assessment'
*/
Controllercb5a69541b20685d1ed664d21d4eec41Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercb5a69541b20685d1ed664d21d4eec41.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllercb5a69541b20685d1ed664d21d4eec41.form = Controllercb5a69541b20685d1ed664d21d4eec41Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/expert-inspection'
*/
const Controller50a3e572e907c9d988298a8b7d433db8 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller50a3e572e907c9d988298a8b7d433db8.url(options),
    method: 'get',
})

Controller50a3e572e907c9d988298a8b7d433db8.definition = {
    methods: ["get","head"],
    url: '/risk/expert-inspection',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/expert-inspection'
*/
Controller50a3e572e907c9d988298a8b7d433db8.url = (options?: RouteQueryOptions) => {
    return Controller50a3e572e907c9d988298a8b7d433db8.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/expert-inspection'
*/
Controller50a3e572e907c9d988298a8b7d433db8.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller50a3e572e907c9d988298a8b7d433db8.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/expert-inspection'
*/
Controller50a3e572e907c9d988298a8b7d433db8.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller50a3e572e907c9d988298a8b7d433db8.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/expert-inspection'
*/
const Controller50a3e572e907c9d988298a8b7d433db8Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller50a3e572e907c9d988298a8b7d433db8.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/expert-inspection'
*/
Controller50a3e572e907c9d988298a8b7d433db8Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller50a3e572e907c9d988298a8b7d433db8.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/expert-inspection'
*/
Controller50a3e572e907c9d988298a8b7d433db8Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller50a3e572e907c9d988298a8b7d433db8.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller50a3e572e907c9d988298a8b7d433db8.form = Controller50a3e572e907c9d988298a8b7d433db8Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/risk-warning'
*/
const Controller4e52ccde3afaaf81c032199e92c5f1ab = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller4e52ccde3afaaf81c032199e92c5f1ab.url(options),
    method: 'get',
})

Controller4e52ccde3afaaf81c032199e92c5f1ab.definition = {
    methods: ["get","head"],
    url: '/risk/risk-warning',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/risk-warning'
*/
Controller4e52ccde3afaaf81c032199e92c5f1ab.url = (options?: RouteQueryOptions) => {
    return Controller4e52ccde3afaaf81c032199e92c5f1ab.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/risk-warning'
*/
Controller4e52ccde3afaaf81c032199e92c5f1ab.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller4e52ccde3afaaf81c032199e92c5f1ab.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/risk-warning'
*/
Controller4e52ccde3afaaf81c032199e92c5f1ab.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller4e52ccde3afaaf81c032199e92c5f1ab.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/risk-warning'
*/
const Controller4e52ccde3afaaf81c032199e92c5f1abForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller4e52ccde3afaaf81c032199e92c5f1ab.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/risk-warning'
*/
Controller4e52ccde3afaaf81c032199e92c5f1abForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller4e52ccde3afaaf81c032199e92c5f1ab.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk/risk-warning'
*/
Controller4e52ccde3afaaf81c032199e92c5f1abForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller4e52ccde3afaaf81c032199e92c5f1ab.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller4e52ccde3afaaf81c032199e92c5f1ab.form = Controller4e52ccde3afaaf81c032199e92c5f1abForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/design-units'
*/
const Controller1f601ea8e5bb218ee28f6eda1db540b9 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller1f601ea8e5bb218ee28f6eda1db540b9.url(options),
    method: 'get',
})

Controller1f601ea8e5bb218ee28f6eda1db540b9.definition = {
    methods: ["get","head"],
    url: '/third-party/design-units',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/design-units'
*/
Controller1f601ea8e5bb218ee28f6eda1db540b9.url = (options?: RouteQueryOptions) => {
    return Controller1f601ea8e5bb218ee28f6eda1db540b9.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/design-units'
*/
Controller1f601ea8e5bb218ee28f6eda1db540b9.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller1f601ea8e5bb218ee28f6eda1db540b9.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/design-units'
*/
Controller1f601ea8e5bb218ee28f6eda1db540b9.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller1f601ea8e5bb218ee28f6eda1db540b9.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/design-units'
*/
const Controller1f601ea8e5bb218ee28f6eda1db540b9Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller1f601ea8e5bb218ee28f6eda1db540b9.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/design-units'
*/
Controller1f601ea8e5bb218ee28f6eda1db540b9Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller1f601ea8e5bb218ee28f6eda1db540b9.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/design-units'
*/
Controller1f601ea8e5bb218ee28f6eda1db540b9Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller1f601ea8e5bb218ee28f6eda1db540b9.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller1f601ea8e5bb218ee28f6eda1db540b9.form = Controller1f601ea8e5bb218ee28f6eda1db540b9Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/construction-units'
*/
const Controller9380dca6c49302c2ee76e6228e538c16 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller9380dca6c49302c2ee76e6228e538c16.url(options),
    method: 'get',
})

Controller9380dca6c49302c2ee76e6228e538c16.definition = {
    methods: ["get","head"],
    url: '/third-party/construction-units',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/construction-units'
*/
Controller9380dca6c49302c2ee76e6228e538c16.url = (options?: RouteQueryOptions) => {
    return Controller9380dca6c49302c2ee76e6228e538c16.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/construction-units'
*/
Controller9380dca6c49302c2ee76e6228e538c16.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller9380dca6c49302c2ee76e6228e538c16.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/construction-units'
*/
Controller9380dca6c49302c2ee76e6228e538c16.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller9380dca6c49302c2ee76e6228e538c16.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/construction-units'
*/
const Controller9380dca6c49302c2ee76e6228e538c16Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller9380dca6c49302c2ee76e6228e538c16.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/construction-units'
*/
Controller9380dca6c49302c2ee76e6228e538c16Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller9380dca6c49302c2ee76e6228e538c16.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/construction-units'
*/
Controller9380dca6c49302c2ee76e6228e538c16Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller9380dca6c49302c2ee76e6228e538c16.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller9380dca6c49302c2ee76e6228e538c16.form = Controller9380dca6c49302c2ee76e6228e538c16Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/technical-services'
*/
const Controller0e40752e8ff2ba7acfe8001538e4d6eb = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller0e40752e8ff2ba7acfe8001538e4d6eb.url(options),
    method: 'get',
})

Controller0e40752e8ff2ba7acfe8001538e4d6eb.definition = {
    methods: ["get","head"],
    url: '/third-party/technical-services',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/technical-services'
*/
Controller0e40752e8ff2ba7acfe8001538e4d6eb.url = (options?: RouteQueryOptions) => {
    return Controller0e40752e8ff2ba7acfe8001538e4d6eb.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/technical-services'
*/
Controller0e40752e8ff2ba7acfe8001538e4d6eb.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller0e40752e8ff2ba7acfe8001538e4d6eb.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/technical-services'
*/
Controller0e40752e8ff2ba7acfe8001538e4d6eb.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller0e40752e8ff2ba7acfe8001538e4d6eb.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/technical-services'
*/
const Controller0e40752e8ff2ba7acfe8001538e4d6ebForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller0e40752e8ff2ba7acfe8001538e4d6eb.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/technical-services'
*/
Controller0e40752e8ff2ba7acfe8001538e4d6ebForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller0e40752e8ff2ba7acfe8001538e4d6eb.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/third-party/technical-services'
*/
Controller0e40752e8ff2ba7acfe8001538e4d6ebForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller0e40752e8ff2ba7acfe8001538e4d6eb.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller0e40752e8ff2ba7acfe8001538e4d6eb.form = Controller0e40752e8ff2ba7acfe8001538e4d6ebForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/disease-report'
*/
const Controllerfc972e8bb979ae574f7fff41328ab6b1 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerfc972e8bb979ae574f7fff41328ab6b1.url(options),
    method: 'get',
})

Controllerfc972e8bb979ae574f7fff41328ab6b1.definition = {
    methods: ["get","head"],
    url: '/health/disease-report',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/disease-report'
*/
Controllerfc972e8bb979ae574f7fff41328ab6b1.url = (options?: RouteQueryOptions) => {
    return Controllerfc972e8bb979ae574f7fff41328ab6b1.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/disease-report'
*/
Controllerfc972e8bb979ae574f7fff41328ab6b1.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerfc972e8bb979ae574f7fff41328ab6b1.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/disease-report'
*/
Controllerfc972e8bb979ae574f7fff41328ab6b1.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerfc972e8bb979ae574f7fff41328ab6b1.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/disease-report'
*/
const Controllerfc972e8bb979ae574f7fff41328ab6b1Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerfc972e8bb979ae574f7fff41328ab6b1.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/disease-report'
*/
Controllerfc972e8bb979ae574f7fff41328ab6b1Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerfc972e8bb979ae574f7fff41328ab6b1.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/health/disease-report'
*/
Controllerfc972e8bb979ae574f7fff41328ab6b1Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerfc972e8bb979ae574f7fff41328ab6b1.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerfc972e8bb979ae574f7fff41328ab6b1.form = Controllerfc972e8bb979ae574f7fff41328ab6b1Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/fire/equipment'
*/
const Controllera6832eb7956939320272db69dba3a866 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllera6832eb7956939320272db69dba3a866.url(options),
    method: 'get',
})

Controllera6832eb7956939320272db69dba3a866.definition = {
    methods: ["get","head"],
    url: '/fire/equipment',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/fire/equipment'
*/
Controllera6832eb7956939320272db69dba3a866.url = (options?: RouteQueryOptions) => {
    return Controllera6832eb7956939320272db69dba3a866.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/fire/equipment'
*/
Controllera6832eb7956939320272db69dba3a866.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllera6832eb7956939320272db69dba3a866.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/fire/equipment'
*/
Controllera6832eb7956939320272db69dba3a866.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllera6832eb7956939320272db69dba3a866.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/fire/equipment'
*/
const Controllera6832eb7956939320272db69dba3a866Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera6832eb7956939320272db69dba3a866.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/fire/equipment'
*/
Controllera6832eb7956939320272db69dba3a866Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera6832eb7956939320272db69dba3a866.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/fire/equipment'
*/
Controllera6832eb7956939320272db69dba3a866Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera6832eb7956939320272db69dba3a866.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllera6832eb7956939320272db69dba3a866.form = Controllera6832eb7956939320272db69dba3a866Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/fire/facilities'
*/
const Controller26dcb1a9e73d9c46d8ec1f08de7de20c = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller26dcb1a9e73d9c46d8ec1f08de7de20c.url(options),
    method: 'get',
})

Controller26dcb1a9e73d9c46d8ec1f08de7de20c.definition = {
    methods: ["get","head"],
    url: '/fire/facilities',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/fire/facilities'
*/
Controller26dcb1a9e73d9c46d8ec1f08de7de20c.url = (options?: RouteQueryOptions) => {
    return Controller26dcb1a9e73d9c46d8ec1f08de7de20c.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/fire/facilities'
*/
Controller26dcb1a9e73d9c46d8ec1f08de7de20c.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller26dcb1a9e73d9c46d8ec1f08de7de20c.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/fire/facilities'
*/
Controller26dcb1a9e73d9c46d8ec1f08de7de20c.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller26dcb1a9e73d9c46d8ec1f08de7de20c.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/fire/facilities'
*/
const Controller26dcb1a9e73d9c46d8ec1f08de7de20cForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller26dcb1a9e73d9c46d8ec1f08de7de20c.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/fire/facilities'
*/
Controller26dcb1a9e73d9c46d8ec1f08de7de20cForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller26dcb1a9e73d9c46d8ec1f08de7de20c.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/fire/facilities'
*/
Controller26dcb1a9e73d9c46d8ec1f08de7de20cForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller26dcb1a9e73d9c46d8ec1f08de7de20c.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller26dcb1a9e73d9c46d8ec1f08de7de20c.form = Controller26dcb1a9e73d9c46d8ec1f08de7de20cForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/report'
*/
const Controller5095ce3dd4931a582d1b106234287d92 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller5095ce3dd4931a582d1b106234287d92.url(options),
    method: 'get',
})

Controller5095ce3dd4931a582d1b106234287d92.definition = {
    methods: ["get","head"],
    url: '/accident/report',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/report'
*/
Controller5095ce3dd4931a582d1b106234287d92.url = (options?: RouteQueryOptions) => {
    return Controller5095ce3dd4931a582d1b106234287d92.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/report'
*/
Controller5095ce3dd4931a582d1b106234287d92.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller5095ce3dd4931a582d1b106234287d92.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/report'
*/
Controller5095ce3dd4931a582d1b106234287d92.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller5095ce3dd4931a582d1b106234287d92.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/report'
*/
const Controller5095ce3dd4931a582d1b106234287d92Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller5095ce3dd4931a582d1b106234287d92.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/report'
*/
Controller5095ce3dd4931a582d1b106234287d92Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller5095ce3dd4931a582d1b106234287d92.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/report'
*/
Controller5095ce3dd4931a582d1b106234287d92Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller5095ce3dd4931a582d1b106234287d92.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller5095ce3dd4931a582d1b106234287d92.form = Controller5095ce3dd4931a582d1b106234287d92Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency'
*/
const Controllera776b56c843d3a3b51d2ed29d47bbe72 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllera776b56c843d3a3b51d2ed29d47bbe72.url(options),
    method: 'get',
})

Controllera776b56c843d3a3b51d2ed29d47bbe72.definition = {
    methods: ["get","head"],
    url: '/accident/emergency',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency'
*/
Controllera776b56c843d3a3b51d2ed29d47bbe72.url = (options?: RouteQueryOptions) => {
    return Controllera776b56c843d3a3b51d2ed29d47bbe72.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency'
*/
Controllera776b56c843d3a3b51d2ed29d47bbe72.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllera776b56c843d3a3b51d2ed29d47bbe72.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency'
*/
Controllera776b56c843d3a3b51d2ed29d47bbe72.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllera776b56c843d3a3b51d2ed29d47bbe72.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency'
*/
const Controllera776b56c843d3a3b51d2ed29d47bbe72Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera776b56c843d3a3b51d2ed29d47bbe72.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency'
*/
Controllera776b56c843d3a3b51d2ed29d47bbe72Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera776b56c843d3a3b51d2ed29d47bbe72.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency'
*/
Controllera776b56c843d3a3b51d2ed29d47bbe72Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllera776b56c843d3a3b51d2ed29d47bbe72.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllera776b56c843d3a3b51d2ed29d47bbe72.form = Controllera776b56c843d3a3b51d2ed29d47bbe72Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/teams'
*/
const Controller1599ae2fdadec1c0124d7a7c33efec5b = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller1599ae2fdadec1c0124d7a7c33efec5b.url(options),
    method: 'get',
})

Controller1599ae2fdadec1c0124d7a7c33efec5b.definition = {
    methods: ["get","head"],
    url: '/accident/emergency/teams',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/teams'
*/
Controller1599ae2fdadec1c0124d7a7c33efec5b.url = (options?: RouteQueryOptions) => {
    return Controller1599ae2fdadec1c0124d7a7c33efec5b.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/teams'
*/
Controller1599ae2fdadec1c0124d7a7c33efec5b.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller1599ae2fdadec1c0124d7a7c33efec5b.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/teams'
*/
Controller1599ae2fdadec1c0124d7a7c33efec5b.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller1599ae2fdadec1c0124d7a7c33efec5b.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/teams'
*/
const Controller1599ae2fdadec1c0124d7a7c33efec5bForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller1599ae2fdadec1c0124d7a7c33efec5b.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/teams'
*/
Controller1599ae2fdadec1c0124d7a7c33efec5bForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller1599ae2fdadec1c0124d7a7c33efec5b.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/teams'
*/
Controller1599ae2fdadec1c0124d7a7c33efec5bForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller1599ae2fdadec1c0124d7a7c33efec5b.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller1599ae2fdadec1c0124d7a7c33efec5b.form = Controller1599ae2fdadec1c0124d7a7c33efec5bForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/supplies'
*/
const Controller93550ff8800eaaf91ef08224c1c678dd = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller93550ff8800eaaf91ef08224c1c678dd.url(options),
    method: 'get',
})

Controller93550ff8800eaaf91ef08224c1c678dd.definition = {
    methods: ["get","head"],
    url: '/accident/emergency/supplies',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/supplies'
*/
Controller93550ff8800eaaf91ef08224c1c678dd.url = (options?: RouteQueryOptions) => {
    return Controller93550ff8800eaaf91ef08224c1c678dd.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/supplies'
*/
Controller93550ff8800eaaf91ef08224c1c678dd.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller93550ff8800eaaf91ef08224c1c678dd.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/supplies'
*/
Controller93550ff8800eaaf91ef08224c1c678dd.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller93550ff8800eaaf91ef08224c1c678dd.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/supplies'
*/
const Controller93550ff8800eaaf91ef08224c1c678ddForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller93550ff8800eaaf91ef08224c1c678dd.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/supplies'
*/
Controller93550ff8800eaaf91ef08224c1c678ddForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller93550ff8800eaaf91ef08224c1c678dd.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/supplies'
*/
Controller93550ff8800eaaf91ef08224c1c678ddForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller93550ff8800eaaf91ef08224c1c678dd.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller93550ff8800eaaf91ef08224c1c678dd.form = Controller93550ff8800eaaf91ef08224c1c678ddForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/drill'
*/
const Controller2f9e705b2aec86c00af0abdde41f6aa4 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller2f9e705b2aec86c00af0abdde41f6aa4.url(options),
    method: 'get',
})

Controller2f9e705b2aec86c00af0abdde41f6aa4.definition = {
    methods: ["get","head"],
    url: '/accident/emergency/drill',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/drill'
*/
Controller2f9e705b2aec86c00af0abdde41f6aa4.url = (options?: RouteQueryOptions) => {
    return Controller2f9e705b2aec86c00af0abdde41f6aa4.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/drill'
*/
Controller2f9e705b2aec86c00af0abdde41f6aa4.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller2f9e705b2aec86c00af0abdde41f6aa4.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/drill'
*/
Controller2f9e705b2aec86c00af0abdde41f6aa4.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller2f9e705b2aec86c00af0abdde41f6aa4.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/drill'
*/
const Controller2f9e705b2aec86c00af0abdde41f6aa4Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller2f9e705b2aec86c00af0abdde41f6aa4.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/drill'
*/
Controller2f9e705b2aec86c00af0abdde41f6aa4Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller2f9e705b2aec86c00af0abdde41f6aa4.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/accident/emergency/drill'
*/
Controller2f9e705b2aec86c00af0abdde41f6aa4Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller2f9e705b2aec86c00af0abdde41f6aa4.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller2f9e705b2aec86c00af0abdde41f6aa4.form = Controller2f9e705b2aec86c00af0abdde41f6aa4Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/notices'
*/
const Controller03594eab5c2ebeb813e15f5f2a9a3a6b = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller03594eab5c2ebeb813e15f5f2a9a3a6b.url(options),
    method: 'get',
})

Controller03594eab5c2ebeb813e15f5f2a9a3a6b.definition = {
    methods: ["get","head"],
    url: '/tasks/notices',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/notices'
*/
Controller03594eab5c2ebeb813e15f5f2a9a3a6b.url = (options?: RouteQueryOptions) => {
    return Controller03594eab5c2ebeb813e15f5f2a9a3a6b.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/notices'
*/
Controller03594eab5c2ebeb813e15f5f2a9a3a6b.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller03594eab5c2ebeb813e15f5f2a9a3a6b.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/notices'
*/
Controller03594eab5c2ebeb813e15f5f2a9a3a6b.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller03594eab5c2ebeb813e15f5f2a9a3a6b.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/notices'
*/
const Controller03594eab5c2ebeb813e15f5f2a9a3a6bForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller03594eab5c2ebeb813e15f5f2a9a3a6b.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/notices'
*/
Controller03594eab5c2ebeb813e15f5f2a9a3a6bForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller03594eab5c2ebeb813e15f5f2a9a3a6b.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/notices'
*/
Controller03594eab5c2ebeb813e15f5f2a9a3a6bForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller03594eab5c2ebeb813e15f5f2a9a3a6b.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller03594eab5c2ebeb813e15f5f2a9a3a6b.form = Controller03594eab5c2ebeb813e15f5f2a9a3a6bForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/manuscripts'
*/
const Controller421d2a8f3aeb1b2e342bef6c9b4f52e5 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller421d2a8f3aeb1b2e342bef6c9b4f52e5.url(options),
    method: 'get',
})

Controller421d2a8f3aeb1b2e342bef6c9b4f52e5.definition = {
    methods: ["get","head"],
    url: '/tasks/manuscripts',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/manuscripts'
*/
Controller421d2a8f3aeb1b2e342bef6c9b4f52e5.url = (options?: RouteQueryOptions) => {
    return Controller421d2a8f3aeb1b2e342bef6c9b4f52e5.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/manuscripts'
*/
Controller421d2a8f3aeb1b2e342bef6c9b4f52e5.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller421d2a8f3aeb1b2e342bef6c9b4f52e5.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/manuscripts'
*/
Controller421d2a8f3aeb1b2e342bef6c9b4f52e5.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller421d2a8f3aeb1b2e342bef6c9b4f52e5.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/manuscripts'
*/
const Controller421d2a8f3aeb1b2e342bef6c9b4f52e5Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller421d2a8f3aeb1b2e342bef6c9b4f52e5.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/manuscripts'
*/
Controller421d2a8f3aeb1b2e342bef6c9b4f52e5Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller421d2a8f3aeb1b2e342bef6c9b4f52e5.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/manuscripts'
*/
Controller421d2a8f3aeb1b2e342bef6c9b4f52e5Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller421d2a8f3aeb1b2e342bef6c9b4f52e5.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller421d2a8f3aeb1b2e342bef6c9b4f52e5.form = Controller421d2a8f3aeb1b2e342bef6c9b4f52e5Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/monthly-report'
*/
const Controller8a91a2034574762b1b48654125caeb46 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller8a91a2034574762b1b48654125caeb46.url(options),
    method: 'get',
})

Controller8a91a2034574762b1b48654125caeb46.definition = {
    methods: ["get","head"],
    url: '/tasks/monthly-report',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/monthly-report'
*/
Controller8a91a2034574762b1b48654125caeb46.url = (options?: RouteQueryOptions) => {
    return Controller8a91a2034574762b1b48654125caeb46.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/monthly-report'
*/
Controller8a91a2034574762b1b48654125caeb46.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller8a91a2034574762b1b48654125caeb46.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/monthly-report'
*/
Controller8a91a2034574762b1b48654125caeb46.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller8a91a2034574762b1b48654125caeb46.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/monthly-report'
*/
const Controller8a91a2034574762b1b48654125caeb46Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller8a91a2034574762b1b48654125caeb46.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/monthly-report'
*/
Controller8a91a2034574762b1b48654125caeb46Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller8a91a2034574762b1b48654125caeb46.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/monthly-report'
*/
Controller8a91a2034574762b1b48654125caeb46Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller8a91a2034574762b1b48654125caeb46.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller8a91a2034574762b1b48654125caeb46.form = Controller8a91a2034574762b1b48654125caeb46Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/annual-report'
*/
const Controller929e879457a8031e25756388a6fe8942 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller929e879457a8031e25756388a6fe8942.url(options),
    method: 'get',
})

Controller929e879457a8031e25756388a6fe8942.definition = {
    methods: ["get","head"],
    url: '/tasks/annual-report',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/annual-report'
*/
Controller929e879457a8031e25756388a6fe8942.url = (options?: RouteQueryOptions) => {
    return Controller929e879457a8031e25756388a6fe8942.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/annual-report'
*/
Controller929e879457a8031e25756388a6fe8942.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller929e879457a8031e25756388a6fe8942.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/annual-report'
*/
Controller929e879457a8031e25756388a6fe8942.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller929e879457a8031e25756388a6fe8942.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/annual-report'
*/
const Controller929e879457a8031e25756388a6fe8942Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller929e879457a8031e25756388a6fe8942.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/annual-report'
*/
Controller929e879457a8031e25756388a6fe8942Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller929e879457a8031e25756388a6fe8942.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/tasks/annual-report'
*/
Controller929e879457a8031e25756388a6fe8942Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller929e879457a8031e25756388a6fe8942.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller929e879457a8031e25756388a6fe8942.form = Controller929e879457a8031e25756388a6fe8942Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/change-management'
*/
const Controller97869e24dedf9bea8f73c16937915605 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller97869e24dedf9bea8f73c16937915605.url(options),
    method: 'get',
})

Controller97869e24dedf9bea8f73c16937915605.definition = {
    methods: ["get","head"],
    url: '/change-management',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/change-management'
*/
Controller97869e24dedf9bea8f73c16937915605.url = (options?: RouteQueryOptions) => {
    return Controller97869e24dedf9bea8f73c16937915605.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/change-management'
*/
Controller97869e24dedf9bea8f73c16937915605.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller97869e24dedf9bea8f73c16937915605.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/change-management'
*/
Controller97869e24dedf9bea8f73c16937915605.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller97869e24dedf9bea8f73c16937915605.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/change-management'
*/
const Controller97869e24dedf9bea8f73c16937915605Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller97869e24dedf9bea8f73c16937915605.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/change-management'
*/
Controller97869e24dedf9bea8f73c16937915605Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller97869e24dedf9bea8f73c16937915605.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/change-management'
*/
Controller97869e24dedf9bea8f73c16937915605Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller97869e24dedf9bea8f73c16937915605.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller97869e24dedf9bea8f73c16937915605.form = Controller97869e24dedf9bea8f73c16937915605Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/knowledge/checklists'
*/
const Controller6e7265da7faf052741f8de13edbbea3b = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller6e7265da7faf052741f8de13edbbea3b.url(options),
    method: 'get',
})

Controller6e7265da7faf052741f8de13edbbea3b.definition = {
    methods: ["get","head"],
    url: '/knowledge/checklists',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/knowledge/checklists'
*/
Controller6e7265da7faf052741f8de13edbbea3b.url = (options?: RouteQueryOptions) => {
    return Controller6e7265da7faf052741f8de13edbbea3b.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/knowledge/checklists'
*/
Controller6e7265da7faf052741f8de13edbbea3b.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller6e7265da7faf052741f8de13edbbea3b.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/knowledge/checklists'
*/
Controller6e7265da7faf052741f8de13edbbea3b.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller6e7265da7faf052741f8de13edbbea3b.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/knowledge/checklists'
*/
const Controller6e7265da7faf052741f8de13edbbea3bForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller6e7265da7faf052741f8de13edbbea3b.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/knowledge/checklists'
*/
Controller6e7265da7faf052741f8de13edbbea3bForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller6e7265da7faf052741f8de13edbbea3b.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/knowledge/checklists'
*/
Controller6e7265da7faf052741f8de13edbbea3bForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller6e7265da7faf052741f8de13edbbea3b.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller6e7265da7faf052741f8de13edbbea3b.form = Controller6e7265da7faf052741f8de13edbbea3bForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/assessment/filing'
*/
const Controllerbe1bcb06c4f801f544587db5920cc9ba = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerbe1bcb06c4f801f544587db5920cc9ba.url(options),
    method: 'get',
})

Controllerbe1bcb06c4f801f544587db5920cc9ba.definition = {
    methods: ["get","head"],
    url: '/assessment/filing',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/assessment/filing'
*/
Controllerbe1bcb06c4f801f544587db5920cc9ba.url = (options?: RouteQueryOptions) => {
    return Controllerbe1bcb06c4f801f544587db5920cc9ba.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/assessment/filing'
*/
Controllerbe1bcb06c4f801f544587db5920cc9ba.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerbe1bcb06c4f801f544587db5920cc9ba.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/assessment/filing'
*/
Controllerbe1bcb06c4f801f544587db5920cc9ba.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerbe1bcb06c4f801f544587db5920cc9ba.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/assessment/filing'
*/
const Controllerbe1bcb06c4f801f544587db5920cc9baForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerbe1bcb06c4f801f544587db5920cc9ba.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/assessment/filing'
*/
Controllerbe1bcb06c4f801f544587db5920cc9baForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerbe1bcb06c4f801f544587db5920cc9ba.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/assessment/filing'
*/
Controllerbe1bcb06c4f801f544587db5920cc9baForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerbe1bcb06c4f801f544587db5920cc9ba.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerbe1bcb06c4f801f544587db5920cc9ba.form = Controllerbe1bcb06c4f801f544587db5920cc9baForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/account'
*/
const Controller510f52c4102bebd37c0d8d2bd0bda989 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller510f52c4102bebd37c0d8d2bd0bda989.url(options),
    method: 'get',
})

Controller510f52c4102bebd37c0d8d2bd0bda989.definition = {
    methods: ["get","head"],
    url: '/account',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/account'
*/
Controller510f52c4102bebd37c0d8d2bd0bda989.url = (options?: RouteQueryOptions) => {
    return Controller510f52c4102bebd37c0d8d2bd0bda989.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/account'
*/
Controller510f52c4102bebd37c0d8d2bd0bda989.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller510f52c4102bebd37c0d8d2bd0bda989.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/account'
*/
Controller510f52c4102bebd37c0d8d2bd0bda989.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller510f52c4102bebd37c0d8d2bd0bda989.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/account'
*/
const Controller510f52c4102bebd37c0d8d2bd0bda989Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller510f52c4102bebd37c0d8d2bd0bda989.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/account'
*/
Controller510f52c4102bebd37c0d8d2bd0bda989Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller510f52c4102bebd37c0d8d2bd0bda989.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/account'
*/
Controller510f52c4102bebd37c0d8d2bd0bda989Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller510f52c4102bebd37c0d8d2bd0bda989.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller510f52c4102bebd37c0d8d2bd0bda989.form = Controller510f52c4102bebd37c0d8d2bd0bda989Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/special-equipment'
*/
const Controller0a68e3adf451fdf023b851319bd5513c = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller0a68e3adf451fdf023b851319bd5513c.url(options),
    method: 'get',
})

Controller0a68e3adf451fdf023b851319bd5513c.definition = {
    methods: ["get","head"],
    url: '/risk-monitoring/special-equipment',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/special-equipment'
*/
Controller0a68e3adf451fdf023b851319bd5513c.url = (options?: RouteQueryOptions) => {
    return Controller0a68e3adf451fdf023b851319bd5513c.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/special-equipment'
*/
Controller0a68e3adf451fdf023b851319bd5513c.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller0a68e3adf451fdf023b851319bd5513c.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/special-equipment'
*/
Controller0a68e3adf451fdf023b851319bd5513c.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller0a68e3adf451fdf023b851319bd5513c.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/special-equipment'
*/
const Controller0a68e3adf451fdf023b851319bd5513cForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller0a68e3adf451fdf023b851319bd5513c.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/special-equipment'
*/
Controller0a68e3adf451fdf023b851319bd5513cForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller0a68e3adf451fdf023b851319bd5513c.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-monitoring/special-equipment'
*/
Controller0a68e3adf451fdf023b851319bd5513cForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller0a68e3adf451fdf023b851319bd5513c.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller0a68e3adf451fdf023b851319bd5513c.form = Controller0a68e3adf451fdf023b851319bd5513cForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/monthly-investigation'
*/
const Controller6e23cb11d764b5c6456e065297a4acc8 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller6e23cb11d764b5c6456e065297a4acc8.url(options),
    method: 'get',
})

Controller6e23cb11d764b5c6456e065297a4acc8.definition = {
    methods: ["get","head"],
    url: '/risk-control/monthly-investigation',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/monthly-investigation'
*/
Controller6e23cb11d764b5c6456e065297a4acc8.url = (options?: RouteQueryOptions) => {
    return Controller6e23cb11d764b5c6456e065297a4acc8.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/monthly-investigation'
*/
Controller6e23cb11d764b5c6456e065297a4acc8.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller6e23cb11d764b5c6456e065297a4acc8.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/monthly-investigation'
*/
Controller6e23cb11d764b5c6456e065297a4acc8.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller6e23cb11d764b5c6456e065297a4acc8.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/monthly-investigation'
*/
const Controller6e23cb11d764b5c6456e065297a4acc8Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller6e23cb11d764b5c6456e065297a4acc8.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/monthly-investigation'
*/
Controller6e23cb11d764b5c6456e065297a4acc8Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller6e23cb11d764b5c6456e065297a4acc8.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/risk-control/monthly-investigation'
*/
Controller6e23cb11d764b5c6456e065297a4acc8Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller6e23cb11d764b5c6456e065297a4acc8.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller6e23cb11d764b5c6456e065297a4acc8.form = Controller6e23cb11d764b5c6456e065297a4acc8Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training/relevant-party'
*/
const Controller2f78b8d5c566025525721122f4a14782 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller2f78b8d5c566025525721122f4a14782.url(options),
    method: 'get',
})

Controller2f78b8d5c566025525721122f4a14782.definition = {
    methods: ["get","head"],
    url: '/management/training/relevant-party',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training/relevant-party'
*/
Controller2f78b8d5c566025525721122f4a14782.url = (options?: RouteQueryOptions) => {
    return Controller2f78b8d5c566025525721122f4a14782.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training/relevant-party'
*/
Controller2f78b8d5c566025525721122f4a14782.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller2f78b8d5c566025525721122f4a14782.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training/relevant-party'
*/
Controller2f78b8d5c566025525721122f4a14782.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller2f78b8d5c566025525721122f4a14782.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training/relevant-party'
*/
const Controller2f78b8d5c566025525721122f4a14782Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller2f78b8d5c566025525721122f4a14782.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training/relevant-party'
*/
Controller2f78b8d5c566025525721122f4a14782Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller2f78b8d5c566025525721122f4a14782.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training/relevant-party'
*/
Controller2f78b8d5c566025525721122f4a14782Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller2f78b8d5c566025525721122f4a14782.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller2f78b8d5c566025525721122f4a14782.form = Controller2f78b8d5c566025525721122f4a14782Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/knowledge/checklists'
*/
const Controller12a499929eef7a553e9d06c0921f63cf = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller12a499929eef7a553e9d06c0921f63cf.url(options),
    method: 'get',
})

Controller12a499929eef7a553e9d06c0921f63cf.definition = {
    methods: ["get","head"],
    url: '/management/knowledge/checklists',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/knowledge/checklists'
*/
Controller12a499929eef7a553e9d06c0921f63cf.url = (options?: RouteQueryOptions) => {
    return Controller12a499929eef7a553e9d06c0921f63cf.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/knowledge/checklists'
*/
Controller12a499929eef7a553e9d06c0921f63cf.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller12a499929eef7a553e9d06c0921f63cf.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/knowledge/checklists'
*/
Controller12a499929eef7a553e9d06c0921f63cf.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller12a499929eef7a553e9d06c0921f63cf.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/knowledge/checklists'
*/
const Controller12a499929eef7a553e9d06c0921f63cfForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller12a499929eef7a553e9d06c0921f63cf.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/knowledge/checklists'
*/
Controller12a499929eef7a553e9d06c0921f63cfForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller12a499929eef7a553e9d06c0921f63cf.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/knowledge/checklists'
*/
Controller12a499929eef7a553e9d06c0921f63cfForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller12a499929eef7a553e9d06c0921f63cf.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller12a499929eef7a553e9d06c0921f63cf.form = Controller12a499929eef7a553e9d06c0921f63cfForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/meeting-notices'
*/
const Controllerac1c5f5d630e4641bd1707bf8a32ed73 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerac1c5f5d630e4641bd1707bf8a32ed73.url(options),
    method: 'get',
})

Controllerac1c5f5d630e4641bd1707bf8a32ed73.definition = {
    methods: ["get","head"],
    url: '/management/tasks/meeting-notices',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/meeting-notices'
*/
Controllerac1c5f5d630e4641bd1707bf8a32ed73.url = (options?: RouteQueryOptions) => {
    return Controllerac1c5f5d630e4641bd1707bf8a32ed73.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/meeting-notices'
*/
Controllerac1c5f5d630e4641bd1707bf8a32ed73.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerac1c5f5d630e4641bd1707bf8a32ed73.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/meeting-notices'
*/
Controllerac1c5f5d630e4641bd1707bf8a32ed73.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerac1c5f5d630e4641bd1707bf8a32ed73.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/meeting-notices'
*/
const Controllerac1c5f5d630e4641bd1707bf8a32ed73Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerac1c5f5d630e4641bd1707bf8a32ed73.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/meeting-notices'
*/
Controllerac1c5f5d630e4641bd1707bf8a32ed73Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerac1c5f5d630e4641bd1707bf8a32ed73.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/meeting-notices'
*/
Controllerac1c5f5d630e4641bd1707bf8a32ed73Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerac1c5f5d630e4641bd1707bf8a32ed73.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerac1c5f5d630e4641bd1707bf8a32ed73.form = Controllerac1c5f5d630e4641bd1707bf8a32ed73Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/notices'
*/
const Controller441a422118278822cfddf410836a05d6 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller441a422118278822cfddf410836a05d6.url(options),
    method: 'get',
})

Controller441a422118278822cfddf410836a05d6.definition = {
    methods: ["get","head"],
    url: '/management/tasks/notices',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/notices'
*/
Controller441a422118278822cfddf410836a05d6.url = (options?: RouteQueryOptions) => {
    return Controller441a422118278822cfddf410836a05d6.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/notices'
*/
Controller441a422118278822cfddf410836a05d6.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller441a422118278822cfddf410836a05d6.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/notices'
*/
Controller441a422118278822cfddf410836a05d6.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller441a422118278822cfddf410836a05d6.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/notices'
*/
const Controller441a422118278822cfddf410836a05d6Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller441a422118278822cfddf410836a05d6.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/notices'
*/
Controller441a422118278822cfddf410836a05d6Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller441a422118278822cfddf410836a05d6.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/notices'
*/
Controller441a422118278822cfddf410836a05d6Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller441a422118278822cfddf410836a05d6.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller441a422118278822cfddf410836a05d6.form = Controller441a422118278822cfddf410836a05d6Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/disposal-notices'
*/
const Controller84ddade82b418e05afccc274c652bda7 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller84ddade82b418e05afccc274c652bda7.url(options),
    method: 'get',
})

Controller84ddade82b418e05afccc274c652bda7.definition = {
    methods: ["get","head"],
    url: '/management/tasks/disposal-notices',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/disposal-notices'
*/
Controller84ddade82b418e05afccc274c652bda7.url = (options?: RouteQueryOptions) => {
    return Controller84ddade82b418e05afccc274c652bda7.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/disposal-notices'
*/
Controller84ddade82b418e05afccc274c652bda7.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller84ddade82b418e05afccc274c652bda7.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/disposal-notices'
*/
Controller84ddade82b418e05afccc274c652bda7.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller84ddade82b418e05afccc274c652bda7.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/disposal-notices'
*/
const Controller84ddade82b418e05afccc274c652bda7Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller84ddade82b418e05afccc274c652bda7.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/disposal-notices'
*/
Controller84ddade82b418e05afccc274c652bda7Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller84ddade82b418e05afccc274c652bda7.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/disposal-notices'
*/
Controller84ddade82b418e05afccc274c652bda7Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller84ddade82b418e05afccc274c652bda7.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller84ddade82b418e05afccc274c652bda7.form = Controller84ddade82b418e05afccc274c652bda7Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/manuscripts'
*/
const Controllerf334fa74cfc27e2ed4d9b32a6da3080e = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerf334fa74cfc27e2ed4d9b32a6da3080e.url(options),
    method: 'get',
})

Controllerf334fa74cfc27e2ed4d9b32a6da3080e.definition = {
    methods: ["get","head"],
    url: '/management/tasks/manuscripts',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/manuscripts'
*/
Controllerf334fa74cfc27e2ed4d9b32a6da3080e.url = (options?: RouteQueryOptions) => {
    return Controllerf334fa74cfc27e2ed4d9b32a6da3080e.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/manuscripts'
*/
Controllerf334fa74cfc27e2ed4d9b32a6da3080e.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllerf334fa74cfc27e2ed4d9b32a6da3080e.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/manuscripts'
*/
Controllerf334fa74cfc27e2ed4d9b32a6da3080e.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllerf334fa74cfc27e2ed4d9b32a6da3080e.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/manuscripts'
*/
const Controllerf334fa74cfc27e2ed4d9b32a6da3080eForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerf334fa74cfc27e2ed4d9b32a6da3080e.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/manuscripts'
*/
Controllerf334fa74cfc27e2ed4d9b32a6da3080eForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerf334fa74cfc27e2ed4d9b32a6da3080e.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/tasks/manuscripts'
*/
Controllerf334fa74cfc27e2ed4d9b32a6da3080eForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllerf334fa74cfc27e2ed4d9b32a6da3080e.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllerf334fa74cfc27e2ed4d9b32a6da3080e.form = Controllerf334fa74cfc27e2ed4d9b32a6da3080eForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/accidents/management'
*/
const Controller7cfe7c5d0ed7402dbd15cb776967fcbb = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller7cfe7c5d0ed7402dbd15cb776967fcbb.url(options),
    method: 'get',
})

Controller7cfe7c5d0ed7402dbd15cb776967fcbb.definition = {
    methods: ["get","head"],
    url: '/management/accidents/management',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/accidents/management'
*/
Controller7cfe7c5d0ed7402dbd15cb776967fcbb.url = (options?: RouteQueryOptions) => {
    return Controller7cfe7c5d0ed7402dbd15cb776967fcbb.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/accidents/management'
*/
Controller7cfe7c5d0ed7402dbd15cb776967fcbb.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller7cfe7c5d0ed7402dbd15cb776967fcbb.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/accidents/management'
*/
Controller7cfe7c5d0ed7402dbd15cb776967fcbb.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller7cfe7c5d0ed7402dbd15cb776967fcbb.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/accidents/management'
*/
const Controller7cfe7c5d0ed7402dbd15cb776967fcbbForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller7cfe7c5d0ed7402dbd15cb776967fcbb.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/accidents/management'
*/
Controller7cfe7c5d0ed7402dbd15cb776967fcbbForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller7cfe7c5d0ed7402dbd15cb776967fcbb.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/accidents/management'
*/
Controller7cfe7c5d0ed7402dbd15cb776967fcbbForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller7cfe7c5d0ed7402dbd15cb776967fcbb.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller7cfe7c5d0ed7402dbd15cb776967fcbb.form = Controller7cfe7c5d0ed7402dbd15cb776967fcbbForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/accidents/report'
*/
const Controller6f47e59cd3d5aee91d3bcbe032e9b174 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller6f47e59cd3d5aee91d3bcbe032e9b174.url(options),
    method: 'get',
})

Controller6f47e59cd3d5aee91d3bcbe032e9b174.definition = {
    methods: ["get","head"],
    url: '/management/accidents/report',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/accidents/report'
*/
Controller6f47e59cd3d5aee91d3bcbe032e9b174.url = (options?: RouteQueryOptions) => {
    return Controller6f47e59cd3d5aee91d3bcbe032e9b174.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/accidents/report'
*/
Controller6f47e59cd3d5aee91d3bcbe032e9b174.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller6f47e59cd3d5aee91d3bcbe032e9b174.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/accidents/report'
*/
Controller6f47e59cd3d5aee91d3bcbe032e9b174.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller6f47e59cd3d5aee91d3bcbe032e9b174.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/accidents/report'
*/
const Controller6f47e59cd3d5aee91d3bcbe032e9b174Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller6f47e59cd3d5aee91d3bcbe032e9b174.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/accidents/report'
*/
Controller6f47e59cd3d5aee91d3bcbe032e9b174Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller6f47e59cd3d5aee91d3bcbe032e9b174.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/accidents/report'
*/
Controller6f47e59cd3d5aee91d3bcbe032e9b174Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller6f47e59cd3d5aee91d3bcbe032e9b174.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller6f47e59cd3d5aee91d3bcbe032e9b174.form = Controller6f47e59cd3d5aee91d3bcbe032e9b174Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/experts'
*/
const Controller66b01f163eabd030b9419fd506d25262 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller66b01f163eabd030b9419fd506d25262.url(options),
    method: 'get',
})

Controller66b01f163eabd030b9419fd506d25262.definition = {
    methods: ["get","head"],
    url: '/management/experts',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/experts'
*/
Controller66b01f163eabd030b9419fd506d25262.url = (options?: RouteQueryOptions) => {
    return Controller66b01f163eabd030b9419fd506d25262.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/experts'
*/
Controller66b01f163eabd030b9419fd506d25262.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller66b01f163eabd030b9419fd506d25262.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/experts'
*/
Controller66b01f163eabd030b9419fd506d25262.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller66b01f163eabd030b9419fd506d25262.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/experts'
*/
const Controller66b01f163eabd030b9419fd506d25262Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller66b01f163eabd030b9419fd506d25262.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/experts'
*/
Controller66b01f163eabd030b9419fd506d25262Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller66b01f163eabd030b9419fd506d25262.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/experts'
*/
Controller66b01f163eabd030b9419fd506d25262Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller66b01f163eabd030b9419fd506d25262.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller66b01f163eabd030b9419fd506d25262.form = Controller66b01f163eabd030b9419fd506d25262Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/filing'
*/
const Controllercf5a69d7a5e5b87d2d641f8968f288a2 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllercf5a69d7a5e5b87d2d641f8968f288a2.url(options),
    method: 'get',
})

Controllercf5a69d7a5e5b87d2d641f8968f288a2.definition = {
    methods: ["get","head"],
    url: '/management/assessment/filing',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/filing'
*/
Controllercf5a69d7a5e5b87d2d641f8968f288a2.url = (options?: RouteQueryOptions) => {
    return Controllercf5a69d7a5e5b87d2d641f8968f288a2.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/filing'
*/
Controllercf5a69d7a5e5b87d2d641f8968f288a2.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controllercf5a69d7a5e5b87d2d641f8968f288a2.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/filing'
*/
Controllercf5a69d7a5e5b87d2d641f8968f288a2.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controllercf5a69d7a5e5b87d2d641f8968f288a2.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/filing'
*/
const Controllercf5a69d7a5e5b87d2d641f8968f288a2Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercf5a69d7a5e5b87d2d641f8968f288a2.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/filing'
*/
Controllercf5a69d7a5e5b87d2d641f8968f288a2Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercf5a69d7a5e5b87d2d641f8968f288a2.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/filing'
*/
Controllercf5a69d7a5e5b87d2d641f8968f288a2Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controllercf5a69d7a5e5b87d2d641f8968f288a2.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controllercf5a69d7a5e5b87d2d641f8968f288a2.form = Controllercf5a69d7a5e5b87d2d641f8968f288a2Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/scoring'
*/
const Controller9ac0e423d7f3ebbe9913d45fdad82a6a = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller9ac0e423d7f3ebbe9913d45fdad82a6a.url(options),
    method: 'get',
})

Controller9ac0e423d7f3ebbe9913d45fdad82a6a.definition = {
    methods: ["get","head"],
    url: '/management/assessment/scoring',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/scoring'
*/
Controller9ac0e423d7f3ebbe9913d45fdad82a6a.url = (options?: RouteQueryOptions) => {
    return Controller9ac0e423d7f3ebbe9913d45fdad82a6a.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/scoring'
*/
Controller9ac0e423d7f3ebbe9913d45fdad82a6a.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller9ac0e423d7f3ebbe9913d45fdad82a6a.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/scoring'
*/
Controller9ac0e423d7f3ebbe9913d45fdad82a6a.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller9ac0e423d7f3ebbe9913d45fdad82a6a.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/scoring'
*/
const Controller9ac0e423d7f3ebbe9913d45fdad82a6aForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller9ac0e423d7f3ebbe9913d45fdad82a6a.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/scoring'
*/
Controller9ac0e423d7f3ebbe9913d45fdad82a6aForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller9ac0e423d7f3ebbe9913d45fdad82a6a.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/scoring'
*/
Controller9ac0e423d7f3ebbe9913d45fdad82a6aForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller9ac0e423d7f3ebbe9913d45fdad82a6a.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller9ac0e423d7f3ebbe9913d45fdad82a6a.form = Controller9ac0e423d7f3ebbe9913d45fdad82a6aForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/summary'
*/
const Controller9665b168c37bb54ea2d7f659dead1e45 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller9665b168c37bb54ea2d7f659dead1e45.url(options),
    method: 'get',
})

Controller9665b168c37bb54ea2d7f659dead1e45.definition = {
    methods: ["get","head"],
    url: '/management/assessment/summary',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/summary'
*/
Controller9665b168c37bb54ea2d7f659dead1e45.url = (options?: RouteQueryOptions) => {
    return Controller9665b168c37bb54ea2d7f659dead1e45.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/summary'
*/
Controller9665b168c37bb54ea2d7f659dead1e45.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller9665b168c37bb54ea2d7f659dead1e45.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/summary'
*/
Controller9665b168c37bb54ea2d7f659dead1e45.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller9665b168c37bb54ea2d7f659dead1e45.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/summary'
*/
const Controller9665b168c37bb54ea2d7f659dead1e45Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller9665b168c37bb54ea2d7f659dead1e45.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/summary'
*/
Controller9665b168c37bb54ea2d7f659dead1e45Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller9665b168c37bb54ea2d7f659dead1e45.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/assessment/summary'
*/
Controller9665b168c37bb54ea2d7f659dead1e45Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller9665b168c37bb54ea2d7f659dead1e45.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller9665b168c37bb54ea2d7f659dead1e45.form = Controller9665b168c37bb54ea2d7f659dead1e45Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/monthly-stats'
*/
const Controller368a6ff557a46fa5dd7fa19aad526840 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller368a6ff557a46fa5dd7fa19aad526840.url(options),
    method: 'get',
})

Controller368a6ff557a46fa5dd7fa19aad526840.definition = {
    methods: ["get","head"],
    url: '/management/statistics/monthly-stats',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/monthly-stats'
*/
Controller368a6ff557a46fa5dd7fa19aad526840.url = (options?: RouteQueryOptions) => {
    return Controller368a6ff557a46fa5dd7fa19aad526840.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/monthly-stats'
*/
Controller368a6ff557a46fa5dd7fa19aad526840.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller368a6ff557a46fa5dd7fa19aad526840.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/monthly-stats'
*/
Controller368a6ff557a46fa5dd7fa19aad526840.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller368a6ff557a46fa5dd7fa19aad526840.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/monthly-stats'
*/
const Controller368a6ff557a46fa5dd7fa19aad526840Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller368a6ff557a46fa5dd7fa19aad526840.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/monthly-stats'
*/
Controller368a6ff557a46fa5dd7fa19aad526840Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller368a6ff557a46fa5dd7fa19aad526840.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/monthly-stats'
*/
Controller368a6ff557a46fa5dd7fa19aad526840Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller368a6ff557a46fa5dd7fa19aad526840.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller368a6ff557a46fa5dd7fa19aad526840.form = Controller368a6ff557a46fa5dd7fa19aad526840Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/annual-stats'
*/
const Controller5c9c65e9922e242d3374a628468a34aa = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller5c9c65e9922e242d3374a628468a34aa.url(options),
    method: 'get',
})

Controller5c9c65e9922e242d3374a628468a34aa.definition = {
    methods: ["get","head"],
    url: '/management/statistics/annual-stats',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/annual-stats'
*/
Controller5c9c65e9922e242d3374a628468a34aa.url = (options?: RouteQueryOptions) => {
    return Controller5c9c65e9922e242d3374a628468a34aa.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/annual-stats'
*/
Controller5c9c65e9922e242d3374a628468a34aa.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller5c9c65e9922e242d3374a628468a34aa.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/annual-stats'
*/
Controller5c9c65e9922e242d3374a628468a34aa.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller5c9c65e9922e242d3374a628468a34aa.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/annual-stats'
*/
const Controller5c9c65e9922e242d3374a628468a34aaForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller5c9c65e9922e242d3374a628468a34aa.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/annual-stats'
*/
Controller5c9c65e9922e242d3374a628468a34aaForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller5c9c65e9922e242d3374a628468a34aa.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/annual-stats'
*/
Controller5c9c65e9922e242d3374a628468a34aaForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller5c9c65e9922e242d3374a628468a34aa.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller5c9c65e9922e242d3374a628468a34aa.form = Controller5c9c65e9922e242d3374a628468a34aaForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/notice-stats'
*/
const Controller020afc6709cfc581be0221902dd80483 = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller020afc6709cfc581be0221902dd80483.url(options),
    method: 'get',
})

Controller020afc6709cfc581be0221902dd80483.definition = {
    methods: ["get","head"],
    url: '/management/statistics/notice-stats',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/notice-stats'
*/
Controller020afc6709cfc581be0221902dd80483.url = (options?: RouteQueryOptions) => {
    return Controller020afc6709cfc581be0221902dd80483.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/notice-stats'
*/
Controller020afc6709cfc581be0221902dd80483.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller020afc6709cfc581be0221902dd80483.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/notice-stats'
*/
Controller020afc6709cfc581be0221902dd80483.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller020afc6709cfc581be0221902dd80483.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/notice-stats'
*/
const Controller020afc6709cfc581be0221902dd80483Form = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller020afc6709cfc581be0221902dd80483.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/notice-stats'
*/
Controller020afc6709cfc581be0221902dd80483Form.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller020afc6709cfc581be0221902dd80483.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/statistics/notice-stats'
*/
Controller020afc6709cfc581be0221902dd80483Form.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller020afc6709cfc581be0221902dd80483.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller020afc6709cfc581be0221902dd80483.form = Controller020afc6709cfc581be0221902dd80483Form
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training-center/upload-materials'
*/
const Controller1f5b7e7717418295b5f4adc75f36dfaf = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller1f5b7e7717418295b5f4adc75f36dfaf.url(options),
    method: 'get',
})

Controller1f5b7e7717418295b5f4adc75f36dfaf.definition = {
    methods: ["get","head"],
    url: '/management/training-center/upload-materials',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training-center/upload-materials'
*/
Controller1f5b7e7717418295b5f4adc75f36dfaf.url = (options?: RouteQueryOptions) => {
    return Controller1f5b7e7717418295b5f4adc75f36dfaf.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training-center/upload-materials'
*/
Controller1f5b7e7717418295b5f4adc75f36dfaf.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller1f5b7e7717418295b5f4adc75f36dfaf.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training-center/upload-materials'
*/
Controller1f5b7e7717418295b5f4adc75f36dfaf.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller1f5b7e7717418295b5f4adc75f36dfaf.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training-center/upload-materials'
*/
const Controller1f5b7e7717418295b5f4adc75f36dfafForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller1f5b7e7717418295b5f4adc75f36dfaf.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training-center/upload-materials'
*/
Controller1f5b7e7717418295b5f4adc75f36dfafForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller1f5b7e7717418295b5f4adc75f36dfaf.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training-center/upload-materials'
*/
Controller1f5b7e7717418295b5f4adc75f36dfafForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller1f5b7e7717418295b5f4adc75f36dfaf.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller1f5b7e7717418295b5f4adc75f36dfaf.form = Controller1f5b7e7717418295b5f4adc75f36dfafForm
/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training-center/upload-videos'
*/
const Controller7a27c3c56c21e939ff322471e3c03bbb = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller7a27c3c56c21e939ff322471e3c03bbb.url(options),
    method: 'get',
})

Controller7a27c3c56c21e939ff322471e3c03bbb.definition = {
    methods: ["get","head"],
    url: '/management/training-center/upload-videos',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training-center/upload-videos'
*/
Controller7a27c3c56c21e939ff322471e3c03bbb.url = (options?: RouteQueryOptions) => {
    return Controller7a27c3c56c21e939ff322471e3c03bbb.definition.url + queryParams(options)
}

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training-center/upload-videos'
*/
Controller7a27c3c56c21e939ff322471e3c03bbb.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: Controller7a27c3c56c21e939ff322471e3c03bbb.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training-center/upload-videos'
*/
Controller7a27c3c56c21e939ff322471e3c03bbb.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: Controller7a27c3c56c21e939ff322471e3c03bbb.url(options),
    method: 'head',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training-center/upload-videos'
*/
const Controller7a27c3c56c21e939ff322471e3c03bbbForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller7a27c3c56c21e939ff322471e3c03bbb.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training-center/upload-videos'
*/
Controller7a27c3c56c21e939ff322471e3c03bbbForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller7a27c3c56c21e939ff322471e3c03bbb.url(options),
    method: 'get',
})

/**
* @see \Inertia\Controller::__invoke
* @see vendor/inertiajs/inertia-laravel/src/Controller.php:13
* @route '/management/training-center/upload-videos'
*/
Controller7a27c3c56c21e939ff322471e3c03bbbForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: Controller7a27c3c56c21e939ff322471e3c03bbb.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

Controller7a27c3c56c21e939ff322471e3c03bbb.form = Controller7a27c3c56c21e939ff322471e3c03bbbForm

const Controller = {
    '/': Controller980bb49ee7ae63891f1d891d2fbcf1c9,
    '/login': Controllerb6041c76e8e1cd791f8f89d035d48611,
    '/register': Controllere9819db9819a1d19b38dd89a0c4218c4,
    '/first-login-setup': Controller8faf738c13d56f0246b14ec9e7e25ed8,
    '/dashboard': Controller42a740574ecbfbac32f8cc353fc32db9,
    '/profile': Controlleraed0d18b6699640aa4d85cb0451e9ce2,
    '/change-password': Controllerfdd726fe943ba5011c80134906e72889,
    '/enterprise/current': Controllerd1dbe9162bf030ca8d815a2896726ea8,
    '/enterprise/subsidiaries': Controllerdc63d3ff2633f051c4e6fafefcc65b36,
    '/enterprise/approval': Controllere8686b6a496a704bc4ba429ba6e98199,
    '/enterprise/grouping': Controllerc35191ff34467f1a013ce9a4c98651af,
    '/risk-monitoring/enterprise-map': Controller05aab3f771482e0f736ba647c68e1796,
    '/risk-monitoring/hazard-map': Controller41d0c8523849475797401e3bca6562af,
    '/risk-monitoring/online-inspection': Controller1ff7759eea7e5ed022e6f9862a7742b6,
    '/risk-monitoring/risk-monitor': Controller37f5f03d82b3e02f8541822ce4bb2798,
    '/risk-control/monthly-inspection': Controlleraeeb3a833b5a24551356fc46b501899c,
    '/risk-control/expert-inspection': Controller7172acbc864e98d79046caf7595a5c87,
    '/system/companies': Controller22317c48b944ba855e883e487b7b892e,
    '/system/users': Controller418f97c1428914f8efc179504e0cb069,
    '/system/roles': Controller8e4cd5574fbacbae03adf3fbb312404b,
    '/system/permissions': Controllera91f2d42251113e168610b33abb3d6a3,
    '/system/enterprise-types': Controller044adaaf6c398030da2a3d4b1ded58b3,
    '/system/operational-types': Controllere4b2217aeb9c91d75d471461a9e4519a,
    '/system/dangerous-craft-types': Controllercfa7b1d6544bfa7299202292496bd119,
    '/system/special-types': Controllerbc273d698d993e09426e4f7b4dced41b,
    '/system/process-types': Controllercc5be3c6fdf69815d72033596df0ee33,
    '/system/space-work-types': Controllere4caab3e61c1fa82f402eba68b119b58,
    '/system/safety-management-types': Controllerbc930121bc0a562e140d40c3767083fd,
    '/system/record-types': Controller942b9a21421566f7b3f8dd5be36f9fed,
    '/system/regulation-library-types': Controller6923a40b06398712bb2cc26d54b91ba5,
    '/system/assessment-templates': Controller2490d96901b856cc50b47539eeb24dc8,
    '/safety/buildings': Controller35c989abf185d02a8446cf8abc38b6ca,
    '/safety/job-management': Controller715956940b0d8010718e06152a395138,
    '/safety/personnels': Controllerec095dcfcb6f8fb5a6563900689f141a,
    '/safety/special-equipment': Controller99d56c37532773e1512d1bba444c5af9,
    '/safety/hazardous-chemicals': Controller703b76c5ec2c428d8d5d9020cc82dafa,
    '/safety/major-hazards': Controller403dfddf0437cb6cf052cb1b0209e94d,
    '/safety/fire-safeties': Controllera1dc36a9de9ac5e855e8a9dd7cf132dd,
    '/risk/hazards': Controller32cf2db60b9a6fb418346afc40b53a78,
    '/risk/risk-logs': Controller931dd43bc86e8b5355757ba1bddccc4d,
    '/risk/monthly-investigation': Controllera611b44ef0b191664a928e05a42a8ce4,
    '/risk/rectification-sheets': Controllerd566f42701af8ddf67cda28d9ac9db37,
    '/inspection/daily-checks': Controller848e8b566bd81867beab2986e6102b1d,
    '/inspection/system-checks': Controllera36f412fb20084ab7c9404fa6d8a5f46,
    '/inspection/norm-checks': Controllercf195f51fc323080aba270ea267a22a5,
    '/inspection/accident-checks': Controllere706152e0dcff487ab0e79404ee6f9c3,
    '/health/occupational-hazards': Controllercf17a5673925405b2682957c0642ef97,
    '/health/health-files': Controller00b00e6ebd69280348b8be7b977ff5e2,
    '/health/health-checkups': Controller7a6036087be56a47df63d7013588f45c,
    '/health/hygiene-evaluations': Controlleraa1f88d6fc5920f38d9e2eeb08ca6216,
    '/emergency/fire-safety': Controllerd40c2199ac32bc35ff00293ec9a60196,
    '/emergency/emergency-drills': Controller96c1bb13a3ad3f9b5edb57c77c96c216,
    '/emergency/emergency-supplies': Controllerf907fde89fb2af26dda0618330d90e33,
    '/emergency/emergency-teams': Controllerb5fc8dd61ced1b57e86be40c73008108,
    '/training/training-files': Controllera32e7ab4ac8fc9480b96291a84b03264,
    '/training/training-videos': Controller86903e311a24a371314611c417a8e28d,
    '/training/internal-trainings': Controllerb0d5689953570e04c957bb3ebe226451,
    '/training/external-trainings': Controllercc02a54a8a581f4b995a2da3573af252,
    '/training/other-trainings': Controller5e6f8aff4f45ec9e5a3d2073e62940fa,
    '/training/education': Controller6f7985dd84e48575c462555a011345b1,
    '/meeting/meetings': Controller7f834e4898e9bd3488b8c65a8277d5da,
    '/meeting/meeting-sign-ups': Controller09216ea80d8b90623529eb88c253494c,
    '/meeting/internal-meetings': Controller909bb78633799ac7f9687ea722d32ac2,
    '/statistics/month-statistics': Controllerf32fbe8105905cbebc888eeaca6761b4,
    '/statistics/year-statistics': Controller05d5d86ffee2d15dd9fe346c88e51a3e,
    '/statistics/expert-inspections': Controller3b431e7c6692756fbb94f688636d5142,
    '/statistics/accident-bulletins': Controller3987fe4e864d121bbc5e553a44f3f63e,
    '/regulation/libraries': Controllerb7157957ec9c2ec304b591d2cb37bbee,
    '/space-work/space-works': Controller766be9f9f27b6636ae9ccc917b8410b7,
    '/space-work/three-togethers': Controller8796c98bda408c975eec5ade64700017,
    '/safety-archives/org-structure': Controllerb7cd76195a4a9cf3f9f9d95da3eb09dc,
    '/safety-archives/qualifications': Controllerb33f068b9007af396523a675de9dfd1f,
    '/safety-archives/safety-investment': Controller40c606abc4d95e40c8a1bbe7a7b965f9,
    '/equipment/major-hazardous': Controllerbe8d43bd91af207ce1972a6ea9b07b97,
    '/process/key-hazardous': Controller108e5a2f69ee121654ee34809a401fb1,
    '/process/hazardous': Controller86338a2b8fa3273fd00718be7bfc535e,
    '/risk/safety-assessment': Controllercb5a69541b20685d1ed664d21d4eec41,
    '/risk/expert-inspection': Controller50a3e572e907c9d988298a8b7d433db8,
    '/risk/risk-warning': Controller4e52ccde3afaaf81c032199e92c5f1ab,
    '/third-party/design-units': Controller1f601ea8e5bb218ee28f6eda1db540b9,
    '/third-party/construction-units': Controller9380dca6c49302c2ee76e6228e538c16,
    '/third-party/technical-services': Controller0e40752e8ff2ba7acfe8001538e4d6eb,
    '/health/disease-report': Controllerfc972e8bb979ae574f7fff41328ab6b1,
    '/fire/equipment': Controllera6832eb7956939320272db69dba3a866,
    '/fire/facilities': Controller26dcb1a9e73d9c46d8ec1f08de7de20c,
    '/accident/report': Controller5095ce3dd4931a582d1b106234287d92,
    '/accident/emergency': Controllera776b56c843d3a3b51d2ed29d47bbe72,
    '/accident/emergency/teams': Controller1599ae2fdadec1c0124d7a7c33efec5b,
    '/accident/emergency/supplies': Controller93550ff8800eaaf91ef08224c1c678dd,
    '/accident/emergency/drill': Controller2f9e705b2aec86c00af0abdde41f6aa4,
    '/tasks/notices': Controller03594eab5c2ebeb813e15f5f2a9a3a6b,
    '/tasks/manuscripts': Controller421d2a8f3aeb1b2e342bef6c9b4f52e5,
    '/tasks/monthly-report': Controller8a91a2034574762b1b48654125caeb46,
    '/tasks/annual-report': Controller929e879457a8031e25756388a6fe8942,
    '/change-management': Controller97869e24dedf9bea8f73c16937915605,
    '/knowledge/checklists': Controller6e7265da7faf052741f8de13edbbea3b,
    '/assessment/filing': Controllerbe1bcb06c4f801f544587db5920cc9ba,
    '/account': Controller510f52c4102bebd37c0d8d2bd0bda989,
    '/risk-monitoring/special-equipment': Controller0a68e3adf451fdf023b851319bd5513c,
    '/risk-control/monthly-investigation': Controller6e23cb11d764b5c6456e065297a4acc8,
    '/management/training/relevant-party': Controller2f78b8d5c566025525721122f4a14782,
    '/management/knowledge/checklists': Controller12a499929eef7a553e9d06c0921f63cf,
    '/management/tasks/meeting-notices': Controllerac1c5f5d630e4641bd1707bf8a32ed73,
    '/management/tasks/notices': Controller441a422118278822cfddf410836a05d6,
    '/management/tasks/disposal-notices': Controller84ddade82b418e05afccc274c652bda7,
    '/management/tasks/manuscripts': Controllerf334fa74cfc27e2ed4d9b32a6da3080e,
    '/management/accidents/management': Controller7cfe7c5d0ed7402dbd15cb776967fcbb,
    '/management/accidents/report': Controller6f47e59cd3d5aee91d3bcbe032e9b174,
    '/management/experts': Controller66b01f163eabd030b9419fd506d25262,
    '/management/assessment/filing': Controllercf5a69d7a5e5b87d2d641f8968f288a2,
    '/management/assessment/scoring': Controller9ac0e423d7f3ebbe9913d45fdad82a6a,
    '/management/assessment/summary': Controller9665b168c37bb54ea2d7f659dead1e45,
    '/management/statistics/monthly-stats': Controller368a6ff557a46fa5dd7fa19aad526840,
    '/management/statistics/annual-stats': Controller5c9c65e9922e242d3374a628468a34aa,
    '/management/statistics/notice-stats': Controller020afc6709cfc581be0221902dd80483,
    '/management/training-center/upload-materials': Controller1f5b7e7717418295b5f4adc75f36dfaf,
    '/management/training-center/upload-videos': Controller7a27c3c56c21e939ff322471e3c03bbb,
}

export default Controller