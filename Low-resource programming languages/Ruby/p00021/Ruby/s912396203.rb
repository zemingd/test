n = gets.chop.to_i
delta = 1e-9

n.times do
  x1, y1, x2, y2, x3, y3, x4, y4 = gets.chop.split(/\s+/).map { |s| s.to_f }
  #p [x1, y1, x2, y2, x3, y3, x4, y4].join(", ")

  dxa = x2 - x1
  dya = y2 - y1
  dxb = x4 - x3
  dyb = y4 - y3

  d = (dxa * dyb - dya * dxb).abs

  puts (d <= delta) ? "YES" : "NO"
end