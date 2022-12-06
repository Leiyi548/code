/*
 * @Author: Leiyi548
 * @Date: 2022-12-02 17:17:02
 * @LastEditTime: 2022-12-02 19:43:11
 * @LastEditors: Leiyi548
 * @Description: 泛型学习
 * @FilePath: \typescript\src\ts\5.泛型.ts
 * 任何一个傻瓜都会写能够让机器理解的代码，只有好的程序员才能写出人类可以理解的代码。——Martin Fowler
 * 如在使用中遇到问题，可以联系我，QQ: 1424630446 交流
 */

// 泛型函数
function identity<T>(arg: T): T {
  return arg;
}

// 自己指定
let output = identity<string>('myString');
// 系统判断
// 这一种更加普遍
let output2 = identity(123);
console.log(output); // myString
console.log(output2); // 123

// 泛型类型

interface GenericIdentityFn {
  <T>(arg: T): T;
}
