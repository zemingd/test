n,k,*a=`dd`.split.map &:to_i;45.downto(r=t=0){|i|z=1<<i;c=a.count{|x|x&z>0};r,t=t|z>k||n<=c*2?[r+c*z,t]:[r+(n-c)*z,t|z]};p r