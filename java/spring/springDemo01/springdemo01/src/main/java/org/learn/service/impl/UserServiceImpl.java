package org.learn.service.impl;

import org.learn.dao.UserDao;
import org.learn.dao.impl.UserDaoImpl;
import org.learn.service.UserService;

/**
 * @author jojo
 * @date 2022/10/8 17:01
 */
public class UserServiceImpl implements UserService {
    private UserDao userDao = new UserDaoImpl();

    @Override
    public void getUser() {
        userDao.getUser();
    }

    // 利用set实现
    public void setUserDao(UserDao userDao) {
        this.userDao = userDao;
    }
}
