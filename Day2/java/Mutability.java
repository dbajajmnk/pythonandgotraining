/* 
Write a java Program for Immutablity and Mutability with different operations on String and primitive data types.
ArrayList and Array 
*/
public class Mutability{

public static void main(String[]args)
{
   String totalAmount="1000";
   
   /**
    * name 
    * Memory Allocation
    * value
    */
   System.out.println("Immutability and Mutability in Java");
   // Immutability of String
   String name="John Doe";
   System.out.println("Original Name: " + name);
   name.toUpperCase();
   System.out.println("After toUpperCase(): " + name);
   

   System.out.println("Mutability in Java");
   // Immutability of String
   StringBuilder name1= new StringBuilder("John Doe");
   System.out.println("Original Name: " + name1);
   name1.append(" Smith");
   System.out.println("After append(): " + name1); 
   name1.delete(0, 1); // name remains unchanged
   System.out.println("After delete(): " + name1);

}
}