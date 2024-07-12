.PHONY: build run clean

build:
	go build -v -o ./build/carbon_intensity

run:
	./build/carbon_intensity

clean:
	rm ./build/carbon_intensity

