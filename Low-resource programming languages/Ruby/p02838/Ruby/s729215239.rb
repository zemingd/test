#require 'pp'

n = gets.chomp!.to_i
a_n = gets.split.map(&:to_i)

digit_max = 61
memo = Array.new(digit_max, 0)

n.times do |i|
    tmp = a_n[i]
    dig = 0
    while tmp != 0
        if tmp & 1 == 1
            memo[dig] += 1
        end
        tmp >>= 1
        dig += 1
    end
end