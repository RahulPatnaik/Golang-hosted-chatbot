# app.py
from fastapi import FastAPI
from pydantic import BaseModel
from transformers import pipeline, AutoTokenizer, AutoModelForCausalLM
import torch
from fastapi.middleware.cors import CORSMiddleware  # Import CORS middleware

app = FastAPI()

# Add CORS middleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Allow all origins (you can restrict this to specific origins if needed)
    allow_credentials=True,
    allow_methods=["*"],  # Allow all HTTP methods (GET, POST, etc.)
    allow_headers=["*"],  # Allow all headers
)

# Load model and tokenizer
tokenizer = AutoTokenizer.from_pretrained("microsoft/DialoGPT-medium")
model = AutoModelForCausalLM.from_pretrained("microsoft/DialoGPT-medium")

# Helper function for conversation
async def generate_response(input_text: str, history=[]):
    # Tokenize user input and append to conversation history
    new_user_input_ids = tokenizer.encode(input_text + tokenizer.eos_token, return_tensors="pt")
    bot_input_ids = torch.cat([torch.tensor(history, dtype=torch.long), new_user_input_ids], dim=-1) if history else new_user_input_ids

    # Generate response
    response_ids = model.generate(bot_input_ids, max_length=1000, pad_token_id=tokenizer.eos_token_id)
    response_text = tokenizer.decode(response_ids[:, bot_input_ids.shape[-1]:][0], skip_special_tokens=True)

    # Update history and return response
    history.extend(new_user_input_ids.tolist())
    return response_text, history

class ChatRequest(BaseModel):
    user_input: str

# Endpoint for chat
@app.post("/chat")
async def get_response(request: ChatRequest):
    user_input = request.user_input
    response, _ = await generate_response(user_input)
    return {"response": response}

# Run the server
# uvicorn app:app --reload
