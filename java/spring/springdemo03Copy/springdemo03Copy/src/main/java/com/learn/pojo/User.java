package com.learn.pojo;

/**
 * @author jojo
 * @date 2022/10/9 17:31
 */
public class User {
    private int id;
    private String name;
    private String password;


    @Override
    public String toString() {
        return "User{" +
                "id=" + id +
                ", name='" + name + '\'' +
                ", password='" + password + '\'' +
                '}';
    }
}
