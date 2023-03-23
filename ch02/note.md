# Getting started

## Insertion sort

is efficiency when sorting a small number of elements

pseudocode:

```{r, eval=FALSE}
for j = 2 to A.length
    key = A[j]
    // insert A[j] into the sorted squence A[1..j - 1].
    i = j - 1
    while i > 0 and A[i] > key
        A[i + 1] = A[i]
        i = i - 1
    A[i + 1] = key

```
