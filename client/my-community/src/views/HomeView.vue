<template>
  <div class="home-container">
    <el-container class="elc">
      <el-header class="bc">
        <home-nav />
      </el-header>
      <el-main class="bc">
        <home-body />
      </el-main>
      <el-footer class="footer">Footer</el-footer>
    </el-container>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import HomeNav from "@/components/home/HomeNav.vue";
import HomeBody from "@/components/home/HomeBody.vue";
export default defineComponent({
  name: "HomeView",
  components: {
    HomeNav,
    HomeBody,
  },
});
</script>

<style scoped>
.bc {
  background-color: rgb(51, 56, 77);
}

.footer {
  background-color: rgb(214, 212, 209);
}
</style>