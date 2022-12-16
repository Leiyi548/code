<template>
  <div>
    <h2>第一种写法</h2>
    <button @click="$store.state.num++">count 页面上的数字：{{ num }}</button>
    <h2>第二种写法</h2>
    <button @click="$store.state.num++">
      count 页面上的数字：{{ $store.getters.getNum }}
    </button>
    <h2>使用 $store.commit</h2>
    <button @click="$store.commit('increase', 2)">
      count 页面上的数字（加）：{{ $store.getters.getNum }}
    </button>
    <br />
    <br />
    <button @click="$store.commit('decrease', 2)">
      count 页面上的数字（减）：{{ $store.getters.getNum }}
    </button>
    <h2>延迟一秒减少</h2>
    <button @click="$store.dispatch('decreaseAsync')">点击延迟减 1</button>

    <h2>简单的写法</h2>
    <p>使用 mapState</p>
    <p>{{ num }}</p>
    <p>使用 mapGetters</p>
    <p>{{ getNum }}</p>
    <p>使用 mapMutations</p>
    <button @click="decrease(1)">{{ getNum }}</button>
    <hr />

    <h2>模块的导入</h2>
    <h4>直接在 index.js 中写入</h4>
    <p>{{ $store.state.countryModules.countryName }}</p>
    <h4>分模快导入</h4>
    <p>userName: {{ $store.state.user.name }}</p>
    <p>userAge: {{ $store.state.user.age }}</p>
    <p>countryName: {{ $store.getters.getCountryName }}</p>
  </div>
</template>

<script>
import { mapState, mapGetters, mapMutations } from "vuex";

export default {
  computed: {
    // state,getters 这两个都是解构到 computed 里面
    ...mapState(["num"]),
    ...mapGetters(["getNum"]),
    ...mapGetters(["list"]),
  },
  methods: {
    ...mapMutations(["decrease"]),
  },
};
</script>

<style>
</style>