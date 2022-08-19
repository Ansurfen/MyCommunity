import { RandInt } from "@/utils/rand";
import { TagProps } from "element-plus";

export type Tag = { type: TagProps['type']; label: string }

export type Tags = string[]

export enum EType {
    NONE,
    SUCCESS,
    INFO,
    WARNING,
    DANGER
}

export const RandEType = (): "" | "success" | "warning" | "info" | "danger" => {
    switch (RandInt(4)) {
        case EType.SUCCESS: return "success"
        case EType.INFO: return "info"
        case EType.WARNING: return "warning"
        case EType.DANGER: return "danger"
        case EType.NONE:
        default:
    }
    return ""
}