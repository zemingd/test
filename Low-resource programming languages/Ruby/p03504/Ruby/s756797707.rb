n,c=gets.chomp.split(" ").map(&:to_i)

st=tt=ans=tmp=0
s=[], t=[], j=[]

n.times do |i|
  si,ti,c=gets.chomp.split(" ").map(&:to_i)
  s[i]=[si, c]
  t[i]=[ti, c]
end

s=s.sort
t=t.sort

s.each do |i|
  while t[tt] && i[0] <= t[tt][0] do
	if  i[0] == t[tt][0] && i[1] == t[tt][1] then
  	  t.delete_at(tt)
	  s.delete(i)
	end
    tt+=1
  end
end
tt=0

s.each do |i|
  i[0]-=1 if i
end

j=s+t
j=j.sort

j.each do |i|
  while s[st] && s[st][0] == i[0] do
	tmp+=1
	st+=1
  end
  while t[tt] && t[tt][0] == i[0] do
	tmp-=1
	tt+=1
  end  
  ans = [ans, tmp].max
end

puts ans
