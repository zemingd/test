S = gets.sub("\n", "").chars
T = gets.sub("\n", "").chars

# チェック用
chk = S.uniq

# 数が同じ文字
tmp = Hash.new
S.each do |c|
    break if tmp.keys.length == chk.length
    next if tmp.keys.include?(c)
    
    T.each do |d|
        next if tmp.values.include?(d)
        
        if S.count(c) == T.count(d)
            tmp[c] = d
        end
    end
end

tmp.each do |k, v|
    S.delete(k)
    T.delete(v)
end

# 各文字の数が同じ？(入れ替えで一致できる？)
S.sort!
T.sort!

eq = S == T

puts eq ? 'Yes' : 'No'