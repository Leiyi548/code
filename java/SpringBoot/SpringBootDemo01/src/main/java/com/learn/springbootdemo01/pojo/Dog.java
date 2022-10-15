package com.learn.springbootdemo01.pojo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

/**
 * @author jojo
 * @date 2022/10/12 19:46
 */
@Data
@AllArgsConstructor
@NoArgsConstructor
// 注册bean
@Component
public class Dog {
    @Value("包子")
    private String name;
    @Value("69")
    private Integer age;
}
