a=?1;b=?0;x=y=0;gets.chars{|c|c!=a&&x+=1;c!=b&&y+=1;b,a=a,b};p~-[x,y].min