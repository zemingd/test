s = gets.chomp
t = gets.chomp

if t[0..s.size-1] == s && t.size <= 10
  puts 'Yes'
else
  puts 'No'
end