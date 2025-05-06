"""
Check for Balance
Write a function that takes a string of source code and checks whether the braces/parentheses are balanced. Every ( or { must be closed by a } or ) in the opposite order. Return the index at which an imbalance occurs, or -1 if the string is balanced. If any ( or { are never closed, return the string's length.

Examples
check_balance("if (a(4) > 9) { foo(a(2)); }") ➞ -1
# Returns -1 because it's balanced.

check_balance("for (i=0;i<a(3};i++) { foo{); )") ➞ 14
# Returns 14 because } is out of order.

check_balance("if (x) {")  ➞ 8
# Returns 8 because { is never closed.
"""
def check_balance(code: str) -> int:
    stack = []
    
    for i, ch in enumerate(code):
        if ch in '({':
            stack.append((ch, i))
        elif ch in ')}':
            if not stack:
                return i
            top_char, top_index = stack.pop()
            if (ch == ')' and top_char != '(') or (ch == '}' and top_char != '{'):
                return i
    
    if stack:
        return len(code)

    return -1


print(check_balance("if (a(4) > 9) { foo(a(2)); }")), # -1)
print(check_balance("for (i=0;i<a(3};i++) { foo{); )")), # 14)
print(check_balance("while (true) foo(); }{ ()")), # 20)
print(check_balance("if (x) {")), # 8)
print(check_balance("if (x) }")), # 7)
print(check_balance("(({{}}){}{}())")), # -1)
print(check_balance("({)}")), # 2)
print(check_balance("(")), # 1)
print(check_balance("}")), # 0)
print(check_balance("")), # -1)