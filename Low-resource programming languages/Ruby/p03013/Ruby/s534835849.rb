def fibo(fibo)
  if fibo == 0
    return n = 0
  end
 if fibo == 1
  return  n=1
 end

 return n = fibo(fibo-1) + fibo(fibo-2)
end

line = gets.split(' ').map(&:to_i)
n = line[0]
m = line[1]
a = Array.new(m+1)
box = Array.new(m+1)
a[0] =0
m.times do |i|
  a[i+1] = gets.to_i
  box[i] = fibo(a[i+1]-a[i]-1)
end
box[0] = fibo(a[1])
box[m] = fibo(n-a[m])

multi = 1
box.size.times do |i|
  multi *= box[i]
end

puts multi % 1000000007