<template>
  <create-popup-view :loading="loading" :title="title" @close="handleClose" @save="handleSave">
    <create-popup-view-section title="基本信息">
      <form-create
        v-model:api="formOptions.fApi"
        v-model="form"
        :rule="formOptions.rule"
        :option="formOptions.options"
      />
    </create-popup-view-section>
  </create-popup-view>
</template>
<script>
import CreatePopupView from "@/components/CreatePopupView";
import CreatePopupViewSection from "@/components/CreatePopupViewSection";
import { addCustomer, updateCustomer } from "@/api/crm/customer";
import _ from "lodash";
import { useStore } from "vuex";
import { computed } from "vue";
export default {
  name: "CreateCustomer",
  components: {
    CreatePopupView,
    CreatePopupViewSection
  },
  props: {
    action: {
      type: Object,
      default: () => {
        return {
          type: "add",
          id: "",
          data: {}
        };
      }
    }
  },
  setup() {
    const store = useStore();
    return {
      userInfo: computed(() => store.getters.userInfo),
      publicKey: computed(() => store.getters.publicKey)
    };
  },
  data() {
    return {
      loading: false,
      title: "新增客户",
      form: {},
      formOptions: {
        fApi: {},
        options: {
          submitBtn: false,
          global: {
            //设置所有组件
            "*": {
              props: {
                clearable: true,
                "show-word-limit": true,
                maxlength: 50
              }
            }
          }
        },
        rule: [
          {
            type: "input",
            field: "contacts",
            title: "联系人"
          },

          {
            type: "input",
            field: "tel",
            title: "联系电话",
            validate: [{ min: 7, max: 30, message: "长度在 7 到 30 个数字" }],
            props: {
              minlength: 7,
              maxlength: 30
            },
            on: {
              input: val => {
                this.form.tel = val.replace(/[^\d-]/g, "");
              }
            }
          },
          {
            type: "input",
            field: "email",
            title: "邮箱"
          },
          {
            type: "select",
            field: "ipoFlag",
            title: "是否上市",
            options: [
              {
                label: "否",
                value: false
              },
              {
                label: "是",
                value: true
              }
            ]
          },
          {
            type: "input",
            field: "industryCooperation",
            title: "业内合作情况",
            props: {
              type: "textarea",
              maxlength: 2000,
              rows: 4
            },
            col: {
              span: 48
            }
          }
        ]
      }
    };
  },
  async mounted() {
    this.$nextTick(() => {
      this.getField();
    });
  },
  methods: {
    afterEnter() {},
    handleClose() {},

    /**
     * 保存
     */
    handleSave() {
      this.formOptions.fApi
        .validate(async valid => {
          if (valid !== true) return this.$message.warning("请输入完整信息！");
          const params = _.cloneDeep(this.form);
          const actionApi = this.action.type === "add" ? addCustomer : updateCustomer;
          const res = await actionApi(params);
          if (res && res.code === 0) {
            this.$emit("close");
            this.$emit("load");
            this.$message.success(res.msg);
          } else {
            this.$message.warning(res.msg);
          }
        })
        .catch(() => {});
    },

    /**
     * 编辑-回显数据
     */
    async getField() {
      if (this.action.type === "add") {
        return;
      }
      this.title = "编辑";
      // this.loading = true;
      // 查询数据
    }
  }
};
</script>
<style lang="less" scoped></style>
