a,b,c = gets.chomp.split(' ').map{|n| n}

if a[-1] == b[0] && b[-1] == c[0]
  puts "YES"
else
  puts "NO"
end
