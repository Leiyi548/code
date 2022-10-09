package org.learn.pojo;

/**
 * @author jojo
 * @date 2022/10/9 11:58
 */
public class DogResource {
    private String name;

    public void setName(String name) {
        this.name = name;
    }

    public void shout() {
        System.out.println(name + "wang —— DogSource");
    }
}
