n = gets.chomp.split('').map(&:to_i)

if n[-1] == 2 || n[-1] == 4 || n[-1] == 5 || n[-1] == 7 || n[-1] == 9
  p "hon"
elsif n[-1] == 0 || n[-1] == 1 || n[-1] == 6 || n[-1] == 8
 p "pon"
else
 p "bon"
end