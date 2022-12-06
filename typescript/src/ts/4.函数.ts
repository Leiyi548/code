/*
 * @Author: Leiyi548
 * @Date: 2022-12-02 16:20:43
 * @LastEditTime: 2022-12-02 17:15:24
 * @LastEditors: Leiyi548
 * @Description: 函数学习
 * @FilePath: \typescript\src\ts\4.函数.ts
 * 任何一个傻瓜都会写能够让机器理解的代码，只有好的程序员才能写出人类可以理解的代码。——Martin Fowler
 * 如在使用中遇到问题，可以联系我，QQ: 1424630446 交流
 */

// 和 javascript 一样，Typescript 函数可以创建有名字和匿名函数。你可以随意选择适合应用程序的方式，不论是定义一系列 API 函数还是
// 只使用一次的函数

// Named function（带有名字的函数）
function add(x: number, y: number): number {
  return x + y;
}

// Anonoymous function（匿名函数）
let myAdd = function (x: number, y: number) {
  return x + y;
};

console.log(add(1, 3)); // 4
console.log(myAdd(1, 3)); // 4

// 可选参数和默认参数
// Typescript 里的每个函数参数都是必须的。这不是指不能传递 null 或 undefined 作为参数，而是说编译器检查用户是否会每个参数都传入值。
// 编辑器还会假设只有这些参数会被传递进函数。简短的说，传递给一个函数的参数个数必须与函数期望的参数个数不一致。
function buildName(firstName: string, lastName: string) {
  return firstName + ' ' + lastName;
}

// let result1 = buildName('Bob'); // error，需要两个参数，参数太少了
// let result2 = buildName('Bob', 'Adams', 'Sr.'); // error，需要两个参数，但是有 3 个，太多了
let result3 = buildName('Bob', 'Adams');
console.log(result3); // Bob Adams

// Javascript 里，每个参数都是可选的，可传可不传。没传参的时候，它的值就是 undefined。在 Typescript 里我们可以在参数名旁使用 ?
// 实现可选参数的功能。比如，我们想让 lastName 可选：

function buildName02(firstName: string, lastName?: string) {
  if (lastName) {
    return firstName + ' ' + lastName;
  } else {
    return firstName;
  }
}

let resultAnonoymous01 = buildName02('yeezy'); // 刚好就只有 firstname
let resultAnonoymous02 = buildName02('yeezy', 'nike'); // 就是原来的效果

// 默认参数
function buildNameDefault(firstName: string = 'will', lastName: string) {
  return firstName + ' ' + lastName;
}

// let resultDefault01 = buildNameDefault('Bob'); // error, 需要两个参数，但只获得了一个
// let resultDefault02 = buildNameDefault('Bob', 'Adams', 'Br.'); // error, 太多参数了！
let resultDefault03 = buildNameDefault('Bob', 'Adams'); // okay and returns, "Bob Adams"
let resultDefault04 = buildNameDefault(undefined, 'Adams'); // okay and returns, “Will Adams"

// 剩余参数
// 必要参数，默认参数和可选参数有个共同点：它们表示某一个参数。有时，你想同时操作多个参数，或者你并不知道会有多少参数传递进来。
// 在 Javascript 里，你可以使用 arguments 来访问所有传入的参数。

// 在 Typescript 里，你可以把所有参数收集到一个变量里：
// 剩余参数会被当做个数不限的可选参数。可以一个都没有，同样也可以有任意一个。编辑器创建参数数组，名字是你在省略号 (...) 后面给定的
// 名字，你可以在函数体内使用这个数组。
function buildNameLeft(firstName: string, ...restOfName: string[]) {
  return firstName + ' ' + restOfName.join(' ');
}

let employeeName = buildNameLeft('Joseph', 'Samuel', 'Lucas', 'Mackinzie');

// this
// 学习如何在 javascript 里正确使用 this 就好比一场成年礼。由于 TypeScript 是 JavaScript 的超集，TypeScript 程序员也需要弄清
// this 工作机制并且当有 bug 的时候能够找出错误所在。幸运的是，TypeScript 能通知你错误地使用了 this 地方。如果你想了解 JavaScript
// 里的 this 是如何工作的。

let deck = {
  suits: ['hearts', 'spades', 'clubs', 'diamonds'],
  cards: Array(52),
  createCardPicker: function () {
    // NOTE: the line below is now an arrow function, allowing us to capture 'this' right here
    return () => {
      // Returns a pseudorandom number between 0 and 1. */
      let pickedCard = Math.floor(Math.random() * 52);
      let pickedSuit = Math.floor(pickedCard / 13);

      return { suit: this.suits[pickedSuit], card: pickedCard % 13 };
    };
  },
};

// error，会报错，因为里面的 this 被认为 windows 而不是 deck
let cardPicker = deck.createCardPicker();
let pickedCard = cardPicker();

console.log('card: ' + pickedCard.card + ' of ' + pickedCard.suit);
