### Control Statements ###
def ControlFlowFromPackage():
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

