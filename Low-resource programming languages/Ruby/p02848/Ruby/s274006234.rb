n = gets.chomp.to_i
s = gets.chomp.split(//)

def alp_index(s,n)
    i = n % 26
    alp = "A B C D E F G H I J K L M N O P Q R S T U V W X Y Z".split.map(&:to_s)
    total = (i + alp.find_index(s).to_i + 1) % 26
    alp[total]
end

tmp = []
s.each do |s|
    tmp.push(alp_index(s,n).chomp)
end

puts tmp.join


