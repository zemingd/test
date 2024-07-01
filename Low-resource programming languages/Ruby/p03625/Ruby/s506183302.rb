n = gets.chomp.to_i
sticks = gets.chomp.split(" ").map(&:to_i)

lengths = sticks.uniq.sort{|a,b| b<=>a}

side = nil

lengths.each do |l|
  num = sticks.count(l)
  if num >= 4
    puts l*l
    exit
  end

  if num >= 2
    unless side.nil?
      puts side * l
      exit
    end
    side ||= l
  end
end

puts 0