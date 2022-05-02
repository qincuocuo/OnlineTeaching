<template>
  <div class="detail-visit">
    <div class="table-head-container">
      <div class="query-add-btns-container">
        <div class="add-btns">
          <el-button v-has="'sale_add'" @click="add" type="primary">新增</el-button>
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
            <el-button type="text" @click="edit(scope.row)">编辑</el-button>
          </div>
          <div v-has="'del'" class="table-btn-box">
            <el-button type="text">删除</el-button>
          </div>
        </template>
      </table-view>
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
import { queryCustomerSource } from "@/api/crm/customer";
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
          label: "机会名称",
          prop: "salesLeadName",
          width: 140
        },
        {
          label: "公司",
          prop: "customerName",
          width: 180
        },
        {
          label: "来源",
          slot: "salesLeadSourceId",
          width: 100
        },
        {
          label: "跟进人",
          prop: "salesmanName",
          width: 100
        },
        {
          label: "操作",
          slot: "operate",
          width: 120
        }
      ],
      url: {
        list: "/crm/admin/api/sales_lead/query/list_in_customer",
        type: "get"
      },
      disableMixinInit: true,
      popupShow: false,
      popupType: "CreateOpportunity",
      createAction: {
        type: "add",
        id: "",
        data: {}
      },
      sourceOption: []
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
    add() {
      this.createAction = this.$options.data().createAction;
      this.createAction.data = this.customer;
      this.popupShow = true;
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
  .table-view-container {
    flex: 1;
    width: 100%;
  }
}
</style>
