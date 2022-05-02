<template>
  <el-card
    v-loading="loading"
    :style="style"
    :body-style="{ padding: 0, height: '100%' }"
    class="slide-detail-view-container"
  >
    <el-button class="close-btn" type="primary" @click="close">
      <el-icon name="icon"><close /></el-icon>
    </el-button>

    <slot />
  </el-card>
</template>
<script type="text/javascript">
import { maxZIndex } from "@/utils/index";

export default {
  name: "SlideDetailView",
  components: {},
  props: {
    loading: {
      type: Boolean,
      default: false
    },
    appendToBody: {
      type: Boolean,
      default: false
    },
    fullScreen: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {};
  },
  computed: {
    style() {
      return {
        "z-index": maxZIndex,
        "top": this.fullScreen ? 0 : "60px"
      };
    }
  },
  watch: {},
  mounted() {
    if (this.appendToBody) {
      document.body.appendChild(this.$el);
    }
  },
  beforeUnmount() {
    if (this.appendToBody && this.$el && this.$el.parentNode) {
      this.$el.parentNode.removeChild(this.$el);
    }
  },
  methods: {
    close() {
      this.$emit("close");
    }
  }
};
</script>
<style lang="less" scoped>
.el-card {
  overflow: visible;
}

.slide-detail-view-container {
  position: fixed;
  min-width: 926px;
  width: 75%;
  bottom: 0px;
  right: 0px;
  background-color: white;
}

.close-btn {
  position: absolute;
  top: 160px;
  left: -40px;
  z-index: 0;
  border-top-right-radius: 0;
  border-bottom-right-radius: 0;
  padding: 6px;
  background-color: #ff6a00;
  border-color: #ff6a00;

  &:hover,
  &:focus {
    background: #fc7d63;
    border-color: #fc7d63;
    color: #ffffff;
  }

  :deep(i) {
    font-size: 26px;
    margin-right: 0;
  }
}
</style>
