a,b,c = gets.split.map(&:to_i)
puts a-b >= c ? 0 : c-(a-b)