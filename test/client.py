# import gomarvin.gen
import gomarvin as gomarvin

client = gomarvin.defaultClient

params = gomarvin.OptionalParams(append_url="?name=jim")
get_users = gomarvin.UserEndpoints(client).GetUsers(opt=params)
print(get_users.json())
print(get_users.url)

response_ok = {'status': 200, 'message': '', 'data': {}}

response_ok_with_links = {'status': 200,
                          'message': '', 'data': {}, 'links': {}}

assert get_users.status_code == 200
assert get_users.json() == response_ok_with_links
