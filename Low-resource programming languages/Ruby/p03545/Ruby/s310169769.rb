s=gets.chomp.chars;%w(+ -).repeated_permutation(3){|_|a=s.zip(_)*'';(puts a;exit)if eval(a)==7}