apt = Array.new( 4 ){Aarrey.new(3){Array.new(10.0)}}

gets
while( info = gets )
  b,f,r,v = info.split.map(&:to_i)
  apt[b-1][f-1][r-1] += v
end
apt.each.with_index(0) do |blbg,i|
  puts "#"*20 if i > 0
  blbg.each do |floor|
    puts " " + floor.join(" ")
  end
end