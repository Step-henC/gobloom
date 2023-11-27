### Bloom Filter Go Package

# Bloom Filters

Bloom Filters are a quick way to query for the presence of data, while minimizing memory space. Bloom filters are a data structure that have a high false positive rate, but a low false negative, meaning </br>
they are less likely to miss data that is present. The accuracy tradeoff of bloom filters is counterbalanced by the lower memory space of a bloom filter in comparison to hash tables and other data structures.

# Bloom Filter Mechanism

The bloom filter is a byte array of the hashed values of any input data from a user or service. The data is hashed and stored in an array. To check for a value, the input data is compared to its hash </br>
value and checked against other stored hashes. 

# This Package/Repository

This package creates an optimal bloom filter based on acceptable false positive rate (closer to zero = more accurate) and expected size of input provided by the user. </br>
To use this package for your go project simply type this in the cli: `go get github.com/Step-henC/gobloom`.
