<template>
    <div class="card-container animate__animated animate__fadeInUp">
        <el-skeleton :rows="4" :loading="loading" style="text-align: left;width: 900px;padding-left: 300px;" animated>
            <template #template>
            </template>
            <template #default>
                <el-card class="card">
                    <el-col>
                        <el-row>
                            <el-col :span="5">
                                <el-avatar shape="square" :size="60" :fit="fit" :src="url" />
                            </el-col>
                            <el-col :span="10">
                                <div style="text-align: left;color: #1e1e1e;">{{ community?.name }}</div>
                                <div style="text-align: left;color: #1e1e1e;">所有者: {{ community?.hostname }}</div>
                            </el-col>
                            <el-col :span="8">
                                <div style="border-radius: 30px;background-color: #ebfffc;width: 220px;">
                                    <p style="color:#2e7f74;">
                                        创建时间: {{ FormatTime(community?.timestamp as string) }}
                                    </p>
                                </div>
                            </el-col>
                        </el-row>
                        <el-row style="color: #1e1e1e;height: 100px;">
                            {{ community?.context }}
                        </el-row>
                        <el-row>
                            <el-col :span="19">
                                <el-row>
                                    <div style="margin: 5px;" v-for="(v, i) in community?.tags" :key="i">
                                        <el-tag v-if="i === 0" type="warning" effect="dark">{{ v }}</el-tag>
                                        <el-tag v-else-if="i % 2 === 0" type="" effect="dark">{{ v }}</el-tag>
                                        <el-tag v-else-if="i % 3 === 0" type="danger" effect="dark">{{ v }}</el-tag>
                                        <el-tag v-else-if="i % 5 === 0" type="info" effect="dark">{{ v }}</el-tag>
                                        <el-tag v-else type="success" effect="dark">{{ v }}</el-tag>
                                    </div>
                                </el-row>

                            </el-col>
                            <el-col :span="4">
                                <el-button type="primary" @click="main(community)">查看详情</el-button>
                            </el-col>
                        </el-row>
                    </el-col>
                </el-card>
            </template>
        </el-skeleton>
    </div>
</template>

<script lang="ts" setup>
import { PropType, reactive, ref, toRefs } from 'vue';
import { Community } from '@/models/community'
import { useRouter } from 'vue-router';
import { FormatTime } from '@/utils/time'

defineProps({
    community: Object as PropType<Community>,
    loading: {
        default: function () {
            return true
        },
        type: Boolean
    }
})
const router = useRouter()
const main = (community: Community | undefined) => {
    const cname = community?.name as String
    router.push({ name: "community/main", params: { name: cname.toString(), community: JSON.stringify(community) } })
}
const state = reactive({
    fit: 'fill',
    url: 'https://cube.elemecdn.com/e/fd/0fc7d20532fdaf769a25683617711png.png',
})
const { fit, url } = toRefs(state)
</script>

<style scoped>
.card {
    height: 220px;
    width: 660px;
    z-index: 1;
    overflow: hidden;
    backdrop-filter: blur(10px);
    background-color: rgba(255, 255, 255, 0.72);
}
</style>