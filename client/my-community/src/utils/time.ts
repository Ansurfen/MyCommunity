export const FormatTime = (time: string) => {
    return new Date(Number(time)).toLocaleString().replace(/:\d{1,2}$/, ' ')
}