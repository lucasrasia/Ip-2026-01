nota=float(input('Nota: '))
if nota>=0 and nota<6: 
    print(f'Nota: {nota}, Conceito: D')
elif nota>=6 and nota<7.5: 
    print(f'Nota: {nota}, Conceito: C')
elif nota>=7.5 and nota<9: 
    print(f'Nota: {nota}, Conceito: B')
elif nota>=9 and nota<=10: 
    print(f'Nota: {nota}, Conceito: A')
else:
    print('Nota inválida')