n=float(input(''))
if n<1:
   print('Número inválido')
else:
   if n%1 == 0:
       soma=sum(1/n for n in range (1,int(n+1)))
       print(f'{soma:.6f}')
   else:
       print('Número invalido')

