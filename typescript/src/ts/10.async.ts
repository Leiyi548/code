// 也是用来处理异步的，其实是 Generator 函数的改进，背后原理就是 promise

async function f1() {
  return 'abc';
  // 等价于
  // return new Promise((resolved) => {
  // resolved('abc');
  // });
}

console.log(f1()); // Promise { 'abc' }

async function f3() {
  return new Promise((resolved, rejected) => {
    setTimeout(() => {
      resolved('abc');
    }, 2000);
  });
}

async function f4() {
  var c = await f3();
  console.log('awaiting ...');
  console.log(c);
}

f4();
