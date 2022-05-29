<template>
  <create-popup-view :loading="loading" :title="title" @close="handleClose">
    <create-popup-view-section title="">
      <el-tabs tab-position="left" style="height: 200px" class="demo-tabs">
        <el-tab-pane label="签到">
          <div class="sign-in-box">
            <div v-if="true" class="sign-in-btn">签到</div>
            <div v-else class="signed-btn">已签到</div>
          </div>
        </el-tab-pane>
        <el-tab-pane label="讨论">
          <chat-room></chat-room>
        </el-tab-pane>
        <el-tab-pane label="课后练习">课后练习</el-tab-pane>
      </el-tabs>
    </create-popup-view-section>
  </create-popup-view>
</template>
<script>
import CreatePopupView from "@/components/CreatePopupView";
import CreatePopupViewSection from "@/components/CreatePopupViewSection";
import chatRoom from "../components/chatRoom.vue";
import _ from "lodash";
import { useStore } from "vuex";
import { computed } from "vue";
export default {
  name: "LearnPoup",
  components: {
    CreatePopupView,
    CreatePopupViewSection,
    chatRoom
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
<style lang="less" scoped>
.create-view-section,
/deep/ .create-view-section__content,
.el-tabs,
/deep/.el-tabs__content,
.el-tab-pane {
  height: 100% !important;

  .sign-in-box {
    height: 100%;
    position: relative;
    .sign-in-btn,
    .signed-btn {
      width: 200px;
      height: 200px;
      border: 1px solid var(--el-color-primary);
      color: var(--el-color-primary);
      border-radius: 50%;
      text-align: center;
      line-height: 200px;
      cursor: pointer;
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      &:hover {
        background-color: rgba(204, 218, 255, 0.4);
      }
    }
    .signed-btn {
      background-color: rgba(204, 218, 255, 0.4);
    }
  }
}
</style>
