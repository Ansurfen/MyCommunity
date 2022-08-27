<template>
  <div class="search-container">
    <el-form class="search-form" @submit.native.prevent>
      <el-autocomplete class="search-input" placeholder="Search..." :suffix-icon="Search"
        :fetch-suggestions="querySearchAsync" v-model="msg" @select="handleSelect" @keyup.enter="search(msg)" />
    </el-form>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { Search } from "@element-plus/icons-vue";
import { useRouter } from "vue-router";
import { LinkItem } from '@/models/search'
import { SetStoreWithBoolean } from "@/utils/store";
defineProps({
  msg: {
    type: String,
    default: "",
  },
})
const links = ref<LinkItem[]>([])
const loadAll = () => {
  return [
    { value: 'Community example', link: '' },
    { value: 'Community example2', link: '' },
  ]
}
let timeout: NodeJS.Timeout
const querySearchAsync = (queryString: string, cb: (arg: any) => void) => {
  const results = queryString
    ? links.value.filter(createFilter(queryString))
    : links.value

  clearTimeout(timeout)
  timeout = setTimeout(() => {
    cb(results)
  }, 1000 * Math.random())
}
const createFilter = (queryString: string) => {
  return (restaurant: LinkItem) => {
    return (
      restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0
    )
  }
}

const handleSelect = (item: LinkItem) => {
  console.log()
}

const router = useRouter();
const search = (msg: String) => {
  if (msg.length > 0) {
    console.log(msg)
    router.push({ name: "community/search", params: { str: msg.toString() } })
  }
  SetStoreWithBoolean("search", true)
}

onMounted(() => {
  links.value = loadAll()
})
</script>

<style lang="less" scoped>
/deep/.el-input {
  // 输入框高度
  height: 40px;
  // 边框圆角
  border-radius: 8px;
  --darkreader-bg--el-input-border-color: #1d3557;
  --darkreader-bg--el-input-bg-color: #1d3557;
}

/deep/.el-input__inner {
  height: 40px;
  // 设置字号
  font-size: 14px;
  font-family: PingFangSC-Regular, PingFang SC;
  font-weight: 400;
  // 设置输入字体的颜色
  color: grey;
}

// /deep/.el-autocomplete__popper {
//   background-color: rgba(255, 255, 255, 0.9);
// }
</style>