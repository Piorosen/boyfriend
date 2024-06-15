from model import Database
from environment import get_environment

from datetime import datetime

ip, user, password, db = get_environment()
database = Database(ip, user, password, db)

# 특정 시간대 설정
start_time = datetime(2024, 1, 1, 0, 0, 0)  # 예시: 2023년 1월 1일 00:00:00
end_time = datetime(2024, 12, 31, 23, 59, 59)  # 예시: 2023년 1월 31일 23:59:59
results = database.get_text(start_time, end_time)

for row in results:
    print(row.__dict__)
