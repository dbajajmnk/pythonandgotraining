def myDecoration(fn):
    def wrap():
        print("Hello Decorator")
        fn()
        print("After Funciton call")

    return wrap
@myDecoration
def sampleOfDecorator():
    print("This is my function")

sampleOfDecorator()