all:
	dep
	test

dep:
# writeing some required packages
#	go get -vtd 

test:
	# test backend
	go test -v ./back/...
