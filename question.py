import asyncio
from ollama import AsyncClient
from prompt_toolkit import PromptSession
from prompt_toolkit.shortcuts import clear
from prompt_toolkit.styles import Style
from prompt_toolkit.formatted_text import HTML

# Define a style for the prompt
style = Style.from_dict({
    'question': 'ansigreen',
})

async def chat():
    session = PromptSession(style=style)
    client = AsyncClient()  # Initialize the client outside the loop
    first_run = True  # Flag to check if it's the first iteration
    
    try:
        while True:
            if first_run:
                clear()  # Clear the screen only on the first run
                print("Welcome to the Asynchronous Chat Client!")
                print("Type 'exit' to quit at any time.\n")
                first_run = False  # Set the flag to False after the first iteration
            else :
                # add two  new line after the previous chat message
                print("\n")
            
            # Use HTML for colored prompt text
            user_input = await session.prompt_async(HTML('<ansigreen>Ask a question: </ansigreen>'))
            if user_input.lower() == 'exit':
                break
            
            message = {'role': 'user', 'content': user_input}
            try:
                async for part in await client.chat(model='llama2', messages=[message], stream=True):
                    print(part['message']['content'], end='', flush=True)
            except Exception as e:
                print(f"\nAn error occurred during the chat session: {e}\nPlease try asking something else.")
    finally:
        print("Exiting chat... Thank you for using the chat client.\n")

asyncio.run(chat())
