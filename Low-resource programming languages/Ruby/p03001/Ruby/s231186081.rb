info=gets.chomp.split(" ").map(&:to_i)
def cal_s(x,y,vec,max_x,max_y)#vec...1:x 0:y
	if vec==1
		x2=max_x-x
		x=x2 if(x2<x)
		return x*max_y
	else
		y2=max_y-y
		y=y2 if(y2<y)
		return max_x*y
	end
end
s1=cal_s(info[2],info[3],1,info[0],info[1])
s2=cal_s(info[2],info[3],0,info[0],info[1])
if(s1>s2)
	print("#{s1} 0")
elsif(s1<s2)
	print("#{s2} 0")
else
	if(s1==0 or info[2]==info[3] or info[2]==(info[1]-info[3]).abs or info[3]==(info[0]-info[2]).abs)
		print("#{info[0]*info[1]*0.5} 0")
	else
		print("#{s1} 1")
	end
end
		