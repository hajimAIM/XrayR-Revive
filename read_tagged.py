import os

path = r"C:\Users\user\go\pkg\mod\github.com\xtls\xray-core@v1.251208.0\transport\internet\tagged\tagged.go"
try:
    with open(path, 'r', encoding='utf-8') as f:
        print(f.read())
except Exception as e:
    print(e)
