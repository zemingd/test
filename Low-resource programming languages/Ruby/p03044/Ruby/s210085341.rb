N = gets.chomp.to_i
result = Array.new(N,nil)
 
 
inputs = []
(N - 1).times do
  line = gets.chomp.split(" ").map(&:to_i)
  inputs << line
end
 
inputs = inputs.sort_by{|item| item[0]}
 
result[0] = 0

inputs.each do |line|
    if line[2] % 2 == 0 
      result[line[1]-1] = result[line[0]-1]
    else
      result[line[1]-1] = (result[line[0]-1]+1)%2
    end
end
 
puts result