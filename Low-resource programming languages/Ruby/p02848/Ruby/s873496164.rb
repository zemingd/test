text = gets.chomp
n = gets.to_i

alphabets = ('A'..'Z').to_a
new_chars = []
text.each_char do |t|
  idx = alphabets.index(t)
  new_chars.push(alphabets[idx + n])
end

puts(new_chars.join(''))