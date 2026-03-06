# Steps Decorator 
# Step 1: from functools import wraps
# Step 2: Create Decorator function and use fn (function) as a parameter
# Step 3: use @functools.warps(fn) - showing that this is a decorator
# Step 4: Define the wrapper using the *arguments as parameter 
# Step 5: Write the businesss logic you want to implement before the function
# Step 6: Return the function to the wrapper like fn(*args)
# Step 7: Return the Wrapper value from the decorator

from functools import wraps
import logging
import time

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s - %(levelname)s - %(message)s"
)


def logg(fn):
    @wraps(fn)
    def wrapper(*args,**kwargs):
        # 
        start_time=time.time()
        logging.info(f"Function {fn.__name__} started")
        
        result= fn(*args,**kwargs)
        
        end_time=time.time()
        
        
        
            # print(f"Function {fn} got called successfully")
        logging.info(f"Function {fn.__name__} finished execution")
        logging.info(f"Function started at '{start_time}' and ended at '{end_time}'")

    return wrapper

# def tme(fn):
#     @wraps(fn)
#     def wrapper(*args,**kwargs):
#         start_time=time.time()

#         logging.info()

@logg
def say_Hi():
    """I am the origin"""

say_Hi()




