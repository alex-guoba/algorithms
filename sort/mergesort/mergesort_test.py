import unittest
import random

from mergesort import mergesort 

class TestStringMethods(unittest.TestCase):

    #@unittest.skip("demonstrating skipping")
    def test_split(self):
        src :Vector = []

        for i in range(500):
            n = random.randint(1, 1000)
            if n not in src:
                src.append(n)

        dst = mergesort(src)
        self.assertEqual(sorted(src), dst, "sort failed, check input list")

if __name__ == '__main__':
    unittest.main()
