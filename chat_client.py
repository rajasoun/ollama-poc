import os
import asyncio
import platform 
from ollama import AsyncClient
from prompt_toolkit import PromptSession
from prompt_toolkit.styles import Style
from prompt_toolkit.formatted_text import HTML

class ChatClient:
    def __init__(self):
        self.style = Style.from_dict({'question': 'ansigreen',})
        self.session = PromptSession(style=self.style)
        self.client = AsyncClient()
        self.first_run = True

    def clear(self):
        """Clears the console screen."""
        # Check if the operating system is Windows
        if platform.system() == "Windows":
            os.system('cls')
        else:
            os.system('clear')

    async def start(self):
        if self.first_run:
            self.clear_screen_and_display_welcome_message()
        else:
            print("\n\n")

        user_input = await self.prompt_user()
        if user_input.lower() == 'exit':
            print("Exiting chat... Thank you for using the chat client.\n")
            return False  # Indicates the session should end

        await self.process_user_input(user_input)
        return True  # Indicates the session should continue

    def clear_screen_and_display_welcome_message(self):
        self.clear()
        print("Welcome to the Asynchronous Chat Client!")
        print("Type 'exit' to quit at any time.\n")
        self.first_run = False

    async def prompt_user(self):
        return await self.session.prompt_async(HTML('<ansigreen>Ask a question: </ansigreen>'))

    async def process_user_input(self, user_input):
        message = {'role': 'user', 'content': user_input}
        try:
            async for part in await self.client.chat(model='llama2', messages=[message], stream=True):
                print(part['message']['content'], end='', flush=True)
        except Exception as e:
            print(f"\nAn error occurred during the chat session: {e}\nPlease try asking something else.")
