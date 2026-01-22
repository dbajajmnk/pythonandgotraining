def count_up(n):
    for i in range(n):
        yield i

for x in count_up(5):
    print(x)