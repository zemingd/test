N, K = gets.chop.split.map(&:to_i)
S = gets.to_s.chomp

# if given string can be changed into 111...1 with k times operation, return True
# (10101, 1) -> False
# (10100, 2) -> True
def valid?(s, k)
    if count_zero_lump(s) <= k then
        return true
    else
        return false
    end
end

# count 0 lump
# (00100) -> 2
# (0) -> 1
# (1) -> 0
# (11001) -> 1
def count_zero_lump(s)
    if s==nil or s.length==0 then
        return 0
    end

    ret = 0
    state = false
    s.chars.each do |c|
        if c=="0" and state==false then
            ret = ret + 1
            state = true
        elsif c=="1" then
            state = false
        end
    end
    return ret
end

# Shakutori
head = 1
tail = 0
max_length = 0

while tail <= S.length and head <= S.length do
    if head == S.length then
        tail += 1
    elsif head == tail then
        head += 1
    elsif valid?(S.slice(tail, head+1), K) then
        head += 1
    else
        tail += 1
    end
    if valid?(S.slice(tail, head), K) and head - tail > max_length then
        max_length = head - tail
    end
end
p max_length