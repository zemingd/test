N = gets.to_i
S = gets.chomp
if !(S.match("#"))||!(S.match(".")) then puts 0; exit end
ans = 0
for i in 0...N-1
  if S[i..i+1] == "#." then ans += 1; S[i+1] = "#" end
end
puts ans