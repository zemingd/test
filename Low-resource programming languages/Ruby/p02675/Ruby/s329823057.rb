n = gets.chomp

puts case n[n.length - 1]
when '2', '4', '5', '7', '9' then 'hon'
when '0', '1', '6', '8'      then 'pon'
when '3'                     then 'bon'
end
