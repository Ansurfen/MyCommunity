<template>
    <div>
        <el-container>
            <el-aside width="180px">
                <el-row class="tac">
                    <el-menu default-active="1" style="width: 100%;" class="el-menu-vertical-demo" @select="selectMenu">
                        <el-menu-itme index="0">
                            <img src="../../assets/logo.png" width="90" height="90"
                                style="background-color: #2d3133;" />
                            <span style="background-color: #2d3133;" @click="home">社团云平台</span>
                        </el-menu-itme>
                        <el-menu-item style="background-color:#2d3133" index="1">
                            <el-icon>
                                <icon-menu />
                            </el-icon>
                            <span style="color: aliceblue;background-color:#2d3133;">个人资料</span>
                        </el-menu-item>
                        <el-menu-item index="2">
                            <el-icon>
                                <document />
                            </el-icon>
                            <span style="color: aliceblue;background-color:#2d3133">社团管理</span>
                        </el-menu-item>
                        <el-menu-item index="3">
                            <el-icon>
                                <setting />
                            </el-icon>
                            <span style="color: aliceblue;background-color:#2d3133">用户管理</span>
                        </el-menu-item>
                    </el-menu>
                </el-row>
            </el-aside>
            <el-main>
                <div v-if="ipage == '1'">
                    <edit />
                </div>
                <div v-else-if="ipage == '2'">
                    <commit />
                </div>
                <div v-else-if="ipage == '3'">
                    <user />
                </div>
            </el-main>
        </el-container>
    </div>
</template>

<script lang="ts" setup>
import {
    Document,
    Menu as IconMenu,
    Location,
    Setting,
} from '@element-plus/icons-vue'
import edit from '@/views/user/Edit.vue'
import { useRouter } from 'vue-router'
import User from '@/components/backed/User.vue'
import Commit from '@/components/backed/Commit.vue'
import { ref } from '@vue/reactivity'
// defineProps({
//     vipage: {
//         default: function () {
//             return '1'
//         },
//         type: String
//     }
// })
let ipage = ref('1')
const router = useRouter()
const selectMenu = (index: string) => {
    ipage.value = index
}
const home = () => {
    router.push({ name: "home" })
}
</script>

<style scoped>
.el-aside {
    background-color: #2d3133;
    overflow-y: auto;
    color: #fff;
}

.el-container {
    height: 100vh;
}

.el-menu {
    background-color: #2d3133;
}

.el-menu-item {
    background-color: #2d3133;
}
</style>