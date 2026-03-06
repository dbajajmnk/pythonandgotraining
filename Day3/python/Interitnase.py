class Person:
    def __init__(self,name,age,phone,email):
        self.name= name
        self.age=age
        self.phone=phone
        self.email=email

    def speak(self,language="Hindi"):
        print(f"Language {language}")


class Student(Person):
    def __init__(self,name,age,phone,email,rollNo):
        super().__init__(name,age,phone,email)
        self.rollNo = rollNo

    def study(self):
        print(f"Name {self.name}")



kusuma = Student("Kusuma",20,931334444444,"deepak@gmail.com",59)
kusuma.study()