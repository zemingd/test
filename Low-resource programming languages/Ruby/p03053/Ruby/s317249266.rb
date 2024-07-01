h, w = gets.split.map(&:to_i)
a = []
h.times do
  a << gets.chomp.chars
end
q = []
h.times do |i|
  w.times do |j|
    q << [i, j, 0] if a[i][j] == "#"
  end
end
ans = 0
yet = Array.new(h) { Array.new(w, true) }
d = [[-1, 0], [1, 0], [0, 1], [0, -1]]
until q.empty?
  i, j, cnt = q.shift
  yet[i][j] = false
  ans = cnt if cnt > ans
  d.each do |di, dj|
    ni, nj = i + di, j + dj
    next if ni < 0 || ni > h - 1
    next if nj < 0 || nj > w - 1
    q << [ni, nj, cnt + 1] if yet[ni][nj] && a[ni][nj] == "."
  end
end

p ans
