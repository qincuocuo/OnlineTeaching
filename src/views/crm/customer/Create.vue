<template>
  <create-popup-view :loading="loading" :title="title" @close="handleClose" @save="handleSave">
    <create-popup-view-section title="课程信息">
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
import { addCourse, updateCourse } from "@/api/crm/customer";
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
      title: "新建课程",
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
            field: "course_name",
            title: "课程名称",
            validate: [{ required: true, message: "请输入课程名称", trigger: "blur" }]
          },
          {
            type: "select",
            field: "grade",
            title: "年级",
            options: [
              {
                value: 1,
                label: "一年级"
              },
              {
                value: 2,
                label: "二年级"
              },
              {
                value: 3,
                label: "三年级"
              },
              {
                value: 4,
                label: "四年级"
              },
              {
                value: 5,
                label: "五年级"
              },
              {
                value: 6,
                label: "六年级"
              }
            ],
            validate: [{ required: true, message: "请选择年级", trigger: "change" }]
          },
          {
            type: "input",
            field: "class",
            title: "班级",
            props: {
              placeholder: "仅允许输入数字"
            },
            validate: [{ required: true, message: "请输入班级名称", trigger: "blur" }],
            on: {
              input: val => {
                this.form.class = val.replace(/[^\d]/g, "");
              }
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
          const actionApi = this.action.type === "add" ? addCourse : updateCourse;
          params.class = Number(params.class);
          const res = await actionApi(params);
          if (res && res.code === 200) {
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
      for (let key in this.form) {
        if (Object.prototype.hasOwnProperty.call(this.action.data, key)) {
          this.form[key] = this.action.data[key];
        }
      }
    }
  }
};
</script>
<style lang="less" scoped></style>
