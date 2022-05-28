<template>
  <slide-detail-view :loading="loading" @close="handleClose" @after-enter="afterEnter">
    <div class="detail-box">
      <div class="detail-heaher-box">
        <div class="type-name">课程详情</div>
        <!-- 简要信息展示 -->
        <div class="detail-summary-box">
          <div class="base-item" v-for="item in summaryColumns" :key="item.prop">
            <div class="base-title">{{ item.label }}</div>
            <div class="base-value">
              <el-tooltip
                class="name-box-item"
                effect="dark"
                :content="detailObj[item.prop] || '--'"
                placement="top"
              >
                {{ detailObj[item.prop] || "--" }}
              </el-tooltip>
            </div>
          </div>
        </div>
      </div>
      <!-- tab区域 -->
      <div class="detail-tabs-box">
          <detail-sale :customer="customer"></detail-sale>
      </div>
    </div>
    <create-popup
      :show="popupShow"
      :popup-type="popupType"
      :action="createAction"
      @load="editCompleted"
      @close="popupShow = false"
    />
  </slide-detail-view>
</template>
<script>
import SlideDetailView from "@/components/SlideDetailView";
import { queryCustomerDetail } from "@/api/crm/customer";
import DetailInfo from "@/views/crm/customer/DetailInfo";
import DetailSale from "@/views/crm/customer/DetailSale";
import DetailVisit from "@/views/crm/customer/DetailVIsit";
import CreatePopup from "@/components/CreatePopup";

export default {
  name: "CustomerDetail",
  components: {
    SlideDetailView,
    // DetailInfo,
    DetailSale,
    // DetailVisit,
    CreatePopup
  },
  props: {
    action: {
      type: Object,
      default: () => {
        return {
          id: 1,
          list: [],
          index: 1,
          tabName: "detail",
          data: {}
        };
      }
    }
  },
  setup() {
    return {};
  },
  data() {
    return {
      loading: false,
      detailObj: {},
      customer: {},
      index: 1,
      summaryColumns: [
        {
          label: "课程名称",
          prop: "course_name"
        },
        {
          label: "年级",
          prop: "grade"
        },
        {
          label: "班级",
          prop: "class"
        },
        {
          label: "班级人数",
          prop: "total_member"
        },
        {
          label: "创建时间",
          prop: "create_tm"
        }
      ],
      tabActiveName: "detail",
      popupShow: false,
      popupType: "CreateCustomer",
      createAction: {
        type: "edit",
        id: ""
      }
    };
  },
  created() {
    // 定位tab
    this.tabActiveName = this.action.tabName || this.tabActiveName;
    this.index = this.action.index;
    this.detailObj = this.action;
  },
  async mounted() {
    // this.queryCustomerDetail();
  },
  methods: {
    afterEnter() {},
    handleClose() {
      if (this.$route.params.customerId) delete this.$route.params.customerId;
    },

    /**
     * 编辑客户
     */
    customerEdit() {
      this.createAction = {
        type: "edit",
        id: this.detailObj.customerId
      };
      this.popupShow = true;
    },

    /**
     * 编辑客户完成回调
     */
    editCompleted() {
      this.queryCustomerDetail();
      this.$emit("load");
    },

    /**
     * 更换
     */
    changePage(type) {
      switch (type) {
        case "back":
          if (
            (this.$route.params.customerId && this.index <= 0) ||
            (!this.$route.params.customerId && this.index <= 1)
          )
            return this.$message.error("没有更多了！");
          this.index--;
          break;
        case "next":
          if (this.index >= this.action.list.length) return this.$message.error("没有更多了！");
          this.index++;
          break;
        default:
          break;
      }
      this.queryCustomerDetail();
    },

    /**
     * tab切换
     */
    tabClick() {},

    /**
     * 查询详情
     */
    async queryCustomerDetail() {
      this.loading = true;
      const params = {
        customerId: this.action.list[this.index - 1] || this.action.id
      };
      const res = await queryCustomerDetail(params).finally(() => {
        this.loading = false;
      });
      if (res && res.code === 0) {
        this.detailObj = res.data.obj;
        this.customer = {
          customerId: this.detailObj.customerId,
          customerName: this.detailObj.customerName
        };
      } else {
        this.$message.warning(res.msg);
      }
    }
  },
  watch: {
    /**
     * 点击列表换页
     */
    "action.index": {
      handler(val) {
        if (val !== this.index) {
          this.index = val;
          this.queryCustomerDetail();
        }
      },
      deep: true
    }
  }
};
</script>
<style lang="less">
.detail-box {
  display: flex;
  flex-direction: column;
  height: 100%;
  .detail-heaher-box {
    padding: 24px 24px 16px;

    .type-name {
      margin-bottom: 2px;
      font-size: 14px;
      color: #6b778c;
    }

    .detail-name-operate-box {
      display: flex;
      justify-content: space-between;
      .detail-name-page-box {
        display: flex;
        overflow: hidden;
        .detail-name-box {
          .deatil-name {
            font-size: 24px;
            font-weight: 700;
            overflow: hidden;
            text-overflow: ellipsis;
            display: -webkit-box;
            -webkit-line-clamp: 1;
            -webkit-box-orient: vertical;
          }
        }
        .header-page-btn {
          margin: 0 24px;
          min-width: 56px;
          .el-button {
            width: 28px;
            height: 28px;
            color: #344563;
            background-color: transparent;
            &:hover {
              color: #344563;
              background-color: #ebecf0;
            }
          }
        }
      }
    }

    .detail-summary-box {
      display: flex;
      text-align: center;
      width: 100%;
      padding: 16px;
      margin-top: 12px;
      background-color: #f4f5f7;
      border-radius: 3px;

      .base-item {
        flex: 0 0 20%;
        overflow: hidden;
        .base-title {
          color: #6b778c;
        }
        .base-value {
          color: #172b4d;
          font-weight: 500;
          min-height: 14px;
          margin-top: 8px;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
        }
      }
    }
  }

  .detail-tabs-box {
    flex: 1;
    height: 0;
    padding: 0 24px 8px;
    .el-tabs {
      display: flex;
      flex-direction: column;
      height: 100%;
    }
    .el-tabs__content {
      flex: 1;
      .el-tab-pane {
        height: 100%;
      }
    }
  }
}
</style>
