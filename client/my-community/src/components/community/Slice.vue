<template>
    <div class="slice-container">
        <div v-if="communities.length <= 0">
            <el-card class="null-card animate__animated animate__tada">
                <el-col>
                    <h1>空空如也</h1>
                    <h2>未能找到查询的社团，也可能未能及时刷新搜索栏（尝试使用Enter刷新），或者点击搜索栏的右侧新建社团吧</h2>
                </el-col>
            </el-card>
        </div>
        <div v-else>
            <div v-for="(community, index) in communities" :key="index">
                <!-- <div v-if="(index + 1) % 2 === 0">
                    <el-divider />
                </div> -->
                <card style="margin-top: 30px;" :community="community" :loading="loading" />
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { PropType } from 'vue';
import { Community } from '@/models/community'
import Card from '@/components/community/Card.vue'
const loading = false
defineProps({
    communities: {
        type: Array as PropType<Community[]>,
        default: []
    }
})
</script>

<style scoped>
.slice-container {
    display: flex;
    align-items: center;
    justify-content: center;
}

.null-card {
    margin-top: 100px;
    height: 220px;
    width: 660px;
    z-index: 1;
    overflow: hidden;
    backdrop-filter: blur(10px);
    background-color: rgba(255, 255, 255, 0.72);
    text-align: center;
}
</style>