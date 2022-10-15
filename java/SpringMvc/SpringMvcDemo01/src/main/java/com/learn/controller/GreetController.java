package com.learn.controller;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;

/**
 * @author jojo
 * @date 2022/10/9 23:40
 */
@Controller
@RequestMapping("/greetController")
public class GreetController {


    //真实访问地址 : 项目名/greetingController/greeting
    @RequestMapping("/greet")
    public String greeting(Model model, @RequestParam("username") String name) {
        model.addAttribute("msg", "greet,springmvc");
        System.out.println(name);
        //web-inf/jsp/greet.jsp
        return "greet";
    }
}

// "hello,world"
// 'hello,world'