puts gets.split.join == gets.split("-").map {|s| s.size.to_s}.join ? "Yes" : "No"