 #Atividade 03: Concatenação inteira

numero1=input('')
numero2=input('')
numero3=input('')
if int(numero1)>9:
  print('Dígito Inválido')
elif int(numero2)>9:
  print('Dígito Inválido')
elif int(numero3)>9:
  print('Dígito Inválido')
else:
  conta=(int(numero1)*100)+(int(numero2)*10)+(int(numero3))
  conta2=conta**2
  print(f'{conta}, {conta2}')
