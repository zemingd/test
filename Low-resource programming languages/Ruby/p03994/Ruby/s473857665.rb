s=gets;k=gets.to_i;r=s.bytes.map{|e|x=123-e;e<98||k<x ?e:(k-=x;97)};r[-2]+=k%26;puts r.map(&:chr)*''