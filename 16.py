s=float(input('Salário: '))
if s<=300:
    r=s*1.5
    print(f'Salário com reajuste: {r:.2f}')
else:
    r=s*1.3
    print(f'Salário com reajuste: {r:.2f}')