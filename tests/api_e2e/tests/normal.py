import unittest

from tests.api_e2e.binary import operations


class Normal(unittest.TestCase):
    TOKEN: str = ""

    def test_normal_process(self):
        mobile, password, data = operations.register()
        print(mobile, password)
        self.assertEqual(data["code"], 0)
        self.assertNotEqual(len(data["data"]), 0)
        self.TOKEN = data["data"]

        data = operations.login(mobile, password)
        self.assertEqual(data["code"], 0)
        self.assertNotEqual(len(data["data"]), 0)
        self.assertEqual(data["data"], self.TOKEN)

        title, description, data = operations.create_book(self.TOKEN, gen_title=True, gen_description=True)
        self.assertEqual(data["code"], 0)
        self.assertEqual(data["data"]["title"], title)
        self.assertEqual(data["data"]["description"], description)

        book_id = data["data"]["id"]
        for i in range(10):
            question, answer, data = operations.create_card(self.TOKEN, add_to_book=True, book_id=book_id, gen_answer=True, gen_question=True)
            self.assertEqual(data["code"], 0)
            self.assertEqual(data["data"]["question"], question)
            self.assertEqual(data["data"]["answer"], answer)
