200.times do puts gets.split.map(&:to_i).inject { |s, n| s+=n }.to_s.size  end