until (x=gets.chomp.split.map(&:to_i))==[0,0] do puts x.sort.join(?\s) end