N=gets.to_i;T,A=gets.split.map &:to_i;H=gets.split.map &:to_i;p (1..N).min_by{|x|(A*1000-T*1000+H[x-1]*6).abs};