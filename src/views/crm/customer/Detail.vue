<template>
  <slide-detail-view :loading="loading" @close="handleClose" @after-enter="afterEnter">
    <div class="detail-box">
      <div class="detail-heaher-box">
        <div class="type-name">客户</div>
        <!-- 名称和操作区域 -->
        <div class="detail-name-operate-box">
          <div class="detail-name-page-box">
            <div class="detail-name-box">
              <el-tooltip
                class="name-box-item"
                effect="dark"
                :content="detailObj.customerName"
                placement="top"
              >
                <span class="deatil-name">
                  {{ detailObj.customerName }}
                </span>
              </el-tooltip>
            </div>
            <div class="header-page-btn">
              <el-button-group>
                <el-button type="text" @click="changePage('back')">
                  <el-icon><arrow-left /></el-icon>
                </el-button>
                <el-button type="text" @click="changePage('next')">
                  <el-icon><arrow-right /></el-icon>
                </el-button>
              </el-button-group>
            </div>
          </div>
          <div class="detail-operate-box">
            <el-button v-has="'cus_edit'" type="primary" @click="customerEdit">编辑</el-button>
          </div>
        </div>
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
        <el-tabs v-model="tabActiveName" class="detail" @tab-click="tabClick">
          <el-tab-pane label="详细信息" name="detail" :lazy="true">
            <detail-info :detailObj="detailObj"></detail-info>
          </el-tab-pane>
          <el-tab-pane v-has="'cus_detial_visit'" label="拜访记录" name="visit" :lazy="true">
            <detail-visit :customer="customer"></detail-visit>
          </el-tab-pane>
          <el-tab-pane v-has="'cus_detial_sale'" label="销售机会" name="sale" :lazy="true">
            <detail-sale :customer="customer"></detail-sale>
          </el-tab-pane>
        </el-tabs>
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
    DetailInfo,
    DetailSale,
    DetailVisit,
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
          label: "联系人",
          prop: "contacts"
        },
        {
          label: "联系电话",
          prop: "tel"
        },
        {
          label: "邮箱",
          prop: "email"
        },
        {
          label: "客户归属",
          prop: "salesmanName"
        },
        {
          label: "所属行业",
          prop: "industry"
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
  },
  async mounted() {
    this.queryCustomerDetail();
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
