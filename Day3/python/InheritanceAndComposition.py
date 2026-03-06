print("Example of Inheritance")
class Person:
    def __init__(self,name,age,phone,email,address):
        self.name= name
        self.age=age
        self.phone=phone
        self.email=email
        self.address=address

    def speak(self,language="Hindi"):
        print(f"Language {language}")


class Student(Person):
    def __init__(self,name,age,phone,email,address,rollNo):
        super().__init__(name,age,phone,email,address)
        self.rollNo = rollNo

    def study(self):
        print(f"Name {self.name}")





### Composition #####
class Address:
     def __init__(self,houseNo,street,city,state,country,pinCode):
         self.houseNo= houseNo
         self.street= street
         self.city = city
         self.state=state
         self.country=country
         self.pinCode = pinCode
    
     def printAddress(self):
        print(self.city)



myAddress = Address("1401","Tower M","Delhi","Delhi","India",148033)
myAddress2 = Address("1401","Tower M","Delhi","Delhi","India",148033)
kusuma = Student("Kusuma",20,931334444444,"deepak@gmail.com",myAddress2,59)
kusuma.study()
kusuma.address.printAddress()
















print("Inheritance Example End")
print("*******************************************************")
print("Composition Example Start")

print("Compostion Example End")
