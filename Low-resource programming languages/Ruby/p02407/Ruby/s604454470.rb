n = gets.to_i
arr = gets.split.map(&:to_i)

arr_re = arr.reverse
arr_re.each {|n|
  if n == arr_re.last
    print "#{n}"
    break
  end
  print "#{n} "
}
puts