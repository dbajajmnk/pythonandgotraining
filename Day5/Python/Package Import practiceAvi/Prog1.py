from my_package import sumation as sm
from my_package import subtract as st
from my_package import divide as dd
from my_package import multiply as mt
# from my_package.subtract import subt 

def main1():
    a=int(input("Enter the first number= "))
    b=int(input("Enter the Second number= "))
    print(sm.sumer(a,b))
    print(dd.divd(a,b))
    print(mt.multi(a,b))
    print(st.subt(a,b))

    print(sm.sumerTwo(a,b))
    print(dd.divfloat(a,b))
    print(mt.multiTwo(a,b))
    print(st.subtTwo(a,b))

if __name__=="__main__":
    main1()

