export interface Posts {
    id: string,
    title: string,
    author: string,
    context: string,
    timestamp: string,
    tags: string,
    comments: string
}

// interface comment
export interface IComment {
    host: string,
    context: string,
    timestamp: string,
    sub: string
}

export interface Comment {
    host: string,
    context: string,
    timestamp: string,
    sub: sub[],
    level: number
}

interface sub {
    from: string,
    to: string,
    context: string,
    timestamp: string
}