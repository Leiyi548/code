package org.learn.dao.impl;

import org.learn.dao.UserDao;

/**
 * @author jojo
 * @date 2022/10/8 17:10
 */
public class UserDaoMysqelImpl implements UserDao {
    @Override
    public void getUser() {
        System.out.println("MySql获得用户数据");
    }
}
