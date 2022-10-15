package com.learn.controller;

import com.learn.pojo.User;

/**
 * @author jojo
 * @date 2022/10/10 0:32
 */
public class UserController {
    public String showUser(User user) {
        System.out.println(user);
        return "showUser";
    }
}
