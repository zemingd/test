n = gets.to_i

b = gets.split(' ').map(&:to_i)
a = b.sort
#puts a
carry1 = Array.new(n+1,0)
#line = gets.split(' ').map(&:to_i)
carry2 = Array.new(n+1,0)
ans = 0
count = 1
for i in 0..n-2
  if a[i] == a[i+1]
    count +=1
  else
    if count > 1
      carry1[a[i]] = count * (count-1) /2
      carry2[a[i]] = (count-1) * (count-2) /2
    end
    count = 1
  end
end

if count > 1
  carry1[a[i]] = count * (count-1) /2
  carry2[a[i]] = (count-1) * (count-2) /2
end

ans = 0

for i in 0..n-1
  ans += carry1[i]
end

for i in 0..n-1
  puts ans - carry1[b[i]] + carry2[b[i]]
end
