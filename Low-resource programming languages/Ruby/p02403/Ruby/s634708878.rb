while true
    h, w = gets.chomp.split.map(&:to_i)
    break if h == 0 && w == 0

    h.times{
        print "#" * w
        puts
    }
    puts
end