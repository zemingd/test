n = gets.to_i
s = gets.chomp

if n%2==0 && s[0..n/2-1]==s[n/2..-1]
  puts "Yes"
else
  puts "No"
end