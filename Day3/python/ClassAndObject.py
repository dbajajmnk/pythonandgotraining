class Car:
    wheels = 4
 
    def __init__(self, make, model, year):
        self.make = make      
        self.model = model    
        self.year = year      
 
    def honk(self):
        print(f"{self.make} {self.model} goes Beep Beep!")
 
 
my_car = Car("Honda", "Civic type R", 2022)
my_car.year = 2023  
print(my_car.make)  
print(my_car.wheels)  
my_car.honk()      
print(type(10))
print (10 is int)