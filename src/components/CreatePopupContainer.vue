<template>
  <div class="create-popup-container" :style="createViewStyle">
    <el-card v-loading="loading" :body-style="{ height: '100%' }" class="card-container">
      <slot />
    </el-card>
  </div>
</template>
<script type="text/javascript">
import { maxZIndex } from "@/utils/index";

export default {
  name: "CreateContainer",
  components: {},
  props: {
    loading: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {};
  },
  computed: {
    createViewStyle() {
      return {
        zIndex: maxZIndex + 1
      };
    }
  },
  watch: {},
  mounted() {},
  methods: {}
};
</script>
<style lang="less" scoped>
.create-popup-container {
  position: fixed;
  left: 0;
  top: 0;
  bottom: 0;
  right: 0;
  padding: 40px 0px;
  background-color: #f5f6f9;

  .card-container {
    width: 900px;
    height: 100%;
    margin: 0 auto;

    :deep(.el-card__body) {
      padding: 20px 30px;
    }
  }
}
</style>
