import os

path = r"C:\Users\user\go\pkg\mod\github.com\xtls\xray-core@v1.251208.0\transport\internet\tagged"
try:
    for f in os.listdir(path):
        print(f)
except Exception as e:
    print(e)
