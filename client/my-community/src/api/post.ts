import service from "@/utils/net"

export const addPost = (data: string, jwt: string) => {
    return service.request({ url: '/community/post/add', method: 'post', data, headers: { 'Authorization': "Bearer " + jwt } })
}

export const delPost = (data: string, jwt: string) => {
    return service.request({ url: '/community/post/del', method: 'post', data, headers: { 'Authorization': "Bearer " + jwt } })
}

export const infoPost = (data: FormData) => {
    return service.request({ url: '/community/post/info', method: 'post', data })
}

export const infoPosts = (data: FormData) => {
    return service.request({ url: '/post/info', method: 'post', data })
}

export const addComment = (data: string, jwt: string) => {
    return service.request({ url: '/post/add', method: 'post', data, headers: { 'Authorization': "Bearer " + jwt } })
}

export const appendComment = (data: string, jwt: string) => {
    return service.request({ url: '/post/append', method: 'post', data, headers: { 'Authorization': "Bearer " + jwt } })
}

export const delComment = (data: string, jwt: string) => {
    return service.request({ url: '/post/del', method: 'post', data, headers: { 'Authorization': "Bearer " + jwt } })
}