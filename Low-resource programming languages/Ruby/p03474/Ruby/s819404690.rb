puts gets.split == gets.split("-").map {|s| s.size.to_s} ? "Yes" : "No"