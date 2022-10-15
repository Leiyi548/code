package com.learn.springbootdemo01;

import com.learn.springbootdemo01.pojo.Dog;
import com.learn.springbootdemo01.pojo.Person;
import com.learn.springbootdemo01.pojo.PersonCopy;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
class SpringBootDemo01ApplicationTests {

    @Autowired
    Dog dog;

    @Autowired
    Person person;

    @Autowired
    PersonCopy personCopy;

    @Test
    void contextLoads() {
    }

    @Test
    void dogTest() {
        System.out.println(dog);
    }

    @Test
    void personTest() {
        System.out.println(person);
    }


    @Test
    void personCopyTest() {
        System.out.println(personCopy);
    }
}
