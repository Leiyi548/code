import org.junit.Test;
import org.learn.pojo.*;
import org.learn.service.UserService;
import org.learn.service.impl.UserServiceImpl;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import java.util.Arrays;

/**
 * @author jojo
 * @date 2022/10/8 16:49
 */
public class SpringBeanTest {

    @Test
    public void testUser() {
        UserService userService = new UserServiceImpl();
        userService.getUser();
    }

    @Test
    public void testHello() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        Hello hello = (Hello) context.getBean("hello");
        hello.show();
    }

    @Test
    public void testUserDao() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        UserServiceImpl userServiceImpl = (UserServiceImpl) context.getBean("serviceImpl");
        userServiceImpl.getUser();
    }

    @Test
    public void testUserShow() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        // 在执行getBean的时候,user已经创建好，通过无参构造
        User user = (User) context.getBean("user");
        // 会先执行构造函数
        user.show();
    }

    @Test
    public void testUserArgsIndex() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        UserT usert = (UserT) context.getBean("userTIndex");
        usert.show();
    }


    @Test
    public void testUserArgsName() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        UserT usert = (UserT) context.getBean("userTName");
        usert.show();
    }

    @Test
    public void testUserArgsType() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        UserT usert = (UserT) context.getBean("userTType");
        usert.show();
    }


    @Test
    public void testUserArgsAliasType() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        UserT usert = (UserT) context.getBean("TType");
        usert.show();
    }

    @Test
    public void testStudentName() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        Student student = (Student) context.getBean("studentTest");
        System.out.println(student.getName());
    }

    @Test
    public void testStudent() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        Student student = (Student) context.getBean("student");
        System.out.println(student.getName());
        System.out.println(student.getAddress());
        System.out.println(Arrays.toString(student.getBooks()));
        System.out.println(student.getHobbies());
        student.show();
    }

    @Test
    public void testUserP() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        User user = (User) context.getBean("userP");
        System.out.println(user);
    }

    @Test
    public void testUserC() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        User user = (User) context.getBean("userC");
        System.out.println(user);
    }

    @Test
    public void testMaster01() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        Master master = (Master) context.getBean("masterTest");
        master.getDog().shout();
        master.getCat().shout();
    }

    @Test
    public void testMasterByname() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        Master master = (Master) context.getBean("masterByName");
        master.getDog().shout();
        master.getCat().shout();
    }


    @Test
    public void testMasterByType() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        Master master = (Master) context.getBean("masterByType");
        master.getDog().shout();
        master.getCat().shout();
    }


    @Test
    public void testMasterCopyAutowired() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        MasterCopy masterCopy = (MasterCopy) context.getBean("masterAutowiredCopy");
        masterCopy.getCat().shout();
        masterCopy.getDog().shout();
    }

    @Test
    public void testMasterCopyQualifier() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        MasterCopy2 masterQualifier = (MasterCopy2) context.getBean("masterQualifier");
        masterQualifier.getDog().shout();
        masterQualifier.getCat().shout();
    }

    @Test
    public void testMasterResource() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        MasterResource masterResource = (MasterResource) context.getBean("masterResource");
        masterResource.getDogResource().shout();
        masterResource.getCatResource().shout();
    }

    @Test
    public void testAnnotationTeacher01() {
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        Teacher teacher = (Teacher) context.getBean("teacher01");
        teacher.show();
    }

}

