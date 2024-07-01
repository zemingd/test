n = gets.to_i
ary = gets.split.map(&:to_i)
a_sort = ary.sort
if ary == a_sort
  puts "Yes"
  exit
end

swap = []
n.times do |i|
  swap.push(i) if ary[i] != a_sort[i]
end

if swap.size != 2
  puts "No"
  exit
end

ary[swap[0]],ary[swap[1]] = ary[swap[1]],ary[swap[0]]

if ary == a_sort
  puts "Yes"
else
  puts "No"
end




