import math
h=float(input('Altrura da pirâmide: '))
a=float(input('Tamanho da aresta: '))
s3=math.sqrt(3)
ab=(3*(a**2)*s3)/2
v=ab*h/3
print(f'Volume da Pirâmide: {v:.2f} metros cúbicos')
