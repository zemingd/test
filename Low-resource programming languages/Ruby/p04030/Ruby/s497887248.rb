puts 10.times.reduce(gets.chomp){|s, i| s.sub /(^|.)B/, ""}