"""Find the Pattern and Write the Function
By looking at the inputs and outputs below, try to figure out the pattern and write a function to execute it for any number.

Examples
func(3456) ➞ 2

func(89265) ➞ 5

func(97) ➞ 12

func(2113) ➞ -9"
"""

def func(num: int) -> int:
    
    digits = str(num)
    return sum(map(int, digits)) - len(digits) ** 2
    
    # sum_num = 0
    # for i in list(str(num)):
    #     sum_num += int(i)
    
    # return sum_num - len(str(num)) * len(str(num))


print(func(3456)), # 2)
print(func(89265)), # 5)
print(func(97)), # 12)
print(func(2113)), # -9)
print(func(55)), # 6)
print(func(785428)), # -2)
print(func(439)), # 7)
print(func(55654)), # 0)