w,h,x,y,r = gets.chomp.split.map(&:to_i)
xmax = w - r
ymax = h - r
puts r <= x&&x <= xmax&&r <= y&&y<=ymax ? :Yes : :No