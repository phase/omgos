from __future__ import print_function # Python 2 garbage
import sys

filename = sys.argv[1]
content = open(filename).read()

parsing = False
buf = ""
last = ' '
index = 0
for c in content:
    if parsing:
        if last == '%' and c == '>':
            buf = buf[:-1] # Remove `%` from `%>`
            exec(buf)
            buf = ""
            parsing = False
            continue
        buf = buf + c
        continue
    if last == '<' and c == '%':
        parsing = True
    else:
        if not (c == '<' and content[index+1] == '%'):
            print(c, end='')
    last = c
    index = index + 1