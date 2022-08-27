<template>
    <div class="detail-container">
        <el-container>
            <el-header>
                <home-nav style="background-color: #1d3557;"/>
                <h1>{{ post.title }}</h1>
                <div v-if="post.context.length >= 100 && !showContext">
                    <el-button text @click="showContext = true">{{ post.context.slice(0, 100) + ' 显示全部' }}
                    </el-button>
                </div>
                <div v-if="!showContext" class="context">
                    {{ post.context }}
                </div>
                <el-tag v-for="tag in tags" :key="tag.label" :type="tag.type" effect="dark" style="margin: 10px;">
                    {{ tag.label }}
                </el-tag>
                <el-button @click="dialogFormVisible = true">+</el-button>
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
            </el-header>
            <el-main style="margin-top: 200px;margin-left: 200px;">
                <div style="margin-top: 50px;" v-for="(k, i) in CurPage()" :key="i">
                    <el-card class="box-card">
                        <el-row :gutter="20">
                            <el-col :span="4">
                                <el-popover :width="300"
                                    popper-style="box-shadow: rgb(14 18 22 / 35%) 0px 10px 38px -10px, rgb(14 18 22 / 20%) 0px 10px 20px -15px; padding: 20px;">
                                    <template #reference>
                                        <el-avatar :size="100"
                                            src="https://avatars.githubusercontent.com/u/72015883?v=4" />
                                    </template>
                                    <template #default>
                                        <div class="demo-rich-conent"
                                            style="display: flex; gap: 16px; flex-direction: column">
                                            <el-avatar :size="60"
                                                src="https://avatars.githubusercontent.com/u/72015883?v=4"
                                                style="margin-bottom: 8px" />
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
                                    <div style="text-align: left;">{{ k.host }}</div>
                                    <el-button type="danger" :icon="Delete" style="margin-left:400px;"
                                        @click="del(k.level, k.timestamp)" circle />
                                </el-row>
                                <div style="text-align: left;margin-top: 10px;">{{ k.context }}</div>
                                <div style="margin-top: 30px;">
                                    <el-collapse accordion>
                                        <el-collapse-item>
                                            <template #title>
                                                <el-icon>
                                                    <info-filled />
                                                </el-icon>评论
                                            </template>
                                            <el-timeline>
                                                <el-timeline-item v-for="(_k, _i) in k.sub" :key="_i"
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
                                                                    @click="reply(k.level, _k.from)">发送
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
                                        发布时间: {{ FormatTime(k.timestamp) }}
                                    </div>
                                    <el-popover placement="bottom" :width="400" trigger="click">
                                        <template #reference>
                                            <el-button type="text" style="margin-right: 16px">回复</el-button>
                                        </template>
                                        <el-row>
                                            <el-input v-model="reply_msg" />
                                            <el-button type="primary" @click="reply(k.level, '')">发送</el-button>
                                        </el-row>
                                    </el-popover>
                                </el-row>
                            </el-col>
                        </el-row>
                    </el-card>
                </div>
                <el-pagination style="margin-top: 30px;" background v-model:currentPage="currentPage"
                    v-model:page-size="pageSize" :small="small" :disabled="disabled" :background="background"
                    layout="prev, pager, next, jumper" :total="comments.length" @size-change="handleSizeChange"
                    @current-change="handleCurrentChange" />
            </el-main>
            <el-footer>Copyright By Ansurfen</el-footer>
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
const background = ref(false)
const disabled = ref(false)
const dialogFormVisible = ref(false)

const form = reactive({
    context: '',
})

let post = ref<Posts>({ id: '', title: '', timestamp: '', author: '', context: '', comments: '', tags: '' })
let comments = ref<Comment[]>([])

const syncComment = () => {
    if (route.params['id'].length > 0) {
        let form = new FormData()
        form.append("id", route.params['id'] as string)
        infoPosts(form).then((res) => {
            post.value = JSON.parse(res['data']['data'].data)
            let t: Tags = JSON.parse(post.value.tags)
            comments.value = []
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

<style scoped>
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
</style>