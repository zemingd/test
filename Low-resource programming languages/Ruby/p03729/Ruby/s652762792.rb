puts gets.split.each_cons(2).all?{|a,b| a[-1] == b[0] } ? 'YES' : 'NO'
