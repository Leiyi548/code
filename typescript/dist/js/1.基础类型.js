{
    console.log('===数字字符串===');
    let isDone = false;
    console.log(isDone);
    let decLiteral = 6;
    let hexLiteral = 0xf00d;
    console.log(decLiteral);
    console.log(hexLiteral);
    let name = 'bob';
    name = 'smith';
    console.log(name);
    let template_name = `Gene`;
    let age = 37;
    let sentence_name_template = `Hello, my name is ${template_name}`;
    let sentence_age_template = `How old are you? ${age}`;
    console.log(sentence_name_template);
    console.log(sentence_age_template);
    console.log(template_name + age);
}
{
    console.log('===数组===');
    let list = [1, 2, 3];
    console.log(list);
    let list_number = [1, 2, 3];
    console.log(list_number);
    let x;
    x = ['hello', 10];
    console.log(x);
}
{
    console.log('===枚举===');
    let Color;
    (function (Color) {
        Color[Color["Red"] = 0] = "Red";
        Color[Color["Green"] = 1] = "Green";
        Color[Color["Blue"] = 2] = "Blue";
    })(Color || (Color = {}));
    let c = Color.Green;
    console.log(c);
    let Season;
    (function (Season) {
        Season[Season["Spring"] = 1] = "Spring";
        Season[Season["Summer"] = 2] = "Summer";
        Season[Season["Autumn"] = 3] = "Autumn";
        Season[Season["Winter"] = 4] = "Winter";
    })(Season || (Season = {}));
    let s = Season.Autumn;
    console.log(s);
    let Direction;
    (function (Direction) {
        Direction[Direction["Up"] = 3] = "Up";
        Direction[Direction["Down"] = 8] = "Down";
        Direction[Direction["Left"] = 5] = "Left";
        Direction[Direction["Right"] = 4] = "Right";
    })(Direction || (Direction = {}));
    let d = Direction.Down;
    console.log(d);
}
{
    console.log('===Any===');
    let notSure = 4;
    notSure = 'maybe a string instead';
    console.log(notSure);
    notSure = 10;
    console.log(notSure);
    let list = [1, true, 'free'];
    list[1] = 100;
    console.log(list[1]);
    console.log(list[2]);
}
{
    console.log('===void===');
    function warnUser() {
        console.log('This is my warning messages');
    }
    warnUser();
    let unusable_undefined = undefined;
    console.log(unusable_undefined);
    let unusable_null = null;
    console.log(unusable_null);
}
{
    console.log('null 和 undefined');
}
{
    function error(message) {
        throw new Error(message);
    }
    function fail() {
        return error('Something failed');
    }
    function infiniteLoop() {
        while (true) { }
    }
}
{
}
{
    console.log('===类型断言===');
    let someValue = 'this is a string';
    let strLength = someValue.length;
    console.log(someValue);
    console.log(strLength);
}
