<template>
  <el-container>
    <el-header class="nav-container">
      <navbar />
    </el-header>
    <el-container>
      <el-aside>
        <sidebar />
      </el-aside>
      <el-main>
        <app-main />
      </el-main>
    </el-container>
  </el-container>
</template>
<script>
import Navbar from "./components/Navbar";
import AppMain from "./components/AppMain";
import Sidebar from "./components/Sidebar";
export default {
  name: "CrmLayout",
  components: {
    Navbar,
    AppMain,
    Sidebar
  },
  data() {
    return {};
  }
};
</script>
<style lang="less" scoped>
@import url("./styles/common.less");
.el-aside {
  width: 200px;
}
</style>
