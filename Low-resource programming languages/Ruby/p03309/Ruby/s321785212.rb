n = gets.chomp.to_i
as = gets.strip.split.map(&:to_i)
ass = as.map.with_index{|x,index| x -(index+1) }
med = ass.sort[n/2]
tmp = ass.map{|x| (x-med).abs}
sum = tmp.inject(:+)
puts sum
