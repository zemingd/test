a,b=gets.split.map &:to_i;puts (s=gets.chop).size()==a+b+1&&s[a]==?-&&s=~/[0-9]+-[0-9]+/?:Yes: :No