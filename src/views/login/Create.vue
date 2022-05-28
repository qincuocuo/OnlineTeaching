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
import { register } from "@/api/login";
import _ from "lodash";
import { useStore } from "vuex";
import { computed } from "vue";
export default {
  name: "CreateUser",
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
      title: "注册用户",
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
            field: "username",
            title: "用户名",
            validate: { required: true, message: "请输入用户名", trigger: "change" }
          },
          {
            type: "select",
            field: "role",
            title: "用户角色",
            validate: { required: true, message: "请输入用户角色", trigger: "change" },
            options: [
              {
                label: "教师",
                value: 1
              },
              {
                label: "学生",
                value: 2
              }
            ],
            on: {
              change: val => {
                this.formOptions.fApi.getRule("grade").hidden = val === 1;
                this.formOptions.fApi.getRule("class").hidden = val === 1;
              }
            }
          },
          {
            type: "select",
            field: "grade",
            title: "年级",
            validate: { required: true, message: "请输入年级", trigger: "change" },
            hidden: true,
            options: [
              {
                value: 1
              },
              {
                value: 2
              },
              {
                value: 3
              },
              {
                value: 4
              },
              {
                value: 5
              },
              {
                value: 6
              }
            ]
          },
          {
            type: "select",
            field: "class",
            title: "班级",
            hidden: true,
            validate: { required: true, message: "请输入班级", trigger: "change" },
            options: [
              {
                value: 1
              },
              {
                value: 2
              },
              {
                value: 3
              },
              {
                value: 4
              },
              {
                value: 5
              },
              {
                value: 6
              }
            ]
          },
          {
            type: "input",
            field: "password",
            title: "用户密码",
            validate: { required: true, message: "请输入密码", trigger: "change" },
            props: {
              type: "password"
            }
          },
          {
            type: "input",
            field: "confirm",
            title: "确认密码",
            validate: { required: true, message: "请输入密码", trigger: "change" },
            props: {
              type: "password"
            }
          },
          {
            type: "input",
            field: "user_id",
            title: "学生学号/教师工号",
            validate: { required: true, message: "请输入学生学号/教师工号", trigger: "change" }
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
          const res = await register(params);
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
