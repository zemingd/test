proc{|s| puts s.to_i%s.chomp.chars.map(&:to_i).inject(:+)==0 ? 'Yes' : 'No'}.(gets)