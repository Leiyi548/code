package com.learn.service;

import com.learn.pojo.Book;

import java.util.List;

/**
 * @author jojo
 * @date 2022/10/10 21:15
 */
public interface BookService {
    //增加一个Book
    int addBook(Book book);

    //根据id删除一个Book
    int deleteBookById(int id);

    //更新Book
    int updateBook(Book books);

    //根据id查询,返回一个Book
    Book queryBookById(int id);

    //查询全部Book,返回list集合
    List<Book> queryAllBook();
}
