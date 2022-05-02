<template>
  <div class="table-view" :ref="refTabelView">
    <div class="table-box">
      <el-table
        :data="dataSource"
        :max-height="tableMaxHeight || maxHeight"
        :key="tableKey.value"
        :ref="refDataTable"
        :span-method="row => $emit('arraySpanMethod', row)"
        @selection-change="val => $emit('selection-change', val, $refs[refDataTable])"
        class="n-table--border"
        use-virtual
        stripe
        border
        highlight-current-row
        style="width: 100%"
        header-cell-class-name="table-header"
        :row-key="rowKey"
        :tree-props="treeProps"
      >
        <el-table-column
          v-if="whetherIndex"
          type="index"
          width="60"
          label="序号"
          align="left"
        ></el-table-column>
        <el-table-column
          v-if="whetherSelection"
          :selectable="checkSelectable"
          type="selection"
          width="55"
          align="left"
          :label-class-name="singleSelected ? 'single-choice' : ''"
        ></el-table-column>
        <template v-for="(item, index) in columns" :key="index">
          <el-table-column
            :prop="item.prop"
            :label="item.label"
            :min-width="item.width"
            align="left"
            :fixed="item.fixed"
            :column-key="item.prop"
            :class-name="item.class || ''"
            :show-overflow-tooltip="
              item.showOverflowTooltip === undefined ? true : item.showOverflowTooltip
            "
          >
            <template #default="scope">
              <!-- 插槽-自定义内容 ( 在columns该列对应属性中添加slot属性名。) -->
              <slot v-if="item.slot" :name="item.slot" :row="scope.row"></slot>
              <!-- 默认显示文字 -->
              <span v-else>{{ scope.row[item.prop] || "--" }}</span>
            </template>
          </el-table-column>
        </template>
      </el-table>
    </div>
    <!--分页符-->
    <el-row class="p-contianer" v-if="ipagination && Object.keys(ipagination).length">
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :page-sizes="ipagination.sizes"
        :page-size="ipagination.size"
        layout="total, sizes, prev, pager, next, jumper"
        :total="ipagination.total"
        :pager-count="5"
        class="p-bar"
        background
      ></el-pagination>
    </el-row>
  </div>
</template>
<script>
export default {
  name: "TabelView",
  props: {
    columns: {
      type: Array,
      default: () => {
        return [];
      }
    },
    dataSource: {
      type: Array,
      default: () => {
        return [];
      }
    },
    ipagination: {
      type: Object,
      default: () => {
        return {};
      }
    },
    whetherIndex: {
      // 是否显示序号
      default: false
    },
    whetherSelection: {
      type: Boolean,
      // 是否显示选择框
      default: false
    },
    singleSelected: {
      type: Boolean,
      // 是否单选
      default: false
    },
    queryParam: {
      // 查询条件
      default: () => {
        return {};
      }
    },
    tableFullScreen: {
      // 是否占满全屏
      type: Boolean,
      default: true
    },
    tableMaxHeight: {
      // 最大高度
      type: Number,
      default: 0
    },
    tableKey: {
      default: () => {
        return {
          value: Math.random()
        };
      }
    },
    refDataTable: {
      type: String,
      default: "refDataTable" + Math.random()
    },
    // 树形展示
    treeProps: {
      type: Object,
      default: () => {
        return {};
      }
    },
    rowKey: {
      type: String,
      default: ""
    }
  },
  components: {},
  data() {
    return {
      maxHeight: 400,
      refTabelView: "refTabelView" + Math.random(),
      // 忽略勾选数据change 避免触发chang事件
      ignoreSelectedChange: false
    };
  },
  mounted() {
    // table表格是否一屏内显示
    if (this.tableFullScreen) {
      this.tableHeightCalculate();
      window.addEventListener("resize", this.tableHeightCalculate);
    }
  },
  methods: {
    /**
     * 分页数量变化
     *
     */
    handleSizeChange(newSize) {
      this.$emit("update:ipagination", { ...this.ipagination, size: newSize });
      this.$emit("load");
    },

    /**
     * 页数变化
     *
     */
    handleCurrentChange(newPage) {
      this.$emit("update:ipagination", { ...this.ipagination, page: newPage });
      this.$emit("load");
    },

    /**
     * 是否可勾选
     *
     */
    checkSelectable(row) {
      let selectable = true;
      this.$emit("check-selectable", row, function (val) {
        selectable = val;
      });
      return selectable;
    },

    // table高度全屏
    tableHeightCalculate() {
      if (!this.$refs[this.refTabelView]) return;
      this.$nextTick(() => {
        this.maxHeight =
          this.$refs[this.refTabelView].offsetHeight -
          (Object.keys(this.ipagination).length ? 44 : 0);
      });
    }
  }
};
</script>
<style lang="less" scoped>
.table-view {
  position: relative;
  height: 100%;
  display: flex;
  flex-direction: column;
  .table-box {
    flex: 1;
  }
}
</style>
