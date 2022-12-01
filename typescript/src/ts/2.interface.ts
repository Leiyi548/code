/*
 * @Author: Leiyi548
 * @Date: 2022-11-30 23:30:57
 * @LastEditTime: 2022-11-30 23:37:43
 * @LastEditors: Leiyi548
 * @Description: 这是一个学习接口的文件
 * @FilePath: \typescript\src\ts\2.interface.ts
 * 任何一个傻瓜都会写能够让机器理解的代码，只有好的程序员才能写出人类可以理解的代码。——Martin Fowler
 * 如在使用中遇到问题，可以联系我，QQ: 1424630446 交流
 */

class Student {
  fullName: string;
  constructor(public firstName: string, public middleInitial, public lastName) {
    this.fullName = firstName + ' ' + middleInitial + ' ' + lastName;
  }
}

interface Person {
  firstName: string;
  lastName: string;
}

function greeter_interface(person: Person) {
  return 'Hello, ' + person.firstName + ' ' + person.lastName;
}

let user_interface = { firstName: 'Jane', lastName: 'User' };

console.log(user_interface);
console.log(greeter_interface(user_interface));

let red_student = new Student('Jane', 'M.', 'User');
console.log(greeter_interface(red_student));
