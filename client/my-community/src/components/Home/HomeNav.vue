<template>
  <div class="home-nav">
    <el-row :gutter="20">
      <el-col :span="8" class="left-nav">
        <el-row class="sub-left-nav">
          <img src="../../assets/logo.png" width="100" height="100" />
          <el-button class="font-color logo-text" type="text" @click="home">社团云平台</el-button>
        </el-row>
      </el-col>
      <el-col :span="8"></el-col>
      <el-col :span="8" class="right-nav">
        <el-row class="sub-right-nav">
          <search-nav class="search-nav" :msg="msg" />
          <div v-if="userStore.login">
            <el-dropdown>
              <span class="el-dropdown-link">
                <div class="block">
                  <el-avatar :size="50" :src="circleUrl" />
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
            <el-button class="font-color" type="text" @click="login">Login</el-button>
            <el-button class="font-color" type="text" @click="register">Register</el-button>
          </div>
          <el-button class="font-color" type="text" @click="about">About</el-button>
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
.home-nav {}

.sub-left-nav {
  margin-left: 70px;
}

.font-color {
  color: rgb(118, 220, 240);
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