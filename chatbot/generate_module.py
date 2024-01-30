# chatbot/generate_module.py

import os
from dotenv import load_dotenv
load_dotenv()
import google.generativeai as genai
import pathlib
import textwrap
from IPython.display import Markdown

GOOGLE_API_KEY = os.getenv("GOOGLE_API_KEY")

# Configure the GenerativeAI module with the provided Google API key
genai.configure(api_key=GOOGLE_API_KEY)

def to_markdown(text):
    """
    Converts plain text to Markdown format.

    Parameters:
    - text (str): The plain text to be converted to Markdown.

    Returns:
    - str: The Markdown-formatted text.
    """
    # Replace '•' with Markdown bullet points ('*')
    text = text.replace('•', '  *')

    # Indent the text for proper Markdown formatting
    return Markdown(textwrap.indent(text, '', predicate=lambda _: True))

class Generate:
    def __init__(self):
        """
        Initializes an instance of the class with the following steps:
        1. Initializes a GenerativeModel using the 'gemini-pro' model.
        2. Reads the contents of the 'prompt.txt' file and creates a prompt string.
        3. Starts a chat using the initialized model with an empty history.
        4. Sends the prompt as a message to the chat and retrieves the response.
        """
        # Step 1: Initialize the GenerativeModel from the 'gemini-pro' model
        self.model = genai.GenerativeModel('gemini-pro')

        # Step 2: Read the contents of 'prompt.txt' and create a prompt string
        with open('./chatbot/prompt.txt', "r", encoding="utf-8") as f:
            lines = f.readlines()
        prompt = ' '.join(lines)

        # Step 3: Start a chat using the initialized model with an empty history
        self.chat = self.model.start_chat(history=[])

        # Step 4: Send the prompt as a message and get the response
        response = self.chat.send_message(prompt)

        # If you want to store the data in some then uncomment the below code
        
        # s = to_markdown(response.text).data
        # with open("response.txt", 'a', encoding="utf-8") as res:
        #     res.write(s)
        #     res.write('\n\n')

    def generate_text(self, inp):
        """
        Generates and processes text based on the provided input

        Parameters:
        - self: Instance of the class containing the method.
        - s (str): Input text for generating a response.

        Returns:
        - str: The first line of the Markdown-formatted response.
        """
        try:
            response = self.chat.send_message(inp)
        except:
            response = self.model.generate_content(inp, safety_settings={'HARM_CATEGORY_SEXUALLY_EXPLICIT': 'block_none', 'HARM_CATEGORY_HATE_SPEECH': 'block_none','HARM_CATEGORY_HARASSMENT': 'block_none', 'HARM_CATEGORY_DANGEROUS_CONTENT': 'block_none'}) or "None Found"

        # If you want to store the data in some then uncomment the below code

        # with open("response.txt", 'a', encoding="utf-8") as res:
        #     res.write(inp)
        #     res.write('\n\n')

        out = to_markdown(response.text).data

        # If you want to store the data in some then uncomment the below code

        # with open("response.txt", 'a', encoding="utf-8") as res:
        #     res.write(out)
        #     res.write('\n\n')

        out = out.split('\n')[0]
        return out
