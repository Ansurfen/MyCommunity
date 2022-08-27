import service from '@/utils/net'

export const userLogin = (data: string) => {
    return service.request({ url: '/user/login', method: 'post', data })
}

export const userRegister = (data: FormData) => {
    return service.request({ url: '/user/register', method: 'post', data })
}

export const userInfo = (jwt: string) => {
    return service.request({ url: '/user/info', method: 'get', headers: { 'Authorization': "Bearer " + jwt } })
}

export const userEdit = (data: string, jwt: string) => {
    return service.request({ url: '/user/edit/info', method: 'post', data, headers: { 'Authorization': "Bearer " + jwt } })
}

export const imageUpdate = (data: FormData, jwt: string) => {
    return service.request({ url: '/user/image/update', method: 'post', data, headers: { 'Content-Type': 'multipart/form-data', 'Authorization': "Bearer " + jwt } })
}