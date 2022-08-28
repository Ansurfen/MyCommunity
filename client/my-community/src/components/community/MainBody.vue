<template>
    <div class="detail-body" style="flex:1">
        <el-container>
            <el-header style="flex:1">
                <el-card class="head-card animate__animated animate__pulse">
                    <el-col>
                        <el-row :gutter="15">
                            <el-col :span="2">
                                <el-avatar :size="60" src="https://empty" @error="errorHandler">
                                    <img src="https://cube.elemecdn.com/e/fd/0fc7d20532fdaf769a25683617711png.png" />
                                </el-avatar>
                            </el-col>
                            <el-col :span="4">
                                <p style="text-align: left;color: #1e1e1e;font-size: medium;">{{ community.name }}</p>
                                <p style="text-align: left;color: #1e1e1e;font-size: small;">所有者: {{ community.hostname
                                }}
                                </p>
                            </el-col>
                            <el-col :span="8" style="margin-left: 400px;">
                                <p v-if="status === 0"></p>
                                <el-button v-else-if="status === 1" type="success" @click="add">申请加入</el-button>
                                <el-button v-else-if="status === 2" type="info" disabled>等待审核</el-button>
                                <el-button v-else-if="status >= 3" type="primary" @click="dialogFormVisible = true">
                                    发表帖子</el-button>
                                <el-button v-if="status === 4" @click="dialogEditVisible = true">编辑社团</el-button>
                            </el-col>
                        </el-row>
                        <p style="text-align: left;color: #1e1e1e;font-size: medium;">{{
                                community.context
                        }}</p>
                    </el-col>
                </el-card>
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
                                    + 添加标签
                                </el-button>
                            </div>
                            <el-button type="primary" @click="post(community.id)">提交申请</el-button>
                        </span>
                    </template>
                </el-dialog>
                <!-- 需要编辑描述，标签，图片 -->
                <el-dialog v-model="dialogEditVisible" title="编辑社团">
                    <el-form :model="editForm">
                        <el-form-item label="内容">
                            <el-input type="textarea" :value="community.context" :autosize="{ minRows: 4, maxRows: 8 }" v-model="editForm.context"
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
                                    + 添加标签
                                </el-button>
                            </div>
                            <el-button type="primary" @click="post(community.id)">提交申请</el-button>
                        </span>
                    </template>
                </el-dialog>
            </el-header>
            <el-main>
                <el-row :gutter="20" justify="center">
                    <el-col :span="14" style="padding-left: 180px;">
                        <detail-card v-for="(v, i) in posts" :key="i" :title="v.title" :context="v.context"
                            :timestamp="v.timestamp" :author="v.author" :id="v.id" :tags="v.tags" />
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

defineProps({
    community: {
        type: Object as PropType<Community>,
        default: function () {
            return {}
        }
    },
    status: {
        type: Number,
        default: function () {
            return 0
        }
    }
})

const errorHandler = true
const dialogFormVisible = ref(false)
const dialogEditVisible = ref(false)
const form = reactive({
    title: '',
    context: '',
})
const editForm = reactive({
    context: ''
})
const userStore = useUserStore()
const tips = [{ title: '紧急通知', context: '这是一个紧急通知' }, { title: '紧急通知', context: '这是一个紧急通知' }, { title: '紧急通知', context: '这是一个紧急通知' }, { title: '紧急通知', context: '这是一个紧急通知' }, { title: '紧急通知', context: '这是一个紧急通知' }, { title: '紧急通知', context: '这是一个紧急通知' }]
const inputValue = ref('')
const dynamicTags = ref<string[]>([])
const inputVisible = ref(false)
const InputRef = ref<InstanceType<typeof ElInput>>()
const route = useRoute()

let posts = ref<Post[]>([])

const updatePosts = () => {
    if (route.params['name'].length > 0) {
        let form = new FormData()
        form.append("cname", route.params['name'] as string)
        infoPost(form).then(res => {
            posts.value = JSON.parse(res['data']['data'].data)
            posts.value.forEach((e, i) => {
                posts.value[i].tags = JSON.parse(e.tags.toString())
            })
        }).catch(err => console.log(err))
    }
}

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
        updatePosts()
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

updatePosts()
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

const edit = () => {

}
</script>

<style lang="less" scoped>
.detail-body {
    background-image: url("../../assets/search-bg.png");
    min-height: 650px;
}

.head-card {
    margin-top: 15px;
    margin-left: 190px;
    margin-right: 270px;
    background-color: rgba(255, 255, 255, 0.9);
}

/deep/.el-dialog {
    background-color: rgba(255, 255, 255, 0.9);
}

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