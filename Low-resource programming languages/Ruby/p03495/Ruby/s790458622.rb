n,k=gets.to_s.split(" ").map{|s| s.to_i}
a=gets.to_s.split(" ").map{|s| s.to_i}
count=Array.new
for m in 0..n
  count[m]=a.count(m)
end
count.sort!.reverse!
print n-(count[0..k-1].inject(:+))