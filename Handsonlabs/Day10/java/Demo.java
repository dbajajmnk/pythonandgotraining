public class Demo {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
        Student student = new Student("Alice", 20);
        System.out.println("Student Name: " + student.getName());
        System.out.println("Student Age: " + student.getAge());
        Student student2 = new Student("Bob", 22);
        System.out.println("Student2 Name: " + student2.getName()); 
        System.out.println("Student2 Age: " + student2.getAge());
    }
}

class Student {
    String name;
    int age;

    public Student(String name, int age) {
        this.name = name;
        this.age = age;
    }

    public String getName() {
        return name;
    }

    public int getAge() {
        return age;
    }
}