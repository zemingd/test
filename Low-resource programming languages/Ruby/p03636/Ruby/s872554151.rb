S = gets.chomp.chars
ans = S[0] + (S.size - 2).to_s + S[-1]
puts ans
