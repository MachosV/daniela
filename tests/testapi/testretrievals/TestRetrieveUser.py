import unittest
import requests
import json
import MySQLdb

class TestRetrieveUser(unittest.TestCase):
    def setUp(self):
        db = MySQLdb.connect("localhost","root","123456789","daniela" )
        cursor = db.cursor()
        cursor.execute("INSERT INTO users (name,surname,nickname,isEmployer) values('test','test','daniela',1);")
        db.commit()
        db.close()
        
    def test(self):
        dataDict = dict()
        dataDict["id"] = "1"
        r = requests.post("http://localhost:8000/api/retrieveuser",data=dataDict)
        flag = False
        if '''"isEmployer": true''' in r.text and "daniela" in r.text and '''"stars": 0''' in r.text:
            flag = True
        self.assertEqual(flag,True)
        
    def tearDown(self):
        db = MySQLdb.connect("localhost","root","123456789","daniela" )
        cursor = db.cursor()
        cursor.execute("TRUNCATE users;")
        db.close()