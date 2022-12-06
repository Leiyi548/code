// setTimeout 为延时执行任务
// function setTimeout(handler: TimerHandler, timeout?: number, ...arguments: any[]): number
// 返回值是一个 number
// 延迟 3 s 说你好
const num: number = setTimeout(() => {
  console.log('你好');
}, 3000);

// 虽然这里的类型注解写的是 number，但是返回值出来却不是 number，而是一大堆数据如下：
// Timeout {
//   _idleTimeout: 3000,
//   _idlePrev: [TimersList],
//   _idleNext: [TimersList],
//   _idleStart: 806,
//   _onTimeout: [Function (anonymous)],
//   _timerArgs: undefined,
//   _repeat: null,
//   _destroyed: false,
//   [Symbol(refed)]: true,
//   [Symbol(kHasPrimitive)]: false,
//   [Symbol(asyncId)]: 3,
//   [Symbol(triggerId)]: 1
// }
console.log(num);
