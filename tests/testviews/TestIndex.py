import unittest,urllib2

class TestIndex(unittest.TestCase):
    def test(self):
        data = urllib2.urlopen("http://localhost:8000")
        data = data.read()
        self.assertEqual(data,"hello")
