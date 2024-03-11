import os
import platform 
from ollama import AsyncClient
from prompt_toolkit import PromptSession
from prompt_toolkit.styles import Style
from prompt_toolkit.formatted_text import HTML

class ChatClient:
    """
    A chat client that interfaces with the Ollama chat model to answer questions.
    Utilizes prompt_toolkit for an enhanced asynchronous terminal interface.
    """

    def __init__(self):
        """
        Initializes the ChatClient with a styled prompt session and the Ollama AsyncClient.
        """
        self.style = Style.from_dict({'question': 'ansigreen',})
        self.session = PromptSession(style=self.style)
        self.client = AsyncClient()
        self.first_run = True

    def clear(self):
        """
        Clears the console screen based on the operating system.
        """
        # Check if the operating system is Windows
        if platform.system() == "Windows":
            os.system('cls')
        else:
            os.system('clear')

    async def start(self):
        """
        Starts the chat session. Clears the screen and displays a welcome message on first run.
        Continuously prompts the user for questions until 'exit' is entered.
        """
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
        """
        Clears the screen and displays the initial welcome message for the chat client.
        """
        self.clear()
        print("Welcome to the Asynchronous Chat Client!")
        print("Type 'exit' to quit at any time.\n")
        self.first_run = False

    async def prompt_user(self):
        """
        Asynchronously prompts the user for input using a styled prompt.
        Returns the string input by the user.
        """
        return await self.session.prompt_async(HTML('<ansigreen>Ask a question: </ansigreen>'))

    async def process_user_input(self, user_input):
        """
        Sends the user's input to the Ollama chat model and asynchronously processes the response.
        Prints the response to the terminal.
        Handles any exceptions that occur during communication with the chat model.
        """
        message = {'role': 'user', 'content': user_input}
        try:
            async for part in await self.client.chat(model='llama2', messages=[message], stream=True):
                print(part['message']['content'], end='', flush=True)
        except Exception as e:
            print(f"\nAn error occurred during the chat session: {e}\nPlease try asking something else.")
