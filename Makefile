all:
	go build -buildmode=c-archive -o go.a main.go
	mv go.a ./C/
	mv go.h ./C/
	cd ./C/ && gcc -pthread -o main main.c go.a
	cd ./C/ && ./main

clean:
	rm ./C/go.a
	rm ./C/go.h
	rm ./C/main