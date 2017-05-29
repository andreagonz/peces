import random

v = []
random.seed(0)
for x in range(1000):
    v.append(random.randint(-10000, 10000))
print(list(set(v)))

s = 0
for x in range(len(v)):
    if random.random() > 0.8:
        s += v[x]
print(s)
