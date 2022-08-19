<template>
    <div class="user-edit-container">
        <div v-if="conf.edit">
            <edit />
        </div>
        <div v-else>
            <info />
        </div>
    </div>
</template>

<script lang="ts" setup>
import Info from "@/components/user/Info.vue"
import Edit from "@/components/user/Edit.vue"
import { useConfStore } from "@/stores/conf";
let conf = useConfStore()

</script>

<style scoped>
</style>