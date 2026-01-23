from revision.Variable import isTeamResponsive,membersCount,teamName

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
