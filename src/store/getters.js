/**
 * 使用说明
 * import { mapGetters } from 'vuex'
 * computed: {
    ...mapGetters([
      'userInfo'
    ])
  }
 */
const permission = JSON.parse(window.localStorage.getItem("crmPermission")) || {};
const userInfo = permission.loginUser || null;
const getters = {
  userInfo: state => state.user.userInfo || userInfo
};

export default getters;
