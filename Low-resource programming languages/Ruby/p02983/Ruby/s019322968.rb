li,ri=gets.chomp.split(" ").map(&:to_i)
l= li%2019
r= ri%2019
data = (l..r).to_a.combination(2).map do |ij|
  i=ij[0]     
  j=ij[1]
  i*j % 2019
end
puts data.min
