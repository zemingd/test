n, m = gets.chomp.split.map(&:to_i)
demands = m.times.map do
  gets.chomp.split.map(&:to_i)
end.sort_by{|d| [d[0],d[1]]}

cnt = 0
max = demands[0][1]
(demands.size - 1).times do |i|
  if max <= demands[i + 1][0]
    cnt += 1
    max = Float::INFINITY
  end

  max = [max, demands[i + 1][1]].min
end

puts(cnt + 1)