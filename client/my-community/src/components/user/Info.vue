<template>
    <div class="user-info">
        <el-row>
            <el-col :span="6">
                <el-avatar shape="square" :size="100" :fit="fit" :src="url" />
            </el-col>
            <el-col :span="6" :offset="6" style="margin-top: 30px;">
                <el-button type="success" @click="edit">
                    编辑
                </el-button>
            </el-col>
        </el-row>
        <div class="placeholder" />
        <el-row :gutter="20">
            <el-col :span="6">
                <el-row>
                    <el-tag size="large">用户ID</el-tag>
                    <div style="margin: 10px;" />
                    {{ userStore.info['id'] }}
                </el-row>
            </el-col>
            <el-col :span="6" :offset="6">
                <el-row>
                    <el-tag size="large">用户组</el-tag>
                    <div style="margin: 10px;" />
                    <div>{{ userStore.info['right'] }}</div>
                </el-row>
            </el-col>
        </el-row>
        <div class="placeholder" />
        <el-row :gutter="20">
            <el-col :span="6">
                <el-row>
                    <el-tag size="large">用户名</el-tag>
                    <div style="margin: 10px;" />
                    {{ userStore.info['username'] }}
                </el-row>
            </el-col>
            <el-col :span="6" :offset="6">
                <el-row>
                    <el-tag size="large">昵称</el-tag>
                    <div style="margin: 10px;" />
                    {{ userStore.info['alias'] }}
                </el-row>
            </el-col>
        </el-row>
        <div class="placeholder" />
        <el-row :gutter="20">
            <el-col :span="6">
                <el-row>
                    <el-tag size="large">手机号</el-tag>
                    <div style="margin: 10px;" />
                    {{ userStore.info['telephone'] }}
                </el-row>
            </el-col>
            <el-col :span="6" :offset="6">
                <el-row>
                    <el-tag size="large">邮箱</el-tag>
                    <div style="margin: 10px;" />
                    {{ userStore.info['email'] }}
                </el-row>
            </el-col>
        </el-row>
        <div class="placeholder" />
        <el-row :gutter="20">
            <el-col :span="6">
                <el-row>
                    <el-tag size="large">学校</el-tag>
                    <div style="margin: 10px;" />
                    {{ userStore.info['school'] }}
                </el-row>
            </el-col>
        </el-row>
    </div>
</template>

<script lang="ts" setup>
import { useConfStore } from '@/stores/conf'
import { useUserStore } from '@/stores/user'
import { GetStore, SetStoreWithBoolean } from '@/utils/store';
import { toRefs } from '@vueuse/core'
import { reactive } from 'vue'
const userStore = useUserStore()

if (window.localStorage) {
    let jwt = GetStore("jwt")
    if (typeof (jwt) === 'string' && jwt.length > 0) {
        userStore.login = true
        userStore.jwt = jwt
    }
    if (userStore.login && GetStore("edit").length > 0) {
        userStore.syncInfoWithNet()
        SetStoreWithBoolean("edit", false)
    }
    if (userStore.login && userStore.info.username === '') {
        userStore.syncInfoWithCache()
    }
}

const edit = () => {
    useConfStore().edit = true
}

const state = reactive({
    fit: 'fill',
    url: 'https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png',
})
if (userStore.info.profile === '1') {
    state.url = 'http://localhost:9090/images/' + userStore.info.username + '.png'
}
const { fit, url } = toRefs(state)
</script>

<style scoped>
.placeholder {
    margin: 20px;
}
</style>