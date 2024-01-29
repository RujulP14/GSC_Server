# chatbot/main.py

# Run this file in command line with the input string and it will print the response in the command line

# Example
# python3 main.py Hello World
# Hello World!

from generate_module import Generate
import sys

def main():
    inp = ' '.join(sys.argv[1:])
    obj = Generate()
    out = obj.generate_text(inp)
    print(out)

if __name__ == "__main__":
    main()