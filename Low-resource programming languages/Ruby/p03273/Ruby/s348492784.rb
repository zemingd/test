gets;m=$<.map &:chars;eval"m=m.select{|s|s.find{|c|c=='#'}}.transpose;"*2;puts m.map{|s|s*''}