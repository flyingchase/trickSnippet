import re
from collections import abc


class Sentence:
    def __init__(self, sentence):
        self.sentence = sentence
        self.words = re.findall(r"\w+", sentence)

    def __iter__(self):
        return SentenceIterator(self.words)


class SentenceIterator(abc.Iterator):
    def __int__(self, words):
        self.words = words
        self._index = 0

    def __next__(self):
        try:
            word = self.words[self._index]
        except IndexError:
            raise StopIteration
        else:
            self._index += 1

        return word


if __name__ == "__main__":
    sentence = Sentence(
        "Return a list of all non-overlapping matches in the string.")
    assert isinstance(sentence, abc.Iterable)
    assert isinstance(iter(sentence), abc.Iterator)

    for word in sentence:
        print(word, end=".")
