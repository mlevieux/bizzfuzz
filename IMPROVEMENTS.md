# Improvements

## 1. Using sync.Pool for short-lived small-size objects

There are some data structures that are used often and require allocation.
Those data structures are mostly short-lived and really small size, like the gen structure used to generate contiguous number sequences.

Using a *sync.Pool to avoid allocating those objects everytime one is needed could diminish pressure on GC, which can be a visible performance improvement in high traffic environments.

## 2. Refining approximation of fizzbuzz needed buffer size to get the exact value of it

Currently, I was not able to find an exact formula of the needed buffer size for a given fizzbuzz-sequence set of parameters (int1, int2, limit, str1 and str2). The current function only approximates the needed buffer size. Though it always gives a value in [N-100;N+100] where N is the real wanted value, it might be more easily maintainable to have a function that returns the exact value we're looking for.

This should not in any case improve the overall performance of the program in a visible way.

## 3. Handling values over math.MaxInt32

The current version of the projects can only compute fizzbuzz sequence up to a limit of math.MaxInt32. This might already be a high value for such a task, especially given the memory constraints due to []byte generation, yet in systems with way higher performance hardware, one might want to generate greater limit sequences of fizzbuzz procedure. This could also be the case if we stream the result instead of generating it as a whole before returning it (see [section 4](#streaming-result-sequence) below).

## 4. Streaming result sequence

A good way to improve the program would be to stream the fizzbuzz sequence generated to the client instead of generating the full result before returning it.
I did not try to do this and I think it would probably represent a tremendous amount of work (relative to what's already been done).

This improvement would theoretically enable generating any-size sequence generation.