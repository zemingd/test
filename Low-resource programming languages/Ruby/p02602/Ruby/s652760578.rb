n,k = gets.chomp.split.map(&:to_i)
as = gets.chomp.split.map(&:to_i)
(n-k).times do |i|
  if as[i] < as[i + k]
    puts "Yes"
  else
    puts "No"
  end
end