n,k = gets.chomp.split.map(&:to_i)
hs = gets.chomp.split.map(&:to_i)
dp = Array.new(n,10**18)
dp[0] = 0
(n-1).times do|i|
  1.upto(k) do |j|
    if i+j < n && dp[i+j] > dp[i] + (hs[i] - hs[i+j]).abs
      dp[i+j] = dp[i] + (hs[i] - hs[i+j]).abs
    end
  end
end
puts dp[n-1]