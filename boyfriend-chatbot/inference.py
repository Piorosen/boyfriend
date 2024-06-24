import transformers
import torch

def make_prompt(data, last: int, jinju_id: int):
    result = ""
    for item in data[-last:]:
        # Doremi
        if item.user_id == 1657801720:
            result += f"진주 : {item.text}\n"            
        else:
            result += f"차차핑 : {item.text}\n"
    
    return result


def get_model():
    model_id = "MLP-KTLim/llama-3-Korean-Bllossom-8B"
    pipeline = transformers.pipeline(
        "text-generation",
        model=model_id,
        model_kwargs={"torch_dtype": torch.bfloat16},
        # device_map="auto",
        device=0,
    )
    pipeline.model.eval()   
    return pipeline

def get_next_text(jinju_id, text_set, pipeline, size:int = 100):
    PROMPT = f'''
    채팅하는 인원은 차차핑과 진주 입니다.
    진주는 말을 하는것이 금지 되어 있습니다.
    차차핑은 오직 세 문장만 이야기하여야 하며, 이제 말을 해야할 차례 입니다.
    차차핑은 진주를 엄청나게 좋아하므로, 사랑의 멘트를 섞어서 이야기 합니다.
    차차핑은 최대한 이모티콘과 개발 용어를 섞어서 말을 합니다.
    다음의 대화를 예측하여 차차핑이 할 것 같은 이야기를 진행해주세요.
    {text_set[-1].text}
    '''
                
    instruction = make_prompt(text_set, size, jinju_id)

    messages = [
        {"role": "system", "content": f"{PROMPT}"},
        {"role": "chat", "content": f"{instruction}"}
        ]

    prompt = pipeline.tokenizer.apply_chat_template(
            messages, 
            tokenize=False, 
            add_generation_prompt=True
    )

    terminators = [
        pipeline.tokenizer.eos_token_id,
        pipeline.tokenizer.convert_tokens_to_ids("<|eot_id|>")
    ]

    outputs = pipeline(
        prompt,
        max_new_tokens=512,
        eos_token_id=terminators,
        do_sample=True,
        temperature=0.6,
        top_p=0.9
    )
    
    return print(outputs[0]["generated_text"][len(prompt):])
