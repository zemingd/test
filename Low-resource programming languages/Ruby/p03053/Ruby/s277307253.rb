require 'matrix'

M = 999999
h, w = gets.chomp!.split(" ").map!(&:to_i)
a_hw = Array.new(h) {|index| gets.chomp.chars } 
max_hw = Array.new(h) {|index| Array.new(w, M) }
 
black = []
a_hw.each_with_index do |arr, h2|
    arr.each_with_index do |ch, w2|
        if ch == "#"
            black.unshift([h2, w2, 0])
            max_hw[h2][w2] = 0
        end
    end
end
 
#i = 0
 
while black.empty?.!
    h2, w2, count = black.shift
    count += 1

    if h2 != 0 
        if max_hw[h2-1][w2] == M
            max_hw[h2-1][w2] = count
            black.push([h2 - 1, w2, count])
        end
    end
    if h2 != h - 1
        if max_hw[h2 + 1][w2] == M
            max_hw[h2+1][w2] = count
            black.push([h2 + 1, w2, count])
        end
    end
    if w2 != 0 
        if max_hw[h2][w2 - 1] == M
            max_hw[h2][w2 - 1] = count
            black.push([h2, w2 - 1, count])
        end
    end
    if w2 != w - 1
        if max_hw[h2][w2 + 1] == M
            max_hw[h2][w2 + 1] = count
            black.push([h2, w2 + 1, count])
        end
    end
    i += 1
end
 
#puts i
#puts max_hw.map {|x| x.join(",") }
puts max_hw.map{|x| x.max }.max
