/*
 * @Author: Leiyi548
 * @Date: 2022-12-01 22:13:15
 * @LastEditTime: 2022-12-02 16:19:37
 * @LastEditors: Leiyi548
 * @Description: 类
 * @FilePath: \typescript\src\ts\3.类.ts
 * 任何一个傻瓜都会写能够让机器理解的代码，只有好的程序员才能写出人类可以理解的代码。——Martin Fowler
 * 如在使用中遇到问题，可以联系我，QQ: 1424630446 交流
 */

class Greeter {
  greeting: string;
  constructor(message: string) {
    this.greeting = message;
  }
  greet() {
    return 'Hello, ' + this.greeting;
  }
}

// 构造了一个 Greeter 类的实例
let greeting = new Greeter('world');

console.log(greeting.greet()); // Hello, world

// 类的继承
class Animal {
  // 在 Typescript 里，成员默认为 public
  // 当成员被标记为 private 时，它就不能在声明它的类的外部访问。
  // protected 修饰符与 private 修饰符的行为很相似，但有一点不同，protected 成员在派生类中仍然可以访问。
  // readonly 关键字将属性设置为只读的。只读属性必须声明或构造函数里被初始化。
  public name: string;
  public constructor(theName: string) {
    this.name = theName;
  }
  public move(distanceInMeters: number = 0) {
    console.log(`${this.name} moved ${distanceInMeters}`);
  }
}

class Snake extends Animal {
  constructor(name: string) {
    // 基类必须要调用下父类的构造函数
    super(name);
  }
  move(distanceInMeters: number = 5): void {
    console.log('Slithering');
    super.move(distanceInMeters);
  }
}

class Horse extends Animal {
  constructor(name: string) {
    super(name);
  }
  move(distanceInMeters: number = 45): void {
    console.log('Galloping...');
    super.move(distanceInMeters);
  }
}

let sam = new Snake('Sammy the Python');
let tom: Animal = new Horse('Tommy the Palomino');

sam.move();
tom.move(34);

// 存取器

// 这个变量在类内居然能读取到！
let passcode = 'secret passcode';
class Employee {
  private _fullname: string;

  public get fullName(): string {
    return this._fullname;
  }

  public set fullName(newName: string) {
    if (passcode && passcode == 'secret passcode') {
      this._fullname = newName;
    } else {
      console.log('Error: Unauthorized update of employee');
    }
  }
}

let employee = new Employee();
employee.fullName = 'Bob Smith';
if (employee.fullName) {
  console.log(employee.fullName);
}

// 静态属性
class Grid {
  // 静态成员，这些属性存在于类本身上面而不是类的实例上
  // 在这个例子里，我们使用 static 定义 origin，因为它是所有网格都会用到的属性。
  // 每个实例想要访问这个属性的时候，都要在 origin 前面加上类名。
  // 如同在实例属性上使用 this. 前缀来访问属性一样，这里我们使用 Grid. 来访问静态属性。
  static origin = { x: 0, y: 0 };
  public calculateDistanceFromOrigin(point: { x: number; y: number }) {
    let xDist = point.x - Grid.origin.x;
    let yDist = point.y - Grid.origin.y;
    return Math.sqrt(xDist * xDist + yDist * yDist) / this.scale;
  }
  constructor(public scale: number) {}
}

let grid1 = new Grid(1.0);
let grid2 = new Grid(5.0);

console.log(grid1.calculateDistanceFromOrigin({ x: 10, y: 10 }));
console.log(grid2.calculateDistanceFromOrigin({ x: 10, y: 10 }));

// 抽象类
// 抽象类作为其他派生类的基本使用。他们一般不会直接将实例化。
// 不同于接口，抽象类可以包含成员的实现细节。
abstract class AnimalAbstract {
  abstract makeSound(): void;
  move(): void {
    console.log('roaming the earch...');
  }
}

// 抽象类的抽象方法不包含具体实现并且必须在派生类中实现。抽象方法的语法与接口方法相似。
// 两者都是定义方法签名但不包含方法体。
// 然而，抽象方法必须包含 abstract 关键字并且可以包含访问修饰符

abstract class Department {
  constructor(public name: string) {}

  printName(): void {
    console.log('Department name: ' + this.name);
  }
  abstract printMeeting(): void;
}

class AccountingDepartment extends Department {
  constructor() {
    super('Accounting and Anditing');
  }
  printMeeting(): void {
    console.log('The Accounting Department meets each Monday at 10am.');
  }

  generateReports(): void {
    console.log('Generating accounting reports...');
  }
}

let department: Department; // 允许创建一个对抽象类的引用
// department = new Department(); // 错误；不能创建一个抽象类的实例
department = new AccountingDepartment(); // 允许对一个抽象子类进行实例化和赋值
department.printName();
department.printMeeting();
// 在抽象父类中没有 generateReports 方法
// department.generateReports();

// 构造函数
class GreeterConstructor {
  greeting: string;
  constructor(message: string) {
    this.greeting = message;
  }
  greet() {
    return 'Hello, ' + this.greeting;
  }
}

let greeterCons: GreeterConstructor;
greeterCons = new Greeter('world');
console.log(greeterCons.greet()); // Hello, world
