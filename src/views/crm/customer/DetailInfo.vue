<template>
  <div class="detail-info">
    <div v-for="item in detailColumns" :key="item.prop" :class="['detail-info-item', item.prop]">
      <div class="info-item-label">{{ item.label }}</div>
      <div class="info-item-value" v-if="item.prop === 'ipoFlag'">
        {{
          [undefined, ""].includes(detailObj[item.prop]) ? "--" : detailObj[item.prop] ? "是" : "否"
        }}
      </div>
      <div class="info-item-value" v-else>{{ detailObj[item.prop] || "--" }}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: "DetailInfo",
  props: {
    detailObj: {
      type: Object,
      default: () => {
        return {};
      }
    }
  },
  setup() {},
  data() {
    return {
      detailColumns: [
        {
          label: "客户名称",
          prop: "customerName"
        },
        {
          label: "联系人",
          prop: "contacts"
        },
        {
          label: "公司职位",
          prop: "position"
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
          label: "所属行业",
          prop: "industry"
        },
        {
          label: "生命周期",
          prop: "lifecycle"
        },
        {
          label: "是否上市",
          prop: "ipoFlag"
        },
        {
          label: "客户来源",
          prop: "customerSource"
        },
        {
          label: "客户归属",
          prop: "salesmanName"
        },
        {
          label: "地址",
          prop: "address"
        },
        {
          label: "详细地址",
          prop: "addressDetail"
        },
        {
          label: "业内合作情况",
          prop: "industryCooperation"
        }
      ]
    };
  },
  mounted() {},
  methods: {}
};
</script>

<style lang="less" scoped>
.detail-info {
  display: flex;
  flex-wrap: wrap;
  max-height: 100%;
  overflow: auto;
  .detail-info-item {
    flex: 0 0 50%;
    display: flex;
    width: auto;
    padding: 14px 8px;
    .info-item-label {
      width: 100px;
      flex-shrink: 0;
      margin-right: 10px;
      font-size: 13px;
      line-height: 21px;
      color: #777;
    }
    .info-item-value {
      flex: 1;
      font-size: 14px;
      line-height: 21px;
      white-space: pre-wrap;
      word-break: break-all;
    }
  }
  .industryCooperation {
    flex: 0 0 100%;
  }
}
</style>
