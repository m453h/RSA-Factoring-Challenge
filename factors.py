#!/usr/bin/python3

import sys

if len(sys.argv) < 2:
    print("Usage: factors <file>")
    sys.exit()

file_path = sys.argv[1]

file = open(file_path, 'r')
lines = file.readlines()

for line in lines:
    num = int(line)
    flag = False
    for i in range(2, num // 2 + 1):
        if num % i == 0:
            print(f"{num}={num // i}*{i}")
            flag = True
            break
