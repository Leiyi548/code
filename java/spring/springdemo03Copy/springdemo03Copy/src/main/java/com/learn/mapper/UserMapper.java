package com.learn.mapper;

import com.learn.pojo.User;

import java.util.List;

/**
 * @author jojo
 * @date 2022/10/9 17:32
 */
public interface UserMapper {
    public List<User> selectUser();
}
