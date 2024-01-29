# ml_model/main.py

from generate_module import Generate

def main():
    inp = ' '.join(sys.argv[1:])
    obj = Generate()
    out = obj.generate_text(inp)
    print(out)

if __name__ == "__main__":
    main()