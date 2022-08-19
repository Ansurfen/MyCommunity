import service from '@/utils/net'

export const newCommunity = (data: string, jwt: string) => {
    return service.request({ url: '/community/new', method: 'post', data, headers: { 'Authorization': "Bearer " + jwt } })
}

export const searchCommunity = (data: FormData) => {
    return service.request({ url: '/community/search', method: 'post', data })
}

export const infoCommunity = (data: FormData) => {
    return service.request({ url: '/community/info', method: 'post', data })
}