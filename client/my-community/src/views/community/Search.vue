<template>
    <div class="search-container">
        <el-container>
            <el-header>
                <home-nav></home-nav>
                <el-row justify="center">
                    <el-input v-model="input" placeholder="Please input" style="width: 800px;height: 50px;"
                        @submit.native.prevent @keyup.enter="search">
                        <template #prepend>
                            <el-button :icon="Search" />
                        </template>
                        <template #append>
                            <el-button @click="dialogFormVisible = true">+</el-button>
                            <el-dialog v-model="dialogFormVisible" title="新建社团申请">
                                <el-form :model="form">
                                    <el-form-item label="社团名称">
                                        <el-input v-model="form.name" autocomplete="off" />
                                    </el-form-item>
                                    <el-form-item label="社团简介">
                                        <el-input type="textarea" :autosize="{ minRows: 4, maxRows: 8 }"
                                            v-model="form.context" placeholder="输入不少于20字的简介" autocomplete="off" />
                                    </el-form-item>
                                </el-form>
                                <template #footer>
                                    <span class="dialog-footer">
                                        <el-button @click="dialogFormVisible = false" style="margin-right: 30px;">返回
                                        </el-button>
                                        <el-button type="primary" @click="commit">提交申请</el-button>
                                    </span>
                                </template>
                            </el-dialog>
                        </template>
                    </el-input>
                </el-row>
            </el-header>
            <el-main style="margin-top: 50px;">
                <slice :communities="communities" />
            </el-main>
            <el-footer>Copyright By Ansurfen</el-footer>
        </el-container>
    </div>
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router'
import Slice from '@/components/community/Slice.vue'
import { reactive, ref } from 'vue'
import { ElNotification } from 'element-plus'
import { newCommunity, searchCommunity } from '@/api/community'
import { useUserStore } from '@/stores/user'
import { Flags } from '@/models/constant'
import HomeNav from '@/components/Home/HomeNav.vue'
import { Search } from '@element-plus/icons-vue'
import { Community } from '@/models/community'

const router = useRoute()
const userStore = useUserStore()
const dialogFormVisible = ref(false)
const input = ref(router.params['str'].toString())
let communities = ref<Community[]>([])

const form = reactive({
    name: '',
    context: '',
})

const commit = () => {
    // 需要校验合法性，例如名字等，已经检测提交的申请数量
    dialogFormVisible.value = false
    newCommunity(JSON.stringify({
        type: Flags.NEW,
        first: form.name,
        second: userStore.info['username'],
        context: form.context
    }), userStore.jwt).then((res) => {
        ElNotification({
            title: 'Success',
            message: res['data']['msg'],
            type: 'success'
        })
    }).catch((err) => {
        ElNotification({
            title: 'Error',
            message: err['response']['data']['msg'],
            type: 'error'
        })
    })
}

const search = () => {
    let form = new FormData()
    form.append("cname", input.value)
    searchCommunity(form).then(res => {
        communities.value = JSON.parse(res['data']['data'].data)
    }).catch(err => console.log(err))
}
</script>

<style scoped>
.search-container {
    margin-top: 80px;
}
</style>