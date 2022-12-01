/*
 * @Author: Leiyi548
 * @Date: 2022-12-01 14:19:13
 * @LastEditTime: 2022-12-01 20:53:41
 * @LastEditors: Leiyi548
 * @Description:
 * @FilePath: \typescript\src\ts\1.基础类型.ts
 * 任何一个傻瓜都会写能够让机器理解的代码，只有好的程序员才能写出人类可以理解的代码。——Martin Fowler
 * 如在使用中遇到问题，可以联系我，QQ: 1424630446 交流
 */
{
  console.log('===数字字符串===');
  let isDone: boolean = false;
  console.log(isDone); // fasle

  let decLiteral: number = 6;
  // 16 进制
  let hexLiteral: number = 0xf00d;
  console.log(decLiteral); // 6
  console.log(hexLiteral); // 61453

  let name: string = 'bob';
  name = 'smith';
  console.log(name); // smith

  let template_name: string = `Gene`;
  let age: number = 37;
  let sentence_name_template: string = `Hello, my name is ${template_name}`;
  let sentence_age_template: string = `How old are you? ${age}`;
  console.log(sentence_name_template);
  console.log(sentence_age_template);
  // 也可以直接字符串拼接
  console.log(template_name + age);
}

{
  console.log('===数组===');
  // 数组
  let list: number[] = [1, 2, 3];
  console.log(list); // [ 1, 2, 3]

  // 或者我们可以使用数组泛型，Array<元素类型>
  let list_number: Array<number> = [1, 2, 3];
  console.log(list_number);

  // 元组
  // 元组类型允许表示一个已知元素数量和类型的数组，各元素的类型不必相同。比如，你可以定义一对值分别为 string 和 number 类型的元组
  // 声明一个元组类型
  let x: [string, number];
  // 初始化
  x = ['hello', 10];
  console.log(x); // [ 'hello', 10 ]

  // 当然我们要弄对顺序，如果顺序是错的话，那就没有意义了。
  // x = [10, 'hello'];

  // 如果我们访问的是一个越界的元素，会使用联合类型替代：
  // 教程上说可以，但是不知道为什么我不行，不能将类型""world""分配给类型"undefined"
  // x[3] = 'world';
}

{
  console.log('===枚举===');
  // 枚举
  // 默认是从 0 开始
  enum Color {
    Red,
    Green,
    Blue,
  }
  let c: Color = Color.Green;
  console.log(c); // 1

  // 可以修改成从 1 开始
  enum Season {
    Spring = 1,
    Summer,
    Autumn,
    Winter,
  }
  let s: Season = Season.Autumn;
  console.log(s); // 3

  // 当然我们可以手动进行赋值
  enum Direction {
    Up = 3,
    Down = 8,
    Left = 5,
    Right = 4,
  }
  let d: Direction = Direction.Down; // 8
  console.log(d);
}

{
  // Any
  console.log('===Any===');
  // 有时候，我们会想要为那些编程阶段还不清楚类型的变量指定一个类型。这些值可能来自于动态的内容，比如来自用户输入或第三方代码库。
  // 这种情况下，我们不希望类型检查器对这些值进行检查而是直接让它们通过编译阶段的检查。
  // 那么我们可以使用 any 类型来标记这些变量：
  let notSure: any = 4;
  notSure = 'maybe a string instead';
  // 没有指定所以可以赋值给 string
  console.log(notSure); // maybe a string instead
  // 当然我们也可以把这个赋值给 number
  notSure = 10;
  console.log(notSure); // 10

  // 当你只知道一部分的类型时，any 类型也是有用的。比如，你有一个数组，它不包含了不同的类型的数据：
  let list: any[] = [1, true, 'free'];
  list[1] = 100;
  console.log(list[1]); // 100
  console.log(list[2]); // free
}

{
  // void
  // 某种类型程度上来说，void 类型像是于 any 类型相反，它表示没有任何类型，当一个函数没有返回值时，你通常会见到其返回值类型是 void
  console.log('===void===');
  function warnUser(): void {
    console.log('This is my warning messages');
  }
  warnUser();

  // 声明一个 void 类型的变量没有什么大用，因为你只能为它赋予 undefined 和 null
  let unusable_undefined: void = undefined;
  console.log(unusable_undefined); // undefined
  let unusable_null: void = null;
  console.log(unusable_null); // null
}

{
  // null 和 undefined
  // Typescript 里，undefined 和 null 两者各自有自己的类型叫做 undefined 和 null。和 void 相似，它们的本身的类型用处不是很大。
  console.log('null 和 undefined');

  // 默认情况下 null 和 undefined 是所有类型的子类型。就是说你可以把 null 和 undefined 赋值给 number 类型的变量。
  // 然而，当你指定了--strictNullChecks 标记，null 和 undefined 只能赋值给 void 和它们各自。这能避免 很多常见的问题。
  // 也许在某处你想传入一个 string 或 null 或 undefined，你可以使用联合类型 string | null | undefined。
}

{
  // never 类型表示的是那些永不存在的值的类型。例如，nerver 类型是那些总是会抛出异常或根本就不会有返回值的函数表达式或箭头函数表达
  // 式的返回值类型；变量也可能是 never 类型，当它们被永不为真的类型保护所约束时。
  // never 类型是任何类型的子类型，也可以赋值给任何类型；然而，没有类型是 never 的子类型或可以赋值给 never 类型（除了 never 本身
  // 之外）。即使 any 也不可以赋值给 never。

  // 返回 never 的函数必须存在无法达到的终点
  function error(message: string): never {
    throw new Error(message);
  }

  // 推断的返回值类型为 never
  function fail() {
    return error('Something failed');
  }

  // 返回 never 的函数必须存在无法达到的终点
  function infiniteLoop(): never {
    while (true) {}
  }
}

{
  // object
  // object 表示非原始类型，也就是除 number,string,boolean,symbol,null 或 undefined 之外的类型。
  // 使用 object 类型，就可以更好的表示像 Object.create 这样的 API。例如：
}

{
  // 类型断言
  console.log('===类型断言===');
  let someValue: any = 'this is a string';
  let strLength: number = (<string>someValue).length;
  console.log(someValue);
  console.log(strLength);
}
