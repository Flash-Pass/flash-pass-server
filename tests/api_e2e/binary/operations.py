from tests.api_e2e.binary.api import API
from tests.api_e2e.binary import generator
import requests

api = API()


def register(mobile: str = "", password: str = "", gen_mobile: bool = True, gen_password: bool = True) -> (
        str, str, dict):
    if gen_mobile:
        mobile = generator.user_mobile()
    if gen_password:
        password = generator.user_password()

    resp: requests.Response = api.register(mobile, password)
    data = resp.json()

    return mobile, password, data


def login(mobile, password) -> dict:
    resp: requests.Response = api.login(mobile, password)
    data = resp.json()
    return data


def create_book(token: str, title: str = "", gen_title: bool = False, description: str = "",
                gen_description: bool = False) -> (str, str, dict):
    if gen_title:
        title = generator.text(10)
    if gen_description:
        description = generator.text(25)
    resp: requests.Response = api.create_book(token, title, description)
    data = resp.json()
    return title, description, data


def create_card(token: str, book_id: int, add_to_book: bool = False, question: str = "", gen_question: bool = False,
                answer: str = "", gen_answer: bool = False) -> (str, str, dict):
    if gen_question:
        question = generator.text(10)
    if gen_answer:
        answer = generator.text(25)
    resp: requests.Response = api.create_card(token, book_id, add_to_book, question, answer)
    data = resp.json()
    return question, answer, data


def get_book_card_list(token: str, book_id: int) -> dict:
    resp: requests.Response = api.get_book_card_list(token, book_id)
    data = resp.json()
    return data
