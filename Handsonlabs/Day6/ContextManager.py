## Import required packages 
    ## Threading
    # contextmanager from contextLib
    # __future__ from annotations
# Create thread lock object
# Use contextmanager dactorator on your fuction to make it context aware 
    # need to right code to work on file as per passed parameter and open it
    # Use try /catch finally to track the file operations status
    
# Write main and apply context manager with Lock under local , for this call funciton with to create a file 
# read the context of file and print on terminal

from __future__ import annotations
import threading
from contextlib import contextmanager

LOCK = threading.Lock()

@contextmanager
def doFileOperation(path,mode):
    f = open(path,mode,encoding="utf-8")
    try:
        yield f
        print("File Commited")
    except Exception as e:
        print("Roll Back Happend")
        raise
    finally:
        f.close()
        print("File Closed")

def main():
    try:
        with LOCK:
            with doFileOperation("contextmanager.txt","w") as f:
                f.write("Simple\nHarivardan\nSunil\n Manoj")

    except RuntimeError:
        print("Handled error; notice cleanup still happened.")             


if __name__ == "__main__":
    main()