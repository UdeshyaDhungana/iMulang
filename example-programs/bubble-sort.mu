
$ Program to demonstrate bubble sort algorithm $
thoos bubbleSort = kaam_gar(array) {
    thoos n = lambai(array);
    thoos i = 0;
    thoos temp = 0;

    ghuma(i = 0; i < n - 1; i = i + 1) {
        thoos j = 0;
        ghuma(j = 0; j < n - 1 - i; j = j + 1) {
            yedi(array[j] > array[j+1]) {
                temp = array[j];
                array[j] = array[j+1];
                array[j+1] = temp;
            }
        }
    }
};


thoos myArray = [4,3,5,1,9];
bubbleSort(myArray);
bhan(myArray);