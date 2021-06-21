CC = g++
CFLAGS = -Wall -g
OUTDIR = bin/
OUTDIR_OBJ = $(OUTDIR)obj/

all: all_base all_pointers all_structs all_structs_cache

all_base: directories base
	$(CC) $(CFLAGS) -o $(OUTDIR)base-benchmark $(OUTDIR_OBJ)bubbleSortInts.o $(OUTDIR_OBJ)bubbleSortPointers.o $(OUTDIR_OBJ)bubbleSortStructs.o $(OUTDIR_OBJ)insertionSortInts.o $(OUTDIR_OBJ)insertionSortPointers.o $(OUTDIR_OBJ)insertionSortStructs.o $(OUTDIR_OBJ)radixSortInts.o $(OUTDIR_OBJ)radixSortPointers.o $(OUTDIR_OBJ)radixSortStructs.o $(OUTDIR_OBJ)base.o

all_pointers: directories pointers
		$(CC) $(CFLAGS) -o $(OUTDIR)pointer-benchmark $(OUTDIR_OBJ)bubbleSortInts.o $(OUTDIR_OBJ)bubbleSortPointers.o $(OUTDIR_OBJ)bubbleSortStructs.o $(OUTDIR_OBJ)insertionSortInts.o $(OUTDIR_OBJ)insertionSortPointers.o $(OUTDIR_OBJ)insertionSortStructs.o $(OUTDIR_OBJ)radixSortInts.o $(OUTDIR_OBJ)radixSortPointers.o $(OUTDIR_OBJ)radixSortStructs.o $(OUTDIR_OBJ)pointers.o

all_structs: directories structs
		$(CC) $(CFLAGS) -o $(OUTDIR)struct-benchmark $(OUTDIR_OBJ)bubbleSortInts.o $(OUTDIR_OBJ)bubbleSortPointers.o $(OUTDIR_OBJ)bubbleSortStructs.o $(OUTDIR_OBJ)insertionSortInts.o $(OUTDIR_OBJ)insertionSortPointers.o $(OUTDIR_OBJ)insertionSortStructs.o $(OUTDIR_OBJ)radixSortInts.o $(OUTDIR_OBJ)radixSortPointers.o $(OUTDIR_OBJ)radixSortStructs.o $(OUTDIR_OBJ)structs.o

all_structs_cache: directories structs_cache
		$(CC) $(CFLAGS) -o $(OUTDIR)struct-cache-benchmark $(OUTDIR_OBJ)bubbleSortInts.o $(OUTDIR_OBJ)bubbleSortPointers.o $(OUTDIR_OBJ)bubbleSortStructs.o $(OUTDIR_OBJ)insertionSortInts.o $(OUTDIR_OBJ)insertionSortPointers.o $(OUTDIR_OBJ)insertionSortStructs.o $(OUTDIR_OBJ)radixSortInts.o $(OUTDIR_OBJ)radixSortPointers.o $(OUTDIR_OBJ)radixSortStructs.o $(OUTDIR_OBJ)structs_cache.o

base: bubblesort insertionsort radixsort
	$(CC) $(CFLAGS) -c benchmark/base.cpp -o $(OUTDIR_OBJ)base.o

pointers: bubblesort insertionsort radixsort
	$(CC) $(CFLAGS) -c benchmark/pointers.cpp -o $(OUTDIR_OBJ)pointers.o

structs: bubblesort insertionsort radixsort
	$(CC) $(CFLAGS) -c benchmark/structs.cpp -o $(OUTDIR_OBJ)structs.o

structs_cache: bubblesort insertionsort radixsort
	$(CC) $(CFLAGS) -c benchmark/structs_cache.cpp -o $(OUTDIR_OBJ)structs_cache.o

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
