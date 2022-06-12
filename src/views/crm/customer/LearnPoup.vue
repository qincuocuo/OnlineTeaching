<template>
  <create-popup-view :loading="loading" :title="title" @close="handleClose">
    <create-popup-view-section title="">
      <el-tabs tab-position="left" style="height: 200px" class="demo-tabs">
        <el-tab-pane label="签到">
          <div class="sign-in-box">
            <div v-if="!action.data.registered" class="sign-in-btn" @click="register">签到</div>
            <div v-else class="signed-btn">已签到</div>
          </div>
        </el-tab-pane>
        <el-tab-pane label="讨论">
          <chat-room :id="action.data.content_id"></chat-room>
        </el-tab-pane>
        <el-tab-pane label="课后练习">
          <div class="exercises-content-box">
            <el-form v-if="exercises.length" ref="homeworkRef" label-width="100px">
              <div v-for="(item, index) in exercises" :key="index" style="padding-bottom: 20px">
                <el-form-item prop="index" label-width="40px">
                  <h4>第{{ index + 1 }}题</h4>
                </el-form-item>
                <el-form-item label="题目：" prop="question">
                  <span>{{ exercises[index].question }}</span>
                </el-form-item>
                <el-form-item label="题目类型：" prop="type">
                  <span>{{ item.type === 1 ? "判断题" : "选择题" }}</span>
                </el-form-item>
                <el-form-item label="选项：" prop="option">
                  <el-radio-group v-if="item.type === 2" v-model="answers[index]" class="ml-4">
                    <el-radio
                      v-for="(optionItem, index) in item.options"
                      :key="optionItem"
                      :label="selectOption[index].value"
                    >
                      {{ selectOption[index].label }} : {{ optionItem }}
                    </el-radio>
                  </el-radio-group>
                  <el-radio-group v-else v-model="answers[index]" class="ml-4">
                    <el-radio
                      v-for="answerItem in judgeOption"
                      :key="answerItem.value"
                      :label="answerItem.value"
                    >
                      {{ answerItem.label }}
                    </el-radio>
                  </el-radio-group>
                </el-form-item>
                <el-form-item label="结果：" v-if="exercisesResult.length">
                  <div v-if="exercisesResult[index].correct" class="results-right-text">正确</div>
                  <div v-else class="results-error-text">
                    错误。正确答案：{{ getJudgeOption(exercisesResult[index].answer, item.type) }}
                  </div>
                </el-form-item>
              </div>
            </el-form>
            <div v-else class="exercises-box">暂无练习</div>
            <el-button type="primary" @click="postExercises">提交</el-button>
          </div>
        </el-tab-pane>
      </el-tabs>
    </create-popup-view-section>
  </create-popup-view>
</template>
<script>
import CreatePopupView from "@/components/CreatePopupView";
import CreatePopupViewSection from "@/components/CreatePopupViewSection";
import chatRoom from "../components/chatRoom.vue";
import { register, exercises, postExercises } from "@/api/crm/customer";
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
      form: {},
      exercises: [
        {
          question: "",
          type: 2,
          options: ["", ""],
          answer: "1"
        }
      ],
      selectOption: [
        {
          value: "1",
          label: "A"
        },
        {
          value: "2",
          label: "B"
        },
        {
          value: "3",
          label: "C"
        },
        {
          value: "4",
          label: "D"
        },
        {
          value: "5",
          label: "E"
        }
      ],
      judgeOption: [
        {
          value: "1",
          label: "对"
        },
        {
          value: "2",
          label: "错"
        }
      ],
      answers: [],
      exercisesResult: []
    };
  },
  mounted() {
    this.postRlease();
  },
  methods: {
    afterEnter() {},
    handleClose() {},
    register() {
      register({
        content_id: this.action.data.content_id
      }).then(res => {
        if (res.code === 200) {
          this.$message.success("成功");
          // eslint-disable-next-line vue/no-mutating-props
          this.action.data.registered = true;
        } else {
          this.$message.warning(res.message);
        }
      });
    },
    // 获取题目

    postRlease() {
      exercises({
        content_id: this.action.data.content_id
      }).then(res => {
        if (res.code === 200 && res.data) {
          this.exercises = res.data.results;
        } else {
          this.exercises = [];
        }
      });
    },
    // 提交答题
    postExercises() {
      const answer = this.answers.map((item, index) => {
        return {
          id: this.exercises[index].id,
          answer: item
        };
      });
      postExercises({
        content_id: this.action.data.content_id,
        answers: answer
      }).then(res => {
        if (res.code === 200 && res.data) {
          //
          this.exercisesResult = res.data.results;
        } else {
          //
        }
      });
    },
    // 提交答题
    getJudgeOption(answer, type) {
      let res;
      const options = type === 1 ? this.judgeOption : this.selectOption;
      options.forEach(item => {
        if (item.value === answer) res = item.label;
      });
      return res;
    }
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
  .exercises-content-box {
    height: 100%;
    overflow: auto;
    .results-right-text {
      color: green;
    }
    .results-error-text {
      color: red;
    }
  }
  .exercises-box {
    height: 100%;
    margin-top: 160px;
    font-size: 20px;
    text-align: center;
  }
}
</style>
