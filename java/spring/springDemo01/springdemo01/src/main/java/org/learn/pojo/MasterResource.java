package org.learn.pojo;

import javax.annotation.Resource;

/**
 * @author jojo
 * @date 2022/10/9 11:58
 */
public class MasterResource {
    //    @Resource(name = "dogResource01")
    // 如果这个名字设置为dog，那么他会找dog类，所以我们尽量设置和类名一样的名字！
//    @Resource
    @Resource(name = "dogResource01")
    DogResource dogResource;
    //    @Resource(name = "catResource01")
    @Resource
    CatResource catResource;

    public DogResource getDogResource() {
        return dogResource;
    }

    public CatResource getCatResource() {
        return catResource;
    }
}
