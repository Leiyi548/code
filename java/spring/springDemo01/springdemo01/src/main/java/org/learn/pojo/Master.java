package org.learn.pojo;

/**
 * @author jojo
 * @date 2022/10/8 22:06
 */
public class Master {
    private Cat cat;
    private Dog dog;
    private String str;

    public void setCat(Cat cat) {
        this.cat = cat;
    }

    public void setStr(String str) {
        this.str = str;
    }

    public void setDog(Dog dog) {
        this.dog = dog;
    }

    public Dog getDog() {
        return dog;
    }

    public Cat getCat() {
        return cat;
    }

    @Override
    public String toString() {
        return "Master{" +
                "cat=" + cat +
                ", dog=" + dog +
                ", str='" + str + '\'' +
                '}';
    }
}
