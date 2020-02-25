.PHONY: all run clean

CMD_PATH = ./cmd
CMD_DIRS=  $(wildcard $(CMD_PATH)/*)
BIN = $(patsubst $(CMD_PATH)/%, %, $(CMD_DIRS))

all: $(BIN)

$(BIN):	
	go build $(patsubst %, $(CMD_PATH)/%, $@)

run: 
	./compareto

clean:
	rm compareto