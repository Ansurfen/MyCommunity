<template>
  <div class="home-nav">
    <el-row :gutter="20">
      <el-col :span="8" class="left-nav">
        <el-row class="sub-left-nav">
          <img src="../../assets/logo.png" width="90" height="90" />
          <el-button class="font-color logo-text" type="text"
            style="margin-left: 20px;padding-top: 45px;font-size: 20px;" @click="home">社团云平台</el-button>
        </el-row>
      </el-col>
      <el-col :span="8"></el-col>
      <el-col :span="8" class="right-nav">
        <el-row class="sub-right-nav" style="margin-bottom: 10px;">
          <search-nav class="search-nav" :msg="msg" />
          <div v-if="userStore.login">
            <el-dropdown>
              <span class="el-dropdown-link">
                <div class="block">
                  <el-avatar :size="45" :src="circleUrl" />
                </div>
              </span>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item @click="edit">个人信息</el-dropdown-item>
                  <el-dropdown-item @click="backed">后台</el-dropdown-item>
                  <el-dropdown-item @click="exit">退出登录</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
          <div v-else>
            <el-button class="font-color" style="margin-top: 6px;" type="text" @click="login">登录</el-button>
            <el-button class="font-color" style="margin-top: 6px;" type="text" @click="register">注册</el-button>
          </div>
          <el-button style="margin-left: 15px;margin-top: 6px;" class="font-color" type="text" @click="about">关于我们
          </el-button>
        </el-row>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts" setup>
import { useUserStore } from "@/stores/user"
import { useRouter } from "vue-router"
import SearchNav from "../search/SearchNav.vue"
import { reactive, toRefs } from 'vue'

const state = reactive({
  circleUrl:
    'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png',
})
const userStore = useUserStore()
if (userStore.info.profile === '1') {
  state.circleUrl = 'http://localhost:9090/images/' + userStore.info.username + '.png'
}
if (window.localStorage) {
  let storage = window.localStorage
  let jwt = storage.getItem('jwt')
  if (typeof (jwt) === 'string' && jwt.length > 0) {
    userStore.login = true
    userStore.jwt = jwt
  }
  if (userStore.login) {
    userStore.syncInfoWithNet()
  }
}
const { circleUrl } = toRefs(state)
let msg = ''
const router = useRouter()
const login = () => {
  router.push({ path: "/user/login" });
}
const register = () => {
  router.push({ path: "/user/register" });
}
const edit = () => {
  router.push({ name: "backed" })
}
const backed = () => {
  router.push({ name: "backed" })
}
const home = () => {
  router.push({ name: "home" })
}
const about = () => {
  router.push({ path: "about" });
}
const exit = () => {
  userStore.login = false
  userStore.jwt = ""
  window.localStorage.removeItem('jwt')
}
</script>

<style scoped>
.sub-left-nav {
  margin-left: 70px;
}

.font-color {
  color: whitesmoke
}

.logo-text {
  padding-top: 25px;
}

.sub-right-nav {
  margin-top: 30px;
}

.search-nav {
  margin-right: 20px;
}

.block:not(:last-child) {
  border-right: 1px solid var(--el-border-color);
}
</style>