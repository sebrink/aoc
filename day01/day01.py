#!/usr/local/bin/python3
import itertools

# Part 1
data=[]
file = open("input.txt","r")
for i in file:
    data.append(int(i))
print("Part 1 Answer: "+str(sum(data)))

# Part 2
freq = 0
seen = {0}
for num in itertools.cycle(data):
    freq += num
    if freq in seen:
        print("Part 2 Answer: "+(str(freq)))
        exit()
    seen.add(freq)
