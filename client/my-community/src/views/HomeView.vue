<template>
  <div class="home-container">
    <el-container class="elc">
      <el-header style="background-color: #1d3557;flex: 1;">
        <home-nav />
      </el-header>
      <el-main class="bc">
        <home-body />
      </el-main>
      <el-footer class="footer">
        <home-bottom />
      </el-footer>
    </el-container>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import HomeNav from "@/components/home/HomeNav.vue"
import HomeBody from "@/components/home/HomeBody.vue"
import HomeBottom from "@/components/home/HomeBottom.vue"
export default defineComponent({
  name: "HomeView",
  components: {
    HomeNav,
    HomeBody,
    HomeBottom
  },
})
</script>

<style scoped>
.bc {
  background-color: rgb(51, 56, 77);
}

.footer {
  background-color: rgb(214, 212, 209);
}
</style>