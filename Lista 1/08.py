#custo por m2 = 100 reais
r=float(input('Raio da lata: '))
h=float(input('Altura da lata: '))
pi=3.14159
at=2*(pi*(r**2))+2*pi*r*h
valor=at*100
print(f'Valor de custo: {valor:.2f}')
