# Daily Coding Problem: Problem #1802 [Hard]

## Problem Statement

This problem was asked by Airbnb.

Given a list of words,
find all pairs of unique indices such that the concatenation of the two
words is a palindrome.

For example, given the list
["code", "edoc", "da", "d"],
return [(0, 1), (1, 0), (2, 3)].

## Analysis

There's two algorithms to deal with,
finding out whether a string is a palindrome,
and iterating over the strings in the array correctly.

The statement doesn't require any limits,
like O(n) time, or O(1) additional space,
but it does want pairs of indices.
This could be a hint to use array indexing to find pairs of string to test,
or it could be a deliberately misleading cue.

The example gives extra info, "(0, 1)" and "(1, 0)" indicate
they want you to find "codeedoc" and "edoccode".
You could easily read the text of the problem statement to mean that
"(0, 1)" and "(1, 0)" would be some kind of equivalent.

#### Palindromes

Commonly, "palindromes" are taken to ignore whitespace,
so "race car" is often cited as a palindrome.
This is a great place for a job candidate given this problem
to ask a clarifying question.
I assumed that that the input strings do not contain whitespace,
but I'm not solving the problem in an interview,
I have the luxury of ignoring this question.

I think it would be worth asking if a zero length string
can occur in the input.
It's the identity element in string concatenation,
so the possibility of zero length strings in the input could have consequences.

My code converts two input strings into Go `[]rune` type slices,
then concatenates the slices, so as to be able to use a simple
palindrome check.

It might be possible to write some super clever indexing-into-the-strings
algorithm that could account for strings of different lengths to avoid
the concatentation.
I didn't bother.

Assuming Unicode strings does lead you to some dark places.
You should consider performing ["NFC" normalization](https://dencode.com/en/string/unicode-normalization)
on Unicode input strings.
It might be worthwhile to ask if byte-level or code point level palindromes are desired,
and find out what character encoding is used.

#### Complexity

The number of concatenated strings to check to see if they're palindromes
should be the combinatoric "N P 2" - N strings, permuted 2 at a time.
A function implementing this problem would  make n!/(n - 2)! or n<sup>2</sup> - n
checks on the palindrome property.

I think that comes down to O(n<sup>2</sup>), where n is the number of input strings.

I think checking whether or not a string is a palindrome
is O(m/2), where m is the number of characters in the string.
I'm not sure how to combine these two measures of complexity.

## Interview Analysis

This does not merit a "[Hard]" rating.
There are two algorithms involved, checking for palindrome property,
and iterating through the permutations of 2 strings out of N,
but they're both fairly  simple.
There are no data structures involved, other than a couple of arrays.

I think an interviewer might get a feel if a job candidate
could write loops,
and understood the difference between "2 strings out of N"
and "permutations of 2 strings out of N"
Candidates with a deeper understanding of character encoding in
the implementation language might end up asking a lot more clarifying questions.
If the interviewer is testing for Unicode knowledge,
this problem might reveal something.
