 #Atividade 01: Calculo de média

print('Cálculo de Média')
nota_1=float(input('Primeira nota: '))
nota_2=float(input('Segunda nota: '))
nota_3=float(input('Terceira nota: '))
media=(nota_1+nota_2+nota_3)/3
print(f'Média: {media:.2f}')
if media>=6:
    print('Aprovado')
else:
    print('Reprovado')
