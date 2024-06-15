from sqlalchemy import Column, String, Integer, BigInteger, Text, Sequence
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import relationship
from sqlalchemy.sql import func
from sqlalchemy.types import TIMESTAMP
from sqlalchemy import create_engine, desc
from sqlalchemy.orm import sessionmaker

Base = declarative_base()

class message(Base):
    __tablename__ = 'messages'

    id = Column(Integer, primary_key=True, autoincrement=True)
    created_at = Column(TIMESTAMP, server_default=func.now())
    updated_at = Column(TIMESTAMP, server_default=func.now(), onupdate=func.now())
    deleted_at = Column(TIMESTAMP, nullable=True)
    first_name = Column(String(100), nullable=True)
    last_name = Column(String(100), nullable=True)
    user_name = Column(String(100), nullable=True)
    user_id = Column(BigInteger, nullable=False)
    text_id = Column(BigInteger, nullable=False)
    text = Column(Text, nullable=True)

class Database():
    def __init__(self, ip, user, password, db):
        URL = f'postgresql://{user}:{password}@{ip}:{5432}/{db}'
        engine = create_engine(URL, client_encoding='utf8')
        # 세션 생성
        self.Session = sessionmaker(bind=engine)
    def get_text(self, start_time, end_time):
        session = self.Session()
        return session.query(message)\
            .filter(message.created_at >= start_time, message.created_at <= end_time)\
            .order_by(desc(message.created_at))\
            .all()
