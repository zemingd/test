n=gets.chomp.to_i
A=gets.chomp.split.map(&:to_i)
y=A.inject(:+)
x=0
ans=(y-2*A[0])

A.each do |e|
   x+=e
   y-=e
   ans=[ans, (y-x).abs].min
end
puts ans