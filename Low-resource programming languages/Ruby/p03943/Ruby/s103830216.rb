input = gets.split.map(&:to_i).sort
p input
if input[0] == input[1] + input[2] || input[0] + input[1] == input[2]
  puts 'Yes'
else
  puts 'No'
end