## Variable and Data types Done
## print method Done
## Function and Argument Done
## Control Statement 
    #For
    #While
    #if
    #if else
    #if elif
    #match
    #continue
    #break
## Inheritance && Composition
## Collections 
## Exception 
## Custom Exception
## Duck Typing 
## Classes 
## type,id,instancIn
## Comprehension

teamName="Rocking Team with Go,Java and Python Knowledge"
membersCount=21000000000000000
isTeamResponsive = True

##Var args
def printValue(*args):
    print(args)
##Positional 
def add(a,b)-> int:
    return a+b
## Key argument : Order doesn't matter
def printStudentInfo(name,rollNo):
    print(name,rollNo)

printStudentInfo(rollNo=10,name="Deepak")

## Defaults
def multi(a,b=10):
        print("Multipicaiton Result",a*b)

##Key args
def printClassData(** studentDetails):
        for item in studentDetails.items():
                print("Item",item)

printValue(teamName,membersCount,isTeamResponsive)
print(add(10,30))
print(multi(30))
print(multi(40,20))
printClassData(studentDetails={"1":"A","2":"B"})

### Control Statements ###
num1 = int(input("Enter Number 1"))
num2 = int(input("Enter Number 2"))
num3 = int(input("Enter Number 3"))

if num1 > num2:
    print("I am Working Fine23")

if num1>num2:
    print("Number 1 is Greater",num1)
else:
    print("Number 2 is Greater",num2)


if num1>num2:
    print("Number 1 is Greater",num1)
    if num1>num3:
        print("Number 3 is Greater",num3)
else:
    if num2>num3:
        print("Number 2 is Greater",num2)
    else:
        print("Number 3 is Greater",num3)


count = 1
while count<=10:
     print(num1*count)  
     count+=1 





