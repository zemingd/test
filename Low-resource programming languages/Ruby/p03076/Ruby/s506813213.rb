t = []
m = []

5.times do |i|
  v = gets.chomp.to_i
  vm = v % 10
  m.push(vm == 0 ? 0: 10 - vm)
  t.push(((v + 9) / 10 * 10))
end

r = t.sum - m.max
puts r

