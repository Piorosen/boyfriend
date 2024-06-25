import os
from dotenv import load_dotenv
from collection import Collection

def get_environment():
    load_dotenv(override=True)
    postgresql_ip = os.getenv('POSTGRES_IP')
    assert(postgresql_ip != '')

    postgresql_port = os.getenv('POSTGRES_PORT')
    assert(postgresql_ip != '')

    postgresql_user = os.getenv('POSTGRES_USER')
    assert(postgresql_user != '')

    postgresql_password = os.getenv('POSTGRES_PASSWORD')
    assert(postgresql_password != '')

    postgresql_db = os.getenv('POSTGRES_DB')
    assert(postgresql_db != '')
    
    # jinju_id = os.getenv('TELEGRAM_USER_JINJU_ID')
    # assert(jinju_id != '')

    return postgresql_ip, postgresql_port, postgresql_user, postgresql_password, postgresql_db #, jinju_id

def next_tokne(data):
    result = []
    for item in data:
        if len(result) == 0:
            result.append(item)
        elif item.user_id == result[-1].user_id:
            result[-1].text += ', ' + item.text
        else:
            result.append(item)
    return result

def get_text(ip, port, user, pw, db):
    coll = Collection(ip, port, user, pw, db)
    coll.open()
    data = coll.get_text()
    coll.close()
    data = reversed(data)
    data = next_tokne(data)
    # data = list(map(lambda x: {'name': x.user_name, 
    #             'text': x.text,
    #             'created_at': x.created_at,
    #             'text_id': x.text_id,
    #             'user_id': x.user_id}, data))
    return data


