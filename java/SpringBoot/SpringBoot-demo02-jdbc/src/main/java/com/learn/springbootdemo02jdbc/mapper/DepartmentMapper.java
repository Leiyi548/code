package com.learn.springbootdemo02jdbc.mapper;

import com.learn.springbootdemo02jdbc.pojo.Department;
import org.springframework.stereotype.Repository;

import java.util.List;

/**
 * @author jojo
 * @date 2022/10/12 22:21
 */
@Repository
@Mapper
public interface DepartmentMapper {

    // 获取所有部门信息
    List<Department> getDepartments();

    // 通过id获得部门
    Department getDepartmentById(Integer id);
}
