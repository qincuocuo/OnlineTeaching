<template>
  <create-container>
    <div v-loading="loading" class="create-popup-view">
      <div class="create-popup-view__header">
        <div v-if="!$slots.header" class="title">{{ title }}</div>
        <slot name="header" />
        <el-icon :size="25" @click="close"><close /></el-icon>
      </div>
      <div class="create-popup-view__body">
        <slot />
      </div>
      <div class="create-popup-view__footer">
        <slot name="footer-left" />
        <el-button v-if="showCancelButton" @click="close">
          {{ cancelButtonText || "取消" }}
        </el-button>
        <slot name="footer-center" />
        <el-button
          v-if="showConfirmButton && !$slots['footer-save']"
          v-clickdebounce="save"
          type="primary"
        >
          {{ confirmButtonText || "保存" }}
        </el-button>
        <slot name="footer-save" />
        <slot name="footer-right" />
      </div>
    </div>
  </create-container>
</template>

<script>
import CreateContainer from "@/components/CreatePopupContainer";

export default {
  name: "CreateView",
  components: {
    CreateContainer
  },
  props: {
    title: {
      type: String,
      default: ""
    },
    loading: {
      type: Boolean,
      default: false
    },
    appendToBody: {
      type: Boolean,
      default: false
    },
    showConfirmButton: {
      type: Boolean,
      default: true
    },
    showCancelButton: {
      type: Boolean,
      default: true
    },
    confirmButtonText: {
      type: String,
      default: "保存"
    },
    cancelButtonText: {
      type: String,
      default: "取消"
    }
  },
  data() {
    return {};
  },
  computed: {},
  watch: {},
  created() {},
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
    save() {
      this.$emit("save");
    },
    close() {
      this.$emit("close");
    }
  }
};
</script>

<style lang="less" scoped>
.create-popup-view {
  position: relative;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: stretch;
  flex-direction: column;

  &__header {
    display: flex;
    text-align: left;
    align-items: center;
    width: 100%;
    height: 40px;
    margin-bottom: 15px;
    padding: 0 10px;
    flex-shrink: 0;
    .title {
      flex: 1;
      font-size: 17px;
      color: #333;
      font-weight: bold;
    }
    .el-icon {
      color: #909399;
      cursor: pointer;
    }
    .el-icon:hover {
      color: @color-primary;
    }
  }

  &__body {
    position: relative;
    overflow-x: hidden;
    overflow-y: auto;
    flex: 1;
  }

  &__footer {
    position: relative;
    text-align: right;
    padding: 60px 20px 0;
    .el-button {
      width: 80px;
    }
  }
}
</style>
