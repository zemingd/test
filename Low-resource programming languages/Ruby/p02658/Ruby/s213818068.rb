n = gets.to_i
a = gets.split.map(&:to_i)
ans = 1
if a.include?(0)
	ans = 0
else
		a.each do |i|
		ans *= i
	 if ans > 10**18
	 		ans = "-1"
	 		break
	end
	end
end
puts ans