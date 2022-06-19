n = int(input())

res_l = [ 0 for i in range(n + 1)]
res_l[0] = 0
res_l[1] = 1
res_l[2] = 2
for i in range(3,n+1):
    res_l[i] = res_l[i-1] + res_l[i-2] # 세로 1개 추가 , 가로 2개 추가

print(res_l[n] % 10007)