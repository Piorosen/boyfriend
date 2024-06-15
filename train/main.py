#%%
import torch
from transformers import GPT2LMHeadModel, GPT2Tokenizer, Trainer, TrainingArguments
from datasets import load_dataset
#%%
# 데이터셋 로드
dataset = load_dataset('daily_dialog')
train_dataset = dataset['train']
test_dataset = dataset['test']

# 모델과 토크나이저 로드
model_name = 'gpt2'
model = GPT2LMHeadModel.from_pretrained(model_name)
tokenizer = GPT2Tokenizer.from_pretrained(model_name)

# 데이터셋 토큰화
def tokenize_function(examples):
    return tokenizer(examples['dialog'], padding="max_length", truncation=True, max_length=128)

tokenized_train_dataset = train_dataset.map(tokenize_function, batched=True)
tokenized_test_dataset = test_dataset.map(tokenize_function, batched=True)

# 데이터셋 포맷 변경
train_dataset = tokenized_train_dataset.remove_columns(['dialog'])
train_dataset.set_format('torch')

test_dataset = tokenized_test_dataset.remove_columns(['dialog'])
test_dataset.set_format('torch')

# 학습 파라미터 설정
training_args = TrainingArguments(
    output_dir='./results',
    evaluation_strategy='epoch',
    num_train_epochs=1,
    per_device_train_batch_size=2,
    per_device_eval_batch_size=2,
    logging_dir='./logs',
    logging_steps=10,
)

# Trainer 설정
trainer = Trainer(
    model=model,
    args=training_args,
    train_dataset=train_dataset,
    eval_dataset=test_dataset,
)

# 학습
trainer.train()

# 모델 저장
model.save_pretrained('./gpt2-girlfriend-bot')
tokenizer.save_pretrained('./gpt2-girlfriend-bot')

# 추론 예제
def generate_text(prompt, model, tokenizer, max_length=50):
    inputs = tokenizer(prompt, return_tensors="pt")
    outputs = model.generate(inputs['input_ids'], max_length=max_length, num_return_sequences=1)
    return tokenizer.decode(outputs[0], skip_special_tokens=True)

prompt = "How was your day?"
generated_text = generate_text(prompt, model, tokenizer)
print(generated_text)