N, K = gets.split(' ').map(&:to_i)
l = gets.split(' ').map(&:to_i)

p l.sort.slice(l.length - K, K).inject(&:+)