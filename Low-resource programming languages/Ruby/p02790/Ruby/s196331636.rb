a, b = gets.chomp.split.map(&:to_i)

result1 = a.to_s * b
result2 = b.to_s * a

puts result1 if result1.to_i < result2.to_i
puts result2