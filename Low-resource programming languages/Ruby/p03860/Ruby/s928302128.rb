s = gets.chomp

array = []
s.split(' ').each do |word|
  array << word[0]
end

p array.join('')
