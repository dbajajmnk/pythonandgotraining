public class Test{
    public static void main(String[] args){
        Car myCar = new Car("Toyota", "Corolla", 2020,new Engine("40HRS"));
        myCar.year = 2021;
        System.out.println(myCar.make);
        System.out.println(Car.wheels); 
        System.out.println( myCar.year);
        myCar.honk();
        Student Student = new Student("Deepak","10");
        
    }
}