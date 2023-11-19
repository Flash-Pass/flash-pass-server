import unittest

from tests.api_e2e.binary import operations


class Register(unittest.TestCase):

    TOKEN: str = ""

    def setUp(self) -> None:
        pass

    def tearDown(self) -> None:
        pass

    def test_login_success(self):
        mobile, password, data = operations.register()
        self.assertEqual(data["code"], 0)
        self.assertNotEqual(len(data["data"]), 0)
        self.TOKEN = data["data"]

        data = operations.login(mobile, password)
        self.assertEqual(data["code"], 0)
        self.assertNotEqual(len(data["data"]), 0)
        self.assertEqual(data["data"], self.TOKEN)

    def test_login_without_mobile(self):
        data = operations.login("", "root")
        self.assertNotEqual(data["code"], 0)
        self.assertEqual(data["data"], None)

    def test_login_without_password(self):
        data = operations.login("13935651548", "")
        self.assertNotEqual(data["code"], 0)
        self.assertEqual(data["data"], None)

    def test_login_with_incorrect_password(self):
        mobile, password, data = operations.register()
        self.assertEqual(data["code"], 0)
        self.assertNotEqual(len(data["data"]), 0)
        self.TOKEN = data["data"]

        data = operations.login(mobile, password + "123")
        self.assertNotEqual(data["code"], 0)
        self.assertEqual(data["data"], None)