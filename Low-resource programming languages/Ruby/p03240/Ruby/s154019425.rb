nn = gets.to_i
xyh =[]
$stdin.each_line{|line|
xyh.push((line.split).map!(&:to_i))}
if nn == 1
  puts "#{xyh[0][0]} #{xyh[0][1]} #{xyh[0][2]}"
else
h = Array.new(nn)
for i in (0..100)
  for j in (0..100)
    for k in (0..nn-1)
      h[k] = (xyh[k][0] - i).abs + (xyh[k][1] - j).abs + xyh[k][2]
    end
    if h.uniq.size == 1
      puts "#{i} #{j} #{h[0]}"
    end
  end
end
end