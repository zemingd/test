def calcs(x,y,ax,ay,bx,by) ((ax-x)*(by-y)-(bx-x)*(ay-y))/2.0 end
def inpf() a=gets.chomp.split(" ").map(&:to_f)end
def inps() a=gets.chomp.split(" ")end  
def copy(a) Marshal.load(Marshal.dump(a)) end
def factorial(n)(n < 2)? 1 : (2..n).inject(:*) end
def scount(a) b=na(a.max+1);a.each{|n|b[n]+=1};return b end
def na(n,d=0) Array.new(n,d)end
def na2(n,m,d=0) Array.new(n){Array.new(m,d)}end
def na3(n,m,l,d=0) Array.new(n){Array.new(m){Array.new(l,d)}}end
def bit(n) n.to_s(2).split("").map(&:to_i) end
def inp() a=gets.chomp.split(" ").map(&:to_i)end
h,w = inp
t = []
h.times{t << gets.chomp.split("")}
ngh = na(h,false)
ngw = na(h,false)
h.times do |i|
  ngh[i] = true unless(t[i].include?("#"))
end
tz = copy(t.transpose)
w.times do |i|
  ngw[i] = true unless(tz[i].include?("#"))
end
(h).times do |i|
  next if(ngh[i])
  w.times do |j|
    print t[i][j] unless(ngw[j])
  end
  puts
end