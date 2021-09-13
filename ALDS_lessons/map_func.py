#python公式ドキュメント 2021/09/14
#https://docs.python.org/ja/3/tutorial/datastructures.html

def sample(num):
    if num < 2:
        return "2未満"
    elif num == 2:
        return "2です"
    elif num > 2:
        return "2を超えています"
for ans in map(sample, range(10)):
    #map(関数名、シーケンス) 逆にすると機能しないので注意

    print(ans)