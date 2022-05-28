<template>
  <div class="navbar">
    <div class="logo">阳光小学线上教学系统</div>
    <div class="user">
      {{ $store.getters?.userInfo?.username }}（{{ $store.getters?.userInfo?.roleName }}）
    </div>
    <div class="mine" @click="toMy"><el-button type="text">我的</el-button></div>
    <div class="logout" @click="logout"><el-button type="text">退出</el-button></div>
    <create-popup
      :show="popupShow"
      :showCancelButton="false"
      :showConfirmButton="false"
      :popup-type="popupType"
      :action="createAction"
      @close="popupShow = false"
    />
  </div>
</template>

<script>
import CreatePopup from "@/components/CreatePopup";
import { logout } from "@/api/login";
export default {
  name: "NavBar",
  components: { CreatePopup },

  data() {
    return {
      popupShow: false,
      popupType: "UserInfo",
      createAction: {
        type: "add",
        id: "",
        data: {}
      }
    };
  },
  methods: {
    // 退出登录
    logout() {
      this.$confirm("您确定退出当前账号的登录状态么？", "退出登录", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          logout().then(res => {
            if (res && res.code === 200) {
              this.$router.push("/login");
              window.localStorage.removeItem("crmPermission");
            }
          });
        })
        .catch(() => {});
    },
    // 我的
    toMy() {
      this.popupShow = true;
    }
  }
};
</script>

<style lang="less" scoped>
.navbar {
  position: relative;
  height: 60px;
  min-height: 60px;
  background-color: #fff;
  padding: 0 30px;
  display: flex;
  align-items: center;

  .logo {
    font-size: 28px;
    font-weight: 500;
  }

  .user,
  .mine,
  .logout {
    position: absolute;
    right: 120px;
  }
  .mine {
    right: 80px;
  }

  .logout {
    cursor: pointer;
    position: absolute;
    right: 30px;
  }
}
</style>
