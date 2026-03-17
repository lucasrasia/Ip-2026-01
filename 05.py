 #Atividade 05: Conta de Água

conta=input('Conta: ')
consumo=float(input( 'Consumo de água: '))
tipo=input("Tipo de consumidor(C/I/R): ")
if tipo == 'R':
	valor1=consumo*0.05+5
	print(f'Valor da conta: {valor1}')
elif tipo == 'C':
	if consumo<=80:   #6,25 p/metro cubico
		valor2=6.25*consumo
		print (f'Valor da conta: (valor2)')
	else:
		valor2= 500 + 0.25*(consumo-80)
		print (f'Valor da conta: {valor2}')
elif tipo == 'I':
	if consumo<=100:  #8,00 p/metro cubico
		valor3=consumo*8
		print (f'Valor da conta: {valor3}')
	else:
		valor3=800+0.04*(consumo-100)
		print (f'Valor da conta: {valor3}')
else:
	print('Tipo de conta não encontrado')
