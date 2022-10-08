package org.learn.pojo;

/**
 * @author jojo
 * @date 2022/10/8 17:44
 */
public class User {
    private String name;
    private int age;

    public User() {
//        System.out.println("user 无参构造函数");
    }

    public User(String name, int age) {
        this.name = name;
        this.age = age;
    }

    public String getName() {
        return name;
    }

    public int getAge() {
        return age;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setAge(int age) {
        this.age = age;
    }

    public void show() {
        System.out.println("useranme=" + name);
    }

    @Override
    public String toString() {
        return "User{" +
                "name='" + name + '\'' +
                ", age=" + age +
                '}';
    }
}
