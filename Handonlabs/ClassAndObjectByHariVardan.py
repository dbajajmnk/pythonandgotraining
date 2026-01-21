class User:
    companyName = "testing"

    def __init__(self,name,email,age,phoneNumber):
        self.name = name
        self.email = email
        self.age = age
        self.phoneNumber = phoneNumber

    def speak(self):
        print(f"{self.name} is speaking")


user = User("Hari","h@gmail.com","hari@gmail.com",9345)
print(user.name)
user.speak()





a = None
print(a)