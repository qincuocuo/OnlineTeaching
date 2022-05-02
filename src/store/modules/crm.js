const crm = {
  state: {
    message: null // 消息
  },

  mutations: {
    setState(state, { key, value }) {
      if (!key) return;
      state[key] = value;
    }
  },

  actions: {}
};

export default crm;
