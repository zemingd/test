a=gets(p).split.map &:to_i;b=[0];2.upto(a.shift){|i|b<<(1..a[0]).map{|j|j<i ?b[-j]+(a[i]-a[i-j]).abs: 9**9}.min};p b[-1]