a , b = gets.split(' ').map!(&:to_i)

if a >= b
  puts b.to_s * a
else
  a.to_s * b
end
