n,m,*a=`dd`.split.map(&:to_i)
puts a.sort[-m]<a.inject(:+)/4*m ?:No: :Yes
