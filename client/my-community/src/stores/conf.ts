import { defineStore } from "pinia";

export const useConfStore = defineStore("conf", {
    state: () => ({
        isLoading: false,
        edit: false
    }),
})