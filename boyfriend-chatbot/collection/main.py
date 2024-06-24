from .model import Database
from datetime import datetime

class Collection():
    def __init__(self, ip: str, port: int, id: str, pw: str, db: str):
        self.ip = ip
        self.port = port
        self.user = id
        self.password = pw
        self.db = db

    def open(self):
        self.database = Database(self.ip, self.port, self.user, self.password, self.db)
    
    def close(self):
        self.database = None
    
    def get_text(self, start: datetime = datetime(2024, 1, 1, 0, 0, 0), end: datetime = datetime(2024, 12, 31, 23, 59, 59)):
        # 특정 시간대 설정
        # start_time = datetime(2024, 1, 1, 0, 0, 0)  # 예시: 2023년 1월 1일 00:00:00
        # end_time = datetime(2024, 12, 31, 23, 59, 59)  # 예시: 2023년 1월 31일 23:59:59
        results = self.database.get_text(start, end)
        return results