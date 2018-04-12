import unittest
import requests
import json
import MySQLdb

class TestCreateUser(unittest.TestCase):
    def test(self):
        data = {}
        data['name'] = 'testname'
        data['surname'] = 'test surname'
        data['nickname'] = 'daniela'
        data['isEmployer'] = True
        json_data = json.dumps(data)
        r = requests.post(
            "http://localhost:8000/api/createuser",
            data = json_data
        )
        self.assertEqual(r.text,u"OK")
    
    def tearDown(self):
        db = MySQLdb.connect("localhost","root","123456789","daniela" )
        cursor = db.cursor()
        cursor.execute("TRUNCATE users;")
        db.close()