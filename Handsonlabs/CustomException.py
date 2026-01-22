# class CheckOutError(Exception):

#   class OutOfStock(CheckOutError):
#       pass

#   class InvalidCouponError(CheckOutError):
#       pass

# def apply_Coupon(coupon_Code,total):
#     if coupon_Code != "HariIsRock":
#        raise InvalidCouponError(coupon_Code)
#        return total

# try:
#     apply_Coupon("Team",100)
#     expect InvalidCouponError as e:
#      print("Coupon Error",e)

class CheckOutError(Exception):
    """"Base Class of execption"""

class OutOfStock(CheckOutError):
    pass

class InvalidCoupon(OutOfStock):
    pass

def apply_Coupon(coupon_Code,total):
    if coupon_Code!="HariIsRock":
        raise InvalidCoupon("Sorry! Coupon Code is not Valid")
        
    return total

try:
    print(apply_Coupon("HariIsRock",100))
except InvalidCoupon as e:
    print("Invalid Coupon Error : ",e)




