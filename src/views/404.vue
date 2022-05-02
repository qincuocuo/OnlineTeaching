<template>
  <div class="wrapper">
    <div class="http404-container">
      <div class="title">抱歉，您要访问的页面丢失了</div>
      <div class="desc">请检查您输入的网址是否正确，点击以下按钮返回首页</div>
      <div class="btn" @click="handleBack">返回首页</div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Page-404",
  methods: {
    handleBack() {
      this.$router.push("/");
    }
  }
};
</script>

<style lang="less" scoped>
.wrapper {
  position: relative;
  height: 100%;
}

.http404-container {
  position: relative;
  top: 50%;
  transform: translateY(-50%);
  .title {
    width: 100%;
    text-align: center;
    font-size: 24px;
    color: #111;
    line-height: 38px;
  }
  .desc {
    width: 100%;
    font-size: 18px;
    color: #777;
    text-align: center;
    line-height: 38px;
    margin-top: 8px;
  }
  .btn {
    width: 125px;
    height: 44px;
    text-align: center;
    font-size: 14px;
    color: #fefefe;
    line-height: 44px;
    background: #2362fb;
    border-radius: 6px;
    margin: 50px auto 0;
    cursor: pointer;
  }
}
</style>
