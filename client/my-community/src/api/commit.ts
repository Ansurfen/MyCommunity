import service from "@/utils/net"

export const passCommit = (data: string, jwt: string) => {
    return service.request({ url: '/commit/pass', method: 'post', data, headers: { 'Authorization': "Bearer " + jwt } })
}

export const rejectCommit = (data: string, jwt: string) => {
    return service.request({ url: '/commit/reject', method: 'post', data, headers: { 'Authorization': "Bearer " + jwt } })
}

export const communityInfoCommit = (data: string, jwt: string) => {
    return service.request({ url: '/commit/info', method: 'post', data, headers: { 'Authorization': "Bearer " + jwt } })
}

export const userInfoCommit = (data: string, jwt: string) => {
    return service.request({ url: '/commit/info/user', method: 'post', data, headers: { 'Authorization': "Bearer " + jwt } })
}

export const changeCommit = (data: string, jwt: string) => {
    return service.request({ url: '/commit/change', method: 'post', data, headers: { 'Authorization': "Bearer " + jwt } })
}
