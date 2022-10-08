package org.learn.dao.impl;

import org.learn.dao.UserDao;

/**
 * @author jojo
 * @date 2022/10/8 16:57
 */
public class UserDaoImpl implements UserDao {
    @Override
    public void getUser() {
        System.out.println("获得用户数据");
    }
}
