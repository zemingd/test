n = gets.to_i
srray = gets.split("").map(&:to_s)
answer = []
answer.push(srray.count("."))
i=0
brray = []
crray = srray
answer.push(brray.count("#")+crray.count("."))
while i<n
  brray.push(crray.shift)
  # j=0
  # while j<i+1
  #   brray.push(srray[j])
  #   j+=1
  # end
  # while j>i && j<n
  #   crray.push(srray[j])
  #   j+=1
  # end
  answer.push(brray.count("#")+crray.count("."))
  i+=1
end
puts answer.min