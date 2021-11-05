#for jが 1から n-1まで
#for iが 0から j-1まで
#maxv = (maxvと R[j]-R[i]のうち大きい方)
#minv = (minvと R[j]のうち小さい方)

n = int(input())
prev = 1000000000
ans = -1000000000
for i in range(n):
    r = int(input())
    if i:
        ans = max(ans, r - prev) #大きい方を格納　
    prev = min(prev, r) #小さい方を格納

print(ans)



