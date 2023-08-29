import unittest
import random

from heapsort import heapsort

class TestStringMethods(unittest.TestCase):

    #@unittest.skip("demonstrating skipping")
    def test_split(self):
        src :Vector = []

        for i in range(10):
            n = random.randint(1, 1000)
            if n not in src:
                src.append(n)

        dst = heapsort(src)
        self.assertEqual(sorted(src), dst, "failed, check input list")

if __name__ == '__main__':
    unittest.main()
