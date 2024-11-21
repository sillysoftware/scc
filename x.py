prog = "3 4 +"

stack = []

toks = prog.split()
for tok in toks:
    if tok != "+":
        stack.append(int(tok))
    elif tok == "+":
        lhs = stack.pop(0)
        rhs = stack.pop(0)
        stack.append(lhs + rhs)
print(stack)
