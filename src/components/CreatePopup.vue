<template>
  <transition name="opacity-fade" @after-enter="afterEnter">
    <component v-if="show" ref="createView" :is="popupType" v-bind="$attrs" />
  </transition>
</template>
<script>
import CreateCustomer from "@/views/crm/customer/Create";
import CreateUser from "@/views/login/Create.vue";
export default {
  name: "CreatePopup",
  components: {
    CreateCustomer,
    CreateUser
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
