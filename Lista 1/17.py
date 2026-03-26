a=int(input(''))
b=int(input(''))
if a%2==0:
    c=a+2*(b-1)
    i=a
    while i<=c:
        print(i)
        i+=2
else:
    print('O primeiro número não é par')
    