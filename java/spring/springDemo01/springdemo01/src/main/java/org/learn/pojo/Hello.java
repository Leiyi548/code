package org.learn.pojo;

/**
 * @author jojo
 * @date 2022/10/8 16:44
 */
public class Hello {
    private String name;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void show() {
        System.out.println("Hello" + name);
    }
}
