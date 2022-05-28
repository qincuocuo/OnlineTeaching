<template>
  <div class="detail-visit">
    <div class="table-head-container">
      <div class="query-add-btns-container">
        <div class="add-btns">
          <el-input
            v-model="searchContent"
            class="content-search"
            placeholder="请输入学习内容"
            prefix-icon="Search"
          />
          <el-button class="content-add" @click="addContentVisible = true" type="primary">
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
      >
        <template v-slot:salesLeadSourceId="scope">
          {{
            gainAppoint(sourceOption, scope.row.salesLeadSourceId, "customerSourceId")
              .customerSource || "--"
          }}
        </template>
        <template v-slot:operate="scope">
          <div v-has="'visit_edit'" class="table-btn-box">
            <el-button type="text" @click="edit(scope.row)">查看学习内容</el-button>
          </div>
          <div v-has="'del'" class="table-btn-box">
            <el-button type="text" @click="learningDetail = true">查看学习情况</el-button>
          </div>
          <div v-has="'visit_edit'" class="table-btn-box">
            <el-button type="text" @click="qiandaoVisible = true">发起签到</el-button>
          </div>
          <div v-has="'del'" class="table-btn-box">
            <el-button type="text" @click="talkVisible = true">发起讨论</el-button>
          </div>
          <div v-has="'visit_edit'" class="table-btn-box">
            <el-button type="text" @click="qiandaoDetail = true">查看签到结果</el-button>
          </div>
          <div v-has="'del'" class="table-btn-box">
            <el-button type="text">查看讨论情况</el-button>
          </div>
        </template>
      </table-view>
    </div>
    <!-- 新增课程内容 -->
    <el-dialog
      v-model="addContentVisible"
      title="新增课程内容"
      width="40%"
      :before-close="handleClose"
    >
      <el-form ref="contentFormRef" :model="content_form" :rules="contentfFormRules">
        <el-form-item prop="title" class="login-email" label="名称">
          <el-input placeholder="" v-model.trim="content_form.title" class="email-input"></el-input>
        </el-form-item>
        <el-form-item prop="content" class="password" label="内容">
          <el-input placeholder="" v-model.trim="content_form.content" type="text"></el-input>
        </el-form-item>
        <el-upload
          class="upload-demo"
          drag
          action="https://jsonplaceholder.typicode.com/posts/"
          multiple
        >
          <el-icon class="el-icon--upload"><upload-filled /></el-icon>
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
    <el-dialog v-model="qiandaoVisible" title="发起签到" width="40%" :before-close="handleClose">
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
    <el-dialog v-model="talkVisible" title="讨论话题" width="40%" :before-close="handleClose">
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
          slot: "unlearned",
          width: 50
        },
        {
          label: "操作",
          slot: "operate",
          width: 250
        }
      ],
      url: {
        list: "/crm/admin/api/sales_lead/query/list_in_customer",
        type: "get"
      },
      disableMixinInit: true,
      popupShow: false,
      addContentVisible: false,
      qiandaoVisible: false,
      talkVisible: false,
      qiandaoDetail: false,
      learningDetail: false,
      popupType: "CreateOpportunity",
      createAction: {
        type: "add",
        id: "",
        data: {}
      },
      sourceOption: [],
      searchContent: "",
      content_form: {
        title: "",
        content: ""
      },
      register_tm: "", //签到时间限制
      talk: "", // 讨论话题
      //表单验证规则
      contentfFormRules: {
        title: [{ required: true, message: "请输入名称", trigger: "change" }],
        content: [{ required: true, message: "请输入内容", trigger: "change" }]
      },
      qiandaoDetailActiveName: "qiandao",
      finalishList: ["zhangsan", "lisi"],
      notFinalishList: ["zhangsan", "lisi"],
      learningDetailActiveName: 'learning',

    };
  },

  watch: {
    "customer.customerId": {
      handler(val) {
        this.queryParam.customerId = val;
        if (val) this.loadData();
      },
      immediate: true,
      deep: true
    }
  },
  methods: {
    addContent() {
      this.$refs.contentFormRef.validate(async valid => {
        if (!valid) return;
        this.addContentVisible = false;
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
</style>
