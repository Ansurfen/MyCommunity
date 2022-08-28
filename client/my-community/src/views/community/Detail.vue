<template>
    <div class="detail-container">
        <el-container>
            <el-header style="flex:1">
                <home-nav style="background-color: #1d3557;" />
            </el-header>
            <el-main style="margin: 0px;padding: 0px;">
                <el-card class="title-card animate__animated animate__pulse">
                    <el-col>
                        <p style="font-size: 25px;font-weight: 900;text-align: left;">{{ post.title }}</p>
                        <div v-if="post.context.length >= 100 && !showContext">
                            <el-button style="text-align:left" text @click="showContext = true">{{ post.context.slice(0,
                                    100) + ' 显示全部'
                            }}
                            </el-button>
                        </div>
                        <div style="text-align:left;margin-left:70px;" v-if="!showContext" class="context">
                            {{ post.context }}
                        </div>
                        <el-tag v-for="tag in tags" :key="tag.label" :type="tag.type" effect="dark"
                            style="margin: 10px;">
                            {{ tag.label }}
                        </el-tag>
                        <!-- 需要一个帖子所有者的编辑功能 -->
                    </el-col>
                </el-card>
                <div style="margin-top: 50px;margin-left: 200px;" v-for="(v, i) in CurPage()" :key="i">
                    <el-card class="box-card animate__animated animate__backInLeft">
                        <el-row :gutter="20">
                            <el-col :span="4">
                                <el-popover :width="300"
                                    popper-style="box-shadow: rgb(14 18 22 / 35%) 0px 10px 38px -10px, rgb(14 18 22 / 20%) 0px 10px 20px -15px; padding: 20px;">
                                    <template #reference>
                                        <el-avatar :size="100" :src="profiles[i]" />
                                    </template>
                                    <template #default>
                                        <div class="demo-rich-conent"
                                            style="display: flex; gap: 16px; flex-direction: column">
                                            <el-avatar :size="60" :src="profiles[i]" style="margin-bottom: 8px" />
                                            <div>
                                                <p class="demo-rich-content__name" style="margin: 0; font-weight: 500">
                                                    Element Plus
                                                </p>
                                                <p class="demo-rich-content__mention"
                                                    style="margin: 0; font-size: 14px; color: var(--el-color-info)">
                                                    @element-plus
                                                </p>
                                            </div>
                                            <p class="demo-rich-content__desc" style="margin: 0">
                                                Element Plus, a Vue 3 based component library for developers,
                                                designers and product managers
                                            </p>
                                        </div>
                                    </template>
                                </el-popover>
                            </el-col>
                            <el-col :span="16">
                                <el-row>
                                    <div style="text-align: left;">{{ v.host }}</div>
                                    <el-button type="danger" :icon="Delete" style="margin-left:400px;"
                                        @click="del(v.level, v.timestamp)" circle />
                                </el-row>
                                <div style="text-align: left;margin-top: 10px;">{{ v.context }}</div>
                                <div style="margin-top: 30px;">
                                    <el-collapse accordion>
                                        <el-collapse-item>
                                            <template #title>
                                                <el-icon>
                                                    <info-filled />
                                                </el-icon>评论
                                            </template>
                                            <el-timeline>
                                                <el-timeline-item v-for="(_k, _i) in v.sub" :key="_i"
                                                    :tiemstamp="_k.timestamp">
                                                    <div style="text-align: left;">{{ FormatTime(_k.timestamp) }}</div>
                                                    <el-row>
                                                        <div v-if="_k.to === ''">
                                                            {{ _k.from }} : {{ _k.context }}
                                                        </div>
                                                        <div v-else>
                                                            {{ _k.from }} 回复 {{ _k.to }} : {{ _k.context }}
                                                        </div>
                                                        <el-popover placement="bottom" :width="400" trigger="click">
                                                            <template #reference>
                                                                <el-button type="text" style="margin-right: 16px">回复
                                                                </el-button>
                                                            </template>
                                                            <el-row>
                                                                <el-input v-model="reply_msg" />
                                                                <el-button type="primary"
                                                                    @click="reply(v.level, _k.from)">发送
                                                                </el-button>
                                                            </el-row>
                                                        </el-popover>
                                                    </el-row>
                                                </el-timeline-item>
                                            </el-timeline>
                                        </el-collapse-item>
                                    </el-collapse>
                                </div>
                                <el-row>
                                    <div class="time">
                                        发布时间: {{ FormatTime(v.timestamp) }}
                                    </div>
                                    <el-popover placement="bottom" :width="400" trigger="click">
                                        <template #reference>
                                            <el-button type="text" style="margin-right: 16px">回复</el-button>
                                        </template>
                                        <el-row>
                                            <el-input v-model="reply_msg" />
                                            <el-button type="primary" @click="reply(v.level, '')">发送</el-button>
                                        </el-row>
                                    </el-popover>
                                </el-row>
                            </el-col>
                        </el-row>
                    </el-card>
                </div>
                <el-row style="margin-top: 30px;margin-left: 330px;">
                    <el-pagination v-model:currentPage="currentPage" v-model:page-size="pageSize" :small="small"
                        :disabled="disabled" background="true" layout="prev, pager, next, jumper"
                        :total="comments.length" @size-change="handleSizeChange"
                        @current-change="handleCurrentChange" />
                    <el-button @click="dialogFormVisible = true" type="primary" style="margin-left: 30px;">发表评论
                    </el-button>
                </el-row>
                <el-dialog v-model="dialogFormVisible" title="发表评论">
                    <el-form :model="form">
                        <el-form-item label="内容">
                            <el-input type="textarea" :autosize="{ minRows: 4, maxRows: 8 }" v-model="form.context"
                                autocomplete="off" />
                        </el-form-item>
                    </el-form>
                    <template #footer>
                        <span class="dialog-footer">
                            <el-button @click="dialogFormVisible = false" style="margin-right: 30px;">返回
                            </el-button>
                            <el-button type="primary" @click="send">发表</el-button>
                        </span>
                    </template>
                </el-dialog>
            </el-main>
            <el-footer style="padding-top: 20px;background-color: #1d3557;color: whitesmoke;">©Copyright MyCommunity.org
                2022-2023</el-footer>
        </el-container>
    </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from '@vue/reactivity'
import { Posts, IComment, Comment } from '@/models/posts'
import { InfoFilled } from '@element-plus/icons-vue'
import { Tag, Tags, RandEType } from '@/models/common'
import { useRoute } from 'vue-router'
import { infoPosts, addComment, appendComment, delComment } from '@/api/post'
import { useUserStore } from '@/stores/user'
import HomeNav from '@/components/home/HomeNav.vue'
import { Delete } from '@element-plus/icons-vue'
import { ElNotification } from 'element-plus'
import { FormatTime } from '@/utils/time'

const route = useRoute()
const userStore = useUserStore()
const tags = ref<Array<Tag>>([])
const reply_msg = ref('')
const showContext = ref(false)
const currentPage = ref(1)
const pageSize = ref(10)
const small = ref(false)
const disabled = ref(false)
const dialogFormVisible = ref(false)

const form = reactive({
    context: '',
})

let post = ref<Posts>({ id: '', title: '', timestamp: '', author: '', context: '', comments: '', tags: '' })
let comments = ref<Comment[]>([])
let profiles = ref<String[]>([])

const syncComment = () => {
    if (route.params['id'].length > 0) {
        let form = new FormData()
        form.append("id", route.params['id'] as string)
        infoPosts(form).then((res) => {
            post.value = JSON.parse(res['data']['data'].data)
            let t: Tags = JSON.parse(post.value.tags)
            comments.value = []
            profiles.value = []
            tags.value = []
            t.forEach(v => tags.value.push({ type: RandEType(), label: v }))
            let tmpl: IComment[] = JSON.parse(post.value.comments)
            tmpl.forEach((v, i) => {
                let t: Comment = {
                    host: v.host,
                    context: v.context,
                    timestamp: v.timestamp,
                    sub: JSON.parse(v.sub),
                    level: i
                }
                comments.value.push(t)
                profiles.value.push("http://localhost:9090/images/" + t.host + ".png")
            })
        }).catch(err => console.log(err))
    }
}

syncComment()

const handleSizeChange = (val: number) => {
    console.log(`${val} items per page`)
}

const handleCurrentChange = (val: number) => {
    console.log(`current page: ${val}`)
}

const CurPage = (): Comment[] => {
    if (currentPage.value <= comments.value.length) {
        let curPage = currentPage.value - 1
        return comments.value.slice(curPage * 10, curPage * 10 + 10)
    }
    return []
}

const send = () => {
    addComment(JSON.stringify({
        context: form.context,
        host: route.params['id']
    }), userStore.jwt).then(res => {
        ElNotification({
            title: 'Success',
            message: '发帖成功',
            type: 'success'
        })
        syncComment()
    }).catch(err => console.log(err))
    dialogFormVisible.value = false
}

const reply = (level: number, to: string) => {
    appendComment(JSON.stringify({
        from: route.params['id'],
        to: to,
        context: reply_msg.value,
        timestamp: level.toString()
    }), userStore.jwt).then(res => {
        ElNotification({
            title: 'Success',
            message: '评论成功',
            type: 'success'
        })
        syncComment()
    }).catch(err => console.log(err))
    reply_msg.value = ''
}

const del = (level: number, time: string) => {
    delComment(JSON.stringify({
        from: route.params['id'],
        to: time,
        context: '',
        timestamp: level.toString()
    }), userStore.jwt).then(res => {
        ElNotification({
            title: 'Success',
            message: '删除成功',
            type: 'success'
        })
        syncComment()
        console.log(comments.value)
    }).catch(err => console.log(err))
}

// 后端的数据转成 remark的数组 通过curpage 拿到当前map渲染
</script>

<style lang="less" scoped>
.el-header {
    --el-header-padding: 0 0px;
}

.box-card {
    width: 800px;
}

.time {
    font-size: 12px;
    color: #999;
}

.context {
    word-wrap: break-word;
    position: absolute;
    left: 100px;
    width: 800px;
}

.hidden-context {
    word-wrap: break-word;
    position: absolute;
    left: 100px;
    width: 500px;
}

/deep/.el-dialog {
    background-color: rgba(255, 255, 255, 0.9);
}

/deep/.el-card {
    background-color: rgba(255, 255, 255, 0.72);
}

/deep/.el-collapse-item__header {
    background-color: rgba(255, 255, 255, 0.20);
}

/deep/.el-timeline-item__wrapper {
    background-color: rgba(255, 255, 255, 0.20);
}

.el-main {
    background-image: url("../../assets/search-bg.png");
    min-height: 750px;
}

.title-card {
    margin-top: 30px;
    margin-left: 150px;
    margin-right: 150px;
}

.box-card {
    width: 700px;
}
</style>