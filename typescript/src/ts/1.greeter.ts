/*
 * @Author: Leiyi548
 * @Date: 2022-11-30 19:45:01
 * @LastEditTime: 2022-11-30 23:14:57
 * @LastEditors: Leiyi548
 * @Description:
 * @FilePath: \typescript\src\1.greeter.ts
 * 任何一个傻瓜都会写能够让机器理解的代码，只有好的程序员才能写出人类可以理解的代码。——Martin Fowler
 * 如在使用中遇到问题，可以联系我，QQ: 1424630446 交流
 */

function greeter(person: string) {
  return 'Hello, ' + person;
}

let user = 'Jane User';

console.log(greeter(user));
document.body.innerHTML = greeter(user);
