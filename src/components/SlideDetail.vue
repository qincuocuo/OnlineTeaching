<template>
  <transition name="slide-fade" @after-enter="afterEnter">
    <component v-if="show" ref="detailView" :is="slideType" v-bind="$attrs" />
  </transition>
</template>
<script>
import CustomerDetail from "@/views/crm/customer/Detail";
export default {
  name: "SlideDetail",
  components: {
    CustomerDetail
  },
  inheritAttrs: false,
  props: {
    slideType: {
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
      this.$refs.detailView.afterEnter();
    }
  }
};
</script>
<style lang="less" scoped>
.slide-fade-enter-active,
.slide-fade-leave-active {
  will-change: transform;
  transition: all 0.35s ease;
}
.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(100%);
}
</style>
