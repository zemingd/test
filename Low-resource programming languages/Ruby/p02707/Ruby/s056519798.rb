#復習用
n = gets.to_i
array = gets.split(" ").map(&:to_i)
hash = {}
for i in 1..n
  hash[i] = 0
end
array.each_with_index do |a|
  hash[a] += 1
end
puts hash.values