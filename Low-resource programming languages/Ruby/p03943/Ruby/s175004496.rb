puts gets.split.map(&:to_i).sort.reverse.reduce(:-) == 0 ? "Yes" : "No"