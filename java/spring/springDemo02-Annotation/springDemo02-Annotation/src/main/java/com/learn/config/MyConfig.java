package com.learn.config;

import com.learn.pojo.Dog;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Import;

/**
 * @author jojo
 * @date 2022/10/9 13:57
 */
//将这个类标注为Spring的一个组件，放到容器中！
@Configuration
// 导入其它配置
@Import(MyConfig2.class)
public class MyConfig {
    @Bean // 通过方法注册一个bean，这里的返回值就Bean的类型，方法名就是bean的id！
    public Dog dog(){
        return new Dog();
    }
}
