k = gets.to_i
a,b = gets.split.map(&:to_i)
if a%k == 0 || b%k == 0
  puts "OK"
elsif a/k + 1 <= b/k
  puts "OK"
else
  puts "NG"
end