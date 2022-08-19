export const SetStore = (k: string, v: string) => {
    if (window.localStorage) {
        let storage = window.localStorage
        storage.setItem(k, v)
    } else {
        alert("该平台不支持此浏览器")
    }
}

export const GetStore = (k: string): string => {
    if (window.localStorage) {
        let storage = window.localStorage
        let value = storage.getItem(k)
        if (typeof (value) === 'string' && value.length > 0) {
            return value as string
        }
    } else {
        alert("该平台不支持此浏览器")
    }
    return ""
}