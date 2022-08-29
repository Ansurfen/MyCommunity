import { userInfo } from "@/api/user"
import { UserInfo } from "@/models/user"
import { GetStore, SetStore, SetStoreWithBoolean } from "@/utils/store"
import { defineStore } from "pinia"

export const useUserStore = defineStore("user", {
  state: () => ({
    jwt: "",
    login: false,
    info: {
      "school": '',
      "username": '',
      "alias": '',
      "telephone": '',
      "email": '',
      "id": '',
      "right": '',
      "profile": '0'
    }
  }),
  actions: {
    syncInfoWithNet() {
      userInfo(this.jwt).then((res) => {
        const tres = res['data']['data']
        let right = ''
        switch (tres['right'] as number) {
          case 0: right = "超级管理员"
          case 1: right = "管理员"
          case 2: right = "普通用户"
          default: right = "未知用户组"
        }
        this.info = {
          "username": tres['username'],
          "alias": tres['alias'],
          "telephone": tres['telephone'] === "" ? "未绑定手机号" : tres['telephone'],
          "email": tres['email'] === "" ? "未绑定邮箱" : tres['email'],
          "school": tres['school'] === "" ? "未知" : tres['school'],
          "id": tres['id'].toString(),
          "right": right,
          "profile": tres['profile'].toString()
        }
        SetStore('jwt', this.jwt)
        SetStore('info', JSON.stringify(this.info))
      }).catch((err) => console.log(err))
    },
    syncInfoWithCache() {
      let userInfo: UserInfo = JSON.parse(GetStore("info"))
      let right = ''
      switch (userInfo.right as number) {
        case 0: right = "超级管理员"
        case 1: right = "管理员"
        case 2: right = "普通用户"
        default: right = "未知用户组"
      }
      this.info = {
        "username": userInfo.username,
        "alias": userInfo.alias,
        "telephone": userInfo.telephone === "" ? "未绑定手机号" : userInfo.telephone,
        "email": userInfo.email === "" ? "未绑定邮箱" : userInfo.email,
        "school": userInfo.school === "" ? "未知" : userInfo.school,
        "id": userInfo.id.toString(),
        "right": right,
        "profile": userInfo.profile,
      }
    },
    forceSyncUserInfo() {
      if (this) {
        let jwt = GetStore("jwt")
        if (typeof (jwt) === 'string' && jwt.length > 0) {
            this.login = true
            this.jwt = jwt
        }
        if (this.login && GetStore("edit").length > 0) {
            this.syncInfoWithNet()
            SetStoreWithBoolean("edit", false)
        }
        if (this.login && this.info.username === '') {
            this.syncInfoWithCache()
        }
    }
    }
  },
})