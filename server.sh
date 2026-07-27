PORT=5999

echo "Server starting on port $PORT"

while true; do
  printf 'HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 2\r\nConnection: close\r\n\r\nOK' | nc -l -p "$PORT" -q 1
done
