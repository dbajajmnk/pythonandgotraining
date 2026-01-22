interface Calc{int add(int a,int b);} 


class Day6_Java_Lambda_Demo{
    
    public static void main(String[] a)
    {
        
        Calc c=(x,y)->x+y; 
        System.out.println(c.add(2,3));
    
    
    }}