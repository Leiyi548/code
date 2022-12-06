/*
 * @Author: Leiyi548
 * @Date: 2022-12-02 23:43:58
 * @LastEditTime: 2022-12-03 00:27:24
 * @LastEditors: Leiyi548
 * @Description: 高级类型学习
 * @FilePath: \typescript\src\ts\7.高级类型.ts
 * 任何一个傻瓜都会写能够让机器理解的代码，只有好的程序员才能写出人类可以理解的代码。——Martin Fowler
 * 如在使用中遇到问题，可以联系我，QQ: 1424630446 交流
 */

// 联合类型

// 如果我们设置 padding 为 any 的话，我们传入的参数不是 number 或者 string 就会丢出错误
function padLeft(value: string, padding: string | number) {
  if (typeof padding === 'number') {
    return Array(padding + 1).join(' ') + value;
  }
  if (typeof padding === 'string') {
    return padding + value;
  }
  throw new Error(`Expectd string or number,got ${padding}`);
}

console.log(padLeft('Hello world', 4));
// padLeft('Hello world', true);
