n, q = gets.split.map(&:to_i)

t = Array.new(n) { [] }
(n-1).times do
  a, b = gets.split.map { |i| i.to_i - 1 }
  t[a] << b
  t[b] << a
end

c = Array.new(n) { 0 }
q.times do
  p, x = gets.split.map(&:to_i)
  c[p-1] += x
end

stk = Array.new(n)
stk[0] = [0, -1]
x = 1
until x == 0
  x -= 1
  i, p = stk[x]
  stk = stk[0...x]
  t[i].each do |j|
    next if j == p
    c[j] += c[i]
    stk[x] = [j, i]
    x += 1
  end
end

puts c.join(" ")