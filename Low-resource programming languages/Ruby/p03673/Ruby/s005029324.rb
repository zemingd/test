n = gets.to_i
a = gets.chomp.split(" ").map!{|i| i.to_i}
b = []
n.times do |j|
  if j % 2 == 0
    b.push(a[j])
  else
    b.unshift(a[j])
  end
end
if n % 2 == 0
  puts b.join(" ")
else
  puts b.reverse.join(" ")
end
