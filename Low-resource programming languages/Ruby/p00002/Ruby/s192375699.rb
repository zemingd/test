puts $_.split.map(&:to_i).inject(&:+).to_s.size while gets
