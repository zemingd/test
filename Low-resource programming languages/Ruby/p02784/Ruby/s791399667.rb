h, n = gets.split.map(&:to_i)
a = gets.split.map(&:to_i)
s = a.inject(:+)
if h =< s
  puts "Yes"
else
  puts "No"
end
