<template>
  <create-popup-view :loading="loading" :title="title" @close="handleClose" @save="handleSave">
    <create-popup-view-section title="我的信息">
      <div class="my-information">
        <div v-for="item in fields" :key="item" class="my-information-item">
          <div class="label-box">{{ item.label + "：" }}</div>
          <div>
            {{ infoDeal(item.value) }}
          </div>
          <div>
            <el-button
              v-if="item.value === 'password'"
              type="text"
              class="edit-password"
              @click="editPassword"
            >
              修改密码
            </el-button>
          </div>
        </div>
      </div>
    </create-popup-view-section>
    <create-nest-popup
      :show="nestPopupShow"
      :popup-type="nestPopupType"
      :action="createNestAction"
      @close="nestPopupShow = false"
    />
  </create-popup-view>
</template>
<script>
import CreatePopupView from "@/components/CreatePopupView";
import CreatePopupViewSection from "@/components/CreatePopupViewSection";
import CreateNestPopup from "@/components/CreateNestPopup.vue";
import { getUser } from "@/api/login";
import { timeStr } from "@/utils/utils";
import _ from "lodash";
import { useStore } from "vuex";
import { computed } from "vue";
export default {
  name: "UserInfo",
  components: {
    CreatePopupView,
    CreatePopupViewSection,
    CreateNestPopup
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
      title: "",
      myInformation: {},
      fields: [
        { value: "user_name", label: "用户名" },
        { value: "password", label: "密码" },
        { value: "role", label: "角色" },
        { value: "user_id", label: "学号/工号" },
        { value: "grade", label: "年级" },
        { value: "class", label: "班级" },
        { value: "login_time", label: "注册时间" },
        {
          value: "last_login_time",
          label: "最近登录时间"
        }
      ],
      nestPopupShow: false,
      nestPopupType: "EditPassword",
      createNestAction: {
        type: "add",
        id: "",
        data: {}
      }
    };
  },
  mounted() {
    this.getUser();
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
        })
        .catch(() => {});
    },
    // 获取用户信息
    getUser() {
      this.loading = true;
      getUser({
        user_id: this.userInfo.user_id
      })
        .then(res => {
          if (res && res.code === 200) {
            this.myInformation = res.data;
          }
        })
        .finally(() => {
          this.loading = false;
        });
    },
    // 字段处理
    infoDeal(val) {
      let str;
      switch (val) {
        case "login_time":
        case "last_login_time":
          str = timeStr(this.myInformation[val]);
          break;
        case "role":
          str = this.myInformation[val] === 1 ? "教师" : "学生";
          break;
        case "password":
          str = "************";
          break;
        default:
          str = this.myInformation[val];
          break;
      }
      return str;
    },
    editPassword() {
      this.nestPopupShow = true;
    }
  }
};
</script>
<style lang="less" scoped>
.my-information-item {
  display: flex;
  padding: 8px 0;
  line-height: 20px;
  .label-box {
    width: 140px;
    text-align: right;
    font-weight: bold;
  }
  .edit-password {
    padding: 0;
    height: 20px;
    margin-left: 8px;
  }
}
</style>
