salario=float(input('Salário miínimo: '))
kw=float(input('Quantidade de kw gasto: '))
ckw=salario*0.007  
cc=kw*ckw
cd=cc*0.9
print(f'Custo por kW: R$ {ckw:.2f}')
print(f'Custo do consumo: R$ {cc:.2f}')
print(f'Custo com desconto: R$ {cd:.2f}')

