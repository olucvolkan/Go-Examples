# Go-Examples

#First Example 
You have 50 bitcoins to distribute to 10 users: Matthew, Sarah, Augustus, Heidi, Emilie, Peter, Giana, Adriano, Aaron, Elizabeth The coins will be distributed based on the vowels contained in each name where:

a: 1 coin e: 1 coin i: 2 coins o: 3 coins u: 4 coins

and a user can’t get more than 10 coins. Print a map with each user’s name and the amount of coins distributed. After distributing all the coins, you should have 2 coins left.

The output should look something like that:

map[Matthew:2 Peter:2 Giana:4 Adriano:7 Elizabeth:5 Sarah:2 Augustus:10 Heidi:5 Emilie:6 Aaron:5]
Coins left: 2
Note that Go doesn’t keep the order of the keys in a map, so your results might not look exactly the same but the key/value mapping should be the same.

