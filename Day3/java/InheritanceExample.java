class InheritanceExample {
    public static int wheels = 4;
    public String make;
    public String model;
    public int year;
    public Car(String make,String model,int year)
    {
        this.make = make;
        this.model = model;
        this.year = year;
    }
        
 
    public void  honk(){
        System.out.println(this.make + " " + this.model + " goes Beep Beep!");
    }
       
 


}