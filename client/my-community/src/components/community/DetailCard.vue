<template>
    <div class="detail-card">
        <el-skeleton :rows="4" :loading="loading" style="text-align: left;width: 500px;" animated>
            <template #template>
            </template>
            <template #default>
                <el-card class="box-card" shadow="hover" @click="detail(id)">
                    <template #header>
                        <div class="card-header">
                            <span>{{ title }}</span>
                            <span>{{ author }}</span>
                            <span>发布时间: {{ timestamp }}</span>
                        </div>
                    </template>
                    <span style="text-align: left;">{{ context }}</span>
                    <div class="card-bottom">
                        <el-tag v-for="tag in tags" :key="tag.label" :type="tag.type" style="float: left;margin: 2px;"
                            effect="dark">
                            {{ tag.label }}
                        </el-tag>
                        <el-rate disabled v-model="value" style="float: right;" allow-half />
                    </div>
                </el-card>
            </template>
        </el-skeleton>
    </div>
</template>

<script lang="ts" setup>
import { Tag } from '@/models/common'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
const loading = false
const value = ref(0)
const router = useRouter()
const tags = ref<Array<Tag>>([
    { type: '', label: '趣味' },
    { type: 'danger', label: '热门' }
])
defineProps({
    title: { type: String, default: '' },
    context: { type: String, default: '' },
    author: { type: String, default: '' },
    timestamp: { type: String, default: '' },
    id: { type: String, default: '' }
})
const detail = (id: string) => {
    if (id.length > 0) {
        router.push({ name: "community/detail", params: { id: id } })
    }
}
</script>

<style scoped>
.box-card {
    width: 480px;
}

.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}
</style>