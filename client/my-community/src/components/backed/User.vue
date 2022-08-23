<template>
    <div>
        <p style="text-align: left;">用户列表</p>
        <el-row justify="center" style="background-color:aliceblue;">
            <el-col :span="6">
                <div class="grid-content ep-bg-purple" />用户名（昵称）/ ID
            </el-col>
            <el-col :span="2">
                <div class="grid-content ep-bg-purple-light" />手机号
            </el-col>
            <el-col :span="2">
                <div class="grid-content ep-bg-purple-light" />邮箱
            </el-col>
            <el-col :span="2">
                <div class="grid-content ep-bg-purple-light" />学校
            </el-col>
            <el-col :span="2">
                <div class="grid-content ep-bg-purple-light" />用户组
            </el-col>
            <el-col :span="6">
                <div class="grid-content ep-bg-purple" />管理员操作
            </el-col>
        </el-row>
        <el-divider />
        <div v-for="(k, v) in users" :key="v">
            <el-row justify="center">
                <el-col :span="5">{{ k.username }}（{{ k.alias }}） / {{ k.id }} </el-col>
                <el-col :span="3">{{ k.telephone }}</el-col>
                <el-col :span="3">{{ k.email }}</el-col>
                <el-col :span="3">{{ k.school }}</el-col>
                <el-col :span="3" v-if="k.right === 0">超级管理员</el-col>
                <el-col :span="3" v-else-if="k.right === 1">管理员</el-col>
                <el-col :span="3" v-else-if="k.right === 2">普通用户</el-col>
                <el-col :span="4">
                    <!-- <el-select :v-model="rights[v]" class="m-2" placeholder="Select" size="large">
                            <el-option v-for="item in options" :key="item.value" :label="item.label"
                                :value="item.value" />
                        </el-select> -->
                    <el-button text @click="change(Flags.ADMIN, k.username)">设置管理</el-button>
                    <el-button text @click="change(Flags.None, k.username)">移除权限</el-button>
                </el-col>
            </el-row>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { userInfoCommit, changeCommit } from '@/api/commit'
import { Flags } from '@/models/constant'
import { useUserStore } from '@/stores/user'
import { UserInfo } from '@/models/user'
import { ref } from 'vue'
const userStore = useUserStore()
let rights = ref<number[]>([])
const options = [
    {
        value: Flags.ROOT,
        label: '超级管理员',
    },
    {
        value: Flags.ADMIN,
        label: '管理员',
    },
    {
        value: Flags.None,
        label: '普通用户',
    }
]
let users = ref<UserInfo[]>([])
userInfoCommit(JSON.stringify({
    type: Flags.INFO,
    status: Flags.USERINFO
}), userStore.jwt).then(res => {
    users.value = JSON.parse(res['data']['data'].data)
    for (let i = 0; i < users.value.length; i++) {
        rights.value.push(users.value[i].right)
    }
}).catch(err => alert("你没有权限操作"))
const change = (right: Flags, username: string) => {
    // 得防止root自己设置自己
    changeCommit(JSON.stringify({
        type: Flags.CHANGE,
        status: right,
        first: username
    }), userStore.jwt).then(res => {

    }).catch(err => console.log(err))
}
</script>