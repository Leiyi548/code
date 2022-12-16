import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default {
  state: {
    name: 'leiyi',
    age: 18,
    gender: 'women',
  },
  getters: {
    getUserName(state) {
      return state.name;
    },
    getUserAge(state) {
      return state.age;
    },
    getUserGender(state) {
      return state.gender;
    },
  },
  mutations: {},
};
