s = gets.split(//).map(&:to_s)

hash = {"a"=>1,"b"=>1,"c"=>1}
shash = Hash.new
s.each do |string|
    string = string.to_s
    shash[string] = 1
end

if hash == shash
    puts "Yes"
else
    puts "No"
end