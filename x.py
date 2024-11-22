prog = "5 7 / 2 + ."

stack = []
segment = {}

toks = iter(prog.split())
for tok in toks:
    if tok.isdigit():
        stack.append(int(tok))
    elif tok.isalpha() and tok not in segment:
        segment[str(tok)] = next(toks, 0)
    elif tok.isalpha() and tok in segment:
        stack.append(segment[str(tok)])
    elif tok == ".":
        print(stack)
    elif tok == "+":
        lhs = stack.pop(0)
        rhs = stack.pop(0)
        stack.append(lhs + rhs)
    elif tok == "-":
        lhs = stack.pop(0)
        rhs = stack.pop(0)
        stack.append(lhs - rhs)
    elif tok == "*":
        lhs = stack.pop(0)
        rhs = stack.pop(0)
        stack.append(lhs * rhs)
    elif tok == "/":
        lhs = stack.pop(0)
        rhs = stack.pop(0)
        stack.append(lhs / rhs)

print(segment)
