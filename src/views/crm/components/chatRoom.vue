<template>
  <div class="chat-room">
    <div class="topic-box">
      <div class="topic-label">话题内容：</div>
      <div>手动改</div>
    </div>
    <div class="content-box">
      <div
        v-for="item in 20"
        :key="item"
        :class="{ 'comment-item': true, 'my-comment': item === 2 }"
      >
        <div v-if="item !== 2" class="personal-information">
          <img src="/head_logo.svg" alt="" />
        </div>
        <div class="comment-box">
          <span>我的</span>
          <div class="comment-content">
            她把请求的参数放在body里了，传了json过去她把请求的参数放在body里了，传了json过去她把请求的参数放在body里了，传了json过去她把请求的参数放在body里了，传了json过去她把请求的参数放在body里了，传了json过去
          </div>
        </div>
        <div v-if="item === 2" class="personal-information">
          <img src="/head_logo.svg" alt="" />
        </div>
      </div>
    </div>
    <div class="operate-box">
      <el-input v-model="comment"></el-input>
      <el-button type="primary" @click="submit">发送</el-button>
    </div>
  </div>
</template>

<script>
import { useStore } from "vuex";
import { computed } from "vue";
export default {
  name: "chatRoom",
  props: {
    id: {
      type: String,
      default: "1"
    }
  },
  setup() {
    const store = useStore();
    return {
      userInfo: computed(() => store.getters.userInfo)
    };
  },
  data() {
    return {
      comment: ""
    };
  },
  mounted() {
    this.scrollToBottom();
  },
  methods: {
    submit() {
      this.scrollToBottom();
    },
    scrollToBottom() {
      const domWrapper = document.querySelector(".content-box");
      (function smoothscroll() {
        const currentScroll = domWrapper.scrollTop;
        const clientHeight = domWrapper.offsetHeight;
        const scrollHeight = domWrapper.scrollHeight;
        if (scrollHeight - 10 > currentScroll + clientHeight) {
          window.requestAnimationFrame(smoothscroll);
          domWrapper.scrollTo(0, currentScroll + (scrollHeight - currentScroll - clientHeight) / 2);
        }
      })();
    }
  }
};
</script>

<style lang="less" scoped>
.chat-room {
  display: flex;
  flex-direction: column;
  height: 100%;
  .topic-box {
    display: flex;
    .topic-label {
      font-weight: bold;
    }
  }
  .content-box {
    flex: 1;
    height: 0;
    margin: 8px 0 20px;
    overflow: auto;
    padding: 10px 0;
    border: 1px solid #eee;
    box-shadow: 10px 0px 10px #eee;
    background-color: #f3f3f3;
    .comment-item {
      display: flex;
      margin: 10px 0;
      .personal-information {
        margin: 0 8px;
        img {
          width: 40px;
          height: 40px;
        }
      }
      .comment-box {
        span {
          color: rgb(109, 109, 109);
        }
        .comment-content {
          max-width: 80%;
          background-color: #fff;
          padding: 8px;
          border-radius: 2px;
          position: relative;
          line-height: 20px;
          &::after {
            content: "";
            position: absolute;
            left: -12px;
            top: 12px;
            border: 6px solid transparent;
            border-right-color: #fff;
          }
        }
      }
    }
    .my-comment {
      .comment-box {
        display: flex;
        flex-direction: column;
        align-items: flex-end;
        text-align: right;
        .comment-content {
          text-align: left;
          background-color: #9feb71;

          &::after {
            left: unset;
            right: -12px;
            border-left-color: #9feb71;
            border-right-color: transparent;
          }
        }
      }
    }
  }
  .operate-box {
    display: flex;
  }
}
</style>
