.PHONY: build clean

build:
	go build -o raytracer

clean:
	rm -f raytracer
