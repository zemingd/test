n = gets.to_i
a = gets.split.map{|m|m.to_i}
temp_color = ["hai","cha","midori","mizu","ao","ki","daidai","aka"]
color = Array.new
extra_color = Array.new
a.each do |rate|
    if rate < 400
        color.push("hai") if !(color.include?("hai"))
    elsif rate < 800
        color.push("cha") if !(color.include?("cha"))         
    elsif rate < 1200
        color.push("midori") if !(color.include?("midori"))         
    elsif rate < 1600
        color.push("mizu") if !(color.include?("mizu")) 
    elsif rate < 2000
        color.push("ao") if !(color.include?("ao"))        
    elsif rate < 2400
        color.push("ki") if !(color.include?("ki"))        
    elsif rate < 2800
        color.push("daidai") if !(color.include?("daidai"))
    elsif rate < 3200
        color.push("aka") if !(color.include?("aka"))         
    else
        if color.length+extra_color.length != temp_color.length
            temp_color.each do |c|
                if !(extra_color.include?(c))
                    extra_color.push(c)
                    break
                end
            end
        end
    end 
end
color.push("aka") if color.length == 0
print color.length.to_s+" "+(color.length+extra_color.length).to_s+"\n"