import { defineStore } from "pinia";

export const useConfStore = defineStore("conf", {
    state: () => ({
        isLoading: false,
        edit: false,
        // http://xxx.xxx.xxx.xxx:9090/
        server: 'http://127.0.0.1:9090/'
    }),
})