/*
 * @Author: Leiyi548
 * @Date: 2022-12-02 19:43:53
 * @LastEditTime: 2022-12-02 19:46:35
 * @LastEditors: Leiyi548
 * @Description: 枚举学习
 * @FilePath: \typescript\src\ts\6.枚举.ts
 * 任何一个傻瓜都会写能够让机器理解的代码，只有好的程序员才能写出人类可以理解的代码。——Martin Fowler
 * 如在使用中遇到问题，可以联系我，QQ: 1424630446 交流
 */

enum Direction {
  Up = 1,
  Down,
  Left,
  Right,
}
// 如上，我们定义了一个数字枚举，UP 使用初始化为 1。其余的成员会从 1 开始自动增长。换句话，
// Direction.Up 的值为 1, Down 为 2，Left 为 3，Right 为 4。
