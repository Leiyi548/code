import com.learn.mapper.UserMapper;
import com.learn.pojo.User;
import org.junit.Test;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

/**
 * @author jojo
 * @date 2022/10/9 17:44
 */
public class Test01 {

    @Test
    public void springMybatisTest(){
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");
        UserMapper userMapper = (UserMapper)context.getBean("userMapper");
        System.out.println(userMapper.selectUser());
    }
}
