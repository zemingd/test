s=gets.chomp
t=gets.chomp
h1 = {}
h2 = {}
0.upto(s.size-1) do |i|
    next if s[i] == t[i]
    if h1.has_key?(s[i])
        if h1[s[i]] != t[i]
            puts "No"
            exit
        end
    elsif h2.has_key?(t[i]) 
        if h2[t[i]] != s[i]
            puts "No"
            exit
        end
    else
        h1[s[i]] = t[i]
        h2[t[i]] = s[i]
    end
end

puts "Yes"