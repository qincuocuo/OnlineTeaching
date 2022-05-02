const common = {
  state: {
    provincesList: [],
    industryList: [],
    publicKey: "" // RSA公钥
  },

  mutations: {
    setState(state, { key, value }) {
      if (!key) return;
      state[key] = value;
    }
  },

  actions: {}
};

export default common;
