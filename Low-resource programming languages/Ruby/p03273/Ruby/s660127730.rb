puts eval"[*$<]"+".grep(/#/).map(&:chars).transpose.map(&:join)"*2