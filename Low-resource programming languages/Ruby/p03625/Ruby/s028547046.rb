n=gets.chomp.to_i
a=gets.split.map(&:to_i)

a=a.sort.reverse
freq, b={}, []
a.each do |e|
  if freq.has_key?(e)
    freq[e]+=1
    b.push(e) if freq[e]%2==0
    break if b.size==2
  else
    freq[e]=1
  end
end

if b.size==2
  puts b[0]*b[1]
else 
  puts 0
end