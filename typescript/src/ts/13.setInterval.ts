// 这个是循环执行的函数

// 每隔一秒数字加一并输出，在第 10s 退出

var count: number = 1;

const timer: number = setInterval(() => {
  console.log(count);
  if (count == 10) {
    clearInterval(timer);
  }
  count++;
}, 1000);
