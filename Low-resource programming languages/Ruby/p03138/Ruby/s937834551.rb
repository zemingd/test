n, k = gets.chomp.split(' ').map(&:to_i)
as = gets.chomp.split(' ').map(&:to_i)
max_digits = k.to_s(2).length
bits = max_digits.times.map{|d|
 n.times.map{|x| (x & (1 << d)) >> d}.reduce(:+)
}
min_bit = bits.min
x = bits.reverse.map{|b| b == min_bit ? 1 : 0}.join.to_i(2)
puts as.map{|n| x ^ n}.reduce(:+)