a, b, c, d = gets.chomp.split(' ').map(&:to_i)

if a < c
  if c < b
    if b < d
      print "#{b-c}"
    else
      print "#{d-c}"
    end
  else
    print 0
  end
else
  if a < d
    print "#{d-a}"
  else
    print 0
  end
end