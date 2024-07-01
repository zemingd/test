W, H, N = gets.split.map(&:to_i)

left, right, top, bottom = 0, W, H, 0

N.times do
	x, y, a = gets.split.map(&:to_i)
	
	if a == 1
		left = [left, x].max
	elsif a == 2
		right = [right, x].min
	elsif a == 3
		bottom = [bottom, y].max
	else
		top = [top, y].min
	end
end

if right - left > 0 && top - bottom > 0
	puts (right - left) * (top - bottom)
else
	puts 0
end
