# プロ幼稚園に通う 2 人の子供がキャンディーの取り合いをしています。
# 3 個のキャンディーパックがあり、
# それぞれのパックにはキャンディーが a , b , c 個入っています。
# えび先生はこの 3 個のパックを、
# キャンディーの個数が等しくなるように 2 人に分けようとしています。
# そのような分け方が可能かどうかを判定してください。
# ただし、キャンディーをパックから取り出すことはできず、
# それぞれのパックをそのままどちらかの子供にあげる必要があります。

# 1 <= a, b, c <= 100

# 3個のキャンディーパックのそれぞれ入っているキャンディーの個数を取得
input = gets.chomp.split
a = input[0].to_i
b = input[1].to_i
c = input[2].to_i

# 1人目の子供に1個目のキャンディーパックを、
# 2人目の子供に2個目のキャンディーパックを渡す
div1 = a
div2 = b

# 持っているキャンディーの個数の少ない子供に、
# 3個目のキャンディーパックを渡す
if div1 > div2
  div2 += c
else
  div1 += c
end

# 結果を格納するための変数を宣言
result = String.new

# 2人の子供の持っているキャンディーの個数が等しいか判定
if div1 == div2
  result = "Yes"
else
  result = "No"
end

# 結果を出力
puts result