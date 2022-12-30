echo "Start"
echo "CheckTestFile"
cd server
go test -v
cd ..
echo "# Run go build -o TCPChat"
go build -o TCPChat
echo "Server is up and listening on port 8989. Please open new terminal and type <nc localhost 8989> to enter the chat" #"Press command and N to open terminal"
./TCPChat   