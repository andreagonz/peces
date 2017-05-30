import random
 
random.seed(0)
d = {}
x = 0
while x < 10000:
    r = random.randint(-4000000, 4000000)
    if d.get(r, None) == None:
        k = {r : r}
        d.update(k)
        x += 1

s = 0
st = "{"
for k, v in d.items():
    st += str(v) + ", "
    if random.random() > 0.8:
        s += v
         
st = st[:len(st) - 2] + "}"
print(st)
print(s)
