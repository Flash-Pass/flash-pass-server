from typing import *
import requests
import json


class API:
    BASE_URL = "http://localhost:8080"
    TOKEN = ""

    def __init__(self, base_url: str = "", token: str = ""):
        if base_url != "":
            self.BASE_URL = base_url
        if token != "":
            self.TOKEN = token

    @property
    def token(self) -> str:
        return self.TOKEN

    @token.setter
    def token(self, token: str):
        self.TOKEN = token

    def request(self, retry: int = 3):
        pass

    def register(self, mobile: str, password: str) -> requests.Response:
        return requests.post(f"{self.BASE_URL}/user/register/mobile", json={
            "mobile": mobile,
            "password": password
        })

    def login(self, mobile: str, password: str) -> requests.Response:
        return requests.post(f"{self.BASE_URL}/user/login", json={
            "mobile": mobile,
            "password": password
        })

    def create_card(self, question: str, answer: str) -> requests.Response:
        return requests.post(f"{self.BASE_URL}/card", json={
            "question": question,
            "answer": answer
        }, headers={
            "Authorization": self.TOKEN
        })

    def get_card(self, id: str) -> requests.Response:
        return requests.get(f"{self.BASE_URL}/card", params={
            "id": id
        }, headers={
            "Authorization": self.TOKEN
        })
