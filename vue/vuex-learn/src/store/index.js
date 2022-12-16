import Vue from 'vue';
import Vuex from 'vuex';
import user from './modules/user';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    num: 0,
    list: [
      { id: 1, name: '111' },
      { id: 2, name: '222' },
      { id: 3, name: '333' },
    ],
  },
  getters: {
    getNum(state) {
      return state.num;
    },
  },
  mutations: {
    increase(state, payload) {
      state.num += payload ? payload : 1;
    },
    decrease(state, payload) {
      state.num -= payload ? payload : 1;
    },
  },
  // actions 专门用来处理异步，实际修改状态值的，依然是 mutations
  actions: {
    // 点击了“减 1”按钮后，放慢一秒再执行减法
    decreaseAsync(context) {
      context.commit('decrease', 5);
    },
  },
  modules: {
    user,
    countryModules: {
      state: {
        countryName: 'china',
      },
      getters: {
        getCountryName(state) {
          return state.countryName;
        },
      },
    },
  },
});
