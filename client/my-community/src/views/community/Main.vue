<template>
    <div class="main-container">
        <el-container class="elc">
            <el-header class="bc">
                <home-nav />
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
import DetailBody from "@/components/community/DetailBody.vue"
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
/* .bc {
    background-color: rgb(51, 56, 77);
} */
</style>