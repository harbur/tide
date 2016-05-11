all:
	go build github.com/harbur/tide/cmd/tide

cross:
	mkdir -p build
	gox --os windows --os linux --os darwin --arch 386 --arch amd64 github.com/harbur/tide/cmd/tide
	mv tide_{darwin,linux,windows}_* build

