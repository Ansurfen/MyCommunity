<template>
    <div class="slice-container">
        <div v-if="communities.length <= 0">
            <h1 style="margin-top: 30px;">空空如也</h1>
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
</style>