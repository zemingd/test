#!/usr/local/bin/ruby

line=gets
rec=line.split(" ") 
x=rec[0].to_i 
y=rec[1].to_i

def big(a,b)
  if a<b
   return a'<'b
  elsif a>b
   return a'>'b
  else
   return a'==b'
end

puts "#{big(x,y)}"