<template>
  <div class="login_container">
    <div class="login_box">
      <el-row class="logo-row">
        <el-col class="logo-col">
          <!-- <img class="logo-img" src="@/assets/images/login_logo.svg" alt="" /> -->
          CRM
        </el-col>
      </el-row>
      <!-- 表单 -->
      <el-form ref="loginFormRef" class="login_form" :model="form" :rules="formRules">
        <el-form-item prop="user_id" class="login-email">
          <el-input placeholder="账户名" v-model.trim="form.user_id" class="email-input"></el-input>
        </el-form-item>
        <el-form-item prop="password" class="password">
          <el-input placeholder="密码" v-model.trim="form.password" type="password"></el-input>
        </el-form-item>
        <el-form-item prop="vcode" class="vcode">
          <el-input placeholder="验证码" v-model.trim="form.vcode"></el-input>
          <img :src="vcodeImage" alt="" class="v-code-img" @click="verifycode" />
        </el-form-item>
        <el-form-item class="login-button">
          <el-button type="primary" @click="login" :loading="loginLoading">登录</el-button>
        </el-form-item>
        <div class="add-box">
          <el-button class="add-btn" type="text" @click="add">注册</el-button>
        </div>
      </el-form>
    </div>
    <create-popup
      :show="popupShow"
      :popup-type="popupType"
      :action="createAction"
      @close="popupShow = false"
    />
  </div>
</template>
<script>
import { verifycode, login } from "@/api/login";
import CreatePopup from "@/components/CreatePopup";

export default {
  name: "LoginLayout",
  components: { CreatePopup },

  data() {
    return {
      form: {
        user_id: "",
        password: "",
        captid: "",
        vcode: ""
      },
      loginLoading: false,
      vcodeImage: "",
      //表单验证规则
      formRules: {
        user_id: [
          { required: true, message: "请输入用户名", trigger: "change" },
          {
            min: 3,
            max: 20,
            message: "长度在 3 到 20 个字符",
            trigger: "change"
          }
        ],
        password: [{ required: true, message: "请输入密码", trigger: "change" }],
        vcode: [{ required: true, message: "请输入验证码", trigger: "change" }]
      },
      initialRoute: "/crm/customer",
      popupShow: false,
      popupType: "CreateUser",
      createAction: {
        type: "add",
        id: "",
        data: {}
      }
    };
  },
  mounted() {
    this.verifycode();
  },
  methods: {
    //验证码
    verifycode() {
      verifycode()
        .then(res => {
          if (res && res.code === 200) {
            this.form.captid = res.data.captid;
            this.vcodeImage = res.data.image;
          } else {
            this.$message.error(res.error);
          }
        })
        .catch(err => {
          console.log(err);
        });
    },
    //登录
    login() {
      this.$refs.loginFormRef.validate(async valid => {
        if (!valid) return;
        this.loginLoading = true;
        login(this.form)
          .then(res => {
            if (res && res.code === 200) {
              let token = res.data.authorization;
              localStorage.setItem("crmToken", token);
              localStorage.setItem(
                "crmPermission",
                JSON.stringify({
                  loginUser: res.data
                })
              );
              this.$store.commit("set_userInfo", res.data);
              this.$message.success("登录成功");
              this.$router.push("/crm/customer");
            } else {
              this.$message.warning(res.message);
            }
          })
          .catch(() => {})
          .finally(() => {
            this.loginLoading = false;
            this.verifycode();
          });
      });
    },
    // 注册
    add() {
      this.createAction = this.$options.data().createAction;
      this.popupShow = true;
    }
  }
};
</script>
<style lang="less" scoped>
.login_container {
  height: 100%;
  background: url(../../assets/images/login_bg.png) no-repeat;
  background-size: cover;
  .login_box {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 640px;
    height: 466px;
    box-sizing: border-box;
    padding: 52px 60px 60px !important;
    background-color: #fff;

    .logo-row {
      .logo-col {
        margin-top: 30px;
        font-size: 28px;
        font-weight: 400;
        display: flex;
        flex-direction: column;
        align-items: center;

        .logo-img {
          width: 144px;
          height: 107px;
          margin-bottom: 24px;
        }
      }
    }
    .vcode {
      .el-form-item__content {
        display: flex;
        .el-input {
          flex: 1;
          margin-right: 12px;
        }
      }
      .v-code-img {
        height: 40px;
      }
    }

    .login-button {
      margin-bottom: 0;
      .el-button {
        width: 100%;
        height: 38px;
      }
    }
    .add-box {
      text-align: right;
    }
    .add-btn {
      font-size: 12px;
    }

    :deep(.el-input__inner) {
      height: 40px;
    }

    :deep(.el-form-item__error) {
      padding-top: 4px;
    }

    :deep(.login_form) {
      text-align: center;
      position: absolute;
      bottom: 82px;
      left: 50%;
      transform: translateX(-50%);
      width: 328px;
      .login-email {
        margin-bottom: 20px;
      }
      .email-input {
        .el-input__suffix-inner {
          color: rgba(0, 0, 0, 0.3);
        }

        .el-input-group__append {
          background-color: #ffffff;
          border-color: rgba(0, 0, 0, 0.3);
        }
      }
    }
  }
}
</style>
