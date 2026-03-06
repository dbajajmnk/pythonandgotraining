from time import sleep

def callBackFun(value):
    print(value)

def apiCall(fn):
    sleep(100)
    fn("World")


print("Hello")
apiCall(callBackFun)
print("I am running without waiting")