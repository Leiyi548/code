package com.learn.service;

import com.learn.pojo.Book;

import java.util.List;

/**
 * @author jojo
 * @date 2022/10/10 21:15
 */
public interface BookService {
    //增加一个 Book
    int addBook(Book book);

    //根据 id 删除一个 Book
    int deleteBookById(int id);

    //更新 BookBook
    int updateBook(Book books);

    //根据 id 查询，返回一个 Book 个 Book
    Book queryBookById(int id);

    //查询全部 Book，返回 list 集合
    List<Book> queryAllBook();
}
