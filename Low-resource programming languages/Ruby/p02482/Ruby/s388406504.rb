a,b= gets.split(" ").map(&:to_i)

if a< b
	puts "#{a} < #{b}" 
elsif a >b
puts "#{a} > #{b}"
else
puts "#{a} == #{b}"
end