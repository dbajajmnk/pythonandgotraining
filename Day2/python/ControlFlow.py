'''
Docstring for Day2.python.ControlFlow

'''
##Loops
    #while
    #
## if else 
## match case
## continue and brek
# age = 20
# if age > 18:
#     print("You are an adult and Eligible to vote.")   
# else:
#     print("You are a minor and not Eligible to vote.")

# marksObtained = 85
# if marksObtained >= 90:
#     print("Grade A")
# elif marksObtained >= 75:
#     print("Grade B")    
# elif marksObtained >= 60:
#     print("Grade C")
# else:
#     print("Grade D")

# match case
# day = 20

# if day == 1:
#     print("Monday")
# elif day == 2:
#     print("Tuesday")
# elif day == 3:
#     print("Wednesday")
# elif day == 4:
#     print("Thursday")
# elif day == 5:
#     print("Friday")
# elif day == 6:
#     print("Saturday")

# elif day == 7:
#     print("Sunday")
# else:
#     print("Invalid day")

# match day:
#     case 1:
#         print("Monday")
#     case 3:
#         print("Wednesday")
#     case 4:
#         print("Thursday")
#     case 5:
#         print("Friday")
#     case 6:
#         print("Saturday")
#     case 7:
#         print("Sunday")
#     case _:
#         print("Invalid day")


## Loops
for i in range(5):
    print("Iteration:", i)
for item in list(range(5)):
    print("Item:", item)
for item in "Hello":
    print("Character:", item)
for tupleitem in (1, 'a'):
    print("Tuple Item:", tupleitem)
for mapItem in {'name': 'John', 'age': 30}:
    print("Map Item:", mapItem)
## While loop
count = 0
while count < 5:
    print("Count:", count)
    count += 1
## Continue and Break
for i in range(10):
    if i % 2 == 0:
        continue  # Skip even numbers
    if i > 7:
        break  # Exit loop if i is greater than 7
    print("Odd Number:", i) 



   