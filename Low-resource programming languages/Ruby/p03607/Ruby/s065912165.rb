puts readlines.drop(1).map(&:to_i).each_with_object(Hash.new(false)){|a, h| h[a] = !h[a]}.values.count(true)