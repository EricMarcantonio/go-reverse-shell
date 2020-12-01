EXE = rshell
SRC = ./src

windows:
	GOOS=windows go build -o dist/$(EXE)_win.exe -ldflags="-s -w" $(SRC)

macos:
	GOOS=darwin go build -o dist/$(EXE)_macos -ldflags="-s -w" $(SRC)

linux:
	GOOS=linux go build -o dist/$(EXE)_linux -ldflags="-s -w" $(SRC)

all: windows macos linux

clean:
	rm -f $(EXE)_win.exe $(EXE)_macos $(EXE)_linux