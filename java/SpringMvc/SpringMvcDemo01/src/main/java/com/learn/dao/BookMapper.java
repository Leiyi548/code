package com.learn.dao;

import com.learn.pojo.Book;

import java.util.List;

/**
 * @author jojo
 * @date 2022/10/10 19:59
 */
public interface BookMapper {

    // 增加一本书
    int addBook(Book book);

    // 根据id删除一本书
    int deleteBookById(int id);

    // 更新book
    int updateBook(Book book);

    // 根据id查询一本书
    Book queryBookById(int id);

    // 查询全部book，返回一个集合
    List<Book> queryAllBooks();
}
