import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\UploadController::upload
* @see app/Http/Controllers/Api/V1/UploadController.php:11
* @route '/api/v1/upload'
*/
export const upload = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: upload.url(options),
    method: 'post',
})

upload.definition = {
    methods: ["post"],
    url: '/api/v1/upload',
} satisfies RouteDefinition<["post"]>

/**
* @see \App\Http\Controllers\Api\V1\UploadController::upload
* @see app/Http/Controllers/Api/V1/UploadController.php:11
* @route '/api/v1/upload'
*/
upload.url = (options?: RouteQueryOptions) => {
    return upload.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\UploadController::upload
* @see app/Http/Controllers/Api/V1/UploadController.php:11
* @route '/api/v1/upload'
*/
upload.post = (options?: RouteQueryOptions): RouteDefinition<'post'> => ({
    url: upload.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadController::upload
* @see app/Http/Controllers/Api/V1/UploadController.php:11
* @route '/api/v1/upload'
*/
const uploadForm = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: upload.url(options),
    method: 'post',
})

/**
* @see \App\Http\Controllers\Api\V1\UploadController::upload
* @see app/Http/Controllers/Api/V1/UploadController.php:11
* @route '/api/v1/upload'
*/
uploadForm.post = (options?: RouteQueryOptions): RouteFormDefinition<'post'> => ({
    action: upload.url(options),
    method: 'post',
})

upload.form = uploadForm

const UploadController = { upload }

export default UploadController