p [*$<][1..-1].reduce{|a,b|b.to_i.lcm(a)}