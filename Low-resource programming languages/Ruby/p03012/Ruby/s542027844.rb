n=gets.chomp.to_i;w=gets.chomp.split(" ").map{|a| a.to_i};ans=0;w.each{|n| ans+=n};for kugiri in 0..n-2 do ans=[ans,(w.slice(0..kugiri).sum-w.slice(kugiri+1..n-1).sum).abs].min end; puts ans;