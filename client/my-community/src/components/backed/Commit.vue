<template>
    <div>
        <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
            <el-tab-pane label="全部" name="first">
                <p style="text-align: left;">社团列表</p>
                <el-row justify="center" style="background-color:aliceblue;">
                    <el-col :span="4">
                        <div class="grid-content ep-bg-purple" />申请类型
                    </el-col>
                    <el-col :span="6">
                        <div class="grid-content ep-bg-purple" />社团 / ID
                    </el-col>
                    <el-col :span="4">
                        <div class="grid-content ep-bg-purple-light" />申请人
                    </el-col>
                    <el-col :span="4">
                        <div class="grid-content ep-bg-purple" />状态
                    </el-col>
                    <el-col :span="6">
                        <div class="grid-content ep-bg-purple" />操作
                    </el-col>
                </el-row>
                <el-divider />
                <div v-for="(k, v) in ApsMineInAll" :key="v">
                    <el-row justify="center">
                        <el-col :span="4">
                            <div class="grid-content ep-bg-purple" />
                            <div v-if="k.type === 7">新建社团</div>
                            <div v-else-if="k.type === 8">加入社团</div>
                        </el-col>
                        <el-col :span="6">
                            <div class="grid-content ep-bg-purple" />{{ k.first }}
                        </el-col>
                        <el-col :span="4">
                            <div class="grid-content ep-bg-purple-light" />{{ k.second }}
                        </el-col>
                        <el-col :span="4">
                            <div class="grid-content ep-bg-purple" />
                            <div v-if="k.status === 3">申请通过</div>
                            <div v-else-if="k.status === 4">未通过</div>
                            <div v-else-if="k.status === 5">等待审核</div>
                        </el-col>
                        <el-col :span="6">
                            <div class="grid-content ep-bg-purple" />
                            <el-row>
                                <el-button text @click="pass(k.first, k.second, k.type)">通过</el-button>
                                <el-button text @click="reject(k.first, k.second, k.type)"> 拒绝</el-button>
                            </el-row>
                        </el-col>
                    </el-row>
                </div>
            </el-tab-pane>
            <el-tab-pane label="我的" name="second">
                <!-- 加入的社团，我的社团审批 到时候在开新的处理申请吧 -->
                <div v-for="(k, v) in ApsMine" style="margin:10px;" :key="v">
                    <el-row>
                        社团名称: {{ k.first }}
                        <el-col v-if="k.status === Flags.END">失败原因: {{ k.context }}</el-col>
                    </el-row>
                    <div v-if="k.status == 2">
                        <el-steps :active="k.status" finish-status="error" simple style="margin-top: 20px">
                            <el-step title="申请成功" />
                            <el-step title="中止" />
                            <el-step title="通过" />
                        </el-steps>
                    </div>
                    <div v-else>
                        <el-steps :active="k.status" finish-status="success" simple style="margin-top: 20px">
                            <el-step title="申请成功" />
                            <el-step title="中止" />
                            <el-step title="通过" />
                        </el-steps>
                    </div>
                </div>
            </el-tab-pane>
            <el-tab-pane label="已审核" name="third">Role</el-tab-pane>
            <el-tab-pane label="未审核" name="fourth">Task</el-tab-pane>
            <el-tab-pane label="我的申请" name="five">Task</el-tab-pane>
        </el-tabs>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import type { TabsPaneContext } from 'element-plus'
import { useUserStore } from '@/stores/user'
import { communityInfoCommit } from '@/api/commit'
import { Flags } from '@/models/constant'
import { Applications } from '@/models/community'
import { passCommit, rejectCommit } from '@/api/commit'
import { GetStoreWithBoolean, SetStoreWithBoolean } from '@/utils/store'
const activeName = ref('first')
let ApsMine = ref<Applications[]>([])
let ApsMineInAll = ref<Applications[]>([])
const userStore = useUserStore()
const handleClick = (tab: Pick<TabsPaneContext, "index">) => {
    if (userStore.login) {
        switch (tab.index) {
            case '0': communityInfoCommit(JSON.stringify({
                type: Flags.INFO,
                status: Flags.ALL,
            }), userStore.jwt).then(res => {
                ApsMineInAll.value = JSON.parse(res['data']['data'].data)
            }).catch(err => console.log(err)); break;
            case '1': communityInfoCommit(JSON.stringify({
                type: Flags.INFO,
                status: Flags.NEW, // ROOT能请求所有的类型
                first: '', // 
                second: '',
                context: ''
            }), userStore.jwt).then(res => {
                ApsMine.value = JSON.parse(res['data']['data'].data)
                for (let i = 0; i < ApsMine.value.length; i++) {
                    switch (ApsMine.value[i].status) {
                        case 5: ApsMine.value[i].status = 1; break;
                        case 4: ApsMine.value[i].status = 2; break;
                        case 3: ApsMine.value[i].status = 3; break;
                    }
                }
            }).catch(err => { console.log(err) }); break;
            case '2': console.log(); break
        }
    }
}
const pass = (cname: string, username: string, type: number) => {
    passCommit(JSON.stringify({
        type: Flags.PASS,
        status: type,
        first: cname,
        second: username
    }), userStore.jwt).then(res => {
        console.log(res)
        handleClick({ "index": '0' })
    }).catch(err => console.log(err))
}
const reject = (cname: string, username: string, type: number) => {
    rejectCommit(JSON.stringify({
        type: Flags.PASS,
        status: type,
        first: cname,
        second: username
    }), userStore.jwt).then(res => {
        console.log(res)
        handleClick({ "index": '0' })
    }).catch(err => console.log(err))
}
if (GetStoreWithBoolean("loadCommit")) {
    handleClick({ "index": '0' })
    SetStoreWithBoolean("loadCommit", false)
}
</script>

<style scoped>
.demo-tabs>.el-tabs__content {
    padding: 32px;
    color: #6b778c;
    font-size: 32px;
    font-weight: 600;
}
</style>