p`dd`.split.map(&:to_i)[1,50].combination(2).to_a.map{|i|i.inject(:*)}.inject(:+)