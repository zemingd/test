puts (((gets.chomp.split(' ').reduce(:+).to_i) % 4 == 0) ? 'YES' : 'NO')