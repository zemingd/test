s = gets.strip.split("")
n = 700
for i in 0..2 do
  if s[i] == "o"
    n += 100
  end
end
  print n