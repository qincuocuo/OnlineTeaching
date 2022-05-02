const user = {
  state: {
    userInfo: null, // 用户信息
    activePath: "/crm/customer"
  },
  mutations: {
    set_userInfo(state, value) {
      state.userInfo = value;
    },
    activePathChange(state, newVal) {
      state.activePath = newVal;
    }
  },

  actions: {}
};

export default user;
