n,*a=`dd`.split.map &:to_i;h=Hash.new 0;a.map{|e|h[e]+=1};p [(h.keys.select{|e|h[e]>3}.max||0)**2,h.keys.select{|e|h[e]>1}.max(2).reduce(:*)||0].max