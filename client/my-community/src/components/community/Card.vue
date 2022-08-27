<template>
    <div class="card-container">
        <el-skeleton :rows="4" :loading="loading" style="text-align: left;width: 900px;padding-left: 300px;" animated>
            <template #template>
            </template>
            <template #default>
                <el-card class="card">
                    <el-col>
                        <el-row>
                            <el-col :span="5">
                                <el-skeleton-item v-if="community?.image === ''" variant="image"
                                    style="width: 80px; height: 50px;" />
                                <el-avatar v-else shape="square" :size="80" :fit="fit" :src="url" />
                            </el-col>
                            <el-col :span="10">
                                <div style="text-align: left;">{{ community?.name }}</div>
                                <div style="text-align: left;">创始人: {{ community?.hostname }}</div>
                            </el-col>
                            <el-col :span="8">
                                <span style="border-radius: 30px;color: #4e8e2f;">
                                    <div>
                                        创建时间: {{ FormatTime(community?.timestamp as string) }}
                                    </div>
                                </span>
                            </el-col>
                        </el-row>
                        <el-row style="color: grey;height: 100px;">
                            {{ community?.context }}
                        </el-row>
                        <el-row>
                            <el-col :span="19">
                                <div v-for="(k, v) in community?.tags" :key="v">
                                    {{ k }}
                                </div>
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
import { emit } from 'process';
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
    url: 'https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png',
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