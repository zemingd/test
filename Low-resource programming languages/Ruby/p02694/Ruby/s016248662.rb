#!/usr/bin/env ruby
N = gets.chomp.to_i
a = 0
n = 100
while n < N
  n  = (n * 1.01).to_i
  a += 1
end
puts a
