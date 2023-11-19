import unittest

from tests.api_e2e.binary import operations


class Register(unittest.TestCase):

    TOKEN: str = ""

    def setUp(self) -> None:
        pass

    def tearDown(self) -> None:
        pass

    def test_register_success(self):
        mobile, password, data = operations.register()
        self.assertEqual(data["code"], 0)
        self.assertGreater(len(data["data"]), 0)

    def test_register_without_formatted_mobile(self):
        mobile, password, data = operations.register(mobile="123456789", gen_mobile=False)
        self.assertNotEqual(data["code"], 0)
        self.assertEqual(data["data"], None)

    def test_register_without_mobile(self):
        mobile, password, data = operations.register(gen_mobile=False)
        self.assertNotEqual(data["code"], 0)
        self.assertEqual(data["data"], None)

    def test_register_without_password(self):
        mobile, password, data = operations.register(gen_password=False)
        self.assertNotEqual(data["code"], 0)
        self.assertEqual(data["data"], None)
