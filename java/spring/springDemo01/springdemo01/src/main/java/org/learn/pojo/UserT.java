package org.learn.pojo;

/**
 * @author jojo
 * @date 2022/10/8 17:59
 */
public class UserT {
    private String name;

    public UserT(String name) {
        this.name = name;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void show() {
        System.out.println("userTName=" + name);
    }
}
