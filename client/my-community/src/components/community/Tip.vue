<template>
    <div class="tip">
        <el-button text @click="open(title, context)">{{ title }}</el-button>
    </div>
</template>

<script lang="ts" setup>
import { ElMessageBox } from 'element-plus';

defineProps({
    title: { type: String, default: "" },
    context: { type: String, default: "" }
})
const open = (title: String, context: String) => {
    ElMessageBox.confirm(
        context.toString(),
        title.toString(),
        {
            confirmButtonText: 'ok',
            cancelButtonText: '返回',
            type: 'info',
            center: true,
        }
    )
    // .then(() => {
    //     ElMessage({
    //         type: 'success',
    //         message: 'Delete completed',
    //     })
    // })
    // .catch(() => {
    //     ElMessage({
    //         type: 'info',
    //         message: 'Delete canceled',
    //     })
    // })
}
</script>