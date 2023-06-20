build:
	cd binary && go build -o bin main.go && cd -
	go build main.go

die-run: build
	./main die

timeout-run: build
	./main timeout

exit-run: build
	./main exit

ok-run: build
	./main ok

clean:
	rm main
	rm ./binary/bin
