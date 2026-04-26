import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\AuthController::publicKey
* @see app/Http/Controllers/Api/V1/AuthController.php:25
* @route '/api/v1/public-key'
*/
export const publicKey = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: publicKey.url(options),
    method: 'get',
})

publicKey.definition = {
    methods: ["get","head"],
    url: '/api/v1/public-key',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\AuthController::publicKey
* @see app/Http/Controllers/Api/V1/AuthController.php:25
* @route '/api/v1/public-key'
*/
publicKey.url = (options?: RouteQueryOptions) => {
    return publicKey.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AuthController::publicKey
* @see app/Http/Controllers/Api/V1/AuthController.php:25
* @route '/api/v1/public-key'
*/
publicKey.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: publicKey.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::publicKey
* @see app/Http/Controllers/Api/V1/AuthController.php:25
* @route '/api/v1/public-key'
*/
publicKey.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: publicKey.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::publicKey
* @see app/Http/Controllers/Api/V1/AuthController.php:25
* @route '/api/v1/public-key'
*/
const publicKeyForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: publicKey.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::publicKey
* @see app/Http/Controllers/Api/V1/AuthController.php:25
* @route '/api/v1/public-key'
*/
publicKeyForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: publicKey.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::publicKey
* @see app/Http/Controllers/Api/V1/AuthController.php:25
* @route '/api/v1/public-key'
*/
publicKeyForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: publicKey.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

publicKey.form = publicKeyForm

/**
* @see \App\Http\Controllers\Api\V1\AuthController::login
* @see app/Http/Controllers/Api/V1/AuthController.php:33
* @route '/api/v1/login'
*/
export const login = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: login.url(options),
    method: 'post',
})

login.definition = {
    methods: ["post"],
    url: '/api/v1/login',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\AuthController::login
* @see app/Http/Controllers/Api/V1/AuthController.php:33
* @route '/api/v1/login'
*/
login.url = (options?: RouteQueryOptions) => {
    return login.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AuthController::login
* @see app/Http/Controllers/Api/V1/AuthController.php:33
* @route '/api/v1/login'
*/
login.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: login.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::login
* @see app/Http/Controllers/Api/V1/AuthController.php:33
* @route '/api/v1/login'
*/
const loginForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: login.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::login
* @see app/Http/Controllers/Api/V1/AuthController.php:33
* @route '/api/v1/login'
*/
loginForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: login.url(options),
    method: 'post',
})

login.form = loginForm

/**
* @see \App\Http\Controllers\Api\V1\AuthController::register
* @see app/Http/Controllers/Api/V1/AuthController.php:73
* @route '/api/v1/register'
*/
export const register = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: register.url(options),
    method: 'post',
})

register.definition = {
    methods: ["post"],
    url: '/api/v1/register',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\AuthController::register
* @see app/Http/Controllers/Api/V1/AuthController.php:73
* @route '/api/v1/register'
*/
register.url = (options?: RouteQueryOptions) => {
    return register.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AuthController::register
* @see app/Http/Controllers/Api/V1/AuthController.php:73
* @route '/api/v1/register'
*/
register.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: register.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::register
* @see app/Http/Controllers/Api/V1/AuthController.php:73
* @route '/api/v1/register'
*/
const registerForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: register.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::register
* @see app/Http/Controllers/Api/V1/AuthController.php:73
* @route '/api/v1/register'
*/
registerForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: register.url(options),
    method: 'post',
})

register.form = registerForm

/**
* @see \App\Http\Controllers\Api\V1\AuthController::logout
* @see app/Http/Controllers/Api/V1/AuthController.php:124
* @route '/api/v1/logout'
*/
export const logout = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: logout.url(options),
    method: 'post',
})

logout.definition = {
    methods: ["post"],
    url: '/api/v1/logout',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\AuthController::logout
* @see app/Http/Controllers/Api/V1/AuthController.php:124
* @route '/api/v1/logout'
*/
logout.url = (options?: RouteQueryOptions) => {
    return logout.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AuthController::logout
* @see app/Http/Controllers/Api/V1/AuthController.php:124
* @route '/api/v1/logout'
*/
logout.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: logout.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::logout
* @see app/Http/Controllers/Api/V1/AuthController.php:124
* @route '/api/v1/logout'
*/
const logoutForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: logout.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::logout
* @see app/Http/Controllers/Api/V1/AuthController.php:124
* @route '/api/v1/logout'
*/
logoutForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: logout.url(options),
    method: 'post',
})

logout.form = logoutForm

/**
* @see \App\Http\Controllers\Api\V1\AuthController::userInfo
* @see app/Http/Controllers/Api/V1/AuthController.php:134
* @route '/api/v1/user-info'
*/
export const userInfo = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: userInfo.url(options),
    method: 'get',
})

userInfo.definition = {
    methods: ["get","head"],
    url: '/api/v1/user-info',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\AuthController::userInfo
* @see app/Http/Controllers/Api/V1/AuthController.php:134
* @route '/api/v1/user-info'
*/
userInfo.url = (options?: RouteQueryOptions) => {
    return userInfo.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AuthController::userInfo
* @see app/Http/Controllers/Api/V1/AuthController.php:134
* @route '/api/v1/user-info'
*/
userInfo.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: userInfo.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::userInfo
* @see app/Http/Controllers/Api/V1/AuthController.php:134
* @route '/api/v1/user-info'
*/
userInfo.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: userInfo.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::userInfo
* @see app/Http/Controllers/Api/V1/AuthController.php:134
* @route '/api/v1/user-info'
*/
const userInfoForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: userInfo.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::userInfo
* @see app/Http/Controllers/Api/V1/AuthController.php:134
* @route '/api/v1/user-info'
*/
userInfoForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: userInfo.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::userInfo
* @see app/Http/Controllers/Api/V1/AuthController.php:134
* @route '/api/v1/user-info'
*/
userInfoForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: userInfo.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

userInfo.form = userInfoForm

/**
* @see \App\Http\Controllers\Api\V1\AuthController::updateUserInfo
* @see app/Http/Controllers/Api/V1/AuthController.php:157
* @route '/api/v1/user-info'
*/
export const updateUserInfo = (options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: updateUserInfo.url(options),
    method: 'put',
})

updateUserInfo.definition = {
    methods: ["put"],
    url: '/api/v1/user-info',
} satisfies RouteDefinition<["put"]>

/**
* @see \App\Http\Controllers\Api\V1\AuthController::updateUserInfo
* @see app/Http/Controllers/Api/V1/AuthController.php:157
* @route '/api/v1/user-info'
*/
updateUserInfo.url = (options?: RouteQueryOptions) => {
    return updateUserInfo.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AuthController::updateUserInfo
* @see app/Http/Controllers/Api/V1/AuthController.php:157
* @route '/api/v1/user-info'
*/
updateUserInfo.put = (options?: RouteQueryOptions): RouteDefinition<'put'> => ({
    url: updateUserInfo.url(options),
    method: 'put',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::updateUserInfo
* @see app/Http/Controllers/Api/V1/AuthController.php:157
* @route '/api/v1/user-info'
*/
const updateUserInfoForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: updateUserInfo.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'PUT',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::updateUserInfo
* @see app/Http/Controllers/Api/V1/AuthController.php:157
* @route '/api/v1/user-info'
*/
updateUserInfoForm.put = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: updateUserInfo.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'PUT',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'post',
})

updateUserInfo.form = updateUserInfoForm

/**
* @see \App\Http\Controllers\Api\V1\AuthController::checkUsernameAvailability
* @see app/Http/Controllers/Api/V1/AuthController.php:206
* @route '/api/v1/username-availability'
*/
export const checkUsernameAvailability = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: checkUsernameAvailability.url(options),
    method: 'get',
})

checkUsernameAvailability.definition = {
    methods: ["get","head"],
    url: '/api/v1/username-availability',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\AuthController::checkUsernameAvailability
* @see app/Http/Controllers/Api/V1/AuthController.php:206
* @route '/api/v1/username-availability'
*/
checkUsernameAvailability.url = (options?: RouteQueryOptions) => {
    return checkUsernameAvailability.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AuthController::checkUsernameAvailability
* @see app/Http/Controllers/Api/V1/AuthController.php:206
* @route '/api/v1/username-availability'
*/
checkUsernameAvailability.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: checkUsernameAvailability.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::checkUsernameAvailability
* @see app/Http/Controllers/Api/V1/AuthController.php:206
* @route '/api/v1/username-availability'
*/
checkUsernameAvailability.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: checkUsernameAvailability.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::checkUsernameAvailability
* @see app/Http/Controllers/Api/V1/AuthController.php:206
* @route '/api/v1/username-availability'
*/
const checkUsernameAvailabilityForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: checkUsernameAvailability.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::checkUsernameAvailability
* @see app/Http/Controllers/Api/V1/AuthController.php:206
* @route '/api/v1/username-availability'
*/
checkUsernameAvailabilityForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: checkUsernameAvailability.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::checkUsernameAvailability
* @see app/Http/Controllers/Api/V1/AuthController.php:206
* @route '/api/v1/username-availability'
*/
checkUsernameAvailabilityForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: checkUsernameAvailability.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

checkUsernameAvailability.form = checkUsernameAvailabilityForm

/**
* @see \App\Http\Controllers\Api\V1\AuthController::firstLoginSetup
* @see app/Http/Controllers/Api/V1/AuthController.php:225
* @route '/api/v1/first-login-setup'
*/
export const firstLoginSetup = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: firstLoginSetup.url(options),
    method: 'post',
})

firstLoginSetup.definition = {
    methods: ["post"],
    url: '/api/v1/first-login-setup',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\AuthController::firstLoginSetup
* @see app/Http/Controllers/Api/V1/AuthController.php:225
* @route '/api/v1/first-login-setup'
*/
firstLoginSetup.url = (options?: RouteQueryOptions) => {
    return firstLoginSetup.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AuthController::firstLoginSetup
* @see app/Http/Controllers/Api/V1/AuthController.php:225
* @route '/api/v1/first-login-setup'
*/
firstLoginSetup.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: firstLoginSetup.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::firstLoginSetup
* @see app/Http/Controllers/Api/V1/AuthController.php:225
* @route '/api/v1/first-login-setup'
*/
const firstLoginSetupForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: firstLoginSetup.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::firstLoginSetup
* @see app/Http/Controllers/Api/V1/AuthController.php:225
* @route '/api/v1/first-login-setup'
*/
firstLoginSetupForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: firstLoginSetup.url(options),
    method: 'post',
})

firstLoginSetup.form = firstLoginSetupForm

/**
* @see \App\Http\Controllers\Api\V1\AuthController::resetPassword
* @see app/Http/Controllers/Api/V1/AuthController.php:268
* @route '/api/v1/reset-password'
*/
export const resetPassword = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: resetPassword.url(options),
    method: 'post',
})

resetPassword.definition = {
    methods: ["post"],
    url: '/api/v1/reset-password',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\AuthController::resetPassword
* @see app/Http/Controllers/Api/V1/AuthController.php:268
* @route '/api/v1/reset-password'
*/
resetPassword.url = (options?: RouteQueryOptions) => {
    return resetPassword.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\AuthController::resetPassword
* @see app/Http/Controllers/Api/V1/AuthController.php:268
* @route '/api/v1/reset-password'
*/
resetPassword.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: resetPassword.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::resetPassword
* @see app/Http/Controllers/Api/V1/AuthController.php:268
* @route '/api/v1/reset-password'
*/
const resetPasswordForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: resetPassword.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\AuthController::resetPassword
* @see app/Http/Controllers/Api/V1/AuthController.php:268
* @route '/api/v1/reset-password'
*/
resetPasswordForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: resetPassword.url(options),
    method: 'post',
})

resetPassword.form = resetPasswordForm

const AuthController = { publicKey, login, register, logout, userInfo, updateUserInfo, checkUsernameAvailability, firstLoginSetup, resetPassword }

export default AuthController