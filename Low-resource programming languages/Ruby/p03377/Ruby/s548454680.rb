puts gets.split.tap{|a,b,c| break (a..b).include?(c) } ? 'YES' : 'NO'