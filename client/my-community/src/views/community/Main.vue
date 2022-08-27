<template>
    <div class="main-container">
        <el-container class="elc">
            <el-header class="bc">
                <home-nav style="background-color: #1d3557;" />
            </el-header>
            <el-main class="bc">
                <detail-body :community="community" />
            </el-main>
            <el-footer class="footer">Copyright By Ansurfen</el-footer>
        </el-container>
    </div>
</template>

<script lang="ts" setup>
import HomeNav from "@/components/home/HomeNav.vue";
import DetailBody from "@/components/community/MainBody.vue"
import { useRoute } from "vue-router"
import { Community } from "@/models/community"
import { infoCommunity } from '@/api/community'
import { ref } from "vue"
const route = useRoute()
const name = route.params['name'] as string
let community = ref<Community>()
if (name.length > 0) {
    let form = new FormData()
    form.append("cname", name)
    infoCommunity(form).then(res => {
        community.value = JSON.parse(res['data']['data'].data)
    }).catch(err => console.log(err))
}
</script>

<style scoped>
.el-header {
    --el-header-padding: 0 0px;
    flex: 1;
}
</style>