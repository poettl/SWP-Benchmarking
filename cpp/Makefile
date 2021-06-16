CC = g++
CFLAGS = -Wall -g
OUTDIR = bin/
OUTDIR_OBJ = $(OUTDIR)obj/

all: directories main
	$(CC) $(CFLAGS) -o $(OUTDIR)program $(OUTDIR_OBJ)bubbleSortInts.o $(OUTDIR_OBJ)bubbleSortPointers.o $(OUTDIR_OBJ)bubbleSortStructs.o $(OUTDIR_OBJ)insertionSortInts.o $(OUTDIR_OBJ)insertionSortPointers.o $(OUTDIR_OBJ)insertionSortStructs.o $(OUTDIR_OBJ)radixSortInts.o $(OUTDIR_OBJ)radixSortPointers.o $(OUTDIR_OBJ)radixSortStructs.o $(OUTDIR_OBJ)main.o

main: bubblesort insertionsort radixsort
	$(CC) $(CFLAGS) -c main.cpp -o $(OUTDIR_OBJ)main.o

bubblesort: bubbleSortInts.o bubbleSortPointers.o bubbleSortStructs.o
insertionsort: insertionSortInts.o insertionSortPointers.o insertionSortStructs.o
radixsort: radixSortInts.o radixSortPointers.o radixSortStructs.o

bubbleSortInts.o:
	$(CC) $(CFLAGS) -c bubbleSort/source/bubbleSort.cpp -o $(OUTDIR_OBJ)bubbleSortInts.o

bubbleSortPointers.o:
	$(CC) $(CFLAGS) -c bubbleSort/source/bubbleSortPointers.cpp -o $(OUTDIR_OBJ)bubbleSortPointers.o

bubbleSortStructs.o:
	$(CC) $(CFLAGS) -c bubbleSort/source/bubbleSortStructs.cpp -o $(OUTDIR_OBJ)bubbleSortStructs.o

insertionSortInts.o:
	$(CC) $(CFLAGS) -c insertionSort/source/insertionSort.cpp -o $(OUTDIR_OBJ)insertionSortInts.o

insertionSortPointers.o:
	$(CC) $(CFLAGS) -c insertionSort/source/insertionSortPointers.cpp -o $(OUTDIR_OBJ)insertionSortPointers.o

insertionSortStructs.o:
	$(CC) $(CFLAGS) -c insertionSort/source/insertionSortStructs.cpp -o $(OUTDIR_OBJ)insertionSortStructs.o

radixSortInts.o:
	$(CC) $(CFLAGS) -c radixSort/source/radixSort.cpp -o $(OUTDIR_OBJ)radixSortInts.o

radixSortPointers.o:
	$(CC) $(CFLAGS) -c radixSort/source/radixSortPointers.cpp -o $(OUTDIR_OBJ)radixSortPointers.o

radixSortStructs.o:
	$(CC) $(CFLAGS) -c radixSort/source/radixSortStructs.cpp -o $(OUTDIR_OBJ)radixSortStructs.o

directories:
	mkdir -p $(OUTDIR)
	mkdir -p $(OUTDIR_OBJ)

clean:
	rm -rf bin
