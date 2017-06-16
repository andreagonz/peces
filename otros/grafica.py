import matplotlib.pyplot as plt

file = open("../archivero/6.exp","r")
lst = file.readlines() 
file.close() 

x = []
for i in range(50):
    x.append(int(i))
    
y = []
for s in lst:
    y.append(int(s))


plt.title('6.ss, params1.txt')    
plt.yticks(y)
plt.plot(x, y)
plt.ylabel('Diferencia entre suma buscada y mejor suma encontrada')
plt.xlabel('Semilla')
plt.show()
