#python公式ドキュメント 2021/09/14
#https://docs.python.org/ja/3/tutorial/datastructures.html
#https://python.ms/iterator/#_1-%E3%82%A4%E3%83%86%E3%83%AC%E3%83%BC%E3%82%BF%E3%82%92%E8%A7%A6%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B%E3%80%82

def sample(num):
    if num < 2:
        return "2未満"
    elif num == 2:
        return "2です"
    elif num > 2:
        return "2を超えています"
for ans in map(sample, range(10)):
    #map(function, iterable) 逆にすると機能しないので注意

    print(ans)

#example2

a = [100, 200, 300]
a_iter = iter(a)
print(a_iter) 
print(next(a_iter))
print(next(a_iter))
print(next(a_iter))
print(next(a_iter))
#stopItereationとなる

