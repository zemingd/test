# 構造化プログラミング

# goto 文は、C/C++言語などで使うことのできる文で、実行されると無条件に指定されたラベルに処理が飛びます。
# 例えば goto CHECK_NUM; という文が実行されると、プログラムの中で CHECK_NUM: と書かれた行に処理が移ります。
# この機能を使って、繰り返し処理なども実現することができます。

# 一方、goto 文は自由度が高すぎ、プログラムの可読性に影響するため、可能な限り使わないことが推奨されています。

# 次のC++言語のプログラムと同じ動作をするプログラムを作成してください。
# ただし、goto 文は使わないで実装してみましょう。 

# (意訳) 標準入力によって数字Nが入力されたとき、
#	1からNまでの数字のうち、3の倍数または末尾が3の数全てを出力しなさい。

# Input 30
# Output 3 6 9 12 13 15 18 21 23 24 27 30

# 標準入力
N = gets.chomp.to_i

# 解答を入れる配列
ans_arr = []

# カウント
for i in 1..N
	if i % 3 == 0
		ans_arr += [i]
	elsif i / 10 == 3
		ans_arr += [i]
	end
end

puts ans_arr.join("\s")