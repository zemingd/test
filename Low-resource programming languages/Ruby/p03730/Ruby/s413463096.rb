a,b,c=gets.split.map &:to_i;
b.times{|i|if a%b*i%b==c;puts"YES";exit;end}
puts"NO"