n = gets.split("").map(&:to_i)

if ((n[0] == 7) or (n[1]==7) or (n[2]==7))
  puts 'Yes'
else
  puts 'No'
end
