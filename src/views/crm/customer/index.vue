<template>
  <div class="customer">
    <div class="table-head-container">
      <div class="filter-container">
        <form-create
          v-model:api="formOptions.fApi"
          v-model="queryParam"
          :rule="formOptions.rule"
          :option="formOptions.options"
          @keyup.enter="searchQuery"
        />
        <div class="query-add-btns-container">
          <div class="add-btns">
            <el-button v-has="'teach'" @click="add" :disabled="slideShow" type="primary">
              新增
            </el-button>
          </div>
          <div class="query-btns">
            <el-button type="primary" v-clickdebounce="searchQuery">查询</el-button>
            <el-button v-clickdebounce="searchReset">重置</el-button>
          </div>
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
        <template v-slot:course_name="scope">
          <div class="table-visit--underline" @click="viewDetails(scope.row)">
            {{ scope.row.course_name || "--" }}
          </div>
        </template>
        <template v-slot:operate="scope">
          <div v-has="'teach'" class="table-btn-box">
            <el-button type="text" @click="edit(scope.row)">编辑</el-button>
          </div>
          <div v-has="'teach'" class="table-btn-box">
            <el-button type="text" @click="itemHandle('del', scope.row)">删除</el-button>
          </div>
          <div v-has="'student'" class="table-btn-box">
            <el-button type="text" @click="viewDetails(scope.row)">详情</el-button>
          </div>
        </template>
      </table-view>
    </div>
    <create-popup
      :show="popupShow"
      :popup-type="popupType"
      :action="createAction"
      @load="loadData"
      @close="popupShow = false"
    />
    <slide-detail
      :show="slideShow"
      :action="detailAction"
      :slide-type="slideType"
      @load="loadData"
      @close="slideShow = false"
    ></slide-detail>
  </div>
</template>
<script>
import TableView from "@/views/crm/components/TableView";
import TableMixin from "@/views/crm/mixins/Table";
import CreatePopup from "@/components/CreatePopup";
import SlideDetail from "@/components/SlideDetail";
import { deleteCourse } from "@/api/crm/customer";
import { computed } from "vue";
import { useStore } from "vuex";

export default {
  /** 客户管理 的 客户列表 */
  name: "CustomerIndex",
  mixins: [TableMixin],
  components: { TableView, CreatePopup, SlideDetail },
  setup() {
    const store = useStore();
    return {
      userInfo: computed(() => store.getters.userInfo)
    };
  },
  data() {
    return {
      columns: [
        {
          label: "课程名",
          prop: "course_name",
          slot: "course_name",
          width: 140
        },
        {
          label: "年级",
          prop: "grade",
          width: 120
        },
        {
          label: "班级",
          prop: "class",
          width: 100
        },
        {
          label: "班级人数",
          prop: "total_member",
          width: 100
        },
        {
          label: "创建时间",
          prop: "create_tm",
          width: 200
        },
        {
          label: "操作",
          slot: "operate",
          width: 180
        }
      ],
      url: {
        list: "/api/v1/course"
      },
      formOptions: {
        fApi: {},
        options: {
          submitBtn: false,
          resetBtn: false,
          form: {
            labelWidth: "100px"
          },
          global: {
            //设置所有组件
            "*": {
              props: {
                clearable: true,
                maxlength: 50
              }
            }
          }
        },
        rule: [
          {
            type: "select",
            field: "grade",
            title: "年级",
            options: [
              {
                value: 1,
                label: "一年级"
              },
              {
                value: 2,
                label: "二年级"
              },
              {
                value: 3,
                label: "三年级"
              },
              {
                value: 4,
                label: "四年级"
              },
              {
                value: 5,
                label: "五年级"
              },
              {
                value: 6,
                label: "六年级"
              }
            ]
          },
          {
            type: "select",
            field: "class",
            title: "班级",
            options: [
              {
                value: 1
              }
            ]
          },
          {
            type: "DatePicker",
            field: "section_day",
            title: "创建时间",
            props: {
              type: "datetimerange",
              placeholder: "请选择时间"
            }
          },
          {
            type: "input",
            field: "course",
            title: "课程名",
            props: {
              "suffix-icon": "Search"
            }
          }
        ]
      },
      popupShow: false,
      popupType: "CreateCustomer",
      slideShow: false,
      slideType: "CustomerDetail",
      createAction: {
        type: "add",
        id: "",
        data: {}
      },
      detailAction: {
        id: 1,
        list: [],
        data: {}
      }
    };
  },
  async mounted() {
    // if (hasFun("cus_query")) await this.loadData();
  },
  methods: {
    add() {
      this.createAction = this.$options.data().createAction;
      this.popupShow = true;
    },

    /**
     * 事项操作
     */
    edit(item) {
      this.createAction = {
        type: "edit",
        id: item.customerId,
        data: item
      };
      this.popupShow = true;
    },

    /**
     * 查看详情
     */
    viewDetails(item) {
      this.detailAction = {
        id: item.customerId,
        list: this.dataSource.map(it => {
          return it.customerId;
        }),
        index: item.index,
        tabName: item.tabName
      };
      this.detailAction = Object.assign(this.detailAction, item);
      this.slideShow = true;
    },

    /**
     * 事项操作
     */
    itemHandle(type, item, params) {
      if (type === "del") {
        this.$confirm("您确定要删除这一条数据吗?", "提示", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        })
          .then(async () => {
            const res = await deleteCourse({ customerId: item.customerId }).catch(() => {});
            if (res && res.code === 200) {
              this.$message.success(res.msg);
              this.loadData();
            } else {
              this.$message.warning(res.msg);
            }
          })
          .catch(() => {});
      }
    }
  }
};
</script>
<style lang="less" scoped>
.customer {
  height: 100%;
  display: flex;
  flex-direction: column;
  .table-head-container {
    padding: 8px 0 0;
  }
  .table-view-container {
    flex: 1;
    height: 0;
  }
}
</style>
