import unittest
import requests
import json
import MySQLdb

class TestDeleteUser(unittest.TestCase):
    def setUp(self):
        db = MySQLdb.connect("localhost","root","123456789","daniela" )
        cursor = db.cursor()
        cursor.execute("INSERT INTO users (name,surname,nickname) values('test','test','daniela')")
        db.commit()
        db.close()

    def test(self):
        r = requests.post("http://localhost:8000/api/deleteuser",data={"id":"1"})
        self.assertEqual(r.text,u"OK")
        r = requests.post("http://localhost:8000/api/deleteuser",data={"id":"test"})
        self.assertEqual(r.text,u"Error deleting user")
    
    def tearDown(self):
        db = MySQLdb.connect("localhost","root","123456789","daniela" )
        cursor = db.cursor()
        cursor.execute("TRUNCATE users;")
        db.close()