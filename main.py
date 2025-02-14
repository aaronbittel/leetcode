import importlib
import sys
from typing import Callable, Any

COLOR24 = "\033[38;2;{};{};{}m"
RESET = "\033[m"

FAILURE_COLOR = COLOR24.format(197, 99, 99)
SUCCESS_COLOR = COLOR24.format(101, 137, 132)

FAILURE = f"{FAILURE_COLOR}Failure!{RESET}"
SUCCESS = f"{SUCCESS_COLOR}Success!{RESET}"


def import_module(module_name: str) -> type:
    """Dynamically import a module and return its Solution class."""
    solution_module = importlib.import_module(module_name.split(".")[0])
    return getattr(solution_module, "Solution")


def run_tests(solution: type, method_name: str):
    """Run the dynamically determined method with test cases."""
    method: Callable[..., Any] = getattr(solution(), method_name)

    for tt in solution.TESTS:
        inputs = tt.ins if isinstance(tt.ins, tuple) else (tt.ins,)
        got = method(*inputs)
        if got != tt.out:
            print(f"Ins: {tt.ins} -> {FAILURE} Expected: {tt.out}, got: {got}")
        else:
            print(f"Ins: {tt.ins} -> {got} {SUCCESS}")


def print_usage():
    """Print correct usage instructions."""
    print("Usage: python main.py <module_name>")
    print("Example: python main.py sum3_closest.py")


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print_usage()
        sys.exit(1)

    sol_class = import_module(sys.argv[1])

    methods = [
        method
        for method in dir(sol_class)
        if callable(getattr(sol_class, method))
        and not method.startswith("__")
        and method != "tests"
    ]

    if not methods:
        print("Error: No valid testable method found in the class.")
        sys.exit(1)

    method_name = methods[0]
    print(f"Running tests for problem: {method_name}\n")

    run_tests(sol_class, method_name)
