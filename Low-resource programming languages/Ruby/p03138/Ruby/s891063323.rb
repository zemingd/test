N, K = gets.split.map &:to_i
K_b = K.to_s(2)

Size = K_b.size

As = gets.split.map &:to_i

A_bs_tr = As.map{|a|
	b = a.to_s(2)
	("0"*([Size - b.size, 0].max) + b).chars
}.transpose

Best_ans = [*0...Size].map{|i|
	num_0 = A_bs_tr[i].count('0')
	num_0 <= (N - num_0) ? '0' : '1'
}.join.to_i(2)

X = [Best_ans, K].min

p As.map{|a| a^X }.inject(0, :+)


