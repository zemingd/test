line = gets
line = line.split(" ")
def abs(n)
	if n<0
		-n
	else
		n
	end
end

def simple_calc(x,y)
	if abs(y)>abs(x)
		if y>0
			y-x
		else
			if x>0
				abs(y)-abs(x)+1 
			else
				abs(y)-abs(x)+2
			end
		end
	elsif abs(y)<abs(x)
		if y>0
			if x>0
				x-y+2
			else
				y-x
			end
		else
			if x>0
				abs(x)-abs(y)+1
			else
				abs(x)-abs(y)
			end
		end
	else
		if x==y
			0
		else
			1
		end
	end
end
print( simple_calc(line[0].to_i,line[1].to_i) )