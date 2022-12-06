/*
 * @Author: Leiyi548
 * @Date: 2022-12-03 00:44:01
 * @LastEditTime: 2022-12-03 00:53:08
 * @LastEditors: Leiyi548
 * @Description: 跌代器生成器学习
 * @FilePath: \typescript\src\ts\8.迭代器和生成器.ts
 * 任何一个傻瓜都会写能够让机器理解的代码，只有好的程序员才能写出人类可以理解的代码。——Martin Fowler
 * 如在使用中遇到问题，可以联系我，QQ: 1424630446 交流
 */

let someArray = [1, 'string', false];

for (const entry of someArray) {
  console.log(entry);
}
// 1
// string
// false

for (let i in someArray) {
  console.log(i);
}
// 0
// 1
// 2
