.PHONY: all run clean

all:
	go build

run: 
	./mk8dx_pareto

clean:
	rm mk8dx_pareto