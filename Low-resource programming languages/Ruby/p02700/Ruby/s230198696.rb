a, b, c, d = gets.split(' ').map{|item| item.to_i}

while a > 0 && c > 0
  c = c - b
  if c <= 0
    puts "Yes"
    break
  end
  a = a - d
  if a <= 0
    puts "No"
    break
  end
end