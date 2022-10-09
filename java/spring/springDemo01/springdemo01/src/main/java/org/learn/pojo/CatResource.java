package org.learn.pojo;

/**
 * @author jojo
 * @date 2022/10/9 11:59
 */
public class CatResource {
    private String name;

    public void setName(String name) {
        this.name = name;
    }

    public void shout() {
        System.out.println(name + "miao —— CatResource");
    }
}
