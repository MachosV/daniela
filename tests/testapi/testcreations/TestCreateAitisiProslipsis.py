import unittest
import requests
import json
import MySQLdb

class TestCreateAitisiProslipsis(unittest.TestCase):
    def setUp(self):
        db = MySQLdb.connect("localhost","root","123456789","daniela" )
        cursor = db.cursor()
        cursor.execute("INSERT INTO users (name,surname,nickname) values('test','test','danielaEmployee')")
        db.commit()
        cursor.execute("INSERT INTO users (name,surname,nickname) values('test','test','danielaEmployer')")
        db.commit()
        db.close()

    def test(self):
        data = {}
        data['id'] = '1'
        data['idEmployer'] = '2'
        json_data = json.dumps(data)
        r = requests.post(
            "http://localhost:8000/api/createaitisiproslipsis",
            data = json_data
        )
        self.assertEqual(r.text,u"OK")
    
    def tearDown(self):
        db = MySQLdb.connect("localhost","root","123456789","daniela" )
        cursor = db.cursor()
        cursor.execute("TRUNCATE users;")
        db.commit()
        cursor.execute("TRUNCATE aitiseis_proslipsis;")
        db.commit()
        db.close()