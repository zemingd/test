N = gets.to_i
A = {}

N.times do
  
  a = gets.to_i
  
  if A[a].nil?
    A[a] = 1
  else
    A[a] += 1
  end
end

puts A.count{|k, v| v%2 == 1 }