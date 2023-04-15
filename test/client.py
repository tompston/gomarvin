# import gomarvin.gen
from typing import Optional
import gomarvin as gomarvin

client = gomarvin.defaultClient



response_ok = {'status': 200, 'message': '', 'data': {}}

response_ok_with_links = {'status': 200,
                          'message': '', 'data': {}, 'links': {}}


def test_get_users():
    params = gomarvin.OptionalParams(append_url="?name=jim")
    get_users = gomarvin.UserEndpoints(client).GetUsers(opt=params)
    print(get_users.json())
    print(get_users.url)
    assert get_users.status_code == 200
    assert get_users.json() == response_ok_with_links


# def my_function(param1: str, param2: Optional[int] = None, param3: Optional[float] = None):
#     # Function code here
#     pass