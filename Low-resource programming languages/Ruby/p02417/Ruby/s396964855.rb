al = Hash[*[[*'a'..'z'], Array.new(26, 0)].transpose.flatten]

while w = gets.chomp
  w.split('').each { |e| al[e] += 1 }
end

al.each { |a, n| puts "#{a} : #{n}" }