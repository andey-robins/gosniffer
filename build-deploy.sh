CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ GOOS=linux CGO_ENABLED=1 go build main.go
scp -i /Users/andey/andey-JMBP.pem main ubuntu@seekandyouwillbefound.org:/home/ubuntu
scp -i /Users/andey/andey-JMBP.pem -r website ubuntu@seekandyouwillbefound.org:/home/ubuntu
scp -i /Users/andey/andey-JMBP.pem startup.sh ubuntu@seekandyouwillbefound.org:/home/ubuntu
ssh -i /Users/andey/andey-JMBP.pem ubuntu@seekandyouwillbefound.org
