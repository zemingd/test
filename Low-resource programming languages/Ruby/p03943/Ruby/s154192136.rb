s=gets.split.map(&:to_i).sort
puts s[2]==s[0]+s[1] ? :Yes: :No
