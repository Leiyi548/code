## Vuex

Vuex 是一个专为 Vue.js 应用程序开发的**状态管理模式**。

调试工具：vue devtools

> Vuex 就像眼镜：您自会知道什么时候需要它。

### 1、state

在 store 中定义数据，在组件中直接使用：

目录：`store/index.js`

```js
export default new Vuex.Store({
  // state 相当于组件中的 data，专门用来存放全局的数据
  state: {
    num: 0,
  },
  getters: {},
  mutations: {},
  actions: {},
  modules: {},
});
```

目录：`Home.vue`

```html
<template>
  <div class="home">
    <h2>Home 页面的数字：{{$store.state.num}}</h2>
  </div>
</template>

<script>
  export default {};
</script>
```

或者写为：

```html
<template>
  <div class="about">
    <h2>About 页面的数字：{{num}}</h2>
  </div>
</template>

<script>
  export default {
    computed: {
      num() {
        return this.$store.state.num;
      },
    },
  };
</script>
```

### 2、getters

将组件中统一使用的 computed 都放到 getters 里面来操作

目录：`store/index.js`

```js
export default new Vuex.Store({
  // state 相当于组件中的 data，专门用来存放全局的数据
  state: {
    num: 0,
  },
  // getters 相当于组件中的 computed，getters 是全局的，computed 是组件内部使用的
  getters: {
    getNum(state) {
      return state.num;
    },
  },
  mutations: {},
  actions: {},
  modules: {},
});
```

目录：`Home.vue`

```html
<template>
  <div class="home">
    <h2>Home 页面的数字：{{$store.getters.getNum}}</h2>
  </div>
</template>

<script>
  export default {};
</script>
```

### 3、mutations

更改 Vuex 的 store 中的状态的**唯一方法**是提交 mutation。

目录：`store/index.js`

```js
export default new Vuex.Store({
  // state 相当于组件中的 data，专门用来存放全局的数据
  state: {
    num: 0,
  },
  // getters 相当于组件中的 computed，getters 是全局的，computed 是组件内部使用的
  getters: {
    getNum(state) {
      return state.num;
    },
  },
  // mutations 相当于组件中的 methods，但是它不能使用异步方法（定时器、axios）
  mutations: {
    // 让 num 累加
    // payload 是一个形参，如果组件在 commit 时，有传这个参数过来，就存在，如果没有传过来，就是 undefined
    increase(state, payload) {
      state.num += payload ? payload : 1;
    },
  },
  actions: {},
  modules: {},
});
```

目录：`Btn.vue`

```html
<template>
  <div>
    <button @click="$store.commit('increase', 2)">点击加 1</button>
  </div>
</template>
<script>
  export default {
    methods: {
      /* addFn(){
           // 调用 store 中的 mutations 里的 increase 方法
           // 传参的话，使用 payload
           this.$store.commit('increase', 2)
       } */
    },
  };
</script>
```

### 4、actions

actions 是 store 中专门用来处理异步的，实际修改状态值的，还是 mutations

目录：`store/index.js`

```js
// 在 store(仓库) 下的 index.js 这份文件，就是用来做状态管理
import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  // state 相当于组件中的 data，专门用来存放全局的数据
  state: {
    num: 0,
  },
  // getters 相当于组件中的 computed，getters 是全局的，computed 是组件内部使用的
  getters: {
    getNum(state) {
      return state.num;
    },
  },
  // mutations 相当于组件中的 methods，但是它不能使用异步方法（定时器、axios）
  mutations: {
    // 让 num 累加
    // payload 是一个形参，如果组件在 commit 时，有传这个参数过来，就存在，如果没有传过来，就是 undefined
    increase(state, payload) {
      state.num += payload ? payload : 1;
    },
    // 让 num 累减
    decrease(state) {
      state.num--;
    },
  },
  // actions 专门用来处理异步，实际修改状态值的，依然是 mutations
  actions: {
    // 点击了“减 1”按钮后，放慢一秒再执行减法
    decreaseAsync(context) {
      context.commit('decrease');
    },
  },
  modules: {},
});
```

目录：`Btn.vue`

```html
<template>
  <div>
    <button @click="$store.commit('increase', 2)">点击加 1</button>
    <button @click="$store.dispatch('decreaseAsync')">点击延迟减 1</button>
  </div>
</template>
<script>
  export default {
    methods: {
      /* addFn(){
           // 调用 store 中的 mutations 里的 increase 方法
           // 传参的话，使用 payload
           this.$store.commit('increase', 2)
       }
       reduceFn(){
           this.$store.dispatch('decreaseAsync')
       } */
    },
  };
</script>
```

### 5、辅助函数

mapState 和 mapGetters 在组件中都是写在 computed 里面

```html
<template>
  <div>
    <h2>Home 页面的数字：{{num}}</h2>
    <h2>About 页面的数字：{{getNum}}</h2>
  </div>
</template>

<script>
  import { mapState, mapGetters } from 'vuex'

  export default {
    computed: {
      ...mapState(['num'])
      ...mapGetters(['getNum'])
    }
  }
</script>
```

mapMutations 和 mapActions 在组件中都是写在 methods 里面

```html
<template>
  <div>
    <button @click="increase(2)">点击加 1</button>
    <button @click="decreaseAsync()">点击延迟减 1</button>
  </div>
</template>

<script>
  import { mapMutations, mapActions } from 'vuex';

  export default {
    methods: {
      ...mapMutations(['increase']),
      ...mapActions(['decreaseAsync']),
    },
  };
</script>
```

### 6、拆分写法

store 中的所有属性，都可以拆分成单独的 js 文件来书写

### 7、modules

![](.\1111.png)我们的 store 可以认为是一个主模块，它下边可以分解为很多子模块，子模块都可以单独领出来写，写完再导入到主模块中。下面以 `users` 子模块举例：

将 mutations 中所有的方法，归纳起来。

目录：`mutations_type.js`

```js
export const MUTATIONS_TYPE = {
  INCREASE: 'increase',
  DECREASE: 'decrease',
};

export default {
  // 让 num 累加
  // payload 是一个形参，如果组件在 commit 时，有传这个参数过来，就存在，如果没有传过来，就是 undefined
  [MUTATIONS_TYPE.INCREASE](state, payload) {
    state.num += payload ? payload : 1;
  },
  // 让 num 累减
  [MUTATIONS_TYPE.DECREASE](state) {
    state.num--;
  },
};
```

目录：`store/index.js`

```js
import mutations from './mutaions_type'

export default new Vuex.Store({
    ...
    mutations,
    ...
})
```

组件中：

```html
<template>
  <div class="about">
    <h2>About 页面的数字：{{getNum}}</h2>
    <button @click="increase()">About 的按钮，点击加 1</button>
  </div>
</template>
<script>
  import { mapGetters, mapMutations } from 'vuex';
  import { MUTATIONS_TYPE } from '@/store/mutaions_type.js';
  export default {
    computed: {
      ...mapGetters(['getNum']),
    },
    methods: {
      // 方法一：
      ...mapMutations([MUTATIONS_TYPE.INCREASE]),

      // 方法二：
      /* increase(){
      this.$store.commit(MUTATIONS_TYPE.INCREASE)
    } */
    },
  };
</script>
```

## 参考文档

- [手把手教你使用 Vuex，猴子都能看懂的教程](https://juejin.cn/post/6928468842377117709)
- [vuex API 参考](https://vuex.vuejs.org/zh/api/)
