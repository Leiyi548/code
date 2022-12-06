// 根据 MDN 的描述：回调函数是作为参数传给另一个函数的函数，然后通过在外部函数内部调用该回调函数以完成某种操作。

// 让我用人话解释一下，回调函数是一个函数，将会在另一个函数完成执行后立即执行。
// 回调函数是一个作为参数传给另一个 JavaScript 函数的函数。这个回调函数会在传给的函数内部执行。

// 更多详情请看 https://juejin.cn/post/6844903987771097102

function getMessage(msg: string, callback: Function): void {
  setTimeout(() => {
    console.log(msg);
    callback();
  }, 2000);
}

function displayMessage() {
  console.log('Display message');
}

getMessage('Hi,there', displayMessage);
// 执行结果如下：
// Hi,there
// Display message
