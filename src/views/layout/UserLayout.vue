<template>
  <el-container>
    <el-header class="nav-container">个人中心头部导航</el-header>
    <el-container>
      <el-main id="user-main-container" style="padding: 15px">
        <app-main />
      </el-main>
    </el-container>
  </el-container>
</template>
<script>
import AppMain from "./components/AppMain";
export default {
  name: "UserLayout",
  components: {
    AppMain
  },
  data() {
    return {};
  }
};
</script>
<style lang="less">
@import url("./styles/common.less");
</style>
