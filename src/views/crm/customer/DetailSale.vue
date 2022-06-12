<template>
  <div class="detail-visit">
    <div class="table-head-container">
      <div class="query-add-btns-container">
        <div class="add-btns">
          <el-input
            v-model="queryParam.search"
            class="content-search"
            placeholder="请输入学习内容"
            prefix-icon="Search"
          />
          <el-button class="content-add" @click="searchQuery" type="primary">查询</el-button>
          <el-button
            v-has="'teach'"
            class="content-add"
            @click="addContentVisible = true"
            type="primary"
          >
            新增学习内容
          </el-button>
        </div>
      </div>
    </div>
    <div class="table-view-container" v-loading="loading">
      <table-view
        ref="refTableView"
        :columns="columns"
        :dataSource="dataSource"
        :refDataTable="refDataTable"
        v-model:ipagination="ipagination"
        @load="loadData"
      >
        <template v-slot:salesLeadSourceId="scope">
          {{
            gainAppoint(sourceOption, scope.row.salesLeadSourceId, "customerSourceId")
              .customerSource || "--"
          }}
        </template>
        <template v-slot:register="scope">
          <span v-if="scope.row.register">是</span>
          <span v-else>否</span>
        </template>
        <template v-slot:operate="scope">
          <div class="table-btn-box">
            <el-button type="text" @click="viewLearnDetail(scope.row)">查看学习内容</el-button>
          </div>
          <div v-has="'teach'" class="table-btn-box">
            <el-button type="text" @click="getLearningDetail(scope.row)">查看学习情况</el-button>
          </div>
          <div v-has="'teach'" class="table-btn-box">
            <el-button
              type="text"
              @click="
                () => {
                  this.qiandaoVisible = true;
                  this.chooseRow = scope.row;
                }
              "
            >
              发起签到
            </el-button>
          </div>
          <div v-has="'teach'" class="table-btn-box">
            <el-button type="text" @click="getQianDaoDetail(scope.row)">查看签到结果</el-button>
          </div>
          <div v-has="'teach'" class="table-btn-box">
            <el-button type="text" @click="viewDiscussion(scope.row)">查看讨论情况</el-button>
          </div>
          <div v-has="'teach'" class="table-btn-box">
            <el-button type="text" @click="postRlease(scope.row)">发布课后练习</el-button>
          </div>
          <div v-has="'student'" class="table-btn-box">
            <el-button type="text" @click="enterLearning(scope.row)">进入学习</el-button>
          </div>
        </template>
      </table-view>
    </div>
    <!-- 查看学习内容 -->
    <el-dialog
      destroy-on-close
      v-model="viewLearnDetailVisable"
      custom-class="learn-video-box"
      ref="video"
    >
      <video class="learn-video" controls autoplay="autoplay" loop="loop" mute="muted">
        <source :src="learnContent" />
      </video>
    </el-dialog>
    <!-- 新增课程内容 -->
    <el-dialog v-model="addContentVisible" title="新增课程内容" width="40%">
      <el-form
        ref="contentFormRef"
        :model="content_form"
        :rules="contentfFormRules"
        v-loading="addContentLoading"
      >
        <el-form-item prop="title" class="login-email" label="名称">
          <el-input placeholder="" v-model.trim="content_form.title" class="email-input"></el-input>
        </el-form-item>
        <el-upload
          class="upload-demo"
          drag
          action="https://jsonplaceholder.typicode.com/posts/"
          :on-change="beforeUpload"
          :auto-upload="false"
        >
          <el-icon class="el-icon--upload">
            <upload-filled />
          </el-icon>
          <div class="el-upload__text">
            拖拽文件或
            <em>点击图标上传文件</em>
          </div>
        </el-upload>
        <el-form-item class="login-button">
          <el-button type="primary" @click="addContent">提交</el-button>
          <el-button @click="addContentVisible = false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
    <!-- 发起签到 -->
    <el-dialog v-model="qiandaoVisible" title="发起签到" width="40%">
      <div style="text-align: center">
        签到有效时间为
        <el-input-number v-model="register_tm" :min="1" controls-position="right" />
        分钟
      </div>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="qiandaoVisible = false">取消</el-button>
          <el-button type="primary" @click="addQianDaoTask">发起签到</el-button>
        </span>
      </template>
    </el-dialog>
    <!-- 发起讨论 -->
    <el-dialog v-model="talkVisible" title="讨论话题" width="40%">
      <el-input type="textarea" v-model="talk"></el-input>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="talkVisible = false">取消</el-button>
          <el-button type="primary" @click="addtalk">发起讨论</el-button>
        </span>
      </template>
    </el-dialog>
    <!-- 查看讨论情况 -->
    <el-drawer v-model="chatRoomVisable">
      <chat-room :id="chooseRow.content_id"></chat-room>
    </el-drawer>
    <!-- 查看签到情况 -->
    <el-drawer v-model="qiandaoDetailVisable" title="签到结果" direction="rtl" size="20%">
      <el-tabs
        v-model="qiandaoDetailActiveName"
        class="demo-tabs"
        @tab-click="qiandaoDetailTabChange"
      >
        <el-tab-pane :label="`已签到`" name="finished">
          <div v-for="item in resultList" :key="item">{{ item.name }}</div>
        </el-tab-pane>
        <el-tab-pane :label="`未签到`" name="unfinished">
          <div v-for="item in resultList" :key="item">{{ item.name }}</div>
        </el-tab-pane>
      </el-tabs>
    </el-drawer>
    <!-- 查看学习情况 -->
    <el-drawer
      v-model="learningDetailVisable"
      :title="`${chooseRow?.content}学习结果`"
      direction="rtl"
      size="20%"
    >
      <el-tabs
        v-model="learningDetailActiveName"
        class="demo-tabs"
        @tab-click="learningDetailTabChange"
      >
        <el-tab-pane :label="`已学习`" name="learned">
          <div v-for="item in resultList" :key="item">{{ item.name }}</div>
        </el-tab-pane>
        <el-tab-pane :label="`未学习`" name="unlearned">
          <div v-for="item in resultList" :key="item">{{ item.name }}</div>
        </el-tab-pane>
      </el-tabs>
    </el-drawer>
    <!-- 课后练习 -->
    <div class="homework">
      <el-dialog
        v-model="homeworkVisible"
        custom-class="homework-box"
        title="课后练习"
        width="50%"
        destroy-on-close
      >
        <el-form ref="homeworkRef">
          <div v-for="(item, index) in exercises" :key="index" style="padding-bottom: 20px">
            <el-form-item prop="index">
              <h4>第{{ index + 1 }}题</h4>
            </el-form-item>
            <el-form-item label="题目" prop="question">
              <el-input type="textarea" v-model="exercises[index].question"></el-input>
            </el-form-item>
            <el-form-item label="题目类型" prop="type">
              <el-radio-group v-model="item.type" class="ml-4">
                <el-radio :label="1" size="large">判断题</el-radio>
                <el-radio :label="2" size="large">选择题</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item v-if="item.type === 2" label="选项" prop="option">
              <div v-for="(optionItem, optionIndex) in item?.options" :key="optionIndex">
                <span>{{ selectOption[optionIndex]?.label }}:</span>
                <el-input
                  style="padding-right: 10px"
                  v-model="exercises[index].options[optionIndex]"
                  placeholder="请输入选项内容"
                  class="input-with-select"
                >
                  <template #prepend>
                    <el-button icon="Plus" @click="addOption(index)" />
                  </template>
                  <template #append>
                    <el-button icon="Minus" @click="deleteOption(index, optionIndex)" />
                  </template>
                </el-input>
              </div>
            </el-form-item>
            <el-form-item label="答案" prop="answer">
              <el-select v-if="item.type === 2" v-model="item.answer">
                <el-option
                  v-for="item in selectOption.slice(0, item.options.length)"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
              <el-select v-else v-model="item.answer">
                <el-option label="对" value="1" />
                <el-option label="错" value="2" />
              </el-select>
            </el-form-item>

            <el-button type="primary" @click="deleteQuestion(index)">删除题目</el-button>
          </div>
          <el-button type="primary" @click="addQuestion">添加题目</el-button>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="homeworkVisible = false">取消</el-button>
            <el-button type="primary" @click="addExercises">提交</el-button>
          </span>
        </template>
      </el-dialog>
    </div>
    <create-popup
      ref="refCreatePoup"
      :show="popupShow"
      :popup-type="popupType"
      :action="createAction"
      :sourceOption="sourceOption"
      @update="loadData"
      @close="popupShow = false"
    />
  </div>
</template>

<script>
import TableView from "@/views/crm/components/TableView";
import TableMixin from "@/views/crm/mixins/Table";
import chatRoom from "../components/chatRoom.vue";
import CreatePopup from "@/components/CreatePopup";
import { gainAppoint } from "@/utils/utils";
import {
  addContent,
  getLearningDetail,
  getQianDaoDetail,
  addQianDao,
  addtalk,
  addExercises,
  getlearnContentDetail,
  exercises
} from "@/api/crm/customer";
import { useStore } from "vuex";
import { computed } from "vue";
// import VueCoreVideoPlayer from 'vue-core-video-player'
export default {
  name: "DetailSale",
  mixins: [TableMixin],
  components: { TableView, CreatePopup, chatRoom },
  props: {
    customer: {
      type: Object,
      default: () => {
        return {};
      }
    }
  },
  setup() {
    const store = useStore();
    return {
      visitMode: computed(() => store.getters.visitMode)
    };
  },
  created() {
    this.queryParam.course_id = this.customer.course_id;
  },
  data() {
    return {
      columns: [
        {
          label: "学习内容",
          prop: "content",
          width: 140
        },
        {
          label: "需要签到",
          props: "register",
          slot: "register",
          width: 140
        },
        {
          label: "已学人数",
          prop: "learned",
          width: 50
        },
        {
          label: "未学人数",
          prop: "unlearned",
          width: 50
        },
        {
          label: "操作",
          slot: "operate",
          width: 250
        }
      ],
      url: {
        list: "/api/v1/learning_content",
        type: "get"
      },
      popupShow: false,
      addContentVisible: false,
      addContentLoading: false,
      qiandaoVisible: false,
      talkVisible: false,
      qiandaoDetailVisable: false,
      learningDetailVisable: false,
      chooseRow: null,
      homeworkVisible: false,
      viewLearnDetailVisable: false,
      chatRoomVisable: false,
      viewLearnDetailFileUrL: "",
      popupType: "CreateOpportunity",
      createAction: {
        type: "add",
        id: "",
        data: {}
      },
      sourceOption: [],
      search: "",
      content_form: {
        title: "",
        file: []
      },
      register_tm: 2, //签到时间限制
      talk: "", // 讨论话题
      //表单验证规则
      contentfFormRules: {
        title: [{ required: true, message: "请输入名称", trigger: "change" }],
        content: [{ required: true, message: "请输入内容", trigger: "change" }]
      },
      qiandaoDetailActiveName: "finished",
      resultList: ["zhangsan", "lisi"],
      learningDetailActiveName: "learned",
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
      ]
    };
  },
  methods: {
    // 课后作业（先查询再添加）
    postRlease(row) {
      this.homeworkVisible = true;
      this.chooseRow = row;
      exercises({
        content_id: row.content_id
      }).then(res => {
        if (res.code === 200 && res.data) {
          this.exercises = res.data.results;
        } else {
          this.exercises = this.$options.data().exercises;
        }
      });
    },
    // 添加题目
    addQuestion() {
      this.exercises.push({
        question: "",
        type: 2,
        options: ["", ""],
        answer: "1"
      });
    },
    // 删除题目
    deleteQuestion(index) {
      this.exercises.splice(index, 1);
    },
    // 添加选项
    addOption(index) {
      if (this.exercises[index].options.length === 5) {
        this.$message.warning("选项最多为5个");
        return;
      }
      this.exercises[index].options.push("");
    },
    deleteOption(index, optionIndex) {
      if (this.exercises[index].options.length === 2) {
        this.$message.warning("选项至少2个");
        return;
      }
      this.exercises[index].options.splice(optionIndex, 1);
    },
    handleClose() {},
    handleClick() {},
    // 查看学习内容（确认学习）
    viewLearnDetail(row) {
      this.viewLearnDetailVisable = true;
      getlearnContentDetail({ content_id: row.content_id });
      this.learnContent =
        "http://121.199.167.227:5002/api/v1/learning_content/learning_content?content_id=" +
        row.content_id;
    },
    // 文件上传
    beforeUpload(file) {
      this.content_form.file = file.raw;
    },
    // 添加学习内容
    addContent() {
      this.$refs.contentFormRef.validate(async valid => {
        if (!valid) return;
        this.addContentLoading = true;
        const formData = new FormData();
        formData.append("file", this.content_form.file);
        formData.append("title", this.content_form.title);
        formData.append("course_id", this.customer.course_id);
        addContent(formData)
          .then(res => {
            if (res && res.code === 200) {
              this.$message.success(res.message);
              this.addContentVisible = false;
              this.loadData();
            } else {
              this.$message.warning(res.message);
            }
          })
          .finally(() => {
            this.addContentLoading = false;
          });
      });
    },
    // 查看学习情况
    getLearningDetail(row) {
      this.learningDetailVisable = true;
      this.resultList = [];
      this.chooseRow = row;
      this.learningDetailTabChange();
    },
    // 学习情况 切换
    learningDetailTabChange() {
      getLearningDetail({
        course_id: this.customer.course_id,
        content_id: this.chooseRow.content_id,
        status: this.learningDetailActiveName
      }).then(res => {
        if (res && res.code === 200) {
          this.resultList = res.data.student_info || [];
        } else {
          this.$message.warning(res.message);
        }
      });
    },
    // 查看签到结果
    getQianDaoDetail(row) {
      this.qiandaoDetailVisable = true;
      this.resultList = [];
      this.chooseRow = row;
      this.qiandaoDetailTabChange();
    },
    // 查看签到结果 切换
    qiandaoDetailTabChange() {
      this.resultList = [];
      getQianDaoDetail({
        course_id: this.customer.course_id,
        content_id: this.chooseRow.content_id,
        register_result: this.qiandaoDetailActiveName
      }).then(res => {
        if (res && res.code === 200) {
          this.resultList = res.data.student_info || [];
        } else {
          this.$message.warning(res.message);
        }
      });
    },
    // 创建签到任务
    addQianDaoTask() {
      addQianDao({
        content_id: this.chooseRow.content_id,
        register_tm: this.register_tm
      }).then(res => {
        if (res && res.code === 200) {
          this.$message.success(res.message);
          this.qiandaoVisible = false;
        } else {
          this.$message.warning(res.message);
        }
      });
    },
    // 创建讨论话题
    addtalk() {
      addtalk({
        course_id: this.customer.course_id,
        content_id: this.chooseRow.content_id,
        talk: this.talk
      }).then(res => {
        if (res && res.code === 200) {
          this.$message.success(res.message);
          this.talkVisible = false;
        } else {
          this.$message.warning(res.message);
        }
      });
    },
    // 创建课后练习
    addExercises() {
      addExercises({
        content_id: this.chooseRow.content_id,
        learning_content_id: this.chooseRow.content_id,
        exercises: this.exercises
      }).then(res => {
        if (res && res.code === 200) {
          this.$message.success(res.message);
          this.homeworkVisible = false;
        } else {
          this.$message.warning(res.message);
        }
      });
    },
    // 查看讨论情况
    viewDiscussion(row) {
      this.chatRoomVisable = true;
      this.chooseRow = row;
    },
    edit(item) {
      this.createAction = {
        type: "edit",
        id: item.salesLeadId,
        data: item
      };
      this.popupShow = true;
    },
    gainAppoint() {
      return gainAppoint(...arguments);
    },
    enterLearning(item) {
      this.$emit("enterLearning", item);
    }
  }
};
</script>

<style lang="less" scoped>
.detail-visit {
  display: flex;
  flex-direction: column;
  height: 100%;

  .add-btns {
    display: inline-flex;

    .content-search {
      padding-right: 20px;
    }

    .content-add {
      width: 150px;
    }
  }

  .table-view-container {
    flex: 1;
    width: 100%;

    /deep/ .el-table .cell {
      white-space: pre-line;
    }
  }
}
/deep/.learn-video-box {
  .el-dialog__body {
    display: flex;
    justify-content: center;
  }
  .learn-video {
    margin: 0 auto;
    max-height: 60vh;
    max-width: 50vw;
  }
}

.homework {
  /deep/ .el-dialog {
    .el-dialog__body {
      overflow: auto;
      max-height: 60vh;
      overflow: auto;
    }
  }
}
</style>
