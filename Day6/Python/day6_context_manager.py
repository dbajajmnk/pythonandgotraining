import os

print("Current working directory:", os.getcwd())

with open('sample.txt', 'w') as f:
    f.write('Hello Context Manager')

with open('sample.txt','r') as f:
    print(f.read())
