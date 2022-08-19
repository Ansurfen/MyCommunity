<template>
  <div class="register-container">
    <el-form class="register-form" :rules="rules" ref="ruleFormRef" :model="ruleForm" label-width="80px"
      :hide-required-asterisk="true">
      <h1 class="title">Register</h1>
      <el-form-item label="Username" prop="username">
        <el-input v-model="ruleForm.username" />
      </el-form-item>
      <el-form-item label="Password" prop="password">
        <el-input v-model="ruleForm.password" />
      </el-form-item>
      <el-form-item label="Bind" prop="bind">
        <el-input v-model="ruleForm.bind" placeholder="Please input email or telephone" autocomplete="true" />
      </el-form-item>
      <div class="submit-btn">
        <el-button type="primary" style="width: 100%" @click="register(ruleFormRef)">Submit</el-button>
      </div>
      <div class="bottom"></div>
    </el-form>
  </div>
</template>

<script lang="ts" setup>
import { userRegister } from '@/api/user'
import { reactive, ref } from '@vue/reactivity';
import { FormInstance, FormRules } from 'element-plus'
const ruleFormRef = ref<FormInstance>()
const ruleForm = reactive({
  username: "",
  password: "",
  bind: ""
});
const rules = reactive<FormRules>({
  username: [
    { required: true, message: "Please input username", trigger: "blur" },
  ],
  password: [
    { required: true, message: "Please input password", trigger: "blur" },
  ],
  bind: [
    { required: true, message: "Please input email or telephone", trigger: "blur" }
  ]
})
const register = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate((valid, fields) => {
    if (valid) {
      let params = new FormData()
      params.append('username', ruleForm.username)
      params.append('password', ruleForm.password)
      params.append('key', ruleForm.bind)
      userRegister(params)
        .then((res) => console.log(res.data))
        .catch((err) => console.log(err.response.data))
    } else {
      console.log("error submit!", fields);
    }
  })
}
</script>

<style scoped>
.register-container {
  position: absolute;
  width: 100%;
  height: 100%;
  overflow: hidden;
  background-image: url("../../assets/register.png");
  background-size: 100% 100%;
  background-repeat: no-repeat;
  background-color: aqua;
  display: flex;
  align-items: center;
  justify-content: center;
}

.register-form {
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
</style>