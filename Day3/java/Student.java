class Student extends Person
{
    String rollNo;
    public Student(String name,String rollNo)
    {
        super(name);
        System.out.print("Student Intialized");
        this.rollNo=rollNo;
    }
}