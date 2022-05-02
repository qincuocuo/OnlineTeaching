<template>
  <transition name="opacity-fade" @after-enter="afterEnter">
    <component v-if="show" ref="createView" :is="popupType" v-bind="$attrs" />
  </transition>
</template>
<script>
import CreateResource from "@/views/crm/cooperation/CreateResource";

export default {
  // createPopup 中嵌套的 createPopup
  name: "CreateNestPopup",
  components: {
    CreateResource
  },
  inheritAttrs: false,
  props: {
    popupType: {
      type: String,
      default: ""
    },
    show: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {};
  },

  methods: {
    afterEnter() {
      // this.$refs.createView.afterEnter();
    }
  }
};
</script>
<style lang="less" scoped>
.opacity-fade-enter-active,
.opacity-fade-leave-active {
  transition: all 0.3s ease;
}
.opacity-fade-enter-from,
.opacity-fade-leave-to {
  opacity: 0;
}
</style>
