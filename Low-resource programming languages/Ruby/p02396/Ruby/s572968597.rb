b = true
a = Array.new
count = 0

while b
	num = gets.to_i
	b = false if num == 0
	if b
		a.push num
	end
	count += 1
end

a.each{|num|
	puts num
}