n=gets.to_i
button=n.times.map{gets.to_i}
count=0
last=button[0]
n-1.times.each do
    last=button[last-1]
    if last==2
    puts count
    exit
    end
    count+=1
end
puts -1