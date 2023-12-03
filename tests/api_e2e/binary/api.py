from typing import *
import requests
import json


class API:
    BASE_URL = "http://localhost:8080"

    def __init__(self, base_url: str = ""):
        if base_url != "":
            self.BASE_URL = base_url

    def request(self, retry: int = 3):
        pass

    def register(self, mobile: str, password: str) -> requests.Response:
        return requests.post(f"{self.BASE_URL}/user/register/mobile", json={
            "mobile": mobile,
            "password": password
        })

    def login(self, mobile: str, password: str):
        return requests.post(f"{self.BASE_URL}/user/login", json={
            "mobile": mobile,
            "password": password
        })

    def create_book(self, token: str, title: str, description: str = ""):
        return requests.post(f"{self.BASE_URL}/book", json={
            "title": title,
            "description": description
        }, headers={
            "Authorization": token
        })

    def create_card(self, token: str, book_id: int, add_to_book: bool, question: str, answer: str):
        return requests.post(f"{self.BASE_URL}/card", json={
            "question": question,
            "answer": answer,
            "is_add_to_book": add_to_book,
            "book_id": book_id
        }, headers={
            "Authorization": token
        })

    def get_book_card_list(self, token: str, book_id: int):
        return requests.get(f"{self.BASE_URL}/book/card/list", json={
            "book_id": book_id
        }, headers={
            "Authorization": token
        })
