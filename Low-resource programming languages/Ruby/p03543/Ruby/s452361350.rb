s = gets.chomp
if (s[0] == s[1] && s[1] == s[2] ) || (s[0] == s[1] && s[1] == s[3] ) || (s[0] == s[2] && s[2] == s[3] ) || (s[1] == s[2] && s[2] == s[3])
  puts "Yes"
else
  puts "No"
end
