import unittest

import requests

from tests.api_e2e.binary import operations
from tests.api_e2e.binary.api import API


class Card(unittest.TestCase):
    TOKEN: str = ""

    def setUp(self) -> None:
        pass

    def tearDown(self) -> None:
        pass

    def test_create_card(self):
        token = operations.get_token_register_login()
        api = API(token=token)

        question = "How to learn golang"
        answer = "Python is better"
        resp: requests.Response = api.create_card(question, answer)
        create_result = resp.json()
        self.assertEqual(create_result["code"], 0)
        self.assertEqual(create_result["data"]["question"], question)
        self.assertEqual(create_result["data"]["answer"], answer)
        self.assertNotEqual(create_result["data"]["id"], "")

        card_id = create_result["data"]["id"]
        resp: requests.Response = api.get_card(card_id)
        get_result = resp.json()
        self.assertEqual(get_result["code"], 0)
        self.assertEqual(get_result["data"]["question"], question)
        self.assertEqual(get_result["data"]["answer"], answer)
