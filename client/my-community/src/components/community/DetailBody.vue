<template>
    <div class="detail-body">
        <el-container>
            <el-header>
                <el-row :gutter="15">
                    <el-col :span="6">
                        <el-avatar :size="60" src="https://empty" @error="errorHandler">
                            <img src="https://cube.elemecdn.com/e/fd/0fc7d20532fdaf769a25683617711png.png" />
                        </el-avatar>
                    </el-col>
                    <el-col :span="4">
                        <div style="text-align: left;">{{ community.name }}</div>
                    </el-col>
                    <el-col :span="8">
                        <!-- 已关注要变黑啥的 -->
                        <el-button type="success" @click="add">申请加入</el-button>
                        <!-- 已发送申请 -->
                        <!-- 已加入 -->
                    </el-col>
                    <el-button @click="dialogFormVisible = true">+</el-button>
                    <el-dialog v-model="dialogFormVisible" title="发布帖子">
                        <el-form :model="form">
                            <el-form-item label="标题">
                                <el-input v-model="form.title" autocomplete="off" />
                            </el-form-item>
                            <el-form-item label="内容">
                                <el-input type="textarea" :autosize="{ minRows: 4, maxRows: 8 }" v-model="form.context"
                                    placeholder="输入不少于20字的简介" autocomplete="off" />
                            </el-form-item>
                        </el-form>
                        <template #footer>
                            <span class="dialog-footer">
                                <div style="float: left;">
                                    <el-tag v-for="tag in dynamicTags" :key="tag" class="mx-1" closable
                                        :disable-transitions="false" @close="handleClose(tag)">
                                        {{ tag }}
                                    </el-tag>
                                    <el-input v-if="inputVisible" ref="InputRef" v-model="inputValue" class="ml-1 w-20"
                                        size="small" @keyup.enter="handleInputConfirm" @blur="handleInputConfirm" />
                                    <el-button v-else class="button-new-tag ml-1" size="small" @click="showInput">
                                        + New Tag
                                    </el-button>
                                </div>
                                <el-button @click="dialogFormVisible = false" style="margin-right: 30px;">返回
                                </el-button>
                                <el-button type="primary" @click="post(community.id)">提交申请</el-button>
                            </span>
                        </template>
                    </el-dialog>
                </el-row>
            </el-header>
            <el-main>
                <!-- details card -->
                <el-row :gutter="20" justify="center">
                    <el-col :span="14" style="padding-left: 300px;">
                        <div v-for="(k, v) in posts" :key="v">
                            <detail-card :title="k.title" :context="k.context" :timestamp="k.timestamp"
                                :author="k.author" :id="k.id" />
                        </div>
                    </el-col>
                    <el-col :span="8" style="padding-right: 300px;">
                        <el-badge :value="200" :max="99">通知</el-badge>
                        <ul v-infinite-scroll="loadNotification" class="infinite-list" style="overflow: auto">
                            <el-card v-for="(k, v) in tips" class="notification" :key="v" shadow="hover">
                                <tip :title="k.title" :context="k.context" />
                            </el-card>
                        </ul>
                    </el-col>
                </el-row>
            </el-main>
        </el-container>
    </div>
</template>

<script lang="ts" setup>
import DetailCard from '@/components/community/DetailCard.vue'
import Tip from '@/components/community/Tip.vue'
import { Community, Post } from '@/models/community'
import { ElInput, ElNotification } from 'element-plus'
import { addPost } from '@/api/post'
import { nextTick, PropType, reactive, ref } from 'vue'
import { useUserStore } from '@/stores/user'
import { useRoute } from 'vue-router'
import { infoPost } from '@/api/post'
import { addCommunity } from '@/api/community'
import { Flags } from '@/models/constant'
const errorHandler = true
const dialogFormVisible = ref(false)
const form = reactive({
    title: '',
    context: '',
})
const userStore = useUserStore()
const post = (id: string) => {
    addPost(JSON.stringify({
        id: id,
        title: form.title,
        context: form.context,
        tags: JSON.stringify(dynamicTags.value)
    }), userStore.jwt).then(res => {
        ElNotification({
            title: 'Success',
            message: '发帖成功',
            type: 'success'
        })
    }).catch(err => {
        ElNotification({
            title: 'Error',
            message: err,
            type: 'error'
        })
    })
    dialogFormVisible.value = false
}
// 得把参数传进来
const loadNotification = () => {

}
defineProps({
    community: {
        type: Object as PropType<Community>,
        default: function () {
            return {}
        }
    }
})
const tips = [{ title: '紧急通知', context: '这是一个紧急通知' }, { title: '紧急通知', context: '这是一个紧急通知' }, { title: '紧急通知', context: '这是一个紧急通知' }, { title: '紧急通知', context: '这是一个紧急通知' }, { title: '紧急通知', context: '这是一个紧急通知' }, { title: '紧急通知', context: '这是一个紧急通知' }]
const inputValue = ref('')
const dynamicTags = ref<string[]>([])
const inputVisible = ref(false)
const InputRef = ref<InstanceType<typeof ElInput>>()
const route = useRoute()
let posts = ref<Post[]>([])
if (route.params['name'].length > 0) {
    let form = new FormData()
    form.append("cname", route.params['name'] as string)
    infoPost(form).then(res => {
        posts.value = JSON.parse(res['data']['data'].data)
    }).catch(err => console.log(err))
}
const handleClose = (tag: string) => {
    dynamicTags.value.splice(dynamicTags.value.indexOf(tag), 1)
}

const showInput = () => {
    inputVisible.value = true
    nextTick(() => {
        InputRef.value!.input!.focus()
    })
}

const handleInputConfirm = () => {
    if (inputValue.value) {
        dynamicTags.value.push(inputValue.value)
    }
    inputVisible.value = false
    inputValue.value = ''
}

const add = () => {
    addCommunity(JSON.stringify({
        first: userStore.info['username'],
        second: route.params['name'],
        type: Flags.ADD
    }), userStore.jwt).then(res => {

    }).catch(err => console.log(err))
}
</script>

<style scoped>
.infinite-list {
    height: 300px;
    padding: 0;
    margin: 0;
    list-style: none;
}

.infinite-list .notification {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 30px;
    background: var(--el-color-primary-light-9);
    margin: 10px;
    color: var(--el-color-primary);
}

.infinite-list .notification+.list-item {
    margin-top: 10px;
}
</style>