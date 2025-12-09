import os

path = r"C:\Users\user\go\pkg\mod\github.com\xtls\xray-core@v1.251208.0\app\observatory\observer.go"
output = "panic_line.txt"
try:
    with open(path, 'r', encoding='utf-8') as f:
        lines = f.readlines()
        with open(output, 'w', encoding='utf-8') as out:
            for i in range(130, 145):
                if i < len(lines):
                    out.write(f"{i+1}: {lines[i]}")
except Exception as e:
    with open(output, 'w') as out:
        out.write(str(e))
