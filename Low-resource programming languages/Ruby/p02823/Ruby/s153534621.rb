n, a, b = gets.split.map(&:to_i)
if (a - b).even?
  puts (a - b).abs / 2
else
  puts [[a - 1, b - 1].max, [n - a, n - b].max].min
end
