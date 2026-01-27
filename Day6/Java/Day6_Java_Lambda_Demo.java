interface Calc{int add(int a,int b);} 
interface Square{int square(int a);} 


class Day6_Java_Lambda_Demo{
    
    public static void main(String[] a)
    {
        
        // Calc c=(x,y)->x+y; 
        Calc c=(x,y)->x+y; 
        Square s = x->x*x; 
        System.out.println(c.add(2,3));
        System.out.println(s.square(2));

    
    
    }}