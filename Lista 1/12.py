# 3horas = 10 reais
# 5 reais para cada hora ente as 3 horas
#ex 5 horas= 3h + 2h= 10 + 10=20 reais
#   6 horas= 2x3h=20 reais
a=float(input('Quantidade de horas: '))
b=a%3
valor=((a-b)/3)*10 + (b)*5
print(f'Valor a pagar: {valor:.2f}')

