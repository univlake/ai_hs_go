import { queryParams, type RouteQueryOptions, type RouteDefinition, type RouteFormDefinition } from './../../../../../../wayfinder'
/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::personnelCertificate
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:11
* @route '/api/v1/online-inspection-personnel-certificate'
*/
export const personnelCertificate = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: personnelCertificate.url(options),
    method: 'get',
})

personnelCertificate.definition = {
    methods: ["get","head"],
    url: '/api/v1/online-inspection-personnel-certificate',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::personnelCertificate
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:11
* @route '/api/v1/online-inspection-personnel-certificate'
*/
personnelCertificate.url = (options?: RouteQueryOptions) => {
    return personnelCertificate.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::personnelCertificate
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:11
* @route '/api/v1/online-inspection-personnel-certificate'
*/
personnelCertificate.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: personnelCertificate.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::personnelCertificate
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:11
* @route '/api/v1/online-inspection-personnel-certificate'
*/
personnelCertificate.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: personnelCertificate.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::personnelCertificate
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:11
* @route '/api/v1/online-inspection-personnel-certificate'
*/
const personnelCertificateForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: personnelCertificate.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::personnelCertificate
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:11
* @route '/api/v1/online-inspection-personnel-certificate'
*/
personnelCertificateForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: personnelCertificate.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::personnelCertificate
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:11
* @route '/api/v1/online-inspection-personnel-certificate'
*/
personnelCertificateForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: personnelCertificate.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

personnelCertificate.form = personnelCertificateForm

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::personnelCertificateExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:18
* @route '/api/v1/online-inspection-personnel-certificate-expand'
*/
export const personnelCertificateExpand = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: personnelCertificateExpand.url(options),
    method: 'get',
})

personnelCertificateExpand.definition = {
    methods: ["get","head"],
    url: '/api/v1/online-inspection-personnel-certificate-expand',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::personnelCertificateExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:18
* @route '/api/v1/online-inspection-personnel-certificate-expand'
*/
personnelCertificateExpand.url = (options?: RouteQueryOptions) => {
    return personnelCertificateExpand.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::personnelCertificateExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:18
* @route '/api/v1/online-inspection-personnel-certificate-expand'
*/
personnelCertificateExpand.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: personnelCertificateExpand.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::personnelCertificateExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:18
* @route '/api/v1/online-inspection-personnel-certificate-expand'
*/
personnelCertificateExpand.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: personnelCertificateExpand.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::personnelCertificateExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:18
* @route '/api/v1/online-inspection-personnel-certificate-expand'
*/
const personnelCertificateExpandForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: personnelCertificateExpand.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::personnelCertificateExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:18
* @route '/api/v1/online-inspection-personnel-certificate-expand'
*/
personnelCertificateExpandForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: personnelCertificateExpand.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::personnelCertificateExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:18
* @route '/api/v1/online-inspection-personnel-certificate-expand'
*/
personnelCertificateExpandForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: personnelCertificateExpand.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

personnelCertificateExpand.form = personnelCertificateExpandForm

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::equipment
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:25
* @route '/api/v1/online-inspection-equipment'
*/
export const equipment = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: equipment.url(options),
    method: 'get',
})

equipment.definition = {
    methods: ["get","head"],
    url: '/api/v1/online-inspection-equipment',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::equipment
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:25
* @route '/api/v1/online-inspection-equipment'
*/
equipment.url = (options?: RouteQueryOptions) => {
    return equipment.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::equipment
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:25
* @route '/api/v1/online-inspection-equipment'
*/
equipment.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: equipment.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::equipment
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:25
* @route '/api/v1/online-inspection-equipment'
*/
equipment.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: equipment.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::equipment
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:25
* @route '/api/v1/online-inspection-equipment'
*/
const equipmentForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: equipment.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::equipment
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:25
* @route '/api/v1/online-inspection-equipment'
*/
equipmentForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: equipment.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::equipment
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:25
* @route '/api/v1/online-inspection-equipment'
*/
equipmentForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: equipment.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

equipment.form = equipmentForm

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::equipmentExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:32
* @route '/api/v1/online-inspection-equipment-expand'
*/
export const equipmentExpand = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: equipmentExpand.url(options),
    method: 'get',
})

equipmentExpand.definition = {
    methods: ["get","head"],
    url: '/api/v1/online-inspection-equipment-expand',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::equipmentExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:32
* @route '/api/v1/online-inspection-equipment-expand'
*/
equipmentExpand.url = (options?: RouteQueryOptions) => {
    return equipmentExpand.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::equipmentExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:32
* @route '/api/v1/online-inspection-equipment-expand'
*/
equipmentExpand.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: equipmentExpand.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::equipmentExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:32
* @route '/api/v1/online-inspection-equipment-expand'
*/
equipmentExpand.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: equipmentExpand.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::equipmentExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:32
* @route '/api/v1/online-inspection-equipment-expand'
*/
const equipmentExpandForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: equipmentExpand.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::equipmentExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:32
* @route '/api/v1/online-inspection-equipment-expand'
*/
equipmentExpandForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: equipmentExpand.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::equipmentExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:32
* @route '/api/v1/online-inspection-equipment-expand'
*/
equipmentExpandForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: equipmentExpand.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

equipmentExpand.form = equipmentExpandForm

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::monthlyInvestigation
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:39
* @route '/api/v1/online-inspection-monthly-investigation'
*/
export const monthlyInvestigation = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: monthlyInvestigation.url(options),
    method: 'get',
})

monthlyInvestigation.definition = {
    methods: ["get","head"],
    url: '/api/v1/online-inspection-monthly-investigation',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::monthlyInvestigation
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:39
* @route '/api/v1/online-inspection-monthly-investigation'
*/
monthlyInvestigation.url = (options?: RouteQueryOptions) => {
    return monthlyInvestigation.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::monthlyInvestigation
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:39
* @route '/api/v1/online-inspection-monthly-investigation'
*/
monthlyInvestigation.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: monthlyInvestigation.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::monthlyInvestigation
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:39
* @route '/api/v1/online-inspection-monthly-investigation'
*/
monthlyInvestigation.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: monthlyInvestigation.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::monthlyInvestigation
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:39
* @route '/api/v1/online-inspection-monthly-investigation'
*/
const monthlyInvestigationForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: monthlyInvestigation.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::monthlyInvestigation
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:39
* @route '/api/v1/online-inspection-monthly-investigation'
*/
monthlyInvestigationForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: monthlyInvestigation.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::monthlyInvestigation
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:39
* @route '/api/v1/online-inspection-monthly-investigation'
*/
monthlyInvestigationForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: monthlyInvestigation.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

monthlyInvestigation.form = monthlyInvestigationForm

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::monthlyInvestigationExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:46
* @route '/api/v1/online-inspection-monthly-investigation-expand'
*/
export const monthlyInvestigationExpand = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: monthlyInvestigationExpand.url(options),
    method: 'get',
})

monthlyInvestigationExpand.definition = {
    methods: ["get","head"],
    url: '/api/v1/online-inspection-monthly-investigation-expand',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::monthlyInvestigationExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:46
* @route '/api/v1/online-inspection-monthly-investigation-expand'
*/
monthlyInvestigationExpand.url = (options?: RouteQueryOptions) => {
    return monthlyInvestigationExpand.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::monthlyInvestigationExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:46
* @route '/api/v1/online-inspection-monthly-investigation-expand'
*/
monthlyInvestigationExpand.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: monthlyInvestigationExpand.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::monthlyInvestigationExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:46
* @route '/api/v1/online-inspection-monthly-investigation-expand'
*/
monthlyInvestigationExpand.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: monthlyInvestigationExpand.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::monthlyInvestigationExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:46
* @route '/api/v1/online-inspection-monthly-investigation-expand'
*/
const monthlyInvestigationExpandForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: monthlyInvestigationExpand.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::monthlyInvestigationExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:46
* @route '/api/v1/online-inspection-monthly-investigation-expand'
*/
monthlyInvestigationExpandForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: monthlyInvestigationExpand.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::monthlyInvestigationExpand
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:46
* @route '/api/v1/online-inspection-monthly-investigation-expand'
*/
monthlyInvestigationExpandForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: monthlyInvestigationExpand.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

monthlyInvestigationExpand.form = monthlyInvestigationExpandForm

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::productionStandardization
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:53
* @route '/api/v1/online-inspection-production-standardization'
*/
export const productionStandardization = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: productionStandardization.url(options),
    method: 'get',
})

productionStandardization.definition = {
    methods: ["get","head"],
    url: '/api/v1/online-inspection-production-standardization',
} satisfies RouteDefinition<["get","head"]>

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::productionStandardization
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:53
* @route '/api/v1/online-inspection-production-standardization'
*/
productionStandardization.url = (options?: RouteQueryOptions) => {
    return productionStandardization.definition.url + queryParams(options)
}

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::productionStandardization
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:53
* @route '/api/v1/online-inspection-production-standardization'
*/
productionStandardization.get = (options?: RouteQueryOptions): RouteDefinition<'get'> => ({
    url: productionStandardization.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::productionStandardization
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:53
* @route '/api/v1/online-inspection-production-standardization'
*/
productionStandardization.head = (options?: RouteQueryOptions): RouteDefinition<'head'> => ({
    url: productionStandardization.url(options),
    method: 'head',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::productionStandardization
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:53
* @route '/api/v1/online-inspection-production-standardization'
*/
const productionStandardizationForm = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: productionStandardization.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::productionStandardization
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:53
* @route '/api/v1/online-inspection-production-standardization'
*/
productionStandardizationForm.get = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: productionStandardization.url(options),
    method: 'get',
})

/**
* @see \App\Http\Controllers\Api\V1\OnlineInspectionController::productionStandardization
* @see app/Http/Controllers/Api/V1/OnlineInspectionController.php:53
* @route '/api/v1/online-inspection-production-standardization'
*/
productionStandardizationForm.head = (options?: RouteQueryOptions): RouteFormDefinition<'get'> => ({
    action: productionStandardization.url({
        [options?.mergeQuery ? 'mergeQuery' : 'query']: {
            _method: 'HEAD',
            ...(options?.query ?? options?.mergeQuery ?? {}),
        }
    }),
    method: 'get',
})

productionStandardization.form = productionStandardizationForm

const OnlineInspectionController = { personnelCertificate, personnelCertificateExpand, equipment, equipmentExpand, monthlyInvestigation, monthlyInvestigationExpand, productionStandardization }

export default OnlineInspectionController