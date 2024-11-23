prog = "3 4 +"

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
    elif tok in "=-*/":
        lhs = stack.pop(0)
        rhs = stack.pop(0)
        if tok == "+":
            stack.append(lhs + rhs)
        elif tok == "-":
            stack.append(lhs - rhs)
        elif tok == "*":
            stack.append(lhs * rhs)
        elif tok == "/":
            stack.append(lhs // rhs)
