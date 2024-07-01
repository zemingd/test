# frozen_string_literal: true

n, m = gets.split.map(&:to_i)
a = gets.split.map(&:to_i)

min = a.inject(:+) / n

if a.count { |x| x >= min } >= m
  puts 'Yes'
else
  puts 'No'
end
