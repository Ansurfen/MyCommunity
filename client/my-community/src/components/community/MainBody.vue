<template>
    <div class="detail-body" style="flex:1">
        <el-container>
            <el-header style="flex:1">
                <el-card class="head-card animate__animated animate__pulse">
                    <el-col>
                        <el-row :gutter="15">
                            <el-col :span="2">
                                <el-avatar :size="60" src="https://empty" @error="errorHandler">
                                    <img :src="community.image" />
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
                                <el-button v-if="status === 4" @click="editStart(community.name)">编辑社团</el-button>
                                <el-button v-if="status >= 3" @click="dialogNoteVisible = true">发布通知</el-button>
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
                            <el-button type="primary" @click="post(community.id)">发布</el-button>
                        </span>
                    </template>
                </el-dialog>
                <!-- edit dialog -->
                <el-dialog v-model="dialogEditVisible" title="编辑社团">
                    <el-form :model="editForm">
                        <el-col>
                            <el-upload ref="uploadRef" action="http://localhost:9090/community/edit/images"
                                :headers="head" accept="image/png,image/jpg,image/jpeg" class="upload-demo"
                                list-type="picture-card" :auto-upload="false" multiple :on-preview="handlePreview"
                                :on-remove="handleRemove" :before-remove="beforeRemove" :limit="1"
                                :on-exceed="handleExceed">
                                <el-button type="primary">上传图片</el-button>
                                <template #tip>
                                    <div class="el-upload__tip">
                                        jpg/png files with a size less than 500KB.
                                    </div>
                                </template>
                            </el-upload>
                            <el-form-item label="内容" style="margin-top: 20px;">
                                <el-input type="textarea" :autosize="{ minRows: 4, maxRows: 8 }"
                                    v-model="editForm.context" placeholder="输入不少于20字的简介" autocomplete="off" />
                            </el-form-item>
                        </el-col>
                    </el-form>
                    <template #footer>
                        <span class="dialog-footer">
                            <div style="float: left;">
                                <el-tag v-for="tag in dynamicEditTags" :key="tag" class="mx-1" closable
                                    :disable-transitions="false" @close="handleEditClose(tag)">
                                    {{ tag }}
                                </el-tag>
                                <el-input v-if="inputEditVisible" ref="InputEditRef" v-model="inputEditValue"
                                    class="ml-1 w-20" size="small" @keyup.enter="handleInputEditConfirm"
                                    @blur="handleInputEditConfirm" />
                                <el-button v-else class="button-new-tag ml-1" size="small" @click="showInputEdit">
                                    + 添加标签
                                </el-button>
                            </div>
                            <el-button type="primary" @click="edit(community.id)">提交更改</el-button>
                        </span>
                    </template>
                </el-dialog>
                <!-- note -->
                <el-dialog v-model="dialogNoteVisible" title="发布通知">
                    <el-form :model="noteForm">
                        <el-form-item label="标题">
                            <el-input v-model="noteForm.title" autocomplete="off" />
                        </el-form-item>
                        <el-form-item label="内容">
                            <el-input type="textarea" :autosize="{ minRows: 4, maxRows: 8 }" v-model="noteForm.context"
                                placeholder="输入不少于20字的简介" autocomplete="off" />
                        </el-form-item>
                    </el-form>
                    <template #footer>
                        <span class="dialog-footer">
                            <el-button type="primary" @click="note(community.id)">发布</el-button>
                        </span>
                    </template>
                </el-dialog>
            </el-header>
            <el-main>
                <el-row :gutter="20" justify="center">
                    <el-col :span="14" style="padding-left: 180px;">
                        <detail-card v-for="(post, i) in posts" :key="i" :title="post.title" :context="post.context"
                            :timestamp="post.timestamp" :author="post.author" :id="post.id" :tags="post.tags" />
                    </el-col>
                    <el-col :span="8" style="padding-right: 300px;">
                        <el-badge :value="200" :max="99">通知</el-badge>
                        <ul v-infinite-scroll="loadNotification(community.name)" class="infinite-list"
                            style="overflow: auto">
                            <el-card v-for="(tip, i) in tips" class="notification" :key="i" shadow="hover">
                                <tip :title="tip.title" :context="tip.context" />
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
import { Community, Note, Post } from '@/models/community'
import { ElInput, ElMessage, ElMessageBox, ElNotification, UploadInstance, UploadProps } from 'element-plus'
import { addPost } from '@/api/post'
import { nextTick, PropType, reactive, ref } from 'vue'
import { useUserStore } from '@/stores/user'
import { useRoute } from 'vue-router'
import { infoPost } from '@/api/post'
import { addCommunity, addNote, editCommunity, infoCommunity, infoNote } from '@/api/community'
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
const dialogNoteVisible = ref(false)
const form = reactive({
    title: '',
    context: '',
})
const editForm = reactive({
    context: ''
})
const noteForm = reactive({
    title: '',
    context: ''
})
const userStore = useUserStore()
const tips = ref<Note[]>([])
const inputValue = ref('')
const inputEditValue = ref('')
const dynamicTags = ref<string[]>([])
const dynamicEditTags = ref<string[]>([])
const inputVisible = ref(false)
const inputEditVisible = ref(false)
const InputRef = ref<InstanceType<typeof ElInput>>()
const InputEditRef = ref<InstanceType<typeof ElInput>>()
const route = useRoute()
const uploadRef = ref<UploadInstance>()
const head = new Headers()

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

let load = false

const loadNotification = (cname: string) => {
    if (tips.value.length <= 0 || !load) {
        let form = new FormData()
        form.append("cname", cname)
        infoNote(form).then(res => {
            let notes: string = res['data']['data'].data
            if (notes.length > 0) {
                let noteSet: Note[] = JSON.parse(notes)
                noteSet.forEach(v => {
                    tips.value.push({ title: v.title, context: v.context, timestamp: v.timestamp })
                })
            }
        }).catch(err => console.log(err))
        load = true
    }
}

updatePosts()

const handleClose = (tag: string) => {
    dynamicTags.value.splice(dynamicTags.value.indexOf(tag), 1)
}

const handleEditClose = (tag: string) => {
    dynamicEditTags.value.splice(dynamicEditTags.value.indexOf(tag), 1)
}

const showInput = () => {
    inputVisible.value = true
    nextTick(() => {
        InputRef.value!.input!.focus()
    })
}

const showInputEdit = () => {
    inputEditVisible.value = true
    nextTick(() => {
        InputEditRef.value!.input!.focus()
    })
}

const handleInputConfirm = () => {
    if (inputValue.value) {
        dynamicTags.value.push(inputValue.value)
    }
    inputVisible.value = false
    inputValue.value = ''
}

const handleInputEditConfirm = () => {
    if (inputEditValue.value) {
        dynamicEditTags.value.push(inputEditValue.value)
    }
    inputEditVisible.value = false
    inputEditValue.value = ''
}

const add = () => {
    addCommunity(JSON.stringify({
        first: userStore.info['username'],
        second: route.params['name'],
        type: Flags.ADD
    }), userStore.jwt).then(res => {

    }).catch(err => console.log(err))
}

const editStart = (cname: string) => {
    dialogEditVisible.value = true
    let form = new FormData()
    form.append("cname", cname)
    infoCommunity(form, userStore.jwt).then(res => {
        const tmpl: Community = JSON.parse(res['data']['data'].data)
        let tags: string[] = JSON.parse(tmpl.tags.toString() as string)
        dynamicEditTags.value = []
        editForm.context = ""
        tags.forEach(v => {
            dynamicEditTags.value.push(v)
        })
        editForm.context = tmpl.context
    }).catch(err => console.log(err))
}

const edit = (id: string) => {
    head.append('Authorization', JSON.stringify({ first: id }) as string + "Bearer " + userStore.jwt)
    uploadRef.value!.submit()
    editCommunity(JSON.stringify({
        id: id,
        context: JSON.stringify({
            context: editForm.context,
            tags: JSON.stringify(dynamicEditTags.value)
        })
    }), userStore.jwt).then(res => { }).catch(err => console.log(err))
    dialogEditVisible.value = false
}

const handleRemove: UploadProps['onRemove'] = (file, uploadFiles) => {
    console.log(file, uploadFiles)
}

const handlePreview: UploadProps['onPreview'] = (uploadFile) => {
    console.log(uploadFile)
}

const handleExceed: UploadProps['onExceed'] = (files, uploadFiles) => {
    ElMessage.warning(
        `The limit is 1, you selected ${files.length} files this time, add up to ${files.length + uploadFiles.length
        } totally`
    )
}

const beforeRemove: UploadProps['beforeRemove'] = (uploadFile, uploadFiles) => {
    return ElMessageBox.confirm(
        `Cancel the transfert of ${uploadFile.name} ?`
    ).then(
        () => true,
        () => false
    )
}

const note = (id: string) => {
    addNote(JSON.stringify({
        id: id,
        title: noteForm.title,
        context: noteForm.context
    }), userStore.jwt).then(res => { }).catch(err => console.log(err))
    dialogNoteVisible.value = false
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
    background-color: rgba(255, 255, 255, 0.0);
}

.infinite-list .notification {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 30px;
    background: var(--el-color-primary-light-9);
    margin: 10px;
    color: var(--el-color-primary);
    background-color: rgba(255, 255, 255, 0.6);
}

.infinite-list .notification+.list-item {
    margin-top: 10px;
}
</style>