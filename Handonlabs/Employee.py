class Employee:
	companyName = "testing"
	def __init__ (self, id, name, personal_email, official_email, phone):
			self.id = id
			self.name = name
			self.personal_email = personal_email
			self.official_email = official_email
			self.phone = phone

	def speak(self):
		print(f"Hello, my name is {self.name} and my official email is {self.official_email}")



user = Employee(1,"Hari","h@gmail.com","hari@gmail.com",9345)
print(user.name)
user.speak()
