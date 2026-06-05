# Part 2: Machine Learning, Deep Learning & Generative AI

---

## 🤖 Machine Learning

### Beginner

**Q1: What is the difference between supervised, unsupervised, and reinforcement learning?**

| Type | Labels? | Goal | Examples |
|------|---------|------|----------|
| Supervised | ✅ | Predict output from input | Classification, Regression |
| Unsupervised | ❌ | Find patterns/structure | Clustering, Dimensionality Reduction |
| Reinforcement | Reward signal | Maximize cumulative reward | Game AI, Robotics |

**Q2: Explain bias-variance tradeoff.**
- **Bias**: error from overly simple models (underfitting)
- **Variance**: error from overly complex models (overfitting)
- **Goal**: find the sweet spot that minimizes total error
- **Solutions**: cross-validation, regularization, ensemble methods

**Q3: What is overfitting and how do you prevent it?**
- Model memorizes training data, performs poorly on unseen data
- **Prevention**: cross-validation, regularization (L1/L2), dropout, early stopping, more data, data augmentation

**Q4: Explain precision, recall, and F1-score.**
```
Precision = TP / (TP + FP)   → "Of predicted positives, how many are correct?"
Recall    = TP / (TP + FN)   → "Of actual positives, how many did we find?"
F1-Score  = 2 × (P × R) / (P + R)  → harmonic mean
```
- **High precision needed**: spam detection (don't want false positives)
- **High recall needed**: cancer detection (don't want to miss positives)

**Q5: What is cross-validation?**
- **K-Fold**: split data into K folds, train on K-1, test on 1, rotate
- **Stratified K-Fold**: preserves class distribution in each fold
- Gives more reliable performance estimate than single train/test split

### Intermediate

**Q6: Explain gradient descent variants.**
```
θ = θ - α × ∇J(θ)
```
| Variant | Batch Size | Pros | Cons |
|---------|-----------|------|------|
| Batch GD | All data | Stable | Slow, memory heavy |
| Stochastic GD | 1 sample | Fast updates | Noisy |
| Mini-batch GD | 32-512 | Best of both | Needs tuning |

**Optimizers**: SGD → Momentum → RMSProp → **Adam** (most common)

**Q7: Explain regularization (L1 vs L2).**
```
L1 (Lasso): Loss + λ Σ|θᵢ|        → sparse weights (feature selection)
L2 (Ridge): Loss + λ Σθᵢ²         → small weights (prevents overfitting)
Elastic Net: L1 + L2 combined
```

**Q8: What are ensemble methods?**
- **Bagging** (Random Forest): train multiple models on bootstrap samples, average predictions → reduces **variance**
- **Boosting** (XGBoost, LightGBM): train models sequentially, each fixing previous errors → reduces **bias**
- **Stacking**: use predictions of base models as input to a meta-model

**Q9: How do you handle imbalanced datasets?**
- **Data level**: oversampling (SMOTE), undersampling, augmentation
- **Algorithm level**: class weights, cost-sensitive learning
- **Metric level**: use F1, AUC-ROC, precision-recall curve instead of accuracy
- **Ensemble**: balanced bagging, EasyEnsemble

**Q10: Explain feature engineering best practices.**
- **Numerical**: scaling (StandardScaler, MinMaxScaler), log transform, binning
- **Categorical**: one-hot encoding, target encoding, label encoding
- **Text**: TF-IDF, word embeddings, n-grams
- **Feature selection**: correlation analysis, mutual information, recursive feature elimination

### Advanced

**Q11: Explain the math behind logistic regression.**
```
P(y=1|x) = σ(wᵀx + b) = 1 / (1 + e^(-(wᵀx + b)))

Loss = -[y·log(ŷ) + (1-y)·log(1-ŷ)]   (Binary Cross-Entropy)
```
Decision boundary is linear. Coefficients are interpretable as log-odds.

**Q12: How does XGBoost work internally?**
- Additive tree model: `ŷ = Σ fₖ(x)` where each fₖ is a tree
- Objective: `Loss + Ω(f)` where Ω penalizes tree complexity
- Uses **second-order Taylor expansion** of loss for efficient split finding
- Key features: histogram-based splits, regularization, built-in handling of missing values, parallel tree construction

**Q13: Explain AUC-ROC in depth.**
- ROC plots **TPR** (recall) vs **FPR** (1 - specificity) at all thresholds
- AUC = probability that model ranks a random positive higher than a random negative
- AUC = 0.5 → random, AUC = 1.0 → perfect
- Use **PR-AUC** for highly imbalanced datasets

---

## 🧠 Deep Learning

### Beginner

**Q1: What is a neural network?**
- Layers of interconnected nodes (neurons)
- Each neuron: `output = activation(Σ(wᵢxᵢ) + b)`
- **Forward pass**: input → hidden layers → output
- **Backward pass**: compute gradients via chain rule → update weights

**Q2: What are activation functions and why do we need them?**

| Function | Formula | Range | Use Case |
|----------|---------|-------|----------|
| ReLU | max(0, x) | [0, ∞) | Hidden layers (default) |
| Sigmoid | 1/(1+e⁻ˣ) | (0, 1) | Binary classification output |
| Softmax | eˣⁱ/Σeˣʲ | (0, 1) | Multi-class output |
| Tanh | (eˣ-e⁻ˣ)/(eˣ+e⁻ˣ) | (-1, 1) | RNNs |
| GELU | x·Φ(x) | (-0.17, ∞) | Transformers |

Without activation functions, stacking layers collapses to a single linear transformation.

**Q3: What is backpropagation?**
- Algorithm to compute gradients of loss w.r.t. each weight using the **chain rule**
- Goes backward from output to input layer
- Enables efficient gradient computation in deep networks
- `∂L/∂w = ∂L/∂ŷ × ∂ŷ/∂z × ∂z/∂w`

### Intermediate

**Q4: Explain CNNs and their key operations.**
- **Convolution**: sliding filters extract spatial features (edges, textures)
- **Pooling**: downsampling (max/avg) for translation invariance
- **Architecture**: Conv → ReLU → Pool → ... → Flatten → Dense
- **Parameter sharing**: same filter applied across entire input → fewer parameters

**Q5: Explain RNNs, LSTMs, and GRUs.**
```
RNN:  hₜ = tanh(Wₕhₜ₋₁ + Wₓxₜ + b)
      Problem: vanishing/exploding gradients

LSTM: Adds forget gate, input gate, output gate
      → can learn long-range dependencies

GRU:  Simplified LSTM with reset + update gates
      → fewer parameters, similar performance
```

**Q6: What is batch normalization?**
- Normalizes layer inputs to zero mean, unit variance within each mini-batch
- Benefits: faster training, higher learning rates, slight regularization
- Applied **before** or **after** activation (debated)
- **Layer Norm** preferred in Transformers (independent of batch size)

**Q7: Explain dropout.**
- Randomly zeros out neurons with probability `p` during training
- Forces network to learn redundant representations
- At inference: all neurons active, weights scaled by `(1-p)`
- Acts as **ensemble** of sub-networks

### Advanced

**Q8: Explain the Transformer architecture in detail.**
```
Input → Embedding + Positional Encoding
      → [Multi-Head Self-Attention → Add & Norm → FFN → Add & Norm] × N
      → Output
```

**Self-Attention**:
```
Attention(Q, K, V) = softmax(QKᵀ / √dₖ) × V
```
- Q, K, V are linear projections of input
- `√dₖ` scaling prevents softmax saturation
- **Multi-head**: multiple attention heads capture different relationship patterns

**Q9: Explain positional encoding.**
```
PE(pos, 2i)   = sin(pos / 10000^(2i/d))
PE(pos, 2i+1) = cos(pos / 10000^(2i/d))
```
- Transformers have no inherent position awareness (unlike RNNs)
- Sinusoidal encoding allows model to learn relative positions
- Modern models use **learned** positional embeddings or **RoPE** (Rotary Position Embedding)

**Q10: What is the difference between pre-training and fine-tuning?**
- **Pre-training**: train on large unlabeled corpus (self-supervised) → learns general language understanding
- **Fine-tuning**: adapt pre-trained model to specific task with labeled data
- **Types**: full fine-tuning, LoRA, QLoRA, prefix tuning, prompt tuning

---

## ✨ Generative AI

### Beginner

**Q1: What are LLMs?**
- Large Language Models trained on massive text corpora
- Based on Transformer architecture (decoder-only for GPT, encoder-decoder for T5)
- Generate text by predicting next token: `P(xₜ | x₁, ..., xₜ₋₁)`
- Examples: GPT-4, Claude, Gemini, LLaMA, Mistral

**Q2: What is the difference between GPT and BERT?**

| Feature | GPT | BERT |
|---------|-----|------|
| Architecture | Decoder-only | Encoder-only |
| Training | Causal LM (left-to-right) | Masked LM (bidirectional) |
| Best for | Generation | Understanding/classification |
| Attention | Causal (masked future) | Full bidirectional |

**Q3: Explain temperature and top-k/top-p sampling.**
```
temperature = 0.0 → deterministic (greedy)
temperature = 0.7 → balanced creativity
temperature = 1.5 → very creative/random

top-k: sample from top k tokens
top-p (nucleus): sample from smallest set whose cumulative prob ≥ p
```

### Intermediate

**Q4: What is prompt engineering?**
- **Zero-shot**: direct instruction without examples
- **Few-shot**: provide examples in the prompt
- **Chain-of-Thought**: "Let's think step by step" → improves reasoning
- **ReAct**: Reasoning + Acting (interleave thought and tool use)
- **System prompts**: set persona, constraints, output format

**Q5: Explain tokenization in LLMs.**
- **BPE** (Byte-Pair Encoding): iteratively merge most frequent character pairs
- **SentencePiece**: language-agnostic, works on raw text
- **WordPiece**: used by BERT, similar to BPE
- Vocab size tradeoff: larger = more expressiveness but more memory

**Q6: What are embeddings?**
```python
# Dense vector representations of text
from sentence_transformers import SentenceTransformer
model = SentenceTransformer('all-MiniLM-L6-v2')
embeddings = model.encode(["AI is transformative", "ML is powerful"])
# similarity = cosine(embeddings[0], embeddings[1])
```
- Capture semantic meaning in vector space
- Similar concepts → close vectors (high cosine similarity)
- Used for: search, RAG, clustering, recommendation

### Advanced

**Q7: Explain RLHF (Reinforcement Learning from Human Feedback).**
```
1. Supervised Fine-Tuning (SFT): train on high-quality demonstrations
2. Reward Model Training: learn human preferences from comparisons
3. PPO Optimization: fine-tune SFT model to maximize reward
   while staying close to original (KL penalty)
```
- Alternative: **DPO** (Direct Preference Optimization) — simpler, no reward model needed

**Q8: What are the different fine-tuning approaches?**

| Method | Parameters Updated | Memory | Quality |
|--------|-------------------|--------|---------|
| Full Fine-tuning | All | Very High | Best |
| LoRA | Low-rank adapters | Low | Very Good |
| QLoRA | LoRA + quantization | Very Low | Good |
| Prefix Tuning | Prepended vectors | Low | Good |
| Prompt Tuning | Soft prompts only | Minimal | Moderate |

**Q9: Explain model quantization.**
- Reduce precision: FP32 → FP16 → INT8 → INT4
- **GPTQ**: post-training quantization using calibration data
- **GGUF**: format for CPU inference (llama.cpp)
- **AWQ**: activation-aware quantization
- Tradeoff: smaller model + faster inference vs. slight quality loss

**Q10: What are hallucinations and how to mitigate them?**
- Model generates plausible but factually incorrect information
- **Mitigation**: RAG (ground in real data), fine-tuning on domain data, temperature control, constrained decoding, fact-checking chains, citation requirements
