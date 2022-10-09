package com.learn.aop;

import org.springframework.aop.MethodBeforeAdvice;

import java.lang.reflect.Method;

/**
 * @author jojo
 * @date 2022/10/9 14:28
 */
public class Log implements MethodBeforeAdvice {

    /**
     *
     * @param method 要执行的目标对象的方法
     * @param args 被调用的方法的参数
     * @param target 目标对象
     */
    @Override
    public void before(Method method, Object[] args, Object target) throws Throwable {
        assert target != null;
        System.out.println(target.getClass().getName() + "的" + method.getName() + "方法执行了");
    }
}
