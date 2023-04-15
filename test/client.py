# import gomarvin.gen
from typing import Optional
import gomarvin as gomarvin

client = gomarvin.defaultClient


response_ok = {'status': 200, 'message': '', 'data': {}}

response_ok_with_links = {'status': 200,
                          'message': '', 'data': {}, 'links': {}}


def test_get_users():
    data = gomarvin.UserEndpoints(client).GetUsers(append_url="?name=jim")
    assert data.status_code == 200 and data.json() == response_ok_with_links
    print("PASS")


test_get_users()


def test_get_user():
    data = gomarvin.UserEndpoints(client).GetUserById(1)
    assert data.status_code == 200 and data.json() == response_ok
    print("PASS")


test_get_user()


def test_create_user_passing_validation():
    body = gomarvin.CreateUserBody(
        username="jim", password="very-long-password", email="jim@email.com")
    data = gomarvin.UserEndpoints(client).CreateUser(body=body)
    assert data.status_code == 200 and data.json() == response_ok
    print("PASS")


test_create_user_passing_validation()


def test_create_user_failing_validation():
    body = gomarvin.CreateUserBody(
        username="jim", password="short", email="jim@email.com")
    data = gomarvin.UserEndpoints(client).CreateUser(body=body)
    assert data.status_code != 200
    print("PASS")


test_create_user_failing_validation()
