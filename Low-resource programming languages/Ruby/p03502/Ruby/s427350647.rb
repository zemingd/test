# frozen_string_literal: true

N = gets.to_i
if(N < 10)
    puts "Yes"
end
if(N < 100 && N >= 10) 
 a = (N % 10) + (N / 10)
elsif(N >= 100 && N < 1000)
 a = (N / 100) + ((N % 100) / 10) + (N % 10)
elsif(N >= 1000 && N < 10_000)
 a = (N / 1000) + (N % 1000) / 100 + (N % 100) / 10 + (N % 10)
elsif(N >= 10_000 && N < 100_000)
 a = (N / 100_000) + (N % 10_000) / 1000 + (N % 1000) / 100 + (N % 100) / 10 + (N % 10)
elsif(N >= 100_000 && N < 1_000_000)
 a = (N / 100_000) + (N % 100_000) / 10_000 + (N % 10_000) / 1_000 + (N % 1_000) / 100 + (N % 100) / 10 + (N % 10)
 elsif(N / 1_000_000 && N < 10_000_000)
 a = (N / 1_000_000) + (N % 1_000_000) / 100_000 + (N % 100_000) / 10_000 + (N % 10_000) / 1000 + (N % 1_000) / 100 + (N % 100) / 10 + (N % 10)
 elsif(N >= 10_000_000 && N < 100_000_000)
 a = (N / 10_000_000) + (N % 10_000_000) / 1_000_000 + (N % 1_000_000) / 100_000 + (N % 100_000) / 10_000 + (N % 10_000) / 1_000 + (N % 1_000) / 100 + (N % 100) / 10 + (N % 10)
 elsif(N == 100_000_000)
    a = 1
end
if N % a == 0 && N >= 10
    puts "Yes"
elsif N % a != 0 && N >= 10
    puts "No"
end
end
if N % a == 0
    puts "Yes"
elsif N % a != 0
    puts "No"
end