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
          <el-button v-has="'teach'" class="content-add" @click="searchQuery" type="primary">
            查询
          </el-button>
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
        <template v-slot:operate="scope">
          <div v-has="'teach'" class="table-btn-box">
            <el-button type="text" @click="edit(scope.row)">查看学习内容</el-button>
          </div>
          <div v-has="'teach'" class="table-btn-box">
            <el-button type="text" @click="learningDetail = true">查看学习情况</el-button>
          </div>
          <div v-has="'teach'" class="table-btn-box">
            <el-button type="text" @click="qiandaoVisible = true">发起签到</el-button>
          </div>
          <div v-has="'teach'" class="table-btn-box">
            <el-button type="text" @click="talkVisible = true">发起讨论</el-button>
          </div>
          <div v-has="'teach'" class="table-btn-box">
            <el-button type="text" @click="qiandaoDetail = true">查看签到结果</el-button>
          </div>
          <div v-has="'teach'" class="table-btn-box">
            <el-button type="text">查看讨论情况</el-button>
          </div>
          <div v-has="'teach'" class="table-btn-box">
            <el-button type="text" @click="homeworkVisible = true">发布课后练习</el-button>
          </div>
          <div v-has="'student'" class="table-btn-box">
            <el-button type="text" @click="enterLearning(scope.row)">进入学习</el-button>
          </div>
          <div v-has="'student'" class="table-btn-box">
            <el-button type="text">查看通知</el-button>
          </div>
        </template>
      </table-view>
    </div>
    <!-- 新增课程内容 -->
    <el-dialog v-model="addContentVisible" title="新增课程内容" width="40%">
      <el-form ref="contentFormRef" :model="content_form" :rules="contentfFormRules">
        <el-form-item prop="title" class="login-email" label="名称">
          <el-input placeholder="" v-model.trim="content_form.title" class="email-input"></el-input>
        </el-form-item>
        <el-upload
          class="upload-demo"
          drag
          action="https://jsonplaceholder.typicode.com/posts/"
          :before-upload="beforeUpload"
          multiple
        >
          <el-icon class="el-icon--upload">
            <upload-filled />
          </el-icon>
          <div class="el-upload__text">
            拖拽文件或
            <em>点击图标上传文件</em>
          </div>
          <template #tip>
            <div class="el-upload__tip">上传文件限制</div>
          </template>
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
          <el-button type="primary" @click="qiandaoVisible = false">发起签到</el-button>
        </span>
      </template>
    </el-dialog>
    <!-- 发起讨论 -->
    <el-dialog v-model="talkVisible" title="讨论话题" width="40%">
      <el-input type="textarea" v-model="talk"></el-input>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="talkVisible = false">取消</el-button>
          <el-button type="primary" @click="talkVisible = false">发起讨论</el-button>
        </span>
      </template>
    </el-dialog>
    <!-- 查看签到情况 -->
    <el-drawer v-model="qiandaoDetail" title="签到结果" direction="rtl" size="20%">
      <el-tabs v-model="qiandaoDetailActiveName" class="demo-tabs" @tab-click="handleClick">
        <el-tab-pane :label="`已签到(${12})`" name="qiandao">
          <div v-for="item in finalishList" :key="item">{{ item }}</div>
        </el-tab-pane>
        <el-tab-pane :label="`未签到(${23})`" name="notQiandao">
          <div v-for="item in notFinalishList" :key="item">{{ item }}</div>
        </el-tab-pane>
      </el-tabs>
    </el-drawer>
    <!-- 查看学习情况 -->
    <el-drawer v-model="learningDetail" :title="`${'背影'}学习结果`" direction="rtl" size="20%">
      <el-tabs v-model="learningDetailActiveName" class="demo-tabs" @tab-click="handleClick">
        <el-tab-pane :label="`已学习(${12})`" name="learning">
          <div v-for="item in finalishList" :key="item">{{ item }}</div>
        </el-tab-pane>
        <el-tab-pane :label="`未学习(${23})`" name="notLearningDetail">
          <div v-for="item in notFinalishList" :key="item">{{ item }}</div>
        </el-tab-pane>
      </el-tabs>
    </el-drawer>
    <!-- 课后练习 -->
    <div class="homework">
      <el-dialog v-model="homeworkVisible" title="课后练习" width="50%">
        <el-form ref="homeworkRef" :model="homeworkForm">
          <div v-for="(item, index) in exercises" :key="item.question" style="padding-bottom: 20px">
            <el-form-item prop="index">
              <h4>第{{ index + 1 }}题</h4>
            </el-form-item>
            <el-form-item label="题目" prop="question">
              <el-input type="textarea" v-model="item.question"></el-input>
            </el-form-item>
            <el-form-item label="题目类型" prop="type">
              <el-radio-group v-model="item.type" class="ml-4">
                <el-radio label="1" size="large">判断题</el-radio>
                <el-radio label="2" size="large">选择题</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="选项" prop="option">
              <div v-for="(optionItem, optionIndex) in item.option" :key="optionItem">
                <span>{{ selectOption[optionIndex]?.label }}:</span>
                <el-input
                  style="padding-right: 10px"
                  v-model="option"
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
              <el-select v-model="item.answer">
                <el-option
                  v-for="item in selectOption.slice(0, item.option.length)"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
            <el-button type="primary" @click="deleteQuestion(index)">删除题目</el-button>
          </div>
          <el-button type="primary" @click="addQuestion">添加题目</el-button>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="homeworkVisible = false">取消</el-button>
            <el-button type="primary" @click="homeworkVisible = false">提交</el-button>
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
import CreatePopup from "@/components/CreatePopup";
import { gainAppoint } from "@/utils/utils";
import { addContent } from "@/api/crm/customer";
import { useStore } from "vuex";
import { computed } from "vue";

export default {
  name: "DetailSale",
  mixins: [TableMixin],
  components: { TableView, CreatePopup },
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
      qiandaoVisible: false,
      talkVisible: false,
      qiandaoDetail: false,
      learningDetail: false,
      homeworkVisible: false,
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
      register_tm: 1, //签到时间限制
      talk: "", // 讨论话题
      //表单验证规则
      contentfFormRules: {
        title: [{ required: true, message: "请输入名称", trigger: "change" }],
        content: [{ required: true, message: "请输入内容", trigger: "change" }]
      },
      qiandaoDetailActiveName: "qiandao",
      finalishList: ["zhangsan", "lisi"],
      notFinalishList: ["zhangsan", "lisi"],
      learningDetailActiveName: "learning",
      exercises: [
        {
          question: "",
          type: "",
          option: [""],
          answer: ""
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
    // 添加题目
    addQuestion() {
      this.exercises.push({
        question: "",
        type: "",
        option: [""],
        answer: ""
      });
    },
    // 删除题目
    deleteQuestion(index) {
      this.exercises.splice(index, 1);
    },
    // 添加选项
    addOption(index) {
      if (this.exercises[index].option.length === 5) {
        this.$message.warning("选项最多为5个");
        return;
      }
      this.exercises[index].option.push("");
    },
    deleteOption(index, optionIndex) {
      if (this.exercises[index].option.length === 2) {
        this.$message.warning("选项至少2个");
        return;
      }
      this.exercises[index].option.splice(optionIndex, 1);
    },
    handleClose() {},
    handleClick() {},
    // 文件上传
    beforeUpload(file) {
      this.content_form.file = file;
    },
    addContent() {
      this.$refs.contentFormRef.validate(async valid => {
        if (!valid) return;
        const params = { ...this.content_form, course_id: this.customer.course_id };
        addContent(params).then(res => {
          if (res && res.code === 200) {
            this.$message.success(res.message);
            this.addContentVisible = false;
          } else {
            this.$message.warning(res.message);
          }
        });
      });
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

.homework {
  /deep/ .el-dialog {
    height: 565px;
    overflow: auto;
  }
}
</style>
