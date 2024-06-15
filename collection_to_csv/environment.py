from dotenv import load_dotenv
import os

def get_environment():
    load_dotenv(override=True)
    postgresql_ip = os.getenv('POSTGRES_IP')
    assert(postgresql_ip != '')

    postgresql_user = os.getenv('POSTGRES_USER')
    assert(postgresql_user != '')

    postgresql_password = os.getenv('POSTGRES_PASSWORD')
    assert(postgresql_password != '')

    postgresql_db = os.getenv('POSTGRES_DB')
    assert(postgresql_db != '')

    return postgresql_ip, postgresql_user, postgresql_password, postgresql_db
