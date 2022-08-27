<template>
    <div class="search-container" style="flex: 1;">
        <el-container>
            <el-header style="flex: 1;">
                <home-nav style="background-color: #1d3557;" />
            </el-header>
            <el-main style="min-height: 900px;">
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
                                    <div style="float: left;">
                                        <el-tag v-for="tag in dynamicTags" :key="tag" class="mx-1" closable
                                            :disable-transitions="false" @close="handleClose(tag)">
                                            {{ tag }}
                                        </el-tag>
                                        <el-input v-if="inputVisible" ref="InputRef" v-model="inputValue"
                                            class="ml-1 w-20" size="small" @keyup.enter="handleInputConfirm"
                                            @blur="handleInputConfirm" />
                                        <el-button v-else class="button-new-tag ml-1" size="small" @click="showInput">
                                            + New Tag
                                        </el-button>
                                    </div>
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
                <slice :communities="communities" />
            </el-main>
            <el-footer>Copyright By Ansurfen</el-footer>
        </el-container>
    </div>
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router'
import Slice from '@/components/community/Slice.vue'
import { nextTick, reactive, ref } from 'vue'
import { ElInput, ElNotification } from 'element-plus'
import { newCommunity, searchCommunity } from '@/api/community'
import { useUserStore } from '@/stores/user'
import { Flags } from '@/models/constant'
import HomeNav from '@/components/home/HomeNav.vue'
import { Search } from '@element-plus/icons-vue'
import { Community } from '@/models/community'
import { GetStoreWithBoolean, SetStoreWithBoolean } from '@/utils/store'

const inputValue = ref('')
const dynamicTags = ref<string[]>([])
const inputVisible = ref(false)
const InputRef = ref<InstanceType<typeof ElInput>>()
const router = useRoute()
const userStore = useUserStore()
const dialogFormVisible = ref(false)
const input = ref(router.params['str'].toString())
let communities = ref<Community[]>([])

const handleClose = (tag: string) => {
    dynamicTags.value.splice(dynamicTags.value.indexOf(tag), 1)
}

const handleInputConfirm = () => {
    if (inputValue.value) {
        dynamicTags.value.push(inputValue.value)
    }
    inputVisible.value = false
    inputValue.value = ''
}

const showInput = () => {
    inputVisible.value = true
    nextTick(() => {
        InputRef.value!.input!.focus()
    })
}

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
        third: JSON.stringify(dynamicTags.value),
        context: form.context,
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

if (GetStoreWithBoolean("search")) {
    search()
    SetStoreWithBoolean("search", false)
}
</script>

<style lang="less" scoped>
.el-header {
    --el-header-padding: 0 0px;
}

.el-main {
    background-image: url("../../assets/search-bg.png");
}

// /deep/.el-input {
//     --darkreader-bg--el-input-bg-color: rgba(255, 255, 255, 0.72);
// }

/deep/.el-input-group__prepend {
    background-color: rgba(255, 255, 255, 0.72);
}

/deep/.el-input-group__append {
    background-color: rgba(255, 255, 255, 0.72);
}

/deep/.el-dialog {
    background-color: rgba(255, 255, 255, 0.9);
}
</style>