puts (gets.chomp.split.map(&:to_i).inject(:+)<22?"win":"bust")