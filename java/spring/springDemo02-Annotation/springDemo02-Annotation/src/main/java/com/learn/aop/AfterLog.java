package com.learn.aop;

import org.springframework.aop.AfterReturningAdvice;

import java.lang.reflect.Method;

/**
 * @author jojo
 * @date 2022/10/9 14:31
 */
public class AfterLog implements AfterReturningAdvice {
    /**
     *
     * @param returnValue 返回值
     * @param method 被调用的方法
     * @param args 被调用的方法的对象的参数
     * @param target 被调用的目标对象
     */
    @Override
    public void afterReturning(Object returnValue, Method method, Object[] args, Object target) throws Throwable {
        assert target != null;
        System.out.println("执行了" + target.getClass().getName() + "的"
                + method.getName() + "方法，返回值：" + returnValue );
    }
}
