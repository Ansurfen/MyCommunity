<template>
  <div class="login-container">
    <el-form class="login-form" :rules="rules" ref="ruleFormRef" :model="ruleForm" label-width="80px"
      :hide-required-asterisk="true">
      <h1 class="title">Login</h1>
      <el-form-item label="Username" prop="username">
        <el-input v-model="ruleForm.username" autocomplete="off" />
      </el-form-item>
      <el-form-item label="Password" prop="password">
        <el-input type="password" v-model="ruleForm.password" autocomplete="off" />
      </el-form-item>
      <el-form-item class="submit-btn">
        <el-button type="primary" style="width: 100%" @click="submit(ruleFormRef)">Submit</el-button>
      </el-form-item>
      <div class="bottom">
        <el-button type="text" @click="forget">Forget password?</el-button>
        <el-button type="text" @click="register">Register account</el-button>
      </div>
    </el-form>
  </div>
</template>

<script lang="ts" setup>
import { FormInstance, FormRules } from "element-plus"
import { reactive, ref } from "@vue/reactivity"
import { useRouter } from "vue-router"
import { useUserStore } from "@/stores/user"
import { userLogin } from "@/api/user"
import { User } from "@/models/user"
const router = useRouter()
const userStore = useUserStore()
const ruleFormRef = ref<FormInstance>()
const ruleForm = reactive({
  username: "",
  password: "",
})
const rules = reactive<FormRules>({
  username: [
    { required: true, message: "Please input username", trigger: "blur" },
  ],
  password: [
    { required: true, message: "Please input password", trigger: "blur" },
  ],
})
const submit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate((valid, fields) => {
    if (valid) {
      let params: User = {
        username: ruleForm.username,
        password: ruleForm.password,
      }
      userLogin(JSON.stringify(params))
        .then((res) => {
          userStore.jwt = res.data['data']['token']
          userStore.login = true
          if (userStore.login) {
            userStore.syncInfoWithNet()
          }
          router.replace({ name: "home" })
        })
        .catch((err) => console.log(err.response.data)) //登录成功必须去更新用户信息
    } else {
      console.log("error submit!", fields);
    }
  })
}
const forget = () => {
  router.push({ path: "find" })
}
const register = () => {
  router.push({ path: "register" })
}
</script>

<style scoped>
.login-container {
  position: absolute;
  width: 100%;
  height: 100%;
  overflow: hidden;
  background-image: url("../../assets/login.png");
  background-size: 100% 100%;
  background-repeat: no-repeat;
  background-color: aqua;
  display: flex;
  align-items: center;
  justify-content: center;
}

.login-form {
  background: #fff;
  border-radius: 10px;
  position: absolute;
  width: 350px;
  height: 250px;
  padding: 30px;
  background-color: rgba(90, 187, 211, 0.7);
  margin: 0px auto;
  text-align: left;
}

.title {
  color: #505458;
  text-align: center;
}

.submit-btn {
  width: 280px;
}

.bottom {
  margin-left: 65px;
}
</style>