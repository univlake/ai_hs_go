import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::getSignature
* @see app/Http/Controllers/Api/V1/UploadOssController.php:12
* @route '/api/v1/upload/oss-signature'
*/
export const getSignature = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: getSignature.url(options),
    method: 'get',
})

getSignature.definition = {
    methods: ["get","head"],
    url: '/api/v1/upload/oss-signature',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::getSignature
* @see app/Http/Controllers/Api/V1/UploadOssController.php:12
* @route '/api/v1/upload/oss-signature'
*/
getSignature.url = (options?: RouteQueryOptions) => {
    return getSignature.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::getSignature
* @see app/Http/Controllers/Api/V1/UploadOssController.php:12
* @route '/api/v1/upload/oss-signature'
*/
getSignature.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: getSignature.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::getSignature
* @see app/Http/Controllers/Api/V1/UploadOssController.php:12
* @route '/api/v1/upload/oss-signature'
*/
getSignature.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: getSignature.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::getSignature
* @see app/Http/Controllers/Api/V1/UploadOssController.php:12
* @route '/api/v1/upload/oss-signature'
*/
const getSignatureForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: getSignature.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::getSignature
* @see app/Http/Controllers/Api/V1/UploadOssController.php:12
* @route '/api/v1/upload/oss-signature'
*/
getSignatureForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: getSignature.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::getSignature
* @see app/Http/Controllers/Api/V1/UploadOssController.php:12
* @route '/api/v1/upload/oss-signature'
*/
getSignatureForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: getSignature.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

getSignature.form = getSignatureForm

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartInit
* @see app/Http/Controllers/Api/V1/UploadOssController.php:91
* @route '/api/v1/upload/multipart/init'
*/
export const multipartInit = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: multipartInit.url(options),
    method: 'post',
})

multipartInit.definition = {
    methods: ["post"],
    url: '/api/v1/upload/multipart/init',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartInit
* @see app/Http/Controllers/Api/V1/UploadOssController.php:91
* @route '/api/v1/upload/multipart/init'
*/
multipartInit.url = (options?: RouteQueryOptions) => {
    return multipartInit.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartInit
* @see app/Http/Controllers/Api/V1/UploadOssController.php:91
* @route '/api/v1/upload/multipart/init'
*/
multipartInit.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: multipartInit.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartInit
* @see app/Http/Controllers/Api/V1/UploadOssController.php:91
* @route '/api/v1/upload/multipart/init'
*/
const multipartInitForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: multipartInit.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartInit
* @see app/Http/Controllers/Api/V1/UploadOssController.php:91
* @route '/api/v1/upload/multipart/init'
*/
multipartInitForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: multipartInit.url(options),
    method: 'post',
})

multipartInit.form = multipartInitForm

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartSignPart
* @see app/Http/Controllers/Api/V1/UploadOssController.php:118
* @route '/api/v1/upload/multipart/sign-part'
*/
export const multipartSignPart = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: multipartSignPart.url(options),
    method: 'post',
})

multipartSignPart.definition = {
    methods: ["post"],
    url: '/api/v1/upload/multipart/sign-part',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartSignPart
* @see app/Http/Controllers/Api/V1/UploadOssController.php:118
* @route '/api/v1/upload/multipart/sign-part'
*/
multipartSignPart.url = (options?: RouteQueryOptions) => {
    return multipartSignPart.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartSignPart
* @see app/Http/Controllers/Api/V1/UploadOssController.php:118
* @route '/api/v1/upload/multipart/sign-part'
*/
multipartSignPart.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: multipartSignPart.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartSignPart
* @see app/Http/Controllers/Api/V1/UploadOssController.php:118
* @route '/api/v1/upload/multipart/sign-part'
*/
const multipartSignPartForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: multipartSignPart.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartSignPart
* @see app/Http/Controllers/Api/V1/UploadOssController.php:118
* @route '/api/v1/upload/multipart/sign-part'
*/
multipartSignPartForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: multipartSignPart.url(options),
    method: 'post',
})

multipartSignPart.form = multipartSignPartForm

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartComplete
* @see app/Http/Controllers/Api/V1/UploadOssController.php:135
* @route '/api/v1/upload/multipart/complete'
*/
export const multipartComplete = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: multipartComplete.url(options),
    method: 'post',
})

multipartComplete.definition = {
    methods: ["post"],
    url: '/api/v1/upload/multipart/complete',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartComplete
* @see app/Http/Controllers/Api/V1/UploadOssController.php:135
* @route '/api/v1/upload/multipart/complete'
*/
multipartComplete.url = (options?: RouteQueryOptions) => {
    return multipartComplete.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartComplete
* @see app/Http/Controllers/Api/V1/UploadOssController.php:135
* @route '/api/v1/upload/multipart/complete'
*/
multipartComplete.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: multipartComplete.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartComplete
* @see app/Http/Controllers/Api/V1/UploadOssController.php:135
* @route '/api/v1/upload/multipart/complete'
*/
const multipartCompleteForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: multipartComplete.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartComplete
* @see app/Http/Controllers/Api/V1/UploadOssController.php:135
* @route '/api/v1/upload/multipart/complete'
*/
multipartCompleteForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: multipartComplete.url(options),
    method: 'post',
})

multipartComplete.form = multipartCompleteForm

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartAbort
* @see app/Http/Controllers/Api/V1/UploadOssController.php:154
* @route '/api/v1/upload/multipart/abort'
*/
export const multipartAbort = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: multipartAbort.url(options),
    method: 'post',
})

multipartAbort.definition = {
    methods: ["post"],
    url: '/api/v1/upload/multipart/abort',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartAbort
* @see app/Http/Controllers/Api/V1/UploadOssController.php:154
* @route '/api/v1/upload/multipart/abort'
*/
multipartAbort.url = (options?: RouteQueryOptions) => {
    return multipartAbort.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartAbort
* @see app/Http/Controllers/Api/V1/UploadOssController.php:154
* @route '/api/v1/upload/multipart/abort'
*/
multipartAbort.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: multipartAbort.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartAbort
* @see app/Http/Controllers/Api/V1/UploadOssController.php:154
* @route '/api/v1/upload/multipart/abort'
*/
const multipartAbortForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: multipartAbort.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::multipartAbort
* @see app/Http/Controllers/Api/V1/UploadOssController.php:154
* @route '/api/v1/upload/multipart/abort'
*/
multipartAbortForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: multipartAbort.url(options),
    method: 'post',
})

multipartAbort.form = multipartAbortForm

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::signedUrl
* @see app/Http/Controllers/Api/V1/UploadOssController.php:82
* @route '/api/v1/upload/oss-sign'
*/
export const signedUrl = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: signedUrl.url(options),
    method: 'get',
})

signedUrl.definition = {
    methods: ["get","head"],
    url: '/api/v1/upload/oss-sign',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::signedUrl
* @see app/Http/Controllers/Api/V1/UploadOssController.php:82
* @route '/api/v1/upload/oss-sign'
*/
signedUrl.url = (options?: RouteQueryOptions) => {
    return signedUrl.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::signedUrl
* @see app/Http/Controllers/Api/V1/UploadOssController.php:82
* @route '/api/v1/upload/oss-sign'
*/
signedUrl.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: signedUrl.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::signedUrl
* @see app/Http/Controllers/Api/V1/UploadOssController.php:82
* @route '/api/v1/upload/oss-sign'
*/
signedUrl.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: signedUrl.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::signedUrl
* @see app/Http/Controllers/Api/V1/UploadOssController.php:82
* @route '/api/v1/upload/oss-sign'
*/
const signedUrlForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: signedUrl.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::signedUrl
* @see app/Http/Controllers/Api/V1/UploadOssController.php:82
* @route '/api/v1/upload/oss-sign'
*/
signedUrlForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: signedUrl.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::signedUrl
* @see app/Http/Controllers/Api/V1/UploadOssController.php:82
* @route '/api/v1/upload/oss-sign'
*/
signedUrlForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: signedUrl.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

signedUrl.form = signedUrlForm

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::callback
* @see app/Http/Controllers/Api/V1/UploadOssController.php:48
* @route '/api/v1/upload/oss-callback'
*/
export const callback = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: callback.url(options),
    method: 'post',
})

callback.definition = {
    methods: ["post"],
    url: '/api/v1/upload/oss-callback',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::callback
* @see app/Http/Controllers/Api/V1/UploadOssController.php:48
* @route '/api/v1/upload/oss-callback'
*/
callback.url = (options?: RouteQueryOptions) => {
    return callback.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::callback
* @see app/Http/Controllers/Api/V1/UploadOssController.php:48
* @route '/api/v1/upload/oss-callback'
*/
callback.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: callback.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::callback
* @see app/Http/Controllers/Api/V1/UploadOssController.php:48
* @route '/api/v1/upload/oss-callback'
*/
const callbackForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: callback.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadOssController::callback
* @see app/Http/Controllers/Api/V1/UploadOssController.php:48
* @route '/api/v1/upload/oss-callback'
*/
callbackForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: callback.url(options),
    method: 'post',
})

callback.form = callbackForm

const UploadOssController = { getSignature, multipartInit, multipartSignPart, multipartComplete, multipartAbort, signedUrl, callback }

export default UploadOssController