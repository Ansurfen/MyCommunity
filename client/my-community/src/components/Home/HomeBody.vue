<template>
  <div class="home-body">
    <el-alert title="欢 迎  使  用 社 团 云 平 台" type="warning" center show-icon />
    <img src="../../assets/home_body.png" />
    <p class="home-body-detail">开放、多元、共享的社团云平台</p>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "HomeBody",
});
</script>

<style scoped>
.home-body-detail {
  color: rgb(118, 220, 240);
  align-content: center;
  font-size: 1.8rem;
  font-weight: 300;
}

.notice {
  background-color: black;
  color: grey;
}
</style>