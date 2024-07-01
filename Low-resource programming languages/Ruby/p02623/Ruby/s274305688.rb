n,m,k = gets.chomp.split.map &:to_i
a = gets.chomp.split.map &:to_i
b = gets.chomp.split.map &:to_i

A, B = [0], [0]
n.times do |i|
  A.push a[i] + A[i]
end
m.times do |i|
  B.push b[i] + B[i]
end

ans, j = 0, m

(1..n).each do |i|
  break if A[i] > k
  while B[j] > k - A[i]
    j -= 1
  end
  ans = [ans, i+j].max
end

puts ans
