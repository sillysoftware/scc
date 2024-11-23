def getInfix(postfix):
    stack = []
    operators = set("+-*/^")

    for token in postfix.split():
        if token not in operators:
            stack.append(token)
        else:
            operand2 = stack.pop()
            operand1 = stack.pop()
            expression = f"({operand1} {token} {operand2})"
            stack.append(expression)

    return stack[0]


def translate_to_c(program):
    # Split the program into lines
    lines = program.strip().splitlines()

    # Initialize the result with standard includes
    c_code = ["#include <stdio.h>", "int main() {", "    int x;"]

    for line in lines:
        line = line.strip()
        if "->" in line:  # Assignment operation
            parts = line.split(" -> ")
            expression = parts[0].strip()
            variable = parts[1].strip().removesuffix(";")
            # Convert the operation into C syntax
            expression = getInfix(expression)
            c_code.append(f"    {variable} = {expression};")
        elif line.startswith("println:"):  # Print operation
            variable = line.split(":")[1].strip().removesuffix(";")
            c_code.append(f"    printf(\"%d\\n\", {variable});")

    # Close the main function
    c_code.append("    return 0;")
    c_code.append("}")

    return "\n".join(c_code)


# Original Program
program = """
x;
3 4 + -> x;
println: x;
"""

# Translate and display the C code
c_code = translate_to_c(program)
print(c_code)
