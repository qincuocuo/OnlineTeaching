<template>
  <div class="navbar">
    <div class="logo">阳光小学线上教学系统</div>
    <div class="user">
      {{ $store.getters?.userInfo?.username }}（{{ $store.getters?.userInfo?.roleName }}）
    </div>
    <div class="logout" @click="logout">退出</div>
  </div>
</template>

<script>
export default {
  name: "NavBar",
  data() {
    return {};
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
          this.$router.push("/login");
          window.localStorage.removeItem("crmPermission");
        })
        .catch(() => {});
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

  .user {
    position: absolute;
    right: 80px;
  }

  .logout {
    // color: #c0c4cc;
    cursor: pointer;
    position: absolute;
    right: 30px;
  }
  .logout:hover {
    color: #409eff;
  }
}
</style>
