<template>
    <div>
        <el-menu-item :index="index" style="background-color:#2d3133">
            <slot></slot>
            <span>{{ title }}</span>
        </el-menu-item>
    </div>
</template>

<script lang="ts" setup>
defineProps({
    index: {
        default: function () {
            return ''
        },
        type: String
    },
    selected: {
        default: function () {
            return false
        },
        type: Boolean
    },
    title: {
        default: function () {
            return ''
        },
        type: String
    }
})
</script>