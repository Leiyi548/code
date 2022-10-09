package org.learn.pojo;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

/**
 * @author jojo
 * @date 2022/10/9 12:22
 */
@Component("teacher01")
// 相当于配置文件中 <bean id="teacher01" class="当前注解的类"/>
public class Teacher {
    // private String name = "ewell";

    // @Value("ewell")
    // 相当于配置文件中 <property name="name" value="ewell"/>
    private String name;

    // 或者我们可以在set方法上面来设置

    @Value("ewell")
    public void setName(String name) {
        this.name = name;
    }

    public void show() {
        System.out.println("teacher name is " + name);
    }
}
