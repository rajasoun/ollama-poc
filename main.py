# main entry point for the chat application.
import asyncio
from chat_client import ChatClient


async def main():
    """
    Main entry point for the chat application.

    Creates an instance of ChatClient and enters a loop waiting for the user
    to interact with the chat. The loop continues until the user decides to exit
    by typing 'exit'.
    """
    chat_client = ChatClient()
    while await chat_client.start():
        pass  # The loop will continue until the user exits the chat session.


if __name__ == "__main__":
    """
    If this script is executed as the main program, the asyncio event loop
    is run with the main() coroutine.

    This starts the chat client and manages the asynchronous chat session.
    """
    asyncio.run(main())
