<template>
  <div class="chat-room">
    <div class="topic-box">
      <div class="topic-label">课堂讨论：</div>
    </div>
    <div class="content-box">
      <div
        v-for="item in list"
        :key="item"
        :class="{ 'comment-item': true, 'my-comment': item.user_id === userInfo.user_id }"
      >
        <div v-if="item.user_id !== userInfo.user_id" class="personal-information">
          <img src="/head_logo.svg" alt="" />
        </div>
        <div class="comment-box">
          <span>{{ item.name }} ({{ getTime(item.send_time) }})</span>
          <div class="comment-content">{{ item.msg }}</div>
        </div>
        <div v-if="item.user_id === userInfo.user_id" class="personal-information">
          <img src="/head_logo.svg" alt="" />
        </div>
      </div>
    </div>
    <div class="operate-box">
      <el-input v-model="comment"></el-input>
      <el-button type="primary" @click="send">发送</el-button>
    </div>
  </div>
</template>

<script>
import { useStore } from "vuex";
import { computed } from "vue";
import { getTime } from "@/utils/utils";
export default {
  name: "chatRoom",
  props: {
    id: {
      type: Number,
      default: 1
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
      comment: "",
      path: `ws://192.168.3.17:5002/api/v1/learning_content/learning/chat?content_id=${this.id}&user_id=${this.userInfo.user_id}`,
      socket: null,
      list: []
    };
  },
  mounted() {
    // 初始化
    this.init();
  },
  methods: {
    init() {
      if (typeof WebSocket === "undefined") {
        alert("您的浏览器不支持socket");
      } else {
        // 实例化socket
        this.socket = new WebSocket(this.path);
        // 监听socket消息
        this.socket.onmessage = this.getMessage;
        // 监听socket连接
        this.socket.onopen = this.open;
        // 监听socket错误信息
        this.socket.onerror = this.error;
        // 关闭监听
        this.socket.onclose = this.close;
      }
    },
    open() {
      console.log("socket连接成功");
    },
    error() {
      console.log("连接错误");
      // this.init();
    },
    getMessage(msg) {
      console.log(msg.data);
      this.list.push(JSON.parse(msg.data));
    },
    send(msg) {
      this.socket.send(
        JSON.stringify({
          "msg": this.comment || msg
        })
      );
      this.comment = "";
    },
    close() {
      console.log("socket已经关闭");
    },
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
    },
    getTime(time) {
      return getTime(time);
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
