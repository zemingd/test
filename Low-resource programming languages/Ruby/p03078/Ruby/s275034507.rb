x,y,z,k = gets.chomp.split(" ").map{|i|i = i.to_i}.sort
a = gets.chomp.split(" ").map{|i|i = i.to_i}.sort
b = gets.chomp.split(" ").map{|i|i = i.to_i}.sort
c = gets.chomp.split(" ").map{|i|i = i.to_i}.sort
#puts "a=#{a}"
#puts "b=#{b}"
#puts "c=#{c}"
result_arr = []
a.length.times do|a_i|
  b.length.times do|b_i|
    c.length.times do|c_i|
      next if (a_i*b_i*c_i)>k
      aa=a[a_i];bb=b[b_i];cc=c[c_i]
      result_arr << (aa+bb+cc)
    end
  end
end
result_arr.sort
k.times do|i|
  puts result_arr[result_arr.length-1-i]
end