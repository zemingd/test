loop do
    x, y = gets.split.map(&:to_i)

    break if (x == 0 && y == 0)

    if x < y then
        puts "#{x} #{y}"
    else
        puts "#{y} #{x}"
    end
end