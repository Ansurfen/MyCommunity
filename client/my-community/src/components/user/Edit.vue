<template>
    <div class="user-edit">
        <el-row>
            <el-col :span="6">
                <el-avatar style="margin-top: 50px;" shape="square" :size="100" :fit="fit" :src="url" />
            </el-col>
            <el-col :span="6" :offset="6" style="margin-top: 30px;">
                <el-upload ref="uploadRef" action="http://localhost:9090/image/update"
                    accept="image/png,image/jpg,image/jpeg" class="upload-demo" list-type="picture-card"
                    :auto-upload="false" multiple :before-upload="handler" :on-preview="handlePreview"
                    :on-remove="handleRemove" :before-remove="beforeRemove" :limit="1" :on-exceed="handleExceed">
                    <el-icon>
                        <Plus />
                    </el-icon>
                    <template #tip>
                        <div class="el-upload__tip">
                            jpg/png files with a size less than 500KB.
                        </div>
                    </template>
                </el-upload>
            </el-col>
        </el-row>
        <div class="placeholder" />
        <el-row :gutter="20">
            <el-col :span="6">
                <el-row>
                    <el-tag size="large">用户ID</el-tag>
                    <div style="margin: 10px;" />
                    <el-input style="width: 200px;" disabled v-model="id" />
                </el-row>
            </el-col>
            <el-col :span="6" :offset="6">
                <el-row>
                    <el-tag size="large">用户组</el-tag>
                    <div style="margin: 10px;" />
                    <el-input style="width: 200px;" disabled v-model="right" />
                </el-row>
            </el-col>
        </el-row>
        <div class="placeholder" />
        <el-row :gutter="20">
            <el-col :span="6">
                <el-row>
                    <el-tag size="large">用户名</el-tag>
                    <div style="margin: 10px;" />
                    <el-input style="width: 200px;" disabled v-model="username" />
                </el-row>
            </el-col>
            <el-col :span="6" :offset="6">
                <el-row>
                    <el-tag size="large">昵称</el-tag>
                    <div style="margin: 10px;" />
                    <el-input style="width: 200px;" v-model="alias" placeholder="请输入昵称" clearable />
                </el-row>
            </el-col>
        </el-row>
        <div class="placeholder" />
        <el-row :gutter="20">
            <el-col :span="6">
                <el-row>
                    <el-tag size="large">手机号</el-tag>
                    <div style="margin: 10px;" />
                    <el-input style="width: 200px;" v-model="telephone" placeholder="请输入手机号" clearable />
                </el-row>
            </el-col>
            <el-col :span="6" :offset="6">
                <el-row>
                    <el-tag size="large">邮箱</el-tag>
                    <div style="margin: 10px;" />
                    <el-input style="width: 200px;" v-model="email" placeholder="请输入邮箱" clearable />
                </el-row>
            </el-col>
        </el-row>
        <div class="placeholder" />
        <el-row :gutter="20">
            <el-col :span="6">
                <el-row>
                    <el-tag size="large">学校</el-tag>
                    <div style="margin: 10px;" />
                    <el-input style="width: 200px;" v-model="school" placeholder="请输入学校" clearable />
                </el-row>
            </el-col>
        </el-row>
        <div class="placeholder" />
        <div class="placeholder" />
        <div class="placeholder" />
        <el-row :gutter="20">
            <el-col :span="6" />
            <el-col :span="6">
                <el-button type="success" @click="save">
                    保存
                </el-button>
            </el-col>
            <el-col :span="6">
                <el-button type="info" @click="exit">
                    返回
                </el-button>
            </el-col>
            <el-col :span="6" />
        </el-row>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import type { UploadProps, UploadInstance } from 'element-plus'
import { useConfStore } from '@/stores/conf'
import { Plus } from '@element-plus/icons-vue'
import { useUserStore } from '@/stores/user'
import { userEdit } from '@/api/user'
import { SetStore } from '@/utils/store'
defineProps({
    fit: {
        default: function () {
            return 'fill'
        },
        type: String
    },
    url: {
        default: function () {
            return 'https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png'
        },
        type: String
    }
})
const userStore = useUserStore()

const uploadRef = ref<UploadInstance>()
const telephone = ref(userStore.info['telephone'])
const school = ref(userStore.info['school'])
const email = ref(userStore.info['email'])
const username = userStore.info['username']
const id = userStore.info['id']
const alias = ref(userStore.info['alias'])
const right = userStore.info['right']
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
const handler: UploadProps['beforeUpload'] = (file) => {
    // let fd = new FormData()
    // fd.append("file", file)
    // imageUpdate(fd).then((data) => { console.log(data) })
}
const submitUpload = () => {
    uploadRef.value!.submit()
}
const save = () => {
    submitUpload()
    userEdit(JSON.stringify({
        username: username,
        alias: alias.value,
        telephone: telephone.value,
        email: email.value,
        school: school.value
    }), userStore.jwt).then((res) => {
        if (res.status === 200) {
            userStore.info['alias'] = alias.value
            console.log("修改成功")
        }
        SetStore("edit", "1")
    }).catch((err) => console.log(err))
    exit()
}
const exit = () => {
    useConfStore().edit = false
}
</script>

<style scoped>
.placeholder {
    margin: 20px;
}
</style>