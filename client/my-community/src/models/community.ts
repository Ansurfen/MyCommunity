export interface Community {
    id: string,
    name: string,
    context: string,
    tags: string,
    image: string,
    timestamp: string,
    hostname: string,
    admins: string,
    members: string,
    notes: string,
    posts: string,
    status: number
}

export interface FCommunity {
    id: string,
    name: string,
    context: string,
    tags: string[],
    image: string,
    timestamp: string,
    hostname: string,
    admins: string,
    members: string,
    notes: string,
    posts: string,
    status: number
}

export interface Applications {
    first: string,
    second: string,
    context: string,
    status: number,
    type: number
}

export interface Post {
    id: string,
    title: string,
    author: string,
    timestamp: string,
    tags: string,
    context: string,
    score: number
}