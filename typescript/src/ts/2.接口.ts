/*
 * @Author: Leiyi548
 * @Date: 2022-12-01 20:54:30
 * @LastEditTime: 2022-12-01 22:12:42
 * @LastEditors: Leiyi548
 * @Description: 变量声明
 * @FilePath: \typescript\src\ts\2.接口.ts
 * 任何一个傻瓜都会写能够让机器理解的代码，只有好的程序员才能写出人类可以理解的代码。——Martin Fowler
 * 如在使用中遇到问题，可以联系我，QQ: 1424630446 交流
 */

function printLabel(labelledObj: { label: string }) {
  console.log(labelledObj.label);
}

let myObj = { size: 10, label: 'Size 10 Object' };
printLabel(myObj);

// 上面例子我们用接口来进行重写
interface LabelledValue {
  label: string;
}

function printLabelInterface(labelledObj: LabelledValue) {
  console.log(labelledObj.label);
}

printLabel(myObj);
// 上面两个结果都是一样的。

// 可选属性
// 接口里面的属性不全都是必需的。有些只在某些条件下存在，或者根本不存在。可选属性在应用“option bags"模式时很常用，即给函数传入的参
// 数对象中只有部分属性赋值了。
interface SquareConfig {
  color?: string;
  width?: number;
}

function createSquare(config: SquareConfig) {
  let newSquare = { color: 'white', area: 100 };
  if (config.color) {
    newSquare.color = config.color;
  }
  if (config.width) {
    newSquare.area = config.width * config.width;
  }
  return newSquare;
}

let mySquare01 = createSquare({ color: 'black' });
let mySquare02 = createSquare({ width: 50 });
let mySquare03 = createSquare({});
console.log(mySquare01); // { color: 'black', area: 100}
console.log(mySquare02); // { color: 'white', area: 2500}
console.log(mySquare03); // { color: 'white', area: 100}

// 只读属性
interface Point {
  readonly x: number;
  readonly y: number;
}

let p1: Point = { x: 10, y: 20 };
// 无法分配到“x“，因为它是只读属性。
// p1.x = 120;

// 额外的属性检查

interface SquareConfigExtra {
  color?: string;
  width?: number;
  [propName: string]: any;
}
let mySquareExtra01: SquareConfigExtra;
mySquareExtra01 = { color: 'red', hello: 'easy' };
console.log(mySquareExtra01.hello); // easy
console.log(mySquareExtra01.color); // red

// 类接口
interface ClockInterface {
  currentTime: Date;
  setTime(d: Date);
}

// 接口描述了类的公共部分，而不是公共和私有两部分。它不会帮你检查类是否有某些私有成员。
class Clock implements ClockInterface {
  currentTime: Date;
  setTime(d: Date) {
    this.currentTime = d;
  }
  constructor(h: number, m: number) {}
}

// 接口继承
interface Shape {
  color: string;
}

interface Square extends Shape {
  sideLength: number;
}

// 继承了我就可以直接使用子接口的属性
let square = <Square>{};
square.color = 'blue';
square.sideLength = 10;
console.log(square.color); // blue
console.log(square.sideLength); // 10

// 一个接口可以继承多个接口，创建多个接口的合成接口。

interface PenStroke {
  penWidth: number;
}

interface Square extends Shape, PenStroke {
  sideLength: number;
}

let square02 = <Square>{};
square02.color = 'blue';
square02.sideLength = 10;
square02.penWidth = 5.0;
console.log(square02.penWidth); // 5

// 混合类型
interface Counter {
  (start: number): string;
  interval: number;
  reset(): void;
}

function getCounter(): Counter {
  let counter = <Counter>function (start: number) {};
  counter.interval = 123;
  counter.reset = function () {};
  return counter;
}

// 接口继承类
class Control {
  private state: any;
}

interface SelectableControl extends Control {
  select(): void;
}

class Button extends Control implements SelectableControl {
  // ovride function
  select(): void {}
}

class TextBox extends Control {
  select(): void {}
}
