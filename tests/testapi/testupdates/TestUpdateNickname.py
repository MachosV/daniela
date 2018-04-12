import unittest
import requests
import json
import MySQLdb

class TestUpdateNickname(unittest.TestCase):
    def setUp(self):
        db = MySQLdb.connect("localhost","root","123456789","daniela" )
        cursor = db.cursor()
        cursor.execute("INSERT INTO users (name,surname,nickname) values('test','test','daniela')")
        db.commit()
        db.close()

    def test(self):
        dataDict = {}
        dataDict['id'] = '1'
        dataDict['nickname'] = 'new_nickname'
        r = requests.post(
            "http://localhost:8000/api/updateusernickname",
            data = dataDict
        )
        self.assertEqual(r.text,u'OK')
    
    def tearDown(self):
        db = MySQLdb.connect("localhost","root","123456789","daniela" )
        cursor = db.cursor()
        cursor.execute("TRUNCATE users;")
        db.close()