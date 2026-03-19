n=int(input('Número: '))
a=n%3
b=n%5
if a==0 and b==0:
    print('O número é divisível')
else:
    print('O número não é divisível')