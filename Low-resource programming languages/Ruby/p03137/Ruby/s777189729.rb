n, m = gets.chomp.split(' ').map(&:to_i)
xs = gets.chomp.split(' ').map(&:to_i)

xs = xs.sort
ds = {}

(m-1).times do |i|
  ds[i] = xs[i+1] - xs[i]
end
ds = ds.sort { |(k1, v1), (k2, v2)| v2 <=> v1 }
pre = 0
res = 0
if n >= m
  puts 0
else
  n.times do |i|
    if i == 0
      res += xs.slice(pre..ds[i][0]).max - xs.slice(pre..ds[i][0]).min
      pre = ds[i][0]
    elsif i < n - 1
      res += xs.slice((pre + 1)..ds[i][0]).max - xs.slice((pre + 1)..ds[i][0]).min
      pre = ds[i][0]
    else
      res += xs.slice((pre + 1)..(xs.size - 1)).max - xs.slice((pre + 1)..(xs.size - 1)).min
    end
  end
  puts res
end
