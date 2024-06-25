import transformers
import torch

def make_prompt(data, last: int, jinju_id: int):
    result = ""
    format_string = "%H:%M:%S"

    for item in data[-last:]:
        # Doremi
        if item.user_id == jinju_id:
            # result += f"차차 : {item.text}\n"            
            result += f"({item.created_at.strftime(format_string)}) 진주 : {item.text}\n"            
        else:
            # result += f"진주 : {item.text}\n"
            result += f"({item.created_at.strftime(format_string)}) 차차 : {item.text}\n"            
    
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
대화 참여자: 진주와 차차.

조건:
1. 진주는 이 대화에서 말을 할 수 없습니다.
2. 차차가 말하는 모든 문장이나 대화는 개발과 연관지어서 말합니다.
3. 현재 차차가 말할 차례입니다.
4. 최대한 귀엽고, 이모티콘과 C++ 언어 스타일 형태의 문장을 구성하세요.
5. 차차는 진주를 좋아하고 사모합니다.
6. 다음 문장 예측은 오직 한 문장만 수행하세요.

추가 요청: 차차가 할 법한 대화 내용을 예측하여 제시해 주세요. 차차의 대화는 가능한 한 창의적이고 랜덤한 요소를 포함해야 합니다.
대화 구조 : (시간) 참여자 명 : 대화 내용
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
    
    return outputs[0]["generated_text"][len(prompt):]
