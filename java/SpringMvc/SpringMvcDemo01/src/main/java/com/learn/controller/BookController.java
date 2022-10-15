package com.learn.controller;

import com.learn.pojo.Book;
import com.learn.service.BookService;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;

import java.util.List;

/**
 * @author jojo
 * @date 2022/10/12 12:14
 */
@Controller
@RequestMapping("/book")
public class BookController {

    private BookService bookService;

    public String list(Model model) {
        List<Book> list = bookService.queryAllBook();
        model.addAttribute("list", list);
        return "allBook";
    }
}

