X,K,D = gets.split.map(&:to_i)

if (X < 0)
  if (X + K*D <= 0)
    puts (X + K*D).abs
  else
    if (K - (-X.to_f/D).ceil).even?
      puts (X + (-X.to_f/D).ceil*K).abs
    else
      puts (X + (-X.to_f/D).ceil*K - K).abs
    end
  end
else
  if (X - K*D > 0)
    puts (X - K*D).abs
  else
    if (K - (X.to_f/D).ceil).even?
      puts (X - (X.to_f/D).ceil*K).abs
    else
      puts (X - (X.to_f/D).ceil*K + K).abs
    end
  end
end
