package com.learn.springbootdemo01.controller;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author jojo
 * @date 2022/10/12 17:42
 */
@RestController
public class HelloController {
    @RequestMapping("/hello")
    public String helloSpring() {
        return "hello,Spring";
    }
}
