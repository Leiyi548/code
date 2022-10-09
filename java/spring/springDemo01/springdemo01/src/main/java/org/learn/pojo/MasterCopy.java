package org.learn.pojo;

import org.springframework.beans.factory.annotation.Autowired;

/**
 * @author jojo
 * @date 2022/10/9 10:44
 */
public class MasterCopy {
    @Autowired
    private Cat cat;
    @Autowired
    private Dog dog;
    private String str;

    public Cat getCat() {
        return cat;
    }

    public Dog getDog() {
        return dog;
    }

    public String getStr() {
        return str;
    }

    @Override
    public String toString() {
        return "MasterCopy{" +
                "cat=" + cat +
                ", dog=" + dog +
                ", str='" + str + '\'' +
                '}';
    }
}
