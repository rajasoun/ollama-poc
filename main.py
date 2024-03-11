import asyncio
from chat_client import ChatClient

async def main():
    chat_client = ChatClient()
    while await chat_client.start():
        pass

if __name__ == '__main__':
    asyncio.run(main())
