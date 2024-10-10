import torch
from pytorch_transformers import BertModel, BertConfig, BertTokenizer
from torch import nn
import sys

# 自己下载模型相关的文件，并指定路径
config_path = './model/bert_config.json'
model_path = './model/pytorch_model.bin'
vocab_path = './model/vocab.txt'

 
# ——————构造模型——————
class BertTextNet(nn.Module):
    def __init__(self,code_length):
        super(BertTextNet, self).__init__()
 
        modelConfig = BertConfig.from_pretrained(config_path)
        self.textExtractor = BertModel.from_pretrained(
            model_path, config=modelConfig)
        embedding_dim = self.textExtractor.config.hidden_size
 
        self.fc = nn.Linear(embedding_dim, code_length)
        self.tanh = torch.nn.Tanh()
 
    def forward(self, tokens, segments, input_masks):
        output = self.textExtractor(tokens, token_type_ids=segments,
                                    attention_mask=input_masks)
        text_embeddings = output[0][:, 0, :]
        # output[0](batch size, sequence length, model hidden dimension)
 
        features = self.fc(text_embeddings)
        features = self.tanh(features)
        return features
 
req = ""
############# 入口参数检测
# if len(sys.argv) < 1 :
#     sys.exit()

# req = sys.argv[0]
#-----------------------end

 
textNet = BertTextNet(code_length=768)
 
# ——————输入处理——————
tokenizer = BertTokenizer.from_pretrained(vocab_path)
 
tokens, segments, input_masks = [], [], []
tokenized_text = tokenizer.tokenize(req)  # 用tokenizer对句子分词
indexed_tokens = tokenizer.convert_tokens_to_ids([1.2,2.2,3.3,4.4])  # 索引列表
tokens.append(indexed_tokens)
segments.append([0] * len(indexed_tokens))
input_masks.append([1] * len(indexed_tokens))
 
max_len = max([len(single) for single in tokens])  # 最大的句子长度
 
for j in range(len(tokens)):
    padding = [0] * (max_len - len(tokens[j]))
    tokens[j] += padding
    segments[j] += padding
    input_masks[j] += padding
# segments列表全0，因为只有一个句子1，没有句子2
# input_masks列表1的部分代表句子单词，而后面0的部分代表paddig，只是用于保持输入整齐，没有实际意义。
# 相当于告诉BertModel不要利用后面0的部分
 
# 转换成PyTorch tensors
tokens_tensor = torch.tensor(tokens)
segments_tensors = torch.tensor(segments)
input_masks_tensors = torch.tensor(input_masks)
 
print(tokens_tensor.dtype)
# ——————提取文本特征——————
text_hashCodes = textNet(tokens_tensor, segments_tensors, input_masks_tensors)  # text_hashCodes是一个32-dim文本特征
print(text_hashCodes)
