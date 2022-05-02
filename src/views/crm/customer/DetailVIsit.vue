<template>
  <div class="detail-visit">
    <div class="table-head-container">
      <div class="query-add-btns-container">
        <div class="add-btns">
          <el-button v-has="'visit_add'" @click="add" type="primary">新增</el-button>
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
        <template v-slot:visitModeId="scope">
          {{ gainAppoint(visitMode, scope.row.visitModeId).label || "--" }}
        </template>
        <template v-slot:operate="scope">
          <div v-has="'visit_edit'" class="table-btn-box">
            <el-button type="text" @click="edit(scope.row)">编辑</el-button>
          </div>
        </template>
        <template v-slot:tel="scope">
          <div v-loading="scope.row.loading">
            <div class="table-tel-box">
              <el-tooltip effect="dark" :content="scope.row.tel" placement="top">
                {{ scope.row.tel === undefined ? "*******" : scope.row.tel || "--" }}
              </el-tooltip>
            </div>
          </div>
        </template>
      </table-view>
    </div>
    <create-popup
      ref="refCreatePoup"
      :show="popupShow"
      :popup-type="popupType"
      :action="createAction"
      @load="loadData"
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
  name: "DetailVisit",
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
          label: "客户接待",
          prop: "receptionist",
          width: 140
        },
        {
          label: "接待职位",
          prop: "receptionistPosition",
          width: 140
        },
        {
          label: "电话",
          prop: "tel",
          slot: "tel",
          showOverflowTooltip: false,
          width: 100
        },
        {
          label: "拜访方式",
          prop: "visitModeId",
          slot: "visitModeId",
          width: 80
        },
        {
          label: "机会",
          prop: "visitOpportunity",
          width: 80
        },
        {
          label: "销售",
          prop: "salesmanName",
          width: 120
        },
        {
          label: "拜访计划",
          prop: "visitPreparation",
          width: 120
        },
        {
          label: "拜访结果",
          prop: "visitResult",
          width: 120
        },
        {
          label: "操作",
          prop: "operate",
          slot: "operate",
          width: 120
        }
      ],
      url: {
        list: "/crm/admin/api/visit_record/query/list_in_customer",
        type: "get"
      },
      disableMixinInit: true,
      popupShow: false,
      popupType: "CreateVisit",
      createAction: {
        type: "add",
        id: "",
        data: {}
      }
    };
  },
  mounted() {},
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
        id: item.visitId,
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
    :deep(.table-tel-box span) {
      display: inline-block;
      width: 100%;
      overflow: hidden;
      white-space: nowrap;
      text-overflow: ellipsis;
    }
  }
}
</style>
