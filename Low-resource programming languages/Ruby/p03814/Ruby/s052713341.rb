gets.chomp.tap { |str| 
  puts str.rindex("Z") - str.index("A") + 1 
}
