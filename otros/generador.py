import random
import sys

inicio = -10000
final = 10000

if len(sys.argv) < 2:
    print("python3 generador.py <Num. de elementos> [Semilla] [Cota minima] [Cota maxima]")
    sys.exit(0)
    
if len(sys.argv) > 2:
    random.seed(int(sys.argv[2]))
else:
    random.seed(0)

if len(sys.argv) > 3:
    inicio = int(sys.argv[3])

if len(sys.argv) > 4:
    final = int(sys.argv[4])

d = {}
x = 0
while x < int(sys.argv[1]):
    r = random.randint(inicio, final)
    if d.get(r, None) == None:
        k = {r : r}
        d.update(k)
        x += 1

s = 0
st = ""
for k, v in d.items():
    st += str(v) + ", "
    if random.random() > 0.8:
        s += v

st = st[:len(st) - 2]
print(s)
print(st)
