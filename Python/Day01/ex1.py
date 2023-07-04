#!/usr/bin/python3  

import mylib

def main():
    mylib.printName()
    num1 = int(input("Input first number: "))
    num2 = int(input("Input second number: "))
    print(mylib.Addition(num1, num2))

if __name__ == "__main__":
    main()
