n = gets.chomp.to_i
a = gets.chomp.split(" ").map!{|item| item.to_i}
num = 1
ans = 1
a.sort{|a,b| (-1)*(a <=> b)}

for num in 1..n do
    ans = ans * a[num-1]
    num = num + 1
    if ans > 10**18
        ans = -1
        break
    end
end
puts ans