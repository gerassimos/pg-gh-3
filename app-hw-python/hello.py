import sys

print("Hello, World!")

if True:
    print("Inside the if block")

def filter_even_numbers(numbers):
    """Filter and return even numbers from a given list"""
    return [num for num in numbers if num % 2 == 0]

def print_python_version():
    """Print the Python version being used"""
    print(f"Python version: {sys.version}")

# Example usage
result = filter_even_numbers([1, 2, 3, 4, 5, 6, 7, 8, 9, 10])
print(f"Even numbers: {result}")
print_python_version()

