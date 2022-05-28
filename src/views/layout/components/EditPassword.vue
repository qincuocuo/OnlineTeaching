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
import { changePassword } from "@/api/login";
import _ from "lodash";
import { useStore } from "vuex";
import { computed } from "vue";
export default {
  name: "EditPassword",
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
      userInfo: computed(() => store.getters.userInfo)
    };
  },
  data() {
    return {
      loading: false,
      title: "修改密码",
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
            field: "password",
            title: "原密码",
            validate: { required: true, message: "请输入原密码", trigger: "change" },
            props: {
              type: "password"
            }
          },
          {
            type: "input",
            field: "new_password",
            title: "新密码",
            validate: { required: true, message: "请输入密码", trigger: "change" },
            props: {
              type: "password"
            }
          }
        ]
      }
    };
  },
  async mounted() {},
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
          const res = await changePassword(params);
          if (res && res.code === 200) {
            this.$emit("close");
            this.$message.success("成功");
          } else {
            this.$message.warning(res.error);
          }
        })
        .catch(() => {});
    }
  }
};
</script>
<style lang="less" scoped></style>
