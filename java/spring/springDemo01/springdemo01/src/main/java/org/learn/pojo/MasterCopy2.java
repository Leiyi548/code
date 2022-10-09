package org.learn.pojo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;

/**
 * @author jojo
 * @date 2022/10/9 11:28
 */
public class MasterCopy2 {
    //  如果想要使用Autowired，那么你一定要设置这个bean为这个使用Autowired的类！
    @Autowired
    // 这个value是bean的名称
    @Qualifier(value = "catCopy2")
    CatCopy2 cat;
    @Autowired
    @Qualifier(value = "dogCopy2")
    DogCopy2 dog;

    public CatCopy2 getCat() {
        return cat;
    }

    public DogCopy2 getDog() {
        return dog;
    }
}
