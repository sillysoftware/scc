prog = "3 4 +"

stack = []
toks = prog.split()
for tok in toks:
    if tok != "+":
        stack.append(int(tok))
    else:
        lhs = int(stack.pop)
        rhs = int(stack.pop)
        res = lhs + rhs
        stack.append(res)

print(stack)
