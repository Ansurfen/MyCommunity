<template>
    <div class="detail-card animate__animated animate__fadeInUp" style="margin-top:10px;">
        <el-skeleton :rows="4" :loading="loading" style="text-align: left;width: 500px;" animated>
            <template #template>
            </template>
            <template #default>
                <el-card class="box-card" shadow="hover">
                    <template #header>
                        <el-row :gutter="20" style="height: 30px;">
                            <el-col :span="10">
                                <p style="text-align: left;">{{ title }}</p>
                                <el-row style="height: 5px;">
                                    <p style="font-size: small;">作者: {{ author }}</p>
                                    &nbsp; &nbsp; &nbsp;
                                    <p style="font-size: small;">发布时间: {{ FormatTime(timestamp) }}</p>
                                </el-row>
                            </el-col>
                            <el-col :span="4" style="padding-left: 250px;">
                                <el-button @click="detail(id)" type="primary">查看详情</el-button>
                            </el-col>
                        </el-row>
                    </template>
                    <p style="text-align: left;height: 75px;">{{ context }}</p>
                    <div class="card-bottom">
                        <div v-for="(tag, i) in tags" :key="i" style="float: left;margin: 2px;" effect="dark">
                            <el-tag v-if="i === 0" type="warning" effect="dark">{{ tag }}</el-tag>
                            <el-tag v-else-if="i % 2 === 0" type="" effect="dark">{{ tag }}</el-tag>
                            <el-tag v-else-if="i % 3 === 0" type="danger" effect="dark">{{ tag }}</el-tag>
                            <el-tag v-else-if="i % 5 === 0" type="info" effect="dark">{{ tag }}</el-tag>
                            <el-tag v-else type="success" effect="dark">{{ tag }}</el-tag>
                        </div>
                        <el-rate disabled v-model="value" style="float: right;" allow-half />
                    </div>
                </el-card>
            </template>
        </el-skeleton>
    </div>
</template>

<script lang="ts" setup>
import { FormatTime } from '@/utils/time';
import { PropType, ref } from 'vue'
import { useRouter } from 'vue-router'

defineProps({
    title: { type: String, default: '' },
    context: { type: String, default: '' },
    author: { type: String, default: '' },
    timestamp: { type: String, default: '' },
    id: { type: String, default: '' },
    tags: { type: Object as PropType<String[]> }
})

const loading = false
const value = ref(0)
const router = useRouter()

const detail = (id: string) => {
    if (id.length > 0) {
        router.push({ name: "community/detail", params: { id: id } })
    }
}
</script>

<style scoped>
.box-card {
    width: 660px;
    height: 230px;
    z-index: 1;
    overflow: hidden;
    backdrop-filter: blur(10px);
    background-color: rgba(255, 255, 255, 0.72);
}
</style>