n,k = $stdin.gets.split.map(&:to_i)
s = $stdin.gets.chomp.split("").map(&:to_i)
ichi = s[0]
ary=[]
result = []
now = s[0]
count = 0

n.times do |i|
  if now != s[i]
    ary.push(count)
    count = 0
  end
  count += 1
  now = s[i]
end

ary.push(count)

if ichi == 0
  sum = 0
  (k*2).times do |i|
	if !ary[i].nil?
      sum += ary[i]
    end
  end
  result.push(sum)
end

ary.size.times do |i|
  sum = 0
  i = i * 2 + 1 - ichi
  (k*2+1).times do |j|
    if !ary[i+j].nil?
      sum += ary[i+j]
    end
  end
  result.push(sum)
end

puts result.max 