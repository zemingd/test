a,b=`factor`.lines.map &:split;p (a&b).size+(a!=b ?1:0)