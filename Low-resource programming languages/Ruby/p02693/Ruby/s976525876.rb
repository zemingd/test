k = gets.to_i
a, b = gets.split.map(&:to_i)

(a..b).to_a.each do |i|
  if i % k == 0
    puts 'OK'
    exit
  end
end

puts 'NG'
