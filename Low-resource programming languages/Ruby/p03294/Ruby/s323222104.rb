def erastosthenes(n)
  array = (2..n).to_a
  threshold = Math.sqrt(n)
  count = 2
  while threshold > count
    array.reject!{|n| n != count and n % count == 0}
    count += 1
  end
  array
end

lines = $stdin.read
array = lines.split("\n")

def mod(m, arr)
  arr.map do |a|
    m % a
  end.inject(&:+)
end

N   = array[0].to_i
A   = array[1].split(" ").map(&:to_i).sort

PN  = erastosthenes(10**5)

ALT = A.map do |e|
  alt_e = e
  until PN.include?(alt_e)
    alt_e-=1
  end
  alt_e
end

SUM = ALT.uniq.inject(&:*)

ans = 0

for e in SUM..10**5
  f = mod(e, A)
  ans = [f,ans].max
end

p ans