''''
String and List Mutability Test in Python
Mutable Object and Immutable
'''
# list = [1,2,3,5,"Deepak",10.5,True]
# print("Before Change:", list)
# list.append(("hello","Hope you are enjoying"))
# print("After change:",list)

# #Immutablity
# x=5
# #print("Original Value:",id(x))
# print("Original Value:",x)
# def change_value(a):
#     x=a+5


# change_value(x)
# print("Original Value after function call:",x)
    


# # print("After Function Call:", x)
# # x=x+5
# # b=5
# # print("After change:", id(b))
# # print("After change:", id(x))
# # print("After change:", id(x))
# # print("Original Value:",x)
## Mohammand
x = 5
 
print(f"original value of x {x}")
 
def multiply_x(x):
    return x*x
 
multiply_x(x)
print(f"value after operation {x}")

names = ["pavan", "Ashish", "Simple"]
 
print(f'names before {names}')
 
def add_name(x):
    names.append(x)
 
add_name('deepak')
print(f'names after {names}')
 

 


	
	



