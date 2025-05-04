"""Recursion: The Out-Shuffle
An out-shuffle, also known as an out faro shuffle or a perfect shuffle, is a controlled method for shuffling playing cards. It is performed by splitting the deck into two equal halves and interleaving them together perfectly, with the condition that the top card of the deck remains in place.

Using an array to represent a deck of cards, an out-shuffle looks like:

[1, 2, 3, 4, 5, 6, 7, 8] ➞ [1, 5, 2, 6, 3, 7, 4, 8]
// Card 1 remains in the first position.
If we repeat the process, the deck eventually returns to original order:

Shuffle 1:

[1, 2, 3, 4, 5, 6, 7, 8] ➞ [1, 5, 2, 6, 3, 7, 4, 8]
Shuffle 2:

[1, 5, 2, 6, 3, 7, 4, 8] ➞ [1, 3, 5, 7, 2, 4, 6, 8]
Shuffle 3:

[1, 3, 5, 7, 2, 4, 6, 8] ➞ [1, 2, 3, 4, 5, 6, 7, 8]
// Back where we started.
Write a function that takes a positive even integer representing the number of the cards in a deck, and returns the number of out-shuffles required to return the deck to its original order.

Examples
shuffle_count(8) ➞ 3

shuffle_count(14) ➞ 12

shuffle_count(52) ➞ 8
Notes
The number of cards is always even and greater than one. Thus, the smallest possible deck size is two.
An iterative version of this challenge can be found here.
"""
def shuffle_count(num):

    original_deck = list(range(1, num + 1))  # 1からnumまでのリストを作成
    current_deck = original_deck[:] # 現在のデッキを元のデッキと同じにする
    
    shuffle_count = 0

    while True:
        half = len(current_deck) // 2  # デッキの半分の長さを計算
        top_half = current_deck[:half]  # 上半分
        bottom_half = current_deck[half:]  # 下半分

        new_deck = []

        for i in range(half):
            new_deck.append(top_half[i])  # 上半分から1枚追加
            new_deck.append(bottom_half[i])  # 下半分から1枚追加

        # シャッフル後の新しいデッキを現在のデッキとする
        shuffle_count += 1
        current_deck = new_deck

        if current_deck == original_deck:
            return shuffle_count
        
print(shuffle_count(8)) #  ➞ 3
print(shuffle_count(14)) # ➞ 12
print(shuffle_count(52)) # ➞ 8