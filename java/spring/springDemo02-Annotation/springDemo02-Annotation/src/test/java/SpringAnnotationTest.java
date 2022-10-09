import com.learn.config.MyConfig;
import com.learn.pojo.Dog;
import com.learn.service.UserService;
import org.junit.Test;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

/**
 * @author jojo
 * @date 2022/10/9 14:02
 */
public class SpringAnnotationTest {
    @Test
    public void configBeanTest(){
        ApplicationContext applicationContext = new AnnotationConfigApplicationContext(MyConfig.class);
        Dog dog = (Dog)  applicationContext.getBean("dog");
        System.out.println(dog.name);
    }


    @Test
    public void aopTest(){
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        UserService userService = (UserService) context.getBean("userService");
        userService.search();
    }

    @Test
   public void customPointcut(){
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        UserService userService = (UserService) context.getBean("userService");
        userService.delete();
    }

}
