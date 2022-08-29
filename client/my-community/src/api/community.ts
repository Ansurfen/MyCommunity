import service from '@/utils/net'

export const newCommunity = (data: string, jwt: string) => {
    return service.request({ url: '/community/new', method: 'post', data, headers: { 'Authorization': "Bearer " + jwt } })
}

export const searchCommunity = (data: FormData) => {
    return service.request({ url: '/community/search', method: 'post', data })
}

export const infoCommunity = (data: FormData, jwt: string) => {
    return service.request({ url: '/community/info', method: 'post', data, headers: { 'Authorization': "Bearer " + jwt } })
}

export const addCommunity = (data: string, jwt: string) => {
    return service.request({ url: '/community/add', method: 'post', data, headers: { 'Authorization': "Bearer " + jwt } })
}

export const editCommunity = (data: string, jwt: string) => {
    return service.request({ url: '/community/edit', method: 'post', data, headers: { 'Authorization': "Bearer " + jwt } })
}

export const addNote = (data: string, jwt: string) => {
    return service.request({ url: '/community/note/add', method: 'post', data, headers: { 'Authorization': "Bearer " + jwt } })
}

export const infoNote = (data: FormData) => {
    return service.request({ url: '/community/note/info', method: 'post', data })
}