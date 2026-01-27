## Import a funtiontools from annotation\ 
## create a outer function with another fun as paramter
## First step of functiontools.wrap(fn)
## define internal function and pass paramter of outer function (Closure Concept)
## Do the Enchange Things before and after and return passed function in try block
## Finally for End
## last return wrapper (Complete thin
import functools

def myLogger(fn):
    @functools.wraps(fn)
    def wrapper(*a,**kr):
        print(f"Args: {a},Kewargs : {kr}")
        try:
            return fn(*a,**kr)
        finally:
            print("End of funciton")
    return wrapper

@myLogger
def printNumber():
    return 10

@myLogger
def add(a,b):
    return a+b

printNumber()
add(10,20)
add(a=10,b=20)

    

