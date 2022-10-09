package com.learn.aop;

/**
 * @author jojo
 * @date 2022/10/9 14:46
 */
public class CustomPointcut {
   public void before(){
       System.out.println("-------方法执行前-------");
   }

    public void after(){
        System.out.println("-------方法执行后-------");
    }
}
