#ref: https://zenn.dev/joeee/articles/d9bd0983601c4d

n = gets.to_i
a = gets.split.map(&:to_i)

def insertion_sort(n, a)
  puts a.join(" ")

  (1...n).each do |i|
    v = a[i]
    j = i - 1
    # j が配列の範囲内かつ左の値がvより大きい
    while j >= 0 && a[j] > v
      a[j + 1] = a[j]
      j -= 1
    end
    a[j + 1] = v
    puts a.join(" ")
  end
end

insertion_sort(n, a)
