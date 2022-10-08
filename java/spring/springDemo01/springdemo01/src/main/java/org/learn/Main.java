package org.learn;

/**
 * @author jojo
 * @date ${DATE} ${TIME}
 */
public class Main {
    public static void main(String[] args) {
        String s1 = "abc";
        String newS1 = "abc";
        String s2 = new String("abc");
        String s3 = new String("abc");
        System.out.println(s1 == newS1); // true
        System.out.println(s1 == s2); // flase
        System.out.println(s1.equals(s2)); // true
        System.out.println(s2 == s3); // false
        System.out.println(s2.equals(s3)); // true
    }
}