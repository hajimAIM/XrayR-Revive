import os

path = r"C:\Users\user\go\pkg\mod\github.com\xtls\xray-core@v1.251208.0\app\observatory\observer.go"
try:
    with open(path, 'r', encoding='utf-8') as f:
        lines = f.readlines()
        start = 130
        end = 145
        for i in range(start, min(len(lines), end)):
            print(f"{i+1}: {lines[i].rstrip()}")
except Exception as e:
    print(e)
