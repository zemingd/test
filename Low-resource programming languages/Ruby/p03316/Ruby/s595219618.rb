puts gets.to_i%$_.chop.chars.map(&:to_i).inject(:+)==0?:Yes: :No