puts (gets.split.map(&:to_i).all?(&:odd?)) ? "Yes" : "No"