<template>
  <div class="about">
      <el-container class="elc">
          <el-header style="background-color: #1d3557;flex: 1;">
              <home-nav />
          </el-header>
          <el-main style="margin: 0px;padding: 0px;">
              <el-card class="null-card animate__animated animate__tada">
                  <el-col>
                      <h1>关于我们</h1>
                      <h2>施工中...</h2>
                  </el-col>
              </el-card>
          </el-main>
          <el-footer class="footer" style="margin: 0px;padding: 0px;background-color: #1d3557;color: whitesmoke;">
              <home-bottom />
          </el-footer>
      </el-container>
  </div>
</template>

<script lang="ts" setup>
import HomeNav from "@/components/home/HomeNav.vue"
import HomeBottom from "@/components/home/HomeBottom.vue"
</script>

<style lang="less" scoped>
.el-main {
  min-height: 640px;
  background-image: url("../assets/search-bg.png");
}

.null-card {
  margin: auto;
  margin-top: 140px;
  height: 220px;
  width: 660px;
  z-index: 1;
  overflow: hidden;
  backdrop-filter: blur(10px);
  background-color: rgba(255, 255, 255, 0.72);
  text-align: center;
}
</style>