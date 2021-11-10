#include <stdio.h>

/* we should use MAX_SIZE sparingly to avoid unnecessary allocations */
#define MAX_SIZE 100

/* 
 * shiftArr removes the element at the zeroeth index and 
 * shifts the values at consecutive indexes down
 */
void shiftArr(int arr[], size_t arrLen)
{
  for (int i = 0; i < arrLen; ++i)
  {
    arr[i] = arr[i + 1];
  }
}

int maxSeen = 0;
int maxSeenArrLen = MAX_SIZE;
int maxSeenArr[MAX_SIZE];

/*
 * trySubsequences finds subsequences of a given array, checking
 * if the sum of the found subsequence is greater than any previous
 * subsequence. Storing that sum, and the subsequence if so.
 */
void trySubsequences(int arr[], size_t len)
{
  for (int i = 0; i < len; ++i)
  {
    int total = 0;
    int subsq[i + 1];

    for (int j = 0; j < i + 1; ++j)
    {
      total = total + arr[j];
      subsq[j] = arr[j];
    }

    if ((maxSeen == 0) || (total > maxSeen) || (total == maxSeen && len < maxSeenArrLen))
    {
      maxSeen = total;
      maxSeenArrLen = len;

      for (int k = 0; k < i + 1; ++k)
      {
        maxSeenArr[k] = subsq[k];
      }
    }
  }
}

int main()
{
  int n;

  scanf("%d", &n);

  int arr[n];

  for (int i = 0; i < n; ++i)
  {
    scanf("%d", &arr[i]);
  }

  size_t len = sizeof(arr) / sizeof(arr[0]);

  /* run trySubsequences on each subarray of the provided array */
  while (len != 0)
  {
    trySubsequences(arr, len);

    shiftArr(arr, len);
    len = len - 1;
  }

  printf("Largest sum is %d obtained from the subsequence ", maxSeen);

  printf("[ ");
  for (int i = 0; i < maxSeenArrLen; ++i)
  {
    if (maxSeenArr[i] != 0)
    {
      printf("%d ", maxSeenArr[i]);
    }
  }
  printf(" ]\n");

  return 0;
}