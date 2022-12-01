function printLabel(labelledObj) {
    console.log(labelledObj.label);
}
let myObj = { size: 10, label: 'Size 10 Object' };
printLabel(myObj);
function printLabelInterface(labelledObj) {
    console.log(labelledObj.label);
}
printLabel(myObj);
function createSquare(config) {
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
console.log(mySquare01);
console.log(mySquare02);
console.log(mySquare03);
let p1 = { x: 10, y: 20 };
let mySquareExtra01;
mySquareExtra01 = { color: 'red', hello: 'easy' };
console.log(mySquareExtra01.hello);
console.log(mySquareExtra01.color);
class Clock {
    setTime(d) {
        this.currentTime = d;
    }
    constructor(h, m) { }
}
let square = {};
square.color = 'blue';
square.sideLength = 10;
console.log(square.color);
console.log(square.sideLength);
let square02 = {};
square02.color = 'blue';
square02.sideLength = 10;
square02.penWidth = 5.0;
console.log(square02.penWidth);
function getCounter() {
    let counter = function (start) { };
    counter.interval = 123;
    counter.reset = function () { };
    return counter;
}
class Control {
}
class Button extends Control {
    select() { }
}
class TextBox extends Control {
    select() { }
}
