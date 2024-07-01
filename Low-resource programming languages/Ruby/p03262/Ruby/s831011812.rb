N, X = gets.split.map(&:to_i)
towns = gets.split.map(&:to_i).sort
first_point = nil
diff_array = []

if towns.length == 1
  diff = towns[0].dup
  point_diff = towns[0].dup
else
  diff = (towns[-1] - towns[0]).dup
  point_diff = (towns[-1] - towns[0]).dup
end

towns.each_with_index{|i, idx|
  if X > i
    if diff > (X - i)
      first_point = i
      diff = (X - i)
    end
  else
    if diff > (i - X)
      first_point = i
      diff = (i - X)
    end
  end

  unless idx == 0
    diff_array << (towns[idx]-towns[idx-1])
  end
}

diff_array << diff
diff_array.sort!

def gcd(x, y)
  if y > x
    gcd(y, x)
  end

  r = x % y
  if r.zero?
    return y
  end

  gcd(y, r)
end

result = diff_array.reduce{|memo, item|
  gcd(memo, item)
}

puts result
