import { filterObj } from "@/utils/utils";
import { queryList } from "@/api/crm/common";
export default {
  data() {
    return {
      loading: false, // 加载动画
      disableMixinInit: false, // 初始不获取数据
      tableHeight: document.documentElement.clientHeight, // 表的高度
      /* 查询条件 */
      queryParam: {},
      url: {
        list: "",
        field: "list" // response字段
      },
      /* 数据源 */
      dataSource: [{
        course_name:'科目三',
        grade: '1',
        class: '1401',
        total_member: '20',
        create_tm: '2022-5-20 12:50:00'
      }],
      /* 分页参数 */
      ipagination: {
        page: 1,
        page_size: 20,
        sizes: [10, 20, 30],
        total: 0
      },
      queryParamTime: [], // 时间
      pickerOptions: {
        //时间选择器控制条件
        disabledDate(time) {
          return time.getTime() > Date.now();
        }
      },
      selectionList: [], // 勾选列表
      singleSelected: true, // 是否单选
      ignoreSelectedChange: false, // 忽略勾选数据change 避免触发chang事件
      refDataTable: "refDataTable" + Math.random(), // table的ref属性
      whetherSelection: false, // 是否实现勾选列
      whetherIndex: false, // 是否显示序号
      formOptions: {
        fApi: {},
        options: {
          submitBtn: false,
          resetBtn: false,
          form: {
            labelWidth: "80px"
          }
        },
        rule: []
      }
    };
  },
  mounted() {
    if (!this.disableMixinInit) {
      this.loadData();
    }
  },
  methods: {
    /**
     * 获取列表数据
     */
    async loadData(arg) {
      if (!this.url.list) {
        this.$message.error("请设置url.list属性!");
        return;
      }
      this.loading = true;
      //传入 1 则加载第一页的内容
      if (arg === 1) {
        this.ipagination.page = 1;
      }
      const params = this.getQueryParams(); //查询条件
      const res = await queryList(this.url.type || "post", this.url.list, params)
        .catch(() => {
          this.loading = false;
        })
        .finally(() => {
          this.loading = false;
        });
      if (!res) return;
      if (res.code === 200) {
        this.dataSource = res?.data?.result || res?.data[this.url.field] || [];
        this.dataSource.forEach((item, index) => {
          item.index = index + 1 + (params.page - 1) * params.page_size;
        });
        this.ipagination.total = res.data.count;
      } else {
        this.$message.warning(res.msg);
      }
    },

    /**
     * 获取请求参数
     */
    getQueryParams() {
      //获取查询条件
      const param = Object.assign({}, this.queryParam);
      param.create_tm = param?.create_tm?.join('--');
      param.page = this.ipagination.page;
      param.page_size = this.ipagination.page_size;
      return filterObj(param);
    },

    /**
     * 该行是否可勾选
     * @param {*,cb}
     */
    checkSelectable(row, cb) {
      cb(true);
    },

    /**
     * 字段排序（后端）
     */
    sortChange(column) {
      this.queryParam.sortData = column || {};
      this.searchQuery();
    },

    /**
     * 勾选操作
     * @param {*} val
     */
    handleSelectionChange(val, table) {
      if (this.ignoreSelectedChange) {
        return;
      }
      // 单选操作
      if (this.singleSelected && val.length > 1) {
        const lastObj = val[val.length - 1];
        this.ignoreSelectedChange = true;
        table.clearSelection();
        this.$nextTick(() => {
          this.ignoreSelectedChange = false;
          table.toggleRowSelection(lastObj);
        });
        return;
      } else {
        this.selectionList = val;
      }
    },

    /**
     * 执行查询
     */
    searchQuery() {
      this.loadData(1);
    },

    /**
     * 查询条件重置
     */
    searchReset() {
      if (Object.keys(this.formOptions.fApi).length) {
        this.formOptions.fApi.resetFields();
        this.formOptions.fApi.reload();
      } else {
        this.queryParam = {};
      }
      this.$nextTick(() => {
        this.searchQuery();
      });
    }
  }
};
