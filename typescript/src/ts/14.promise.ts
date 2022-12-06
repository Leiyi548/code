// Promise 如果翻译成中文的话那就是承诺
// Promise 实例生成以后，可以用 then 方法分别指定 resolved 状态和 rejected 状态的回调函数

// 关键字翻译
// resolved → 解决  rejected → 拒绝

let promise01 = new Promise(function (resolved, rejected) {
  console.log('Promise01 create');
  resolved('promise01 ok');
  rejected('promise01 error');
});

let promise02 = new Promise(function (resolved, rejected) {
  console.log('Promise02 create');
  resolved('promise02 ok');
  rejected('promise02 error');
});

promise01
  .then(function (res) {
    console.log('resolved ' + res); // resolved ok
  })
  .catch((err) => {
    console.log(err);
  });

// Promise.all 可以将多个 Promise 实例包装成一个新的 Promise 实例。同时，成功和失败的返回值是不同的，
// 成功的时候返回的是一个结果数组，而失败的时候则返回最先返回 reject 失败状态的值
Promise.all([promise01, promise02])
  .then((res) => {
    console.log(res); // ['promise ok','promise 02']
  })
  .catch((error) => {
    console.log(error);
  });

let promise03 = new Promise((resolved, rejected) => {
  console.log('Promise03 create');
  setTimeout(() => {
    resolved('success');
  }, 1000);
});

let promise04 = new Promise((resolved, rejected) => {
  console.log('Promise04 create');
  setTimeout(() => {
    rejected('failed');
  }, 500);
});

// 顾名思义，Promise.race 就是赛跑的意思，意思就是说，Promise.race([p1,p2,p3]) 里面哪个结果获得的快，
// 就返回哪个结果，不管结果本身是成功状态还是失败状态。
Promise.race([promise03, promise04])
  .then((res) => {
    console.log(res);
  })
  .catch((error) => {
    console.log(error);
  });

Promise.race([promise01, promise02])
  .then((res) => {
    console.log(res);
  })
  .catch((error) => {
    console.log(error);
  });

console.log('Hi');
