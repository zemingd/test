n,l = gets.strip.split.map(&:to_i)
numbers = n.times.map{ gets.chomp }.sort.join

puts numbers

