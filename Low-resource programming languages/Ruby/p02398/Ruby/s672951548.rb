a, b, c = gets.split.map(&:to_i)
count = 0

(a..b).each do |i|
  count += 1 if c % i == 0
end

puts count