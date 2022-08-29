<template>
    <div class="main-container">
        <el-container class="elc">
            <el-header class="bc">
                <home-nav style="background-color: #1d3557;" />
            </el-header>
            <el-main class="bc" style="margin:0px;padding: 0px;">
                <main-body :community="community" :status="status" />
            </el-main>
            <el-footer style="padding-top: 20px;background-color: #1d3557;color: whitesmoke;">©Copyright MyCommunity.org
                2022-2023</el-footer>
        </el-container>
    </div>
</template>

<script lang="ts" setup>
import HomeNav from "@/components/home/HomeNav.vue";
import MainBody from "@/components/community/MainBody.vue"
import { useRoute } from "vue-router"
import { Community } from "@/models/community"
import { infoCommunity } from '@/api/community'
import { ref } from "vue"
import { useUserStore } from "@/stores/user"
import { GetStore, SetStoreWithBoolean } from "@/utils/store"

const route = useRoute()
const userStore = useUserStore()
const name = route.params['name'] as string
let community = ref<Community>()
let status = ref(0)

if (name.length > 0) {
    if (window.localStorage) {
        let jwt = GetStore("jwt")
        if (typeof (jwt) === 'string' && jwt.length > 0) {
            userStore.login = true
            userStore.jwt = jwt
        }
        if (userStore.login && GetStore("edit").length > 0) {
            userStore.syncInfoWithNet()
            SetStoreWithBoolean("edit", false)
        }
        if (userStore.login && userStore.info.username === '') {
            userStore.syncInfoWithCache()
        }
    }
    let form = new FormData()
    form.append("cname", name)
    infoCommunity(form, userStore.jwt).then(res => {
        community.value = JSON.parse(res['data']['data'].data)
        if (community.value!.image.length <= 0) {
            community.value!.image = "https://cube.elemecdn.com/e/fd/0fc7d20532fdaf769a25683617711png.png"
        } else {
            community.value!.image = 'http://localhost:9090/' + community.value!.image
        }
        status.value = res['data']['data']['status']
    }).catch(err => console.log(err))
}
</script>

<style scoped>
.el-header {
    --el-header-padding: 0 0px;
    flex: 1;
}
</style>