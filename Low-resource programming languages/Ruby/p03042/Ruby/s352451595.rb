a=gets.chomp;e=false;YM=false;MY=false;puts "NA" if a[0..1] == "00" || a[2..3] == "00";e=true if a[0..1] == "00" || a[2..3] == "00";YM=true if a[0..1].to_i >= 1 && a[2..3].to_i >= 1 && a[2..3].to_i <= 12;MY=true if a[2..3].to_i >= 1 && a[0..1].to_i >= 1 && a[0..1].to_i <= 12;puts "AMBIGUOUS" if YM == true && MY == true;e=true if YM == true && MY == true;puts "YYMM" if YM == true && e == false;puts "MMYY" if MY == true && e == false;