package com.learn.springbootdemo01.pojo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.PropertySource;
import org.springframework.stereotype.Component;

import java.util.Date;
import java.util.List;
import java.util.Map;

/**
 * @author jojo
 * @date 2022/10/12 20:18
 */
@Component
@PropertySource(value = "classpath:person.properties")
@Data
@AllArgsConstructor
@NoArgsConstructor
public class PersonCopy {
    @Value("${name}")
    private String name;
    private Integer age;
    private Boolean happy;
    private Date birth;
    private Map<String, Object> maps;
    private List<Object> lists;
    private Dog dog;
}
