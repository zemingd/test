puts gets.to_i.times.map{gets.to_i}.inject(1){|a,b| a = a.lcm b}