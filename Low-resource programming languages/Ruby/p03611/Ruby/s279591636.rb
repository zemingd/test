n = gets.to_i
arr = gets.split(" ").map{|x|
	x.to_i
}

keys = []
arr.each do |x|
	keys << x-1
	keys << x
	keys << x+1
end
max = 0
cnt = keys.length
keys.uniq.each do |x|
	keys.delete(x)
	cnt = cnt - keys.length
	if max < cnt then
		max = cnt
	end
end
p max
