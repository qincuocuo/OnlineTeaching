<template>
  <create-popup-view :loading="loading" :title="title" @close="handleClose">
    <create-popup-view-section title="课程信息">签到、讨论、课后题</create-popup-view-section>
  </create-popup-view>
</template>
<script>
import CreatePopupView from "@/components/CreatePopupView";
import CreatePopupViewSection from "@/components/CreatePopupViewSection";
import _ from "lodash";
import { useStore } from "vuex";
import { computed } from "vue";
export default {
  name: "LearnPoup",
  components: {
    CreatePopupView,
    CreatePopupViewSection
  },
  props: {
    action: {
      type: Object,
      default: () => {
        return {
          type: "add",
          id: "",
          data: {}
        };
      }
    }
  },
  setup() {
    const store = useStore();
    return {
      userInfo: computed(() => store.getters.userInfo),
      publicKey: computed(() => store.getters.publicKey)
    };
  },
  data() {
    return {
      loading: false,
      title: "新建课程",
      form: {}
    };
  },
  async mounted() {},
  methods: {
    afterEnter() {},
    handleClose() {}
  }
};
</script>
<style lang="less" scoped></style>
