class Shape:
    def __init__(self,title):
        self.title= title

    def area(self):
        print(self.title)

class Circle(Shape):
    def __init__(self,title,width):
        super.__init__(title)
        self.width= width

    def printWidth(self):
        print(self.width)

    


