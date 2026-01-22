def my_decorator(fn):
    def wrap():
        print('Before')
        fn()
        print('After')
    return wrap

@my_decorator
def hello():
    print('Hello')

hello()