io = DATA.eof? ? $stdin : DATA
n=io.gets.to_i
x=io.gets.to_i(2)
def popcount(x)
  b = 0
  while x > 0
    x &= x - 1
    b += 1
  end
  return b
end
def f(n,pc)
  return 0 if pc ==0
  return 1 if n % pc ==0
  rtn=0
  while (r=n % pc) > 0
    n=r
    pc=popcount(n)
    rtn+=1
  end
  rtn+1
end
p_tot=popcount(x)
n.downto(1) do |i|
  bit =  (1 << (i-1))
  pc=  x&bit==0 ? p_tot+1 : p_tot-1
  n = x ^ bit
  puts f(n,pc)
end
__END__
