s=gets;k=gets.to_i;r=s.chop.bytes.map{|e|x=123-e;(e=97;k-=x)if e>97&&k>=x;e};r[-1]+=k%26;puts r.map(&:chr)*''