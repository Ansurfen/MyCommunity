<template>
    <div class="card-container">
        <el-skeleton :rows="4" :loading="loading" style="text-align: left;width: 900px;padding-left: 300px;" animated>
            <template #template>
            </template>
            <template #default>
                <el-card class="card" @click="main(community)">
                    <el-row :gutter="20">
                        <el-col :span="4">
                            <el-avatar shape="square" :size="80" :fit="fit" :src="url" />
                        </el-col>
                        <el-col :span="16">
                            <div style="text-align: left;">{{ community?.name }}</div>
                            <div style="text-align: left;">{{ community?.context }}</div>
                        </el-col>
                    </el-row>
                </el-card>
            </template>
        </el-skeleton>
    </div>
</template>

<script lang="ts" setup>
import { PropType, reactive, toRefs } from 'vue';
import { Community } from '@/models/community'
import { useRouter } from 'vue-router';
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
    height: 150px;
    width: 600px;
}
</style>