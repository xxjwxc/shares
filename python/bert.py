import torch
from pytorch_transformers import BertModel, BertConfig, BertTokenizer
from torch import nn
import sys
from  Ashare import *

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
 

def get_dc_price(code, end_date='',count=10):        #对外暴露只有唯一函数，这样对用户才是最友好的  
    xcode= code.replace('.XSHG','').replace('.XSHE','')                      #证券代码编码兼容处理 
    xcode='sh'+xcode if ('XSHG' in code)  else  'sz'+xcode  if ('XSHE' in code)  else code     
    return get_price_day_dc(xcode,end_date=end_date,count=count,frequency='1d')   #备用                    

code = "sh000001" 
############# 入口参数检测
if len(sys.argv) < 2 :
    sys.exit()

code = sys.argv[1]
#-----------------------end

 
textNet = BertTextNet(code_length=768)
 
df=get_dc_price(code,count=200);      #日线数据获取  1d:1天  4h:4小时   60m: 60分钟    15m:15分钟

# ——————输入处理——————
tokenizer = BertTokenizer.from_pretrained(vocab_path)
 
texts = [df.percent,df.close,df.volume]
tokens, segments, input_masks = [], [], []
for text in texts:
    #tokenized_text = tokenizer.tokenize(text)  # 用tokenizer对句子分词
    indexed_tokens = tokenizer.convert_tokens_to_ids(text)  # 索引列表
    tokens.append(indexed_tokens)
    segments.append([0] * len(indexed_tokens))
    input_masks.append([1] * len(indexed_tokens))#input_masks有什么用？
 
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
 
# ——————提取文本特征——————
text_hashCodes = textNet(tokens_tensor, segments_tensors, input_masks_tensors)  # text_hashCodes是一个32-dim文本特征
data = json.dumps(text_hashCodes.tolist())
print(json.loads(data))
