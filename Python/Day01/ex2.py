#!/usr/bin/python3  

def printName():
    print("Hello world")
    name = input("Enter your name: ")
    print(f"Hi {name}")

def Addition(n1, n2):
    return n1+n2

printName()
num1 = int(input("Input first number: "))
num2 = int(input("Input second number: "))
print(Addition(num1, num2))

