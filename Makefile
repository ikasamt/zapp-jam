
build:
	go build

clean:
	rm -rf example/app/zzz-*

example: build clean
	./zapp-jam example/app
