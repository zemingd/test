n,s = gets.split(" ").map(&:to_i)
a = gets.split(" ").map(&:to_i)
d = Hash.new(0)
a.each {|i|
  d[i] += 1
}

a.each {|i|
  d[i] -= 1
  c = 0
  d.each_key {|j|
    c += d[j]*(d[j]-1)/2
  }
  p c
  d[i] += 1
}