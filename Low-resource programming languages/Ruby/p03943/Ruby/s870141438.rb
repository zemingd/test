puts (a=gets.split.map(&:to_i)).inject(&:+)==2*a.max ? :Yes: :No