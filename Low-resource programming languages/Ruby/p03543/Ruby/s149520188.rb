arr = $stdin.gets.chomp.split(" ")
arr.map! do |a_j|
a_j.to_i
end

if arr[0] == arr[1] and arr[0] == arr[2]
  puts "Yes"
elsif  arr[1] == arr[2] and arr[1] == arr[3]
  puts "Yes"
elsif arr[0] == arr[1] and arr[0] == arr[2] and arr[0] == arr[3]
  puts "Yes"
else
  puts "No"
end