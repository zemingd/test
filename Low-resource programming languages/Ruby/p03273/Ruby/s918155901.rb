#天地はすごい！！！
h,w = gets.strip.split.map(&:to_i)
ary = []
h.times do
  a = gets.strip.to_s
  b = a.split("").map(&:to_s)
  unless b.all?(".")
    ary << b
  end
end
arys = ary.transpose
w.times do |i|
  if arys[i].all?(".")
    arys.delete_at(i)
  end
end
newary = arys.transpose
newary.size.times do |k|
  puts newary[k].join("")
end