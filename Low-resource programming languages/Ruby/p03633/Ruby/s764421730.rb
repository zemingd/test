n = gets.chomp.to_i

t = []

n.times do
  t << gets.chomp.to_i
end

puts t.inject(:lcm)
