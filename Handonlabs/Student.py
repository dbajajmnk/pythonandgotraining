## Duck Typing
# class Animal : 
#     def __init__(self,name,age):
#         self.name=name
#         self.age = age



class Person : 
    def __init__(self,name,age,dob,email,phone):
        self.name=name
        self.age = age
        self.dob = dob
        self.email = email
        self.phone = phone
    def marry(self):
        print("Pooja")




# dog = Animal("Tony",1)
# john = Person("John",30,"30Dec1983","testing","testing 1")



# print(type(dog))
# print(type(john))
# print(dog is john)

class Student(Person):
    def __init__(self, name, age, dob, email, phone,rollNo):
        super().__init__(name, age, dob, email, phone)  
        self.rollNo= rollNo
 
       
student1 = Student("D","10","33","email","93333",10)
student1.marry()
print(student1.name)

print("Composition")
class Engine:
    def __init__(self,power):
        self.power = power

class Car :
    def __init__(self,engine):
        self.engine = engine

    def checkEngine(self):
        print(self.engine)

hondaEngine = Engine("345")
hondacity = Car(hondaEngine)
        




