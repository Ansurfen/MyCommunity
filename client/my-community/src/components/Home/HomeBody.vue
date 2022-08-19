<template>
  <div class="home-body">
    <el-alert title="Welcome to use Could Platform for communities" type="warning" center show-icon />
    <img src="../../assets/home_body.png" />
    <p class="home-body-detail">Open, diverse and eclectic Community Platform</p>
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