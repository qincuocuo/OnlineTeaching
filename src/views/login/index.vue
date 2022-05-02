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
        <el-form-item prop="loginEmail" class="login-email">
          <el-input placeholder="请输入邮箱地址" v-model.trim="form.loginEmail" class="email-input">
            <template #append>{{ mailboxSuffix }}</template>
          </el-input>
        </el-form-item>
        <el-form-item prop="password" class="password">
          <el-input
            placeholder="请输入口令"
            v-model.trim="form.password"
            type="password"
          ></el-input>
        </el-form-item>
        <el-row class="applyForPassword">
          <el-col>
            <el-button
              type="text"
              class="password-countdown"
              v-if="passwordCountdown === 30"
              :loading="applyLoading"
              @click="applyForPassword"
            >
              申请口令
            </el-button>
            <el-button class="prompt-information" type="text" v-else>
              {{ passwordCountdown }}秒后重新获取
            </el-button>
          </el-col>
        </el-row>
        <el-form-item class="login-button">
          <el-button type="primary" @click="login" :loading="loginLoading">登录</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<script>
import { applyForPassword, login } from "@/api/login";
export default {
  name: "LoginLayout",
  data() {
    return {
      form: {
        loginEmail: "",
        password: ""
      },
      mailboxSuffix: "@mycaiwen.com", //邮箱后缀
      passwordCountdown: 30, //口令倒计时
      passwordTimer: null,
      applyLoading: false,
      loginLoading: false,
      //表单验证规则
      formRules: {
        loginEmail: [
          { required: true, message: "请输入邮箱", trigger: "change" },
          {
            min: 3,
            max: 20,
            message: "长度在 3 到 20 个字符",
            trigger: "change"
          }
        ],
        password: [{ required: true, message: "请输入口令", trigger: "change" }]
      },
      initialRoute: "/crm/customer"
    };
  },
  methods: {
    //申请口令
    applyForPassword() {
      let that = this;
      this.$refs.loginFormRef.validateField("loginEmail", async valid => {
        if (valid) return false;
        this.applyLoading = true;
        let params = {
          email: this.form.loginEmail + this.mailboxSuffix
        };
        applyForPassword(params)
          .then(res => {
            if (res && res.code === 0) {
              this.$message.success("发送口令成功！");
              this.passwordTimer = setInterval(function () {
                if (that.passwordCountdown) {
                  that.passwordCountdown--;
                } else {
                  clearInterval(that.passwordTimer);
                  that.passwordTimer = null;
                  that.passwordCountdown = 30;
                }
              }, 1000);
            } else {
              this.$message.error(res.msg);
            }
          })
          .catch(err => {
            console.log(err);
          })
          .finally(() => {
            this.applyLoading = false;
          });
      });
    },
    //登录
    login() {
      this.$refs.loginFormRef.validate(async valid => {
        if (!valid) return;
        const initialRoute = this.initialRoute;
        this.loginLoading = true;
        this.applyLoading = true;
        let params = {
          email: this.form.loginEmail + this.mailboxSuffix,
          code: this.form.password
        };
        login(params)
          .then(res => {
            if (res && res.code === 0) {
              let permission = res.data;
              let token = res.data.token;
              localStorage.setItem("crmToken", token);
              localStorage.setItem("crmPermission", JSON.stringify(permission));
              this.$store.commit("set_userInfo", permission.loginUser);
              let as = permission.permission.routerWhiteList.find(function (item) {
                return item === initialRoute;
              });
              if (as) {
                this.$router.push(initialRoute);
                // window.sessionStorage.setItem("activePath", initialRoute);
              } else {
                this.$router.push(permission.permission.routerWhiteList[0]);
              }
              this.$message.success("登录成功");
            } else {
              this.$message.warning(res.msg);
            }
          })
          .catch(() => {})
          .finally(() => {
            this.loginLoading = false;
            this.applyLoading = false;
          });
      });
    }
  }
};
</script>
<style lang="less" scoped>
.login_container {
  height: 100%;
  // background: url(../../assets/images/login_bg.png) no-repeat;
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

    .password,
    .login-button {
      margin-bottom: 0;
    }

    .applyForPassword {
      text-align: right;
      margin-bottom: 22px;

      .el-button {
        padding-top: 8px;
        font-weight: 400;
      }

      .prompt-information {
        color: rgba(0, 0, 0, 0.7);
      }
    }

    .login-button {
      .el-button {
        width: 100%;
        height: 38px;
      }
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
