from pydantic import BaseModel,Field,EmailStr

class Product(BaseModel):
    name:str = Field(min_length=2,max_length=20,description="Name is required")
    price:int
    discount:int
    email:EmailStr


adidasShoes = Product(name="Deepak",price=10,discount=2,email="dddssadg")
print(adidasShoes)


