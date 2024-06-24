
#%%
# # Import libraries
# from datasets import load_dataset
# from transformers import AutoModelForSequenceClassification, AutoTokenizer
# import torch
# from torch.utils.data import Dataset
# from transformers import TrainingArguments, Trainer
# from datasets import load_metric

# class NSMCDataset(Dataset):
#     def __init__(self, encodings, labels):
#         self.encoding = encodings
#         self.labels = labels

#     def __getitem__(self, idx):
#         data = {key: val[idx] for key, val in self.encoding.items()}
#         data['labels'] = torch.tensor(self.labels[idx]).long()

#         return data

#     def __len__(self):
#         return len(self.labels)

# def compute_metrics(pred):
#     labels = pred.label_ids
#     preds = pred.predictions.argmax(-1)

#     m1 = load_metric('accuracy')
#     m2 = load_metric('f1')

#     acc = m1.compute(predictions=preds, references=labels)['accuracy']
#     f1 = m2.compute(predictions=preds, references=labels)['f1']

#     return {'accuracy':acc, 'f1':f1}


# nsmc = load_dataset('nsmc', trust_remote_code=True)
# train_data = nsmc['train'].shuffle(seed=42).select(range(2000))
# test_data = nsmc['test'].shuffle(seed=42).select(range(2000))

# MODEL_NAME = 'bert-base-multilingual-cased'

# model = AutoModelForSequenceClassification.from_pretrained(MODEL_NAME, num_labels=2)
# tokenizer = AutoTokenizer.from_pretrained(MODEL_NAME)
# tokenizer.tokenize(train_data['document'][0])

# # %%
# train_encoding = tokenizer(
#     train_data['document'],
#     return_tensors='pt',
#     padding=True,
#     truncation=True
# )

# test_encoding = tokenizer(
#     test_data['document'],
#     return_tensors='pt',
#     padding=True,
#     truncation=True
# )
# # len(train_encoding['input_ids']), len(test_encoding['input_ids'])
# train_set = NSMCDataset(train_encoding, train_data['label'])
# test_set = NSMCDataset(test_encoding, test_data['label'])

# # %%
# training_args = TrainingArguments(
#     output_dir = './outputs', # model이 저장되는 directory
#     logging_dir = './logs', # log가 저장되는 directory
#     num_train_epochs = 10, # training epoch 수
#     per_device_train_batch_size=32,  # train batch size
#     per_device_eval_batch_size=32,   # eval batch size
#     logging_steps = 50, # logging step, batch단위로 학습하기 때문에 epoch수를 곱한 전체 데이터 크기를 batch크기로 나누면 총 step 갯수를 알 수 있다.
#     save_steps= 50, # 50 step마다 모델을 저장한다.
#     save_total_limit=2 # 2개 모델만 저장한다.
# )
# # %%
# device = torch.device("cuda" if torch.cuda.is_available() else 'cpu')
# model.to(device)

# trainer = Trainer(
#     model=model,
#     args=training_args,
#     train_dataset=train_set, # 학습 세트
#     eval_dataset=test_set, # 테스트 세트
#     compute_metrics=compute_metrics # metric 계산 함수
# )
# trainer.train()

# %%
# https://nkw011.github.io/nlp/tutorial4_Fine-tune_a_pretrained_model/