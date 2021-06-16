CC = g++
CFLAGS = -Wall -g
OUTDIR = bin/

all: directories bubblesort insertionsort

bubblesort: bubbleSortInts.o bubbleSortPointers.o bubbleSortStructs.o
insertionsort: insertionSortInts.o insertionSortPointers.o insertionSortStructs.o

bubbleSortInts.o:
	$(CC) $(CFLAGS) -c bubbleSort/source/bubbleSort.cpp -o $(OUTDIR)obj/bubbleSortInts.o

bubbleSortPointers.o:
	$(CC) $(CFLAGS) -c bubbleSort/source/bubbleSortPointers.cpp -o $(OUTDIR)obj/bubbleSortPointers.o

bubbleSortStructs.o:
	$(CC) $(CFLAGS) -c bubbleSort/source/bubbleSortStructs.cpp -o $(OUTDIR)obj/bubbleSortStructs.o

insertionSortInts.o:
	$(CC) $(CFLAGS) -c insertionSort/source/insertionSort.cpp -o $(OUTDIR)obj/insertionSortInts.o

insertionSortPointers.o:
	$(CC) $(CFLAGS) -c insertionSort/source/insertionSortPointers.cpp -o $(OUTDIR)obj/insertionSortPointers.o

insertionSortStructs.o:
	$(CC) $(CFLAGS) -c insertionSort/source/insertionSortStructs.cpp -o $(OUTDIR)obj/insertionSortStructs.o

directories:
	mkdir -p $(OUTDIR)
	mkdir -p $(OUTDIR)/obj/

clean:
	rm -rf bin/*
