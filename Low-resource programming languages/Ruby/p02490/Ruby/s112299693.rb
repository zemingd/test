while true
input = STDIN.gets.split(" ")
n = input.collect{|x| x.to_i}
if n[1] == 0 && n[0] == 0 then break end
if n[1] > n[0]
  puts "#{n[0]} #{n[1]}"
else
  puts "#{n[1]} #{n[0]}"
end
end