_n = gets.to_i
a = gets.chomp.split.map(&:to_i)

add = a.map { |i| i + 1 }
sub = a.map { |i| i - 1 }
ans = a + add + sub

puts ans.group_by(&:itself).transform_values(&:size).values.max