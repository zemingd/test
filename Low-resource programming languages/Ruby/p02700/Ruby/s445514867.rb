A, B, C, D = gets.split.map(&:to_i) 

a = C - B
b = A - D

while b > 0 do
  a = a - B
  b = b - D
end

if b == 0
  puts "No"
elsif b < 0
  puts "No"
else
  puts "Yes"
end