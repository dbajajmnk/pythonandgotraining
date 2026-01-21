'''
How to write a Custom Function with Arguments in Python
How to use inbuild Functions with Arguments in Python
Different Aruguments Types in Python Functions
Dynamic Arguments in Python Functions
'''
def add_nums(a,b = 10):
    return a+b 
def add(a, b):
    return a + b
def subtract(a, b):
    return a - b

def multiply(a, b):
    return a * b

print(add(10, 5))
print(subtract(10, 5))
print(multiply(10, 5))
def describe(**kwargs):
    for key, value in kwargs.items():
        print(f"{key} = {value}")
 
def f(a, b=2, c=3):
    print(a, b, c)
 
f(10,c=50)
 
 
def sample(name, age, type):
    print(name, age, type)
 
sample(age=20, name = 'test name', type=1)
 
 
def multi_return():
    return 1,2,3
 
test_return = multi_return()
 
print(test_return)
def add_nums(*args):
    return sum(args)