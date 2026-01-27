squreRoot = lambda x:x*x
multi = lambda x,y:x*y
print(squreRoot(10))
print(multi(10,5))
for num in range(10): print(squreRoot(num))


##Collection 
list = [1,2,3,4,5,6]
print("Simple",list)
list = [x for x in range(1,7) if x%2==0]

list = lambda x:x%2==0
print("Comprension",list)
