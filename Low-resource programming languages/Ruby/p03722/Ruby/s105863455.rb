(n,),*A=$<.map{|l|l.split.map &:to_i};d=0,0,*[-9e99]*n;$:.map{E=d[-2];A.map{A.map{|a,b,c|d[b]=c if d[b]<c+=d[a]}}};puts E<d[-2]?:inf:E